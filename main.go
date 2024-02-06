package main

//ref: https://go.dev/doc/tutorial/web-service-gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// define a estrutura que representa livro
type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Writer string  `json:"writer"`
	Price  float64 `json:"price"`
}

// inicia a lista de livros
var books = []book{
	{ID: "1", Title: "To Kill a Mockingbird", Writer: "Harper Lee", Price: 56.99},
	{ID: "2", Title: "Peter Camenzid", Writer: "Herman Hesse", Price: 17.99},
	{ID: "3", Title: "Doramar ou Odisseia", Writer: "Itamar Vieira Junior", Price: 39.99},
}

// handler que obtem todos os livros
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// handler que adiciona um novo livro a lista
func postBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook) // adiciona o novo livro a lista
	c.IndentedJSON(http.StatusCreated, newBook)
}

// handler para obter um livro pelo seu ID
func getBooksId(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books { //iteração na lista de livros
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBooksId)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}
