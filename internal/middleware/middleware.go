package middleare

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
)

func generateRandomString(length int) string {

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func CSPMiddleware(next http.Handler) http.Handler {
	htmxNonce := generateRandomString(16)
	responseTargetsNonse := generateRandomString(16)
	twNonce := generateRandomString(16)

	// set then in context
	ctx := context.WithValue(context.Background(), "htmxNonce", htmxNonce)
	ctx = context.WithValue(ctx, "twNonce", twNonce)
	ctx = context.WithValue(ctx, "responseTargetsNonse", responseTargetsNonse)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// the hash of the CSS that HTMX injects
		htmxCSSHash := "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg="

		cspHeader := fmt.Sprintf("default-src 'self'; script-src 'nonce-%s' 'nonce-%s'; style-src 'nonce-%s' '%s';", htmxNonce, responseTargetsNonse, twNonce, htmxCSSHash)
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
