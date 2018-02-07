package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"models"
	"response"

	"github.com/julienschmidt/httprouter"
)

var bookstore = make(map[string]*models.Book)

func FillBookStore()  {
	bookstore["123"] = &models.Book{
		ISDN:"123",
		Title:"Silence of the Lambs",
		Author: "Thomas Harris",
		Pages: 367,
	}

	bookstore["124"] = &models.Book{
		ISDN:"124",
		Title:"To kill a Mocking Bird",
		Author:"Harper Lee",
		Pages:320,
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome! \n")
}

func BookCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	book := &models.Book{}
	if err := populateModeFromHandler(w, r, params, book); err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	bookstore[book.ISDN] = book
	writeOkResponse(w, book)
}

func BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//books := []*Book{}
	//for _, book := range bookstore {
	//	books = append(books, book)
	//}
	//response := &JsonResponse{Data:&books}
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	//
	//if err := json.NewEncoder(w).Encode(response); err != nil {
	//	panic(err)
	//}
	books := []*models.Book{}
	for _, book := range bookstore {
		books = append(books, book)
	}
	writeOkResponse(w, books)
}

func BookShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//isdn := params.ByName("isdn")
	//book, ok := bookstore[isdn]
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//
	//if !ok {
	//	w.WriteHeader(http.StatusNotFound)
	//	response := JsonErrorResponse{Error: &ApiError{Status:404, Title:"Record Not Found"}}
	//	if err := json.NewEncoder(w).Encode(response); err != nil {
	//		panic(err)
	//	}
	//}
	//response := JsonResponse{Data: book}
	//if err:=json.NewEncoder(w).Encode(response); err != nil {
	//	panic(err)
	//}

	isdn := params.ByName("isdn")
	book, ok := bookstore[isdn]

	if !ok {
		writeErrorResponse(w, http.StatusNotFound, "Record not Found")
	}

	writeOkResponse(w, book)
}

func AllStudent(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	//stidentId := params.ByName("id")
	//student, ok :
}

func writeOkResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=Utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response.JsonResponse{Data: m}); err != nil {

	}
}

func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(&response.JsonErrorResponse{Error: &response.ApiError{Status: errorCode, Title: errorMsg}})
}

func populateModeFromHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}
