package main

import (
	"YongQing-API/router"
	"net/http"
)

func main() {

	s := http.Server{Addr: "0.0.0.0:5004", Handler: router.GinEngine}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
