package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

type config struct {
	businessSymbol string
}

func main() {
	var conf config

	app := &cli.App{
		Name:  "Finance",
		Usage: "Analyse investment viability of a publicly traded business",
		Action: func(c *cli.Context) error {
			return run(conf)

		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "symbol",
				Aliases:     []string{"s"},
				Required:    true,
				Destination: &conf.businessSymbol,
				Usage:       "Business symbol of the company to analyse. Example: AAPC for Apple",
			},
		},
	}

	app.Run(os.Args)

}

func run(conf config) error {
	// m := YahooAPIClient{}
	m := YahooMockClient{}

	y, err := m.GetIncomeStatement(conf.businessSymbol)

	if err != nil {
		return err
	}

	s, err := m.GetStockInfo(conf.businessSymbol)

	if err != nil {
		return err
	}

	NewIncomeStatement(y, s)

	return nil
}
