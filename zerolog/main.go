package main

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

func main() {

	// ** Example: Simple Log Usage
	// request := "{secret}"
	// log.Info().Msg(request)

	// ** Example: Set the output
	// logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	// logger.Info().Str("foo", "bar").Msg("hello world")

	// ** Example: Uses zerolog ConsoleWriter
	// output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	// output.FormatLevel = func(i interface{}) string {
	// 	return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	// }
	// output.FormatMessage = func(i interface{}) string {
	// 	return fmt.Sprintf("***%s****", i)
	// }
	// output.FormatFieldName = func(i interface{}) string {
	// 	return fmt.Sprintf("%s:", i)
	// }
	// output.FormatFieldValue = func(i interface{}) string {
	// 	return strings.ToUpper(fmt.Sprintf("%s", i))
	// }
	// log := zerolog.New(output).With().Timestamp().Logger()
	// log.Info().Str("foo", "bar").Msg("Hello World")

	// ** Example: Uses zerolog ConsoleWriter
	// consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

	// ** Example: Uses zerolog ConsoleWriter
	// tempFile, err := ioutil.TempFile(os.TempDir(), "deleteme")
	// if err != nil {
	// 	// Can we log an error before we have our logger? :)
	// 	log.Error().Err(err).Msg("there was an error creating a temporary file four our log")
	// }
	// fmt.Printf("The log file is allocated at %s\n", tempFile.Name())
	// fileLogger := zerolog.New(tempFile).With().Logger()

	// ** Example: Uses lumber jack for rolling file
	var lumberJackLogger = &lumberjack.Logger{
		Filename:   "./zerolog.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	// ** Example: Uses zerolog multiwriter
	// multi := zerolog.MultiLevelWriter(fileLogger, os.Stdout)
	multi := zerolog.MultiLevelWriter(lumberJackLogger, os.Stdout)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	logger.Info().Msg("Hello World!")
}

// func fileOutput() {
// 	// create a temp file
// 	tempFile, err := ioutil.TempFile(os.TempDir(), "deleteme")
// 	if err != nil {
// 		// Can we log an error before we have our logger? :)
// 		log.Error().Err(err).Msg("there was an error creating a temporary file four our log")
// 	}
// 	fileLogger := zerolog.New(tempFile).With().Logger()
// 	fileLogger.Info().Msg("This is an entry from my log")

// 	fmt.Printf("The log file is allocated at %s\n", tempFile.Name())
// }
