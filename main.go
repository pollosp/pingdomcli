package main

import (
	"log"

	"os"

	"github.com/pollosp/pingdomcli/internal/pingdomcli"

	"github.com/russellcardullo/go-pingdom/pingdom"
	"github.com/urfave/cli"
)

func main() {
	pingdomEmail := os.Getenv("PINGDOM_EMAIL")
	pingdomPassword := os.Getenv("PINGDOM_PASSWORD")
	pingdomAPIKey := os.Getenv("PINGDOM_APIKEY")

	client := pingdom.NewClient(pingdomEmail, pingdomPassword, pingdomAPIKey)

	app := cli.NewApp()
	app.Name = "pindomcli"
	app.Usage = "Set ENV vars for PINGDOM_EMAIL , PINGDOM_PASSWORD and PINGDOM_APIKEY"
	app.UsageText = "pingdomcli checkName checkHostName userID."
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		var checkName string
		var checkHostName string
		var userID string

		if c.NArg() > 0 {
			checkName = c.Args().Get(0)
			checkHostName = c.Args().Get(1)
			userID = c.Args().Get(2)
		}

		actions.CreateHTTPCheck(*client, checkName, checkHostName, userID)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
