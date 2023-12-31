package service

import (
	"encoding/json"
	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/wcodesoft/mosha-quote-service/data"
	mhttp "github.com/wcodesoft/mosha-service-common/http"
	moshat "github.com/wcodesoft/mosha-service-common/test"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-quote-service/repository"
)

func createHandler() http.Handler {
	memoryDatabase := repository.NewInMemoryDatabase()
	cr := repository.NewFakeClientRepository()
	repo := repository.New(memoryDatabase, cr)
	service := New(repo)
	hs := QuoteService{
		Service: service,
		Port:    "8080",
		Name:    "QuoteService",
	}
	handler := hs.MakeHandler()
	return handler
}

func TestHttp(t *testing.T) {

	authorId := "123"

	Convey("When adding a valid quote", t, func() {
		quote := data.NewQuoteBuilder().WithId(faker.UUID()).WithAuthorId(authorId).WithText(faker.Sentence(10)).Build()
		handler := createHandler()
		req := httptest.NewRequest("POST", "/api/v1/quote", moshat.JsonReaderFactory(quote))
		rr := moshat.ExecuteRequest(req, handler)

		Convey("The response should be 200", func() {
			So(rr.Code, ShouldEqual, http.StatusOK)
		})

		Convey("The response should contain the quote ID", func() {
			var quoteResponse mhttp.IdResponse
			_ = json.NewDecoder(rr.Body).Decode(&quoteResponse)
			So(quoteResponse.ID, ShouldEqual, quote.ID)
		})
	})

	Convey("When adding a quote with invalid quote", t, func() {
		handler := createHandler()
		quote := data.NewQuoteBuilder().WithId(faker.UUID()).WithAuthorId(authorId).WithText(faker.Sentence(10)).Build()

		Convey("When quote already exists the response should be 500", func() {
			req1 := httptest.NewRequest("POST", "/api/v1/quote", moshat.JsonReaderFactory(quote))
			moshat.ExecuteRequest(req1, handler)
			req2 := httptest.NewRequest("POST", "/api/v1/quote", moshat.JsonReaderFactory(quote))
			rr := moshat.ExecuteRequest(req2, handler)
			So(rr.Code, ShouldEqual, http.StatusInternalServerError)
		})

		Convey("When author does not exist the response should be 500", func() {
			quote.AuthorID = "invalid"
			req := httptest.NewRequest("POST", "/api/v1/quote", moshat.JsonReaderFactory(quote))
			rr := moshat.ExecuteRequest(req, handler)
			So(rr.Code, ShouldEqual, http.StatusInternalServerError)
		})

		Convey("When quote is invalid the response should be 500", func() {
			req := httptest.NewRequest("POST", "/api/v1/quote", moshat.JsonReaderFactory("invalid"))
			rr := moshat.ExecuteRequest(req, handler)
			So(rr.Code, ShouldEqual, http.StatusInternalServerError)
		})

	})

	Convey("With a quote in the database", t, func() {
		quote := data.NewQuoteBuilder().WithId(faker.UUID()).WithAuthorId(authorId).WithText(faker.Sentence(10)).Build()
		handler := createHandler()
		req := httptest.NewRequest("POST", "/api/v1/quote", moshat.JsonReaderFactory(quote))
		rr := moshat.ExecuteRequest(req, handler)

		Convey("The response should be 200", func() {
			So(rr.Code, ShouldEqual, http.StatusOK)
		})

		Convey("When getting the quote", func() {
			getReq := httptest.NewRequest("GET", "/api/v1/quote/"+quote.ID, nil)
			getRr := moshat.ExecuteRequest(getReq, handler)

			Convey("The response should be 200", func() {
				So(getRr.Code, ShouldEqual, http.StatusOK)
			})

			Convey("The response should contain the correct ID", func() {
				var quoteResponse data.Quote
				_ = json.NewDecoder(getRr.Body).Decode(&quoteResponse)
				So(quoteResponse.ID, ShouldEqual, quote.ID)
			})

			Convey("When getting a quote that does not exist", func() {
				getReq := httptest.NewRequest("GET", "/api/v1/quote/456", nil)
				getRr := moshat.ExecuteRequest(getReq, handler)

				Convey("The response should be 500", func() {
					So(getRr.Code, ShouldEqual, http.StatusInternalServerError)
				})
			})
		})

		Convey("When updating the quote", func() {
			quote.Text = faker.Sentence(10)
			updateReq := httptest.NewRequest("POST", "/api/v1/quote/update", moshat.JsonReaderFactory(quote))
			updateRr := moshat.ExecuteRequest(updateReq, handler)

			Convey("The response should be 200", func() {
				So(updateRr.Code, ShouldEqual, http.StatusOK)
			})

			Convey("The response should contain the correct ID", func() {
				var quoteResponse data.Quote
				_ = json.NewDecoder(updateRr.Body).Decode(&quoteResponse)
				So(quoteResponse.ID, ShouldEqual, quote.ID)
			})

			Convey("When updating a quote that does not exist", func() {
				quote.ID = "456"
				updateReq := httptest.NewRequest("POST", "/api/v1/quote/update", moshat.JsonReaderFactory(quote))
				updateRr := moshat.ExecuteRequest(updateReq, handler)

				Convey("The response should be 500", func() {
					So(updateRr.Code, ShouldEqual, http.StatusInternalServerError)
				})
			})
		})

		Convey("When deleting the quote", func() {
			deleteReq := httptest.NewRequest("POST", "/api/v1/quote/delete/"+quote.ID, nil)
			deleteRr := moshat.ExecuteRequest(deleteReq, handler)

			Convey("The response should be 200", func() {
				So(deleteRr.Code, ShouldEqual, http.StatusOK)
			})

			Convey("The response should contain the correct ID", func() {
				var quoteResponse mhttp.IdResponse
				_ = json.NewDecoder(deleteRr.Body).Decode(&quoteResponse)
				So(quoteResponse.ID, ShouldEqual, quote.ID)
			})

			Convey("When deleting a quote that does not exist", func() {
				deleteReq := httptest.NewRequest("POST", "/api/v1/quote/delete/456", nil)
				deleteRr := moshat.ExecuteRequest(deleteReq, handler)

				Convey("The response should be 500", func() {
					So(deleteRr.Code, ShouldEqual, http.StatusInternalServerError)
				})
			})
		})

		Convey("When listing all quotes", func() {
			listReq := httptest.NewRequest("GET", "/api/v1/quote/all", nil)
			listRr := moshat.ExecuteRequest(listReq, handler)

			Convey("The response should be 200", func() {
				So(listRr.Code, ShouldEqual, http.StatusOK)
			})

			Convey("The response should contain the correct ID", func() {
				var quoteResponse []data.Quote
				_ = json.NewDecoder(listRr.Body).Decode(&quoteResponse)
				So(len(quoteResponse), ShouldEqual, 1)
			})
		})

		Convey("When listing quotes by author", func() {
			listReq := httptest.NewRequest("GET", "/api/v1/quote/author/"+authorId, nil)
			listRr := moshat.ExecuteRequest(listReq, handler)

			Convey("The response should be 200", func() {
				So(listRr.Code, ShouldEqual, http.StatusOK)
			})

			Convey("The response should contain the correct ID", func() {
				var quoteResponse []data.Quote
				_ = json.NewDecoder(listRr.Body).Decode(&quoteResponse)
				So(len(quoteResponse), ShouldEqual, 1)
			})
		})

		Convey("When getting a random quote", func() {
			getReq := httptest.NewRequest("GET", "/api/v1/quote/random", nil)
			getRr := moshat.ExecuteRequest(getReq, handler)

			Convey("The response should be 200", func() {
				So(getRr.Code, ShouldEqual, http.StatusOK)
			})

			Convey("The response should contain the correct ID", func() {
				var quoteResponse data.Quote
				_ = json.NewDecoder(getRr.Body).Decode(&quoteResponse)
				So(quoteResponse.ID, ShouldEqual, quote.ID)
			})
		})
	})

	Convey("When getting a random quote without any quote", t, func() {
		handler := createHandler()
		getReq := httptest.NewRequest("GET", "/api/v1/quote/random", nil)
		getRr := moshat.ExecuteRequest(getReq, handler)

		Convey("The response should be 500", func() {
			So(getRr.Code, ShouldEqual, http.StatusInternalServerError)
		})
	})
}
