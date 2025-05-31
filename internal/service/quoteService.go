package service

import (
	"encoding/json"
	"io"
	"log/slog"
	"math/rand"
	"os"
)

type QuoteReq struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuoteService struct {
	MaxID   int
	Quotes  map[int]*Quote
	Authors map[string]map[int]bool
}

func NewQuoteService() *QuoteService {
	return &QuoteService{
		Quotes:  map[int]*Quote{},
		Authors: map[string]map[int]bool{},
	}
}

func (s *QuoteService) CreateQuote(req *QuoteReq) *Quote {
	q := &Quote{
		ID:     s.MaxID,
		Author: req.Author,
		Quote:  req.Quote,
	}

	s.MaxID++
	s.Quotes[q.ID] = q
	if _, ok := s.Authors[q.Author]; !ok {
		s.Authors[q.Author] = map[int]bool{}
	}
	s.Authors[q.Author][q.ID] = true

	return q
}

func (s *QuoteService) GetAll() []*Quote {
	quotes := make([]*Quote, 0, len(s.Quotes))
	for _, q := range s.Quotes {
		quotes = append(quotes, q)
	}
	return quotes
}

func (s *QuoteService) GetByAuthor(author string) []*Quote {
	if _, ok := s.Authors[author]; !ok {
		return []*Quote{}
	}
	quotes := make([]*Quote, 0, len(s.Authors[author]))
	for id := range s.Authors[author] {
		quotes = append(quotes, s.Quotes[id])
	}
	return quotes
}

func (s *QuoteService) DeleteByID(id int) bool {
	if quote, ok := s.Quotes[id]; ok {
		delete(s.Quotes, id)
		delete(s.Authors[quote.Author], id)
		if len(s.Authors[quote.Author]) == 0 {
			delete(s.Authors, quote.Author)
		}
		return true
	}
	return false
}

func (s *QuoteService) GetRandomQuote() *Quote {
	if len(s.Quotes) == 0 {
		return nil
	}

	r := rand.Intn(s.MaxID)
	for {
		if quote, ok := s.Quotes[r]; ok {
			return quote
		}
		r = rand.Intn(s.MaxID)
	}
}

func (s *QuoteService) Parse(file *os.File) {
	err := json.NewDecoder(file).Decode(&s.Quotes)
	if err != nil && err != io.EOF {
		slog.Error("decode quotes to file", "err", err)
	}

	for _, quote := range s.Quotes {
		if _, ok := s.Authors[quote.Author]; !ok {
			s.Authors[quote.Author] = map[int]bool{}
		}
		slog.Debug("quote from start", "quote", *quote)
		s.Authors[quote.Author][quote.ID] = true
	}
}

func (s *QuoteService) SaveTo(file *os.File) {
	file.Truncate(0)
	file.Seek(0, 0)
	err := json.NewEncoder(file).Encode(s.Quotes)
	if err != nil {
		slog.Error("encode quotes to file", "err", err)
	}
}
