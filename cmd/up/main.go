package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/antihax/optional"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/peterstace/date"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
	"github.com/wilhelmeek/up/internal/upapi"
	"golang.org/x/text/currency"
)

type Up struct {
	client *upapi.APIClient
	writer *tabwriter.Writer
}

func NewUp() *Up {
	token := os.Getenv("UP_TOK")
	if token == "" {
		log.Fatal("Please make sure you've sourced UP_TOK")
	}

	config := upapi.NewConfiguration()
	config.AddDefaultHeader(
		"Authorization",
		fmt.Sprintf("Bearer %s", os.Getenv("UP_TOK")),
	)
	upClient := upapi.NewAPIClient(config)

	return &Up{
		client: upClient,
		writer: tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight),
	}
}

func main() {
	up := NewUp()
	app := &cli.App{
		Name:  "Unofficial Up CLI",
		Usage: "Some Handy Up Shortcuts",
		Commands: []*cli.Command{
			{
				Name:    "status",
				Aliases: []string{"s"},
				Usage:   "Check API status",
				Action:  up.listAPIStatus,
			},

			{
				Name:    "balances",
				Aliases: []string{"b"},
				Usage:   "List Account Balances",
				Action:  up.listBalances,
			},
			{
				Name:    "transactions",
				Aliases: []string{"t"},
				Usage:   "Fuzzy Find Transactions",
				Action:  up.listTransactions,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (up *Up) listAPIStatus(cliCtx *cli.Context) error {
	pingResp, httpResp, err := up.client.UtilityEndpointsApi.UtilPingGet(context.Background())
	if err != nil {
		if httpResp.StatusCode == http.StatusUnauthorized {
			return fmt.Errorf("UP_TOK not authorized")
		}

		return errors.Wrap(err, "determining api status")
	}

	fmt.Println(fmt.Sprintf("%s", pingResp.Meta.StatusEmoji))
	return nil
}

func (up *Up) listTransactions(cliCtx *cli.Context) error {
	ctx := context.Background()
	accsResp, _, err := up.client.AccountsApi.AccountsGet(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "fetching accounts")
	}

	accs := accsResp.Data
	ind, err := fuzzyfinder.Find(
		accs,
		func(i int) string {
			return accs[i].Attributes.DisplayName
		},
		fuzzyfinder.WithMode(fuzzyfinder.ModeSmart),
	)
	if err != nil {
		return errors.Wrap(err, "selecting account")
	}

	selectedAccountId := accs[ind].Id

	txnsResp, _, err := up.client.TransactionsApi.AccountsAccountIdTransactionsGet(
		ctx,
		selectedAccountId,
		&upapi.AccountsAccountIdTransactionsGetOpts{
			PageSize: optional.NewInt32(100),
		},
	)
	txns := txnsResp.Data

	ind, err = fuzzyfinder.Find(
		txns,
		func(i int) string {
			if i == -1 {
				return ""
			}
			tx := txns[i]
			created := date.FromTime(tx.Attributes.CreatedAt).String()
			return fmt.Sprintf(
				"%s %s %s",
				created,
				tx.Attributes.Description,
				mustMoneyToString(tx.Attributes.Amount),
			)
		},
		fuzzyfinder.WithPreviewWindow(
			func(i, w, h int) string {
				if i == -1 {
					return ""
				}
				return formatTXNAtIndex(i, txns)
			},
		),
	)
	if err != nil {
		return errors.Wrap(err, "selecting transaction")
	}

	fmt.Print(formatTXNAtIndex(ind, txns))
	return nil
}

func formatTXNAtIndex(i int, txns []upapi.TransactionResource) string {
	tx := txns[i]
	amount := mustMoneyToString(tx.Attributes.Amount)
	direction := "IN"
	if tx.Attributes.Amount.ValueInBaseUnits < 0 {
		direction = "OUT"
	}

	roundUp := "No Round-Up Occurred"
	if tx.Attributes.RoundUp != nil {
		roundUp = mustMoneyToString(tx.Attributes.RoundUp.Amount)
	}

	return fmt.Sprintf(
		"%s\n\n🔄 %s\n💵 %s\n👆 %s\n%s %s\n",
		tx.Attributes.Description,
		direction,
		amount,
		roundUp,
		statusToEmoji(tx.Attributes.Status),
		tx.Attributes.Status,
	)
}

func statusToEmoji(s upapi.TransactionStatusEnum) string {
	if s == upapi.SETTLED {
		return "✅"
	}
	return "❓"
}

func mustMoneyToString(m upapi.MoneyObject) string {
	unit := currency.MustParseISO(m.CurrencyCode)
	amountFlt, err := strconv.ParseFloat(m.Value, 64)
	if err != nil {
		log.Fatalf("error parsing money value %s", m.Value)
	}

	return fmt.Sprintf("%s", currency.Symbol(unit.Amount(amountFlt)))
}

func (up *Up) listBalances(cliCtx *cli.Context) error {
	ctx := context.Background()
	accs, _, err := up.client.AccountsApi.AccountsGet(ctx, &upapi.AccountsGetOpts{
		PageSize: optional.NewInt32(10),
	})

	if err != nil {
		return errors.Wrap(err, "fetching accounts")
	} else {
		total := decimal.NewFromInt(0)
		type listItem struct {
			name string
			val  decimal.Decimal
		}
		var list []listItem
		for _, acc := range accs.Data {
			bal, err := decimal.NewFromString(acc.Attributes.Balance.Value)
			if err != nil {
				return errors.Wrap(err, "parsing balance")
			}
			if bal.GreaterThan(decimal.Zero) {
				list = append(list, listItem{name: acc.Attributes.DisplayName, val: bal})
				total = total.Add(bal)
			}
		}

		for _, li := range list {
			fmt.Fprintln(up.writer, fmt.Sprintf("%s:\t$%s", li.name, li.val))
		}
		fmt.Fprintln(up.writer)
		fmt.Fprintln(up.writer, fmt.Sprintf("Total:\t\t$%s", total))
		up.writer.Flush()
	}

	return nil
}
