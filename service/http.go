package service

import (
	mhttp "github.com/wcodesoft/mosha-service-common/http"

	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wcodesoft/mosha-quote-service/data"
	"net/http"
	"time"
)

type HttpRouter struct {
	service     Service
	serviceName string
}

func NewHttpRouter(s Service, serviceName string) HttpRouter {
	return HttpRouter{
		service:     s,
		serviceName: serviceName,
	}
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

func (h HttpRouter) addQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Quote
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	resp, err := h.service.CreateQuote(request)

	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	mhttp.EncodeResponse(w, mhttp.IdResponse{ID: resp})
}

func (h HttpRouter) listAllHandler(w http.ResponseWriter, _ *http.Request) {
	resp := h.service.ListAll()

	mhttp.EncodeResponse(w, resp)
}

func (h HttpRouter) listByAuthorHandler(w http.ResponseWriter, r *http.Request) {
	authorId := chi.URLParam(r, "authorId")

	resp := h.service.GetAuthorQuotes(authorId)

	mhttp.EncodeResponse(w, resp)
}

func (h HttpRouter) getQuoteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := h.service.GetQuote(id)
	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}
	mhttp.EncodeResponse(w, resp)
}

func (h HttpRouter) deleteQuoteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteQuote(id)
	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}
	mhttp.EncodeResponse(w, mhttp.IdResponse{ID: id})
}

func (h HttpRouter) updateQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var request data.Quote
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		mhttp.EncodeError(w, err)
		return
	}

	result, err := h.service.UpdateQuote(request)
	if err != nil {
		mhttp.EncodeError(w, err)
		return
	}
	mhttp.EncodeResponse(w, result)
}

func (h HttpRouter) Start(port string) error {
	log.Infof("Starting %s http on %s", h.serviceName, port)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           h.MakeHandler(),
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("unable to start service %q: %s", h.serviceName, err)
	}
	return nil
}
