package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func RecFib(n int) int {
	// implement recursive fibonacci
	if n <= 1 {
		return n
	}
	return RecFib(n-1) + RecFib(n-2)
}

func Fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain")
	var result string
	parsedPath := strings.Split(r.URL.Path, "/")
	function, nstr := parsedPath[1], parsedPath[2]
	n, err := strconv.Atoi(nstr)

	if err != nil {
		w.Write([]byte(parsedPath[2]))
	}
	switch {
	case function == "fib":
		// tranform n in result
		result = strconv.Itoa(Fib(n))
	case function == "rec_fib":
		result = strconv.Itoa(RecFib(n))
	default:
		result = ""
	}
	w.Write([]byte(result))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &handler{})
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", mux)
}
