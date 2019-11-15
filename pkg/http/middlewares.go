package http

import (
	"context"
	"log"
	"net/http"
	"strings"

	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
	"github.com/rs/cors"
)

// This type and const prevents the
// 		"should not use basic type string as key in context.WithValue"
//																	go-lint
type key int

const (
	userID key = iota
)

// OktaAuth might just do an authorization with Okta. Or it might not. Shrugs.
func OktaAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header["Authorization"]
		jwt, err := validateAccessToken(accessToken)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), userID, jwt.Claims["sub"].(string))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateAccessToken(accessToken []string) (*jwtverifier.Jwt, error) {
	parts := strings.Split(accessToken[0], " ")
	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Issuer:           "{DOMAIN}",
		ClaimsToValidate: map[string]string{"aud": "api://default", "cid": "{CLIENT_ID}"},
	}
	verifier := jwtVerifierSetup.New()
	return verifier.VerifyIdToken(parts[1])
}

// JSONApi lets the calling site know we're returning JSON
func JSONApi(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

// AccessLog logs what methods are called
func AccessLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s", r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

// Cors enable Cross-origin resource sharing checks
func Cors(h http.Handler) http.Handler {
	corsConfig := cors.New(cors.Options{
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		AllowedMethods: []string{"POST", "PUT", "GET", "PATCH", "OPTIONS", "HEAD", "DELETE"},
		Debug:          true,
	})

	return corsConfig.Handler(h)
}

// UseMiddlewares use the middlewares people
func UseMiddlewares(h http.Handler) http.Handler {
	h = JSONApi(h)
	h = OktaAuth(h)
	h = Cors(h)
	return AccessLog(h)
}
