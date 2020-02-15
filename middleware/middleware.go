package middleware

import (
	"context"
	"net/http"
)

type contextKey struct {
	name string
}

var requestCtxKey = &contextKey{"request"}

func AccessTokenForContext(ctx context.Context) *string {
	request := RequestForContext(ctx)
	auth := request.Header.Get("Authorization")
	if len(auth) <= 7 {
		return nil
	}
	authType := auth[:6]
	if authType != "Bearer" {
		return nil
	}
	accessToken := auth[7:len(auth)]
	return &accessToken
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), requestCtxKey, r)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func RequestForContext(ctx context.Context) *http.Request {
	raw, _ := ctx.Value(requestCtxKey).(*http.Request)
	return raw
}
