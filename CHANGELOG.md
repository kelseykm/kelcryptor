# Changelog

## [2.0.0] - 2022-09-23

### Changed

- Remove trailing newline from scanned password
- Use `ReadPassword` function from `term` package to read password
- Change password type from string to []byte
- Bring `colour` package more inline with ANSI escape codes
- Use byte escape sequence in `colour`

## [1.2.0] - 2022-09-18

### Added

- Add `errors` package to contain all errors
- Add "Error" suffix to names of all errors

## [1.1.0] - 2022-09-13

### Added

- Add version flag

## [1.0.0] - 2022-09-10

- First stable release
