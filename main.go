package main

import (
	"bayudsatriyo/belajar-golang-restfull-api/helper"
	"bayudsatriyo/belajar-golang-restfull-api/middleware"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	fmt.Println("Server is running on port 3000")
	fmt.Println("Running Golang REST API with MySQL Database 🚀")

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
