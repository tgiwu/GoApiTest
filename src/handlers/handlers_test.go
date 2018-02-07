package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"models"

	"github.com/julienschmidt/httprouter"
)

func TestBookIndex(t *testing.T) {
	testBook := &models.Book{
		ISDN:   "111",
		Title:  "test title",
		Author: "test author",
		Pages:  42,
	}

	bookstore["111"] = testBook

	req1, err := http.NewRequest("GET", "/books", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr1 := newRequestRecorder(req1, "GET", "/books", BookIndex)
	if rr1.Code != 200 {
		t.Error("Expected response code to be 200")
	}

	er1 := "{\"meta\":null,\"data\":[{\"isdn\":\"111\",\"title\":\"test title\",\"author\":\"test author\",\"pages\":42}]}\n"

	if rr1.Body.String() != er1 {
		t.Error("Response body does not match")
	}
}

func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
