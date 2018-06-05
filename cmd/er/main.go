package main

import (
	"exchange-rates"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const defaultFromt = "USD"

var defaultTo = []string{"CNY", "EUR", "GBP", "CAD", "AUD", "JPY"}

func main() {
	from, amount, to := parseArgs(os.Args)
	core.Query(from, amount, to)
}

func parseArgs(args []string) (string, float32, []string) {
	var from string
	var amount float32
	var to []string

	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			fmt.Println("TODO: display help info")
		} else {
			if len(args[1]) != 3 {
				fmt.Println("Please input a valid currency")
				os.Exit(1)
			} else {
				from = args[1]
				amount = 1
				to = defaultTo
			}
		}
	}

	if len(args) >= 3 {
		from = args[1]
		s, err := strconv.ParseFloat(args[2], 32)
		if err != nil {
			fmt.Println("Please input valid amount")
		}
		amount = float32(s)

		if len(args) == 3 {
			to = defaultTo
		} else if len(args) == 4 {
			to = strings.Split(args[3], ",")
		}
	}

	if len(args) == 1 {
		from = defaultFromt
		amount = 1
		to = defaultTo
	}

	return from, amount, to
}