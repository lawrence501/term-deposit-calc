# term-deposit-calc

Term Deposit Calculator CLI

## Installation

Running this calculator requires the following software to be installed:
- [Go v1.22](https://go.dev/dl/)
- [Golangci-lint](https://golangci-lint.run/usage/install/#local-installation) (only required if you want to run the linter)

## Usage

The following Make commands exist for this codebase:
- `make run`: Runs the calculator. Instructions for using the calculator are provided within the calculator.
- `make test`: Runs the unit tests for the codebase.
- `make lint`: Runs a standard Go linter against the codebase.

## Note when running on Windows

This application has been tested on MacOS (m1 chip) and Windows (using git bash), but performs better on unix systems; on Windows, the CLI prompting package used is known to have some visual quirks (which can rarely cause it to skip printing the final balance - if this occurs please just enter the same inputs again), but the calculator's functionality should be fine. 
