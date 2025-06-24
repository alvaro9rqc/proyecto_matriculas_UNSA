package utils

import (
	"context"
	"net/http"
)

type contextKey string

const SESSION_TOKEN_ID contextKey = "session_token"

func SessionTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie == nil {
			ctx := r.Context()
			ctx = context.WithValue(ctx, SESSION_TOKEN_ID, "")
			r = r.WithContext(ctx)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, SESSION_TOKEN_ID, cookie.Value)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
