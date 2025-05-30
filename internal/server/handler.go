package server

import (
	"github.com/TheTeemka/Test_QuoteManager/internal/service"
	"github.com/TheTeemka/Test_QuoteManager/pkg/utils"
	"net/http"
	"strconv"
)

type QuoteHandler struct {
	quoteService *service.QuoteService
}

func NewQuoteHandler() *QuoteHandler {
	return &QuoteHandler{
		quoteService: service.NewQuoteService(),
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

	if quoteReq.Text == "" {
		http.Error(w, "Author is required", http.StatusBadRequest)
		return
	}

	quote := h.quoteService.CreateQuote(quoteReq)

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, quote)
}

func (h *QuoteHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	quotes := h.quoteService.GetAll()

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, quotes)
}

func (h *QuoteHandler) GeyAuthor(w http.ResponseWriter, req *http.Request) {
	author := req.URL.Query().Get("author")
	if author == "" {
		http.Error(w, "Author query parameter is required", http.StatusBadRequest)
		return
	}

	quotes := h.quoteService.GetByAuthor(author)

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, quotes)
}

func (h *QuoteHandler) DeleteByID(w http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID query parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	h.quoteService.DeleteByID(id)

	w.WriteHeader(http.StatusNoContent)
}
