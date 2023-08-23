package service

import (
	mhttp "github.com/wcodesoft/mosha-service-common/http"

	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wcodesoft/mosha-quote-service/data"
	"net/http"
)

type QuoteService struct {
	Service Service
	Name    string
	Port    string
	mhttp.MoshaHttpService
}

func (qs *QuoteService) GetName() string {
	return qs.Name
}

func (qs *QuoteService) GetPort() string {
	return qs.Port
}

func (qs *QuoteService) MakeHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/api/v1/quote", qs.addQuoteHandler)
	r.Post("/api/v1/quote/delete/{id}", qs.deleteQuoteHandler)
	r.Post("/api/v1/quote/update", qs.updateQuoteHandler)
	r.Get("/api/v1/quote/all", qs.listAllHandler)
	r.Get("/api/v1/quote/author/{authorId}", qs.listByAuthorHandler)
	r.Get("/api/v1/quote/{id}", qs.getQuoteHandler)
	return r
}

func (qs *QuoteService) addQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Quote
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	resp, err := qs.Service.CreateQuote(request)

	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	mhttp.EncodeResponse(w, mhttp.IdResponse{ID: resp})
}

func (qs *QuoteService) listAllHandler(w http.ResponseWriter, _ *http.Request) {
	resp := qs.Service.ListAll()

	mhttp.EncodeResponse(w, resp)
}

func (qs *QuoteService) listByAuthorHandler(w http.ResponseWriter, r *http.Request) {
	authorId := chi.URLParam(r, "authorId")

	resp := qs.Service.GetAuthorQuotes(authorId)

	mhttp.EncodeResponse(w, resp)
}

func (qs *QuoteService) getQuoteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := qs.Service.GetQuote(id)
	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}
	mhttp.EncodeResponse(w, resp)
}

func (qs *QuoteService) deleteQuoteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := qs.Service.DeleteQuote(id)
	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}
	mhttp.EncodeResponse(w, mhttp.IdResponse{ID: id})
}

func (qs *QuoteService) updateQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Quote
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	result, err := qs.Service.UpdateQuote(request)
	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}
	mhttp.EncodeResponse(w, result)
}
