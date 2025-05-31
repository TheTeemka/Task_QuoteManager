package server

import (
	"net/http"
	"strconv"

	"github.com/TheTeemka/Task_QuoteManager/internal/service"
	"github.com/TheTeemka/Task_QuoteManager/pkg/utils"
)

type QuoteHandler struct {
	quoteService *service.QuoteService
}

func NewQuoteHandler(quoteService *service.QuoteService) *QuoteHandler {
	return &QuoteHandler{
		quoteService: quoteService,
	}
}

func (h *QuoteHandler) Create(w http.ResponseWriter, req *http.Request) {
	quoteReq, err := utils.DecodeJson[service.QuoteReq](req.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if quoteReq.Author == "" {
		http.Error(w, "Author is required", http.StatusBadRequest)
		return
	}

	if quoteReq.Quote == "" {
		http.Error(w, "Quote is required", http.StatusBadRequest)
		return
	}

	quote := h.quoteService.CreateQuote(quoteReq)

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, quote)
}

func (h *QuoteHandler) Get(w http.ResponseWriter, req *http.Request) {
	author := req.URL.Query().Get("author")
	var quotes []*service.Quote
	if author == "" {
		quotes = h.quoteService.GetAll()
	} else {
		quotes = h.quoteService.GetByAuthor(author)
	}

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, map[string][]*service.Quote{
		"quotes": quotes,
	})
}

func (h *QuoteHandler) GetRandom(w http.ResponseWriter, req *http.Request) {
	quote := h.quoteService.GetRandomQuote()

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, quote)
}

func (h *QuoteHandler) DeleteByID(w http.ResponseWriter, req *http.Request) {
	idStr := req.PathValue("id")
	if idStr == "" {
		http.Error(w, "ID query parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	deleted := h.quoteService.DeleteByID(id)
	if !deleted {
		http.Error(w, "Quote not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
