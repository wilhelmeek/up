package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
	"github.com/wilhelmeek/help/internal/upapi"
)

func main() {
	token := os.Getenv("UP_TOK")
	if token == "" {
		log.Fatal("Please make sure you've sourced a valid UP_TOK")
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "accounts",
				Aliases: []string{"a"},
				Usage:   "list accounts",
				Action:  listAccounts,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func listAccounts(cliCtx *cli.Context) error {
	config := upapi.NewConfiguration()
	config.AddDefaultHeader(
		"Authorization",
		fmt.Sprintf("Bearer %s", os.Getenv("UP_TOK")),
	)
	ctx := context.Background()
	up := upapi.NewAPIClient(config)

	accs, _, err := up.AccountsApi.AccountsGet(ctx, &upapi.AccountsGetOpts{
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
			fmt.Println(fmt.Sprintf("%s: $%s", li.name, li.val))
		}
		fmt.Println()
		fmt.Println(fmt.Sprintf("Total: $%s", total))
	}

	return nil
}
