package main

import (
	"handlers"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routers := Routes{
		Route{"Index", "GET", "/", handlers.Index},
		Route{"BookIndex", "GET", "/books", handlers.BookIndex},
		Route{"Bookshow", "GET", "/books/:isdn", handlers.BookShow},
		Route{"Bookshow", "POST", "/books", handlers.BookCreate},
	}
	return routers
}
