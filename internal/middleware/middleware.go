package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

type key string

var NonceKey key = "nonces"

type Nonces struct {
	Htmx             string
	ResponseTargets  string
	Tw               string
	WordPressScripts string
	WordPressStyles  string
	HtmxCSSHash      string
}

func generateRandomString(length int) string {

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func CSPMiddleware(next http.Handler) http.Handler {
	// To use the same nonces in all responses, move the Nonces
	// struct creation to here, outside the handler.

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new Nonces struct for every request when here.
		// move to outisde the handler to use the same nonces in all responses
		var nonceSet Nonces
		nonceSet.Htmx = generateRandomString(16)
		nonceSet.ResponseTargets = generateRandomString(16)
		nonceSet.Tw = generateRandomString(16)

		htmxCSSHash := "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg="

		// set nonces in context
		ctx := context.WithValue(r.Context(), NonceKey, nonceSet)
		// insert the nonces into the content security policy header
		cspHeader := fmt.Sprintf("default-src 'self'; script-src 'nonce-%s' 'nonce-%s' ; style-src 'nonce-%s' '%s';", nonceSet.Htmx, nonceSet.ResponseTargets, nonceSet.Tw, htmxCSSHash)
		w.Header().Set("Content-Security-Policy", cspHeader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TextHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

// get the Nonce from the context, it is a struct called Nonces, so we can get the nonce we need by the key, i.e. HtmxNonce
func GetNonces(ctx context.Context) any {

	nonceSet := ctx.Value(NonceKey)
	if nonceSet == nil {
		log.Fatal("nooo no nonces")
	}
	return nonceSet
}

func GetHtmxNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx).(Nonces)
	return nonceSet.Htmx
}

func GetResponseTargetsNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx).(Nonces)
	return nonceSet.ResponseTargets
}

func GetTwNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx).(Nonces)
	return nonceSet.Tw
}
