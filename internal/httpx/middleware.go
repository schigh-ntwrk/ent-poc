package httpx

import (
	"net/http"
)

type middlewareChain []func(http.Handler) http.Handler

func (m middlewareChain) Apply(next http.Handler) http.Handler {
	for x := len(m) - 1; x >= 0; x = x - 1 {
		next = m[x](next)
	}
	return next
}
