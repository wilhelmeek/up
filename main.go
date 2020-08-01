package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/pkg/errors"
	"github.com/wilhelmeek/help/internal/upapi"
)

func main() {
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
    log.Fatal(errors.Wrap(err, "fetching accounts"))
  }
  fmt.Printf("%v", accs)
}
