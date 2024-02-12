package calculator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type calculatorSuite struct {
	suite.Suite
	calculator Calculator
}

func TestCalculator(t *testing.T) {
	suite.Run(t, &calculatorSuite{})
}

func (s *calculatorSuite) SetupTest() {
	s.calculator = *CalculatorFactory(&MockCLI{})
}

func (s *calculatorSuite) TestInvestmentPeriodInterestGain() {
	require := require.New(s.T())

	testCases := []struct {
		testName     string
		original     float64
		periodMonths int
		interestRate float64
		expected     float64
	}{
		{
			"standard",
			float64(10000),
			12,
			float64(2),
			float64(200),
		},
		{
			"high_months",
			float64(10000),
			36,
			float64(2),
			float64(600),
		},
		{
			"low_months",
			float64(10000),
			6,
			float64(2),
			float64(100),
		},
		{
			"odd_months",
			float64(10000),
			7,
			float64(2),
			float64(116.66666666666669),
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.testName, func(t *testing.T) {
			actual := s.calculator.investmentPeriodInterestGain(tc.original, tc.periodMonths, tc.interestRate)
			require.Equal(tc.expected, actual)
		})
	}
}

func (s *calculatorSuite) TestGetInterestFrequencyInMonths() {
	require := require.New(s.T())

	testCases := []struct {
		testName       string
		frequency      string
		investmentTerm int
		expected       int
	}{
		{
			"standard",
			"Quarterly",
			36,
			3,
		},
		{
			"at_maturity",
			"At Maturity",
			36,
			36,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.testName, func(t *testing.T) {
			actual := s.calculator.getInterestFrequencyInMonths(tc.frequency, tc.investmentTerm)
			require.Equal(tc.expected, actual)
		})
	}
}

func (s *calculatorSuite) TestGatherInputs() {
	require := require.New(s.T())

	expectedDeposit := 10000
	expectedInterestRate := float64(1.1)
	expectedInvestmentTerm := 12
	expectedInterestFrequency := "Monthly"

	deposit, interestRate, investmentTerm, interestFrequency, err := s.calculator.gatherInputs()

	require.NoError(err)
	require.Equal(expectedDeposit, deposit)
	require.Equal(expectedInterestRate, interestRate)
	require.Equal(expectedInvestmentTerm, investmentTerm)
	require.Equal(expectedInterestFrequency, interestFrequency)
}

func (s *calculatorSuite) TestRunCalculator() {
	require := require.New(s.T())
	expected := "$10,111"

	actual, err := s.calculator.RunCalculator()

	require.NoError(err)
	require.Equal(expected, actual)
}

type MockCLI struct{}

func (c *MockCLI) PromptWithValidation(promptText string, validatorFunc func(input string) error) (string, error) {
	switch promptText {
	case "Starting deposit amount (in dollars, eg. 10000)":
		return "10000", nil
	case "Investment term (in months, eg. 36)":
		return "12", nil
	case "Interest rate (percent p.a, eg. 1.10)":
		return "1.1", nil
	}
	return "", errors.New("test error")
}

func (c *MockCLI) PromptOptions(promptText string, options []string) (string, error) {
	return options[0], nil
}

func (c *MockCLI) ValidateInt(input string) error {
	return nil
}

func (c *MockCLI) ValidateFloat(input string) error {
	return nil
}
