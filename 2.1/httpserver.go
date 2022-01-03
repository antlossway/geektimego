// client test: curl -i -H "First-Name: Toto" http://localhost:8080/headers

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type MyResponseWriter struct {
	http.ResponseWriter // original http.ResponseWriter
	status              int
	size                int
}

// MyResponseWriter need to implement interface http.ResponseWriter, so need to implement func Write and WriteHeader
func (w *MyResponseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b) // write response using original http.ResponseWriter
	w.size += size
	return size, err
}

func (w *MyResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode) // write status code using original http.ResponseWriter
	w.status = statusCode
}

// decorate handler with logging function
func WithLogging(h http.Handler) http.Handler {
	log_h := func(w http.ResponseWriter, r *http.Request) {

		w2 := MyResponseWriter{
			ResponseWriter: w, //original http.ResponseWriter
			status:         0,
			size:           0,
		}

		//serve the original request
		h.ServeHTTP(&w2, r)

		//3. log client IP, HTTP response code, output to stdout
		log.Printf("Handling request: url: %s, method: %s, remote IP: %s, status code: %d",
			r.RequestURI, r.Method, r.RemoteAddr, w2.status)

	}

	return http.HandlerFunc(log_h)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello")
}

// 4. when visit localhost:8080/healthz, return 200
func healthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "healthz is clicked, return 200\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	//2.  get a Environment variable and set in response header
	os_user := os.Getenv("USER")
	w.Header().Add("OS-USER", os_user)

	// 1. get headers of request, and set the same into response header
	for name, values := range req.Header {
		for _, v := range values {
			w.Header().Add(name, v)
		}
	}
	// set status code before first write
	w.WriteHeader(201)

	// write response body
	for name, values := range req.Header {
		for _, v := range values {
			fmt.Fprintf(w, "[request header] %v: %v\n", name, v)
		}
	}

}

func main() {

	http.Handle("/", WithLogging(http.HandlerFunc(home)))
	http.Handle("/healthz", WithLogging(http.HandlerFunc(healthz)))
	//http.HandleFunc("/headers", headers)
	http.Handle("/headers", WithLogging(http.HandlerFunc(headers)))

	http.ListenAndServe(":8080", nil)
}
