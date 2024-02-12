package cli

import (
	"errors"
	"slices"
	"strconv"
)

var DEFAULT_COMMANDS = []string{"exit"}

func (c *CLI) isDefaultCommand(input string) bool {
	return slices.Contains(DEFAULT_COMMANDS, input)
}

func (c *CLI) ValidateInt(input string) error {
	if c.isDefaultCommand(input) {
		return nil
	}

	if _, err := strconv.Atoi(input); err != nil {
		return errors.New("please enter a whole number")
	}
	return nil
}

func (c *CLI) ValidateFloat(input string) error {
	if c.isDefaultCommand(input) {
		return nil
	}

	if _, err := strconv.ParseFloat(input, 64); err != nil {
		return errors.New("please enter a number")
	}
	return nil
}
