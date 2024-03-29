# Changelog

## [2.0.4] - 2024-02-25

- Fix extra spaces added to erase line when terminal width is less than line to
be written

## [2.0.3] - 2024-02-24

- Rewrite colour package to use functions from fatih/color for terminal
agnosticness

## [2.0.2] - 2023-08-01

- Allow use of flags at any point in arguments, not just the beginning
- Use keyed fields in `errors.GenericError` struct literals
- Update packages
- Use cloud image source for banner in README

## [2.0.1] - 2022-09-30

- Minor code refactors
- Add `install` target to makefile
- Remove third level headings from changelog

## [2.0.0] - 2022-09-23

- Remove trailing newline from scanned password
- Use `ReadPassword` function from `term` package to read password
- Change password type from string to []byte
- Bring `colour` package more inline with ANSI escape codes
- Use byte escape sequence in `colour`

## [1.2.0] - 2022-09-18

- Add `errors` package to contain all errors
- Add "Error" suffix to names of all errors

## [1.1.0] - 2022-09-13

- Add version flag

## [1.0.0] - 2022-09-10

- First stable release
