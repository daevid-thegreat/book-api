package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Subject  string `json:"subject"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Golang", Author: "Google", Subject: "Programming", Quantity: 10},
	{ID: "2", Title: "Python", Author: "Google", Subject: "Programming", Quantity: 10},
	{ID: "3", Title: "Java", Author: "Google", Subject: "Programming", Quantity: 10},
	{ID: "4", Title: "C++", Author: "Google", Subject: "Programming", Quantity: 10},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(id string) (*book, error) {
	for _, b := range books {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, errors.New("book not found")
}

func bookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func checkOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing id"})
		return
	}
	book, err := getBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	if book.Quantity == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not available"})
		return
	}
	book.Quantity--
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing id"})
		return
	}
	book, err := getBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	book.Quantity++
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookByID)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
