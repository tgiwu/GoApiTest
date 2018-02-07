package main

import (
	"log"
	"net/http"

	"db"
	"handlers"
	"models"
)


func main() {
	//router := httprouter.New()
	//router.GET("/", handlers.Index)
	//router.GET("/books", handlers.BookIndex)
	//router.GET("/books/:isdn", handlers.BookShow)
	handlers.FillBookStore()

	router := NewRouter(AllRoutes())

	db.GetAllStudent()

	student:= new(models.Student)
	student.Name = "new student"
	student.Sex = "m"
	student.Age = 14
	student.Tel = "11213451231"
	db.InsertStudent(student)

	db.QueryAllStudent()
	//db.QueryStudent()
	log.Fatal(http.ListenAndServe(":8090", router))
}
