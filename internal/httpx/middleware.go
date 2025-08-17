package httpx

import (
	"context"
	"net/http"

	"github.com/casbin/casbin/v2"
)

type ctxKey string

const userKey ctxKey = "user"

func UserFromContext(ctx context.Context) (string, bool) {
	u, ok := ctx.Value(userKey).(string)
	return u, ok
}

func Authorize(e *casbin.Enforcer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := r.Header.Get("X-User")
			if user == "" {
				http.Error(w, "missing X-User header", http.StatusUnauthorized)
				return
			}

			allowed, err := e.Enforce(user, r.URL.Path, r.Method)
			if err != nil {
				http.Error(w, "authorization error", http.StatusInternalServerError)
				return
			}
			if !allowed {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), userKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
