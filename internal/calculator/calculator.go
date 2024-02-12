package calculator

import (
	"fmt"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var INTEREST_FREQUENCY_TO_MONTHS = map[string]int{
	"Monthly":   1,
	"Quarterly": 3,
	"Annually":  12,
}

// How many months the base interest rate is calculated for (eg. per annum = 12)
var INTEREST_RATE_MONTHS = float64(12.0)

type Calculator struct {
	clier CLIer
}

type CLIer interface {
	PromptWithValidation(promptText string, validatorFunc func(input string) error) (string, error)
	PromptOptions(promptText string, options []string) (string, error)
	ValidateInt(input string) error
	ValidateFloat(input string) error
}

func CalculatorFactory(clier CLIer) *Calculator {
	return &Calculator{clier}
}

func (c Calculator) RunCalculator() (string, error) {
	deposit, interestRate, investmentTerm, interestFrequency, err := c.gatherInputs()
	if err != nil {
		return "", err
	}

	totalMonthsRemaining := investmentTerm
	runningTotal := float64(deposit)
	interestFrequencyMonths := c.getInterestFrequencyInMonths(interestFrequency, investmentTerm)
	for totalMonthsRemaining > 0 {
		monthsThisInterestPeriod := min(totalMonthsRemaining, interestFrequencyMonths)
		totalMonthsRemaining -= monthsThisInterestPeriod

		runningTotal += c.investmentPeriodInterestGain(runningTotal, monthsThisInterestPeriod, interestRate)
	}

	roundedTotal := int(runningTotal + 0.5)

	p := message.NewPrinter(language.English)
	return p.Sprintf("$%d", roundedTotal), nil
}

func (c Calculator) investmentPeriodInterestGain(original float64, periodMonths int, interestRate float64) float64 {
	interestRateDivisor := INTEREST_RATE_MONTHS / float64(periodMonths)
	periodInterestRate := interestRate / interestRateDivisor
	return original * periodInterestRate / 100
}

func (c Calculator) getInterestFrequencyInMonths(frequency string, investmentTerm int) int {
	if frequency == "At Maturity" {
		return investmentTerm
	}
	return INTEREST_FREQUENCY_TO_MONTHS[frequency]
}

func (c Calculator) gatherInputs() (int, float64, int, string, error) {
	depositStr, err := c.clier.PromptWithValidation("Starting deposit amount (in dollars, eg. 10000)", c.clier.ValidateInt)
	if err != nil {
		return 0, 0.0, 0, "", err
	}
	deposit, err := strconv.Atoi(depositStr)
	if err != nil {
		return 0, 0.0, 0, "", fmt.Errorf("failed to parse deposit as a whole dollar amount: %w", err)
	}

	interestRateStr, err := c.clier.PromptWithValidation("Interest rate (percent p.a, eg. 1.10)", c.clier.ValidateFloat)
	if err != nil {
		return 0, 0.0, 0, "", err
	}
	interestRate, err := strconv.ParseFloat(interestRateStr, 64)
	if err != nil {
		return 0, 0.0, 0, "", fmt.Errorf("failed to parse interest rate as a percentage: %w", err)
	}

	investmentTermStr, err := c.clier.PromptWithValidation("Investment term (in months, eg. 36)", c.clier.ValidateInt)
	if err != nil {
		return 0, 0.0, 0, "", err
	}
	investmentTerm, err := strconv.Atoi(investmentTermStr)
	if err != nil {
		return 0, 0.0, 0, "", fmt.Errorf("failed to parse investment term as months: %w", err)
	}

	interestFrequency, err := c.clier.PromptOptions("Interest payment frequency", []string{"Monthly", "Quarterly", "Annually", "At Maturity"})
	if err != nil {
		return 0, 0.0, 0, "", err
	}

	return deposit, interestRate, investmentTerm, interestFrequency, nil
}
