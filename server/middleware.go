package server

import "net/http"

func filterMethod(methods []string, f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		for _, method := range methods {
			if method == r.Method {
				f(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
