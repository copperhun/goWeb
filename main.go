package main

import (
	"copperhun/goWeb/myapp"
	"net/http"
)



func main() {
	

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}