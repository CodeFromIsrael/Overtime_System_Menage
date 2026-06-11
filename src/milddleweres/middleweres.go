package milddleweres

import (
	"log"
	"net/http"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		nextFunc(w, r)
	}
}
