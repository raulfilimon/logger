# Logger Package

This is a simple logging package for Go applications. It provides basic logging functionality with log rotation based on file size and date.

## Features

- Logs messages with [`Info`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fhome%2Frfilimon%2Fgo_projects%2Flogger%2Flogger.go%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A58%2C%22character%22%3A5%7D%5D "logger.go") and `Error` levels.
- Rotates log files based on size (50MB by default) and date (daily).
- Writes logs to a file in the `storage/logs` directory, with the filename format `YYYY-MM-DD.log`.

## Usage

```go
import "github.com/raulfilimon/logger"

func main() {
    logger.Info("This is an informational message")
    logger.Err(errors.New("This is an error message"))
}
```

## Functions

- `Info(info string)`: Logs an informational message. Automatically checks if log rotation is needed before logging the message.
- `Err(_err error)`: Logs an error message. Automatically checks if log rotation is needed before logging the message.

## Log Rotation

The log rotation is handled automatically. When a log message is written, the logger checks the size of the current log file and the current date. If the log file size is greater than 50MB or the date has changed, the logger will close the current log file and open a new one.

## Customization

You can customize the logger by modifying the constants and variables in the `logger.go` file:

- `dateLayout`: The format of the date in the log file name. Default is "2006-01-02".
- `maxSize`: The maximum size of a log file in bytes before it gets rotated. Default is 50,000,000 bytes (50MB).

## Limitations

This logger does not support multiple log levels or log file compression. It also does not delete old log files. If you need these features, consider using a more advanced logging library.