package main

import (
	"errors"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default();
	router.GET("/ping", "Hola, Go!")
	router.Run("localhost:8080")
}
