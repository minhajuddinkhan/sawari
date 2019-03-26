package commands

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/jung-kurt/gofpdf"
	"github.com/minhajuddinkhan/sawari"
	"github.com/urfave/cli"
)

var inputEnvPath string

// Form Creates a fuel receipt form
var Form = cli.Command{
	Name:  "form",
	Usage: "Creates a fuel receipt form",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "env, e",
			Usage:       "flag to enter env path",
			Destination: &inputEnvPath,
		},
	},
	Action: func(c *cli.Context) error {
		var creds sawari.Creds

		serial := 0
		totalAmount := 0
		wantToContinue := true

		entries := []sawari.Entry{}

		err := godotenv.Load("sawari.env")
		if err != nil {
			if !filepath.IsAbs(inputEnvPath) {
				inputEnvPath, err = filepath.Abs(inputEnvPath)
				if err != nil {
					return err
				}
			}
			err := godotenv.Load(inputEnvPath)
			if err != nil {
				color.Yellow("unable to find env vars for your personal info")
			}
		}
		creds = sawari.NewCreds()

		for wantToContinue {
			serial++
			entry := sawari.Entry{Serial: serial}
			entry.Date = sawari.TakeInput("Enter Date:")
			entry.Description = sawari.TakeInput("Enter Description:")

			valid := false
			for !valid {
				entry.Amount = sawari.TakeInput("Enter Amount:")
				amount, err := strconv.Atoi(entry.Amount)
				if err != nil {
					color.Red("Please enter a valid amount :(")
					continue
				}
				totalAmount += amount
				valid = true
			}
			entries = append(entries, entry)
			text := sawari.TakeInput("Do you want add more entries? [y/N]")
			if text != "y" {
				fmt.Println("Thanks!")
				break
			}
		}

		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		err = sawari.MakeHeader(pdf, creds)
		if err != nil {
			return err
		}
		sawari.MakeBody(pdf, entries, totalAmount)
		oFilePath := "form.pdf"
		err = pdf.OutputFileAndClose(oFilePath)
		if err != nil {
			log.Fatal(err)
		}
		fp, _ := filepath.Abs(oFilePath)
		color.Blue("Your pdf can be found at %s", fp)
		return nil
	},
}
