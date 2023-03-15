package middleware

import (
	"fmt"
	"net/http"
)

type LetMiddleware struct {
}

func NewLetMiddleware() *LetMiddleware {
	return &LetMiddleware{}
}

func (m *LetMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		fmt.Println("=======局部中间件之前=======")
		next(w, r)
		fmt.Println("=======局部中间件之后=======")
	}
}
