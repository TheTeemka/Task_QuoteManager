package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/TheTeemka/Task_QuoteManager/internal/server"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	logLevel := flag.String("logLevel", "info", "Log level: debug, info, warn, error")
	port := flag.String("port", ":8080", "Server port")
	filepath := flag.String("file", "temirlan_bayangazy_data.json", "file to save data")
	bePersistent := flag.Bool("persistency", false, "enable persistent storage")
	flag.Parse()

	setSlog(*logLevel)

	slog.Debug("Flags",
		"logLevel", *logLevel,
		"port", *port,
		"filepath", *filepath,
		"bePersistent", *bePersistent,
	)

	srv := server.NewServer(*port, *filepath, *bePersistent)
	srv.Serve()
}

func setSlog(logLevel string) {
	var level slog.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)
}
