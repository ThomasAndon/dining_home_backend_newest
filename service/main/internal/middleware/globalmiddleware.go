package middleware

import "net/http"

type GlobalMiddleware struct {
}

func NewGlobalMiddleware() *GlobalMiddleware {
	return &GlobalMiddleware{}
}

func (m *GlobalMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		r.Header.Add("mwKey", "mwValue")
		// Passthrough to next handler if need
		next(w, r)
	}
}
