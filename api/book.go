package api

import (
	"lab01/api/classes"
	"lab01/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

var books = []classes.Book{
	classes.Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802", Cost: 15.50},
	classes.Book{Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000", Cost: 25.50},
	classes.Book{Title: "Antipatrones en IOS", Author: "Juan Arcos ", ISBN: "1234567890", Cost: 99.99},
	classes.Book{Title: "SAGARPA, punto de quiebre?", Author: "Adrian Gomez", ISBN: "0987654321", Cost: 102.80},
}

func GetBooks(c *gin.Context) {

	var request_books classes.RequestBook

	// llamar metodos de capa modelo

	r := classes.RequestCode{Code: config.OK_PROCESS, Message: "OK process"}
	request_books.Response = r
	request_books.Books = books
	c.JSON(http.StatusOK, request_books)

}

func GetBook(c *gin.Context, isbn string) {
	var request_books classes.RequestBook

	// llamar metodos de capa modelo

	for _, v := range books {
		if v.ISBN == isbn {
			r := classes.RequestCode{Code: config.OK_PROCESS, Message: "OK process"}
			request_books.Response = r
			request_books.Books = []classes.Book{v}
			c.JSON(http.StatusOK, request_books)
			return
		}
	}
	r := classes.RequestCode{Code: config.NOT_FOUND, Message: "Not found book: " + isbn}
	request_books.Response = r
	request_books.Books = []classes.Book{}
	c.JSON(http.StatusNotFound, gin.H{"message": request_books})
	//c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})

	//c.JSON(http.StatusNotFound, "")

}
