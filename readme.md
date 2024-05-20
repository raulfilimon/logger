# Logger Package

This is a simple logging package written in Go. It provides functionality to log informational and error messages to a file.

## Usage

### Initialization

The logger package automatically initializes itself when imported. It creates a new log file in the `storage/logs/` directory with the current date as the filename.

### Logging Information

To log informational messages, use the [`Info()`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fhome%2Frfilimon%2Fgo_projects%2Flogger%2Flogger.go%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A41%2C%22character%22%3A5%7D%5D "logger.go") function:

```go
logger.Info("This is an informational message.")
```

This will append a line to the log file in the following format:

```
INFO This is an informational message.
```

### Logging Errors

To log error messages, use the `Err()` function:

```go
err := errors.New("This is an error message.")
logger.Err(err)
```

This will append a line to the log file in the following format:

```
ERROR This is an error message.
```

## Notes

- The logger package uses the built-in [`log`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fhome%2Frfilimon%2Fgo_projects%2Flogger%2Flogger.go%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A3%2C%22character%22%3A1%7D%5D "logger.go") package to write to the log file.
- The log file is opened in append mode, so new log entries are added to the end of the file.
- If the `storage/logs/` directory does not exist, the logger package will create it.
- If the log file cannot be opened, or if there is an error creating the `storage/logs/` directory, the logger package will log a fatal error and the program will terminate.