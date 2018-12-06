package main

import (
"flag"

"github.com/rs/zerolog"
"github.com/rs/zerolog/log"
)

func main() {
	//zerolog.TimeFieldFormat = ""
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999999-07:00"

	log.Debug().Msg("This message appears only when log level set to Debug")
	log.Info().Msg("This message appears when log level set to Debug or Info")

	if e := log.Debug(); e.Enabled() {
		// Compute log output only if enabled.
		value := "bar"
		e.Str("foo", value).Msg("some debug message")
	}

	subLogger := log.With().Str("component", "foo").Logger()
	subLogger.Info().Msg("hello world")
}
