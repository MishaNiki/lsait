package servers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/MishaNiki/lsait/backend/internal/app/security"
	"github.com/gorilla/mux"
)

// ResponseError ...
func ResponseError(w http.ResponseWriter, r *http.Request, code int, err error) {
	Response(w, r, code, map[string]string{"error :": err.Error()})
}

// Response ...
func Response(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

type ContexKey string // ContexKey
// SecurityMiddleware - cheking accessToken
func SecurityMiddleware(secretKey []byte) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("SecurityMiddleware")
			accessToken := r.Header.Get("Token")
			id, code, err := security.VerificationAccessToken(secretKey, accessToken)
			if err != nil {
				ResponseError(w, r, code, err)
				return
			}
			ctx := context.WithValue(r.Context(), ContexKey("id"), id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
