package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pigeon_carrier/internal/app/httpclient"
	"github.com/pigeon_carrier/internal/model/program"

	tea "charm.land/bubbletea/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	logger, err := newLogger("./log/pigeon_carrier.log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating logger: %v\n", err)
		os.Exit(1)
	}
	log.Logger = logger

	log.Info().Msg("Logger initialized")

	p := tea.NewProgram(program.NewProgram(logger, httpclient.NewHttpClient()))
	if _, err := p.Run(); err != nil {
		log.Error().Err(err).Msg("Error running program")
		os.Exit(1)
	}

	log.Info().Msg("Program exited successfully")
}

func newLogger(logPath string) (zerolog.Logger, error) {
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		return zerolog.Logger{}, fmt.Errorf("create log directory: %w", err)
	}

	fileWriter := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10,   // megabytes
		MaxBackups: 1,    // number of backups to keep
		MaxAge:     20,   // days
		Compress:   true, // compress old logs
	}

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(fileWriter).With().Timestamp().Logger()

	mainLogger := logger.With().Str("module", "main").Logger()

	return mainLogger, nil
}
