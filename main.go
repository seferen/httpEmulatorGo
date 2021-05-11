package main

import (
	"log"
	"net/http"
	"time"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		targetMux.ServeHTTP(w, r)
		requesterIP := r.RemoteAddr

		log.Printf("%s %s %s %v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start))
	})
}

func Health(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Ok"))

}

func Test(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	time.Sleep(500 * time.Millisecond)
	w.Write([]byte("Ок"))

}

func main() {
	log.Println("Application was started")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/test", Test)
	http.ListenAndServe(":8080", RequestLogger(mux))
}
