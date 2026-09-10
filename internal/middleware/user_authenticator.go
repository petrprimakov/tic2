package middleware

import (
	"context"
	"net/http"

	"tic2/internal/application/service"

	"github.com/google/uuid"
)

type ctxKey string

const userIDKey ctxKey = "userID"

type UserAuthenticator struct {
	auth *service.AuthService
}

func NewUserAuthenticator(auth *service.AuthService) *UserAuthenticator {
	return &UserAuthenticator{auth: auth}
}

// Authenticate — middleware, требующая валидный Basic Auth.
func (a *UserAuthenticator) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		login, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		id, err := a.auth.SignIn(r.Context(), login, password)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// прокидываем UUID пользователя в context для хендлеров
		ctx := context.WithValue(r.Context(), userIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	id, ok := ctx.Value(userIDKey).(uuid.UUID)
	return id, ok
}
