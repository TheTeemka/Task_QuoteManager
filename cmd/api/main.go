package main

import (
	"flag"

	"github.com/TheTeemka/Test_QuoteManager/internal/server"
)

func main() {
	port := flag.String("port", ":8000", "Server port")
	flag.Parse()

	srv := server.NewServer(*port)
	srv.Serve()
}
