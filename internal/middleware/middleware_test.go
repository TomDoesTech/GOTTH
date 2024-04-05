package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSPMiddleware(t *testing.T) {

	testCases := []struct {
		name string
	}{
		{
			name: "success",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			assert := assert.New(t)

			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()

				nonces := GetNonces(ctx)
				twNonce := GetTwNonce(ctx)
				htmxNonce := GetHtmxNonce(ctx)
				responseTargetsNonce := GetResponseTargetsNonce(ctx)

				assert.Equal(nonces.Tw, twNonce)
				assert.Len(twNonce, 32)

				assert.Equal(nonces.Htmx, htmxNonce)
				assert.Len(htmxNonce, 32)

				assert.Equal(nonces.ResponseTargets, responseTargetsNonce)
				assert.Len(responseTargetsNonce, 32)

			})

			middleware := CSPMiddleware(next)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/", nil)

			middleware.ServeHTTP(recorder, request)

			csp := recorder.Header().Get("Content-Security-Policy")

			assert.NotEmpty(csp)

		})
	}

}
