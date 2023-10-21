package auth

import (
	"bingo/internal/repositories"
	"bingo/pkg/jwt"
	"log"

	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCtxKey = &contextKey{"id"}

type contextKey struct {
	id string
}

type User struct {
	ID      string `json:"id" bson:"_id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Numbers []int  `json:"numbers,omitempty"`
}

func AuthenticationMiddleware(userRepository *repositories.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header
			// Maybe change to cookie later
			// c, err := r.Cookie("auth-cookie")

			// Allow unauthenticated users in
			if header == nil {
				next.ServeHTTP(w, r)
				return
			}
			tokenStr := header.Get("Authorization")
			if tokenStr == "" {
				log.Printf("Authorization is empty")
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			userClaim, err := jwt.ValidateToken(tokenStr)
			if err != nil {
				log.Printf("Token Validate Error: %s", err)
				next.ServeHTTP(w, r)
				return
			}

			// get user from database
			user, err := userRepository.FindUserByID(userClaim.Id)
			if err == mongo.ErrNoDocuments {
				log.Printf("No document was found with id %s", userClaim.Id)
				next.ServeHTTP(w, r)
				return
			}
			if err != nil {
				log.Printf("MongoDB error: %s", err)
				http.Error(w, "System Error occur.", http.StatusInternalServerError)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &User{user.ID.Hex(), user.Name, user.Role, user.Numbers})

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}
