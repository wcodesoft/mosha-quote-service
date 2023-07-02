package service

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wcodesoft/mosha-quote-service/data"
	"net/http"
)

type idResponse struct {
	ID string `json:"id"`
}

type HttpRouter struct {
	service Service
}

func NewHttpRouter(s Service) HttpRouter {
	return HttpRouter{service: s}
}

func (h HttpRouter) MakeHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/api/v1/quote", h.addQuoteHandler)
	r.Post("/api/v1/quote/delete/{id}", h.deleteQuoteHandler)
	r.Post("/api/v1/quote/update", h.updateQuoteHandler)
	r.Get("/api/v1/quote/all", h.listAllHandler)
	r.Get("/api/v1/quote/author/{authorId}", h.listByAuthorHandler)
	r.Get("/api/v1/quote/{id}", h.getQuoteHandler)
	return r
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte("Error encoding response"))
	}
}

func (h HttpRouter) addQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Quote
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		encodeResponse(w, err)
		return
	}

	resp, err := h.service.CreateQuote(request)

	if err != nil {
		encodeResponse(w, err)
		return
	}

	encodeResponse(w, idResponse{ID: resp})
}

func (h HttpRouter) listAllHandler(w http.ResponseWriter, _ *http.Request) {
	resp := h.service.ListAll()

	encodeResponse(w, resp)
}

func (h HttpRouter) listByAuthorHandler(w http.ResponseWriter, r *http.Request) {
	authorId := chi.URLParam(r, "authorId")

	resp := h.service.GetAuthorQuotes(authorId)

	encodeResponse(w, resp)
}

func (h HttpRouter) getQuoteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := h.service.GetQuote(id)
	if err != nil {
		encodeResponse(w, err)
		return
	}
	encodeResponse(w, resp)
}

func (h HttpRouter) deleteQuoteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteQuote(id)
	if err != nil {
		encodeResponse(w, err)
		return
	}
	encodeResponse(w, idResponse{ID: id})
}

func (h HttpRouter) updateQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Quote
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		encodeResponse(w, err)
		return
	}

	result, err := h.service.UpdateQuote(request)
	if err != nil {
		encodeResponse(w, err)
		return
	}
	encodeResponse(w, result)
}
