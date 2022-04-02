package handlers

import (
	"fmt"
	"net/http"
)

func TestHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("You have logged in")
}
