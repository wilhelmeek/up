package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/AlecAivazis/survey/v2"
	"github.com/antihax/optional"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/peterstace/date"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
	"github.com/wilhelmeek/up/internal/upapi"
)

type Up struct {
	client *upapi.APIClient
}

func NewUp() *Up {
	token := os.Getenv("UP_TOK")
	if token == "" {
		log.Fatal("Please make sure you've sourced a valid UP_TOK")
	}

	config := upapi.NewConfiguration()
	config.AddDefaultHeader(
		"Authorization",
		fmt.Sprintf("Bearer %s", os.Getenv("UP_TOK")),
	)
	upClient := upapi.NewAPIClient(config)

	return &Up{
		client: upClient,
	}
}

func main() {
	up := NewUp()
	app := &cli.App{
		Name:  "Unofficial Up CLI",
		Usage: "Some handy Up shortcuts",
		Commands: []*cli.Command{
			{
				Name:    "balances",
				Aliases: []string{"b"},
				Usage:   "List account balances",
				Action:  up.listBalances,
			},
			{
				Name:    "transactions",
				Aliases: []string{"t"},
				Usage:   "Search for transactions",
				Action:  up.listTransactions,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (up *Up) listTransactions(cliCtx *cli.Context) error {
	ctx := context.Background()
	accs, _, err := up.client.AccountsApi.AccountsGet(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "fetching accounts")
	}

	accountNames := []string{}
	accountNameToId := make(map[string]string)
	for _, acc := range accs.Data {
		accountNames = append(accountNames, acc.Attributes.DisplayName)
		accountNameToId[acc.Attributes.DisplayName] = acc.Id
	}

	// TODO: Use fuzzyfind-go instead
	selectedAccountName := ""
	prompt := &survey.Select{
		Message: "Choose an account:",
		Options: accountNames,
	}
	if survey.AskOne(prompt, &selectedAccountName) != nil {
		return errors.Wrap(err, "selecting account")
	}

	// TODO: Stream pages in while searching
	txns, _, err := up.client.TransactionsApi.AccountsAccountIdTransactionsGet(
		ctx,
		accountNameToId[selectedAccountName],
		&upapi.AccountsAccountIdTransactionsGetOpts{
			PageSize: optional.NewInt32(100),
		},
	)

	lineItems := []string{}
	for _, tx := range txns.Data {
		created := date.FromTime(tx.Attributes.CreatedAt).String()
		lineItems = append(lineItems, fmt.Sprintf(
			"%s %s %s",
			created,
			tx.Attributes.Description,
			tx.Attributes.Amount.Value,
		))
	}

	_, err = fuzzyfinder.Find(
		lineItems,
		func(i int) string {
			return lineItems[i]
		},
		fuzzyfinder.WithMode(fuzzyfinder.ModeSmart),
	)

	return errors.Wrap(err, "fuzzy finding")
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

		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
		for _, li := range list {
			fmt.Fprintln(writer, fmt.Sprintf("%s:\t$%s", li.name, li.val))
		}
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, fmt.Sprintf("Total:\t\t$%s", total))
		writer.Flush()
	}

	return nil
}
