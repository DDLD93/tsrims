package middlewares

import (
	"log"
	"net/http"
	"time"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("[ Jwt Middleware Called ]")
		time.Sleep(6 * time.Second)
		next.ServeHTTP(w, r)
	})
}
