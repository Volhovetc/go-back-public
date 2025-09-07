package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type HalloHandler struct{}

func NewHalloHandler(router *http.ServeMux) {
	handler := &HalloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (handler *HalloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		randomNumber := rand.Intn(100)
		strNumber := strconv.Itoa(randomNumber)
		fmt.Println(randomNumber)
		w.Write([]byte(strNumber))
	}
}
