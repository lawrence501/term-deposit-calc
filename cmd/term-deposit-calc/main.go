package main

import (
	"log"
	"term-deposit-calc/internal/calculator"
	"term-deposit-calc/internal/cli"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	cli := cli.CLIFactory()
	calculator := calculator.CalculatorFactory(cli)

	log.Println("\nWelcome to the term deposit calculator. Please follow the prompts to use the calculator. Input 'exit' at any time to exit the calculator.")
	for {
		result, err := calculator.RunCalculator()
		if err != nil {
			log.Printf("Error running calculator: %s\n\n", err.Error())
			continue
		}
		log.Printf("Final balance: %s\n\n", result)
	}
}
