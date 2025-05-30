package service

import "math/rand"

type QuoteReq struct {
	Author string
	Text   string
}

type Quote struct {
	ID     int
	Author string
	Text   string
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
		Text:   req.Text,
	}
	s.MaxID++
	s.Quotes[q.ID] = q
	s.Authors[q.Author][q.ID] = true
	return q
}

func (s *QuoteService) GetAll() []*Quote {
	quotes := make([]*Quote, len(s.Quotes))
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

func (s *QuoteService) DeleteByID(id int) {
	if quote, ok := s.Quotes[id]; ok {
		delete(s.Authors[quote.Author], id)
		delete(s.Quotes, id)
	}
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
