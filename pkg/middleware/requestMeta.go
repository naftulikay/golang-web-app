package middleware

import (
	"context"
	"github.com/google/uuid"
	"github.com/naftulikay/golang-webapp/pkg/constants"
	"net/http"
)

func RequestMetadata() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			AppendRequestID(w, r)

			next.ServeHTTP(w, r)
		})
	}
}

func AppendRequestID(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()

	ctx := context.WithValue(r.Context(), constants.ContextKeyRequestID, id)

	// update the request, adding the request id to the context
	*r = *r.WithContext(ctx)

	// update the response, adding the request ID as a header
	// NOTE we use Header.Add and not Header.Set because it's possible that there may be multiple request ids from
	//      external sources, e.g. if there are multiple hops to your webapp. for instance, your CDN may create a
	//      request id and attach it, then your edge (NGINX/Envoy/Istio) could append another request id, etc. leaving
	//      a single request with a stack of request ids. we only care about the request id that _we_ created for
	//      internally tracking requests
	w.Header().Add(constants.HeaderKeyRequestID, id.String())
}
