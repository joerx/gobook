package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joerx/gobook/exercises/2.2/massconv"
	"github.com/joerx/gobook/exercises/2.2/tempconv"
	"github.com/urfave/cli/v2"
)

type cvFunc func(f float64) interface{}

type cvCommand struct {
	Name   string
	Usage  string
	Action cvFunc
}

func main() {

	cs := []cvCommand{
		{
			Name:  "ctof",
			Usage: "Convert Celsius to Fahrenheit",
			Action: func(f float64) interface{} {
				return tempconv.CToF(tempconv.Celsius(f))
			},
		},
		{
			Name:  "ftoc",
			Usage: "Convert Fahrenheit to Celsius",
			Action: func(f float64) interface{} {
				return tempconv.FToC(tempconv.Fahrenheit(f))
			},
		},
		{
			Name:  "ktoc",
			Usage: "Convert Kelvin to Celsius",
			Action: func(f float64) interface{} {
				return tempconv.KToC(tempconv.Kelvin(f))
			},
		},
		{
			Name:  "lbtokg",
			Usage: "Convert Pounds to Kilogram",
			Action: func(f float64) interface{} {
				return massconv.LbtoKg(massconv.Pound(f))
			},
		},
		{
			Name:  "kgtolb",
			Usage: "Convert Kilogram to Pounds",
			Action: func(f float64) interface{} {
				return massconv.KgtoLb(massconv.Kilogram(f))
			},
		},
	}

	commands := make([]*cli.Command, len(cs))

	for i, c := range cs {
		commands[i] = &cli.Command{
			Name:   c.Name,
			Usage:  c.Usage,
			Action: fh(c.Action),
		}
	}

	app := cli.App{Commands: commands}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// fh returns a cli.ActionFunc that converts the first argument into a float64, passes it to a cvFunc
// and formats the conversion result
func fh(f cvFunc) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		fVal, err := strconv.ParseFloat(c.Args().First(), 64)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", f(fVal))
		return nil
	}
}
