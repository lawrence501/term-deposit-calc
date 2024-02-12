package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

type CLI struct{}

func CLIFactory() *CLI {
	return &CLI{}
}

func (c *CLI) PromptWithValidation(promptText string, validatorFunc func(input string) error) (string, error) {
	prompt := promptui.Prompt{
		Label:    promptText,
		Validate: validatorFunc,
	}
	input, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("failed to receive user input: %w", err)
	}
	c.handleExit(input)

	return input, nil
}

func (c *CLI) PromptOptions(promptText string, options []string) (string, error) {
	prompt := promptui.Select{
		Label: promptText,
		Items: options,
	}
	_, choice, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("failed to receive user choice: %w", err)
	}

	return choice, err
}

func (c *CLI) handleExit(input string) {
	if strings.ToLower(input) == "exit" {
		os.Exit(0)
	}
}
