package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/TheTeemka/Task_QuoteManager/internal/service"
)

type Server struct {
	Port                string
	saveFile            *os.File
	QuoteHandlerHandler *QuoteHandler
}

func NewServer(Port, filePath string, bePersistent bool) *Server {
	var file *os.File
	if bePersistent {
		var err error
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			slog.Error("file open", "err", err)
		}
	}

	quoteService := service.NewQuoteService()
	if bePersistent {
		quoteService.Parse(file)
	}

	return &Server{
		Port:                Port,
		saveFile:            file,
		QuoteHandlerHandler: NewQuoteHandler(quoteService),
	}
}

func (s *Server) Serve() {
	srv := http.Server{
		Addr:    s.Port,
		Handler: s.Router(),
	}

	shutdown := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, os.Kill)

		slog.Info("Server is shutting down", "cause:", <-quit)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.QuoteHandlerHandler.quoteService.SaveTo(s.saveFile)
		}()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := srv.Shutdown(ctx)

		wg.Wait()
		shutdown <- err
	}()

	slog.Info("Server is starting", "port", s.Port)
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}

	err = <-shutdown
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Server is successfully closed")
}
