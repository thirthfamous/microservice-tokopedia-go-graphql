package middleware

import (
	"fmt"
	"mime"
	"net/http"
)

// type authString string

func EnforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("yomasukyo")
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			if mt != "application/graphql" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// func AuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		auth := r.Header.Get("Authorization")

// 		if auth == "" {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		bearer := "Bearer "
// 		auth = auth[len(bearer):]

// 		validate, err := helper.ValidateToken(auth)
// 		if err != nil || !validate.Valid {
// 			http.Error(w, "Invalid token", http.StatusForbidden)
// 			return
// 		}

// 		customClaim, _ := validate.Claims.(*service.JwtCustomClaim)

// 		ctx := context.WithValue(r.Context(), authString("auth"), customClaim)

// 		r = r.WithContext(ctx)
// 		next.ServeHTTP(w, r)
// 	})
// }

// func CtxValue(ctx context.Context) *service.JwtCustomClaim {
// 	raw, _ := ctx.Value(authString("auth")).(*service.JwtCustomClaim)
// 	return raw
// }
