package middleware

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"
	"thirthfamous/tokopedia-clone-go-graphql/helper"
	"thirthfamous/tokopedia-clone-go-graphql/model/web"
)

// type authString string

func EnforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			if mt != "application/graphql" {
				webResponse := web.WebResponse{
					Code:   http.StatusUnsupportedMediaType,
					Status: "Content-Type header must be application/json",
				}
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnsupportedMediaType)
				encoder := json.NewEncoder(w)
				encoder.Encode(webResponse)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if apiKey := r.Header.Get("X-API-Key"); apiKey != "" && apiKey == "RAHASIA" {
			r = r.WithContext(context.WithValue(r.Context(), "admin", "ok"))
			next.ServeHTTP(w, r)
			return
		}

		auth := r.Header.Get("Authorization")

		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, profile_id, err := helper.ParseToken(auth)
		if err != nil || !validate {
			webResponse := web.WebResponse{
				Code:   http.StatusForbidden,
				Status: "Invalid token",
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			encoder := json.NewEncoder(w)
			encoder.Encode(webResponse)
			return
		}

		ctx := context.WithValue(r.Context(), "profile_id", profile_id)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
