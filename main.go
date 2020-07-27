package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

const UP_TOK = "UP_TOK"

func main() {
	upTok := os.Getenv(UP_TOK)
	if upTok == "" {
		log.Fatalf("%s not set", UP_TOK)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.up.com.au/api/v1/accounts", nil)
  req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", upTok))
	if err != nil {
		log.Fatal("error creating request for accounts")
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(errors.Wrap(err, "fetching accounts"))
	}

	fmt.Printf("%+v\n", res)
}
