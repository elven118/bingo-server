package auth

import (
	"bingo/pkg/jwt"
	"log"

	"context"
	"net/http"
)

var userCtxKey = &contextKey{"id"}

type contextKey struct {
	id string
}

type UserId struct {
	Id   string
	Name string
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		// Allow unauthenticated users in
		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		//validate jwt token
		tokenStr := header
		userClaim, err := jwt.ValidateToken(tokenStr)
		if err != nil {
			log.Printf("Token Validate Error: %s", err)
			http.Error(w, "Invalid Token", http.StatusForbidden)
			return
		}

		// put it in context
		ctx := context.WithValue(r.Context(), userCtxKey, &UserId{userClaim.Id, userClaim.Name})

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *UserId {
	raw, _ := ctx.Value(userCtxKey).(*UserId)
	return raw
}
