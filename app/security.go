// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "pei0804/goa-stater": Application Security
//
// Command:
// $ goagen
// --design=github.com/KokiKono/go_tutorial/design
// --out=$(GOPATH)/src/github.com/KokiKono/go_tutorial
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

type (
	// Private type used to store auth handler info in request context
	authMiddlewareKey string
)

// UseAdminAuthMiddleware mounts the adminAuth auth middleware onto the service.
func UseAdminAuthMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("adminAuth"), middleware)
}

// NewAdminAuthSecurity creates a adminAuth security definition.
func NewAdminAuthSecurity() *goa.APIKeySecurity {
	def := goa.APIKeySecurity{
		In:   goa.LocHeader,
		Name: "X-Authorization",
	}
	def.Description = "トークン(admin)"
	return &def
}

// UseGeneralAuthMiddleware mounts the generalAuth auth middleware onto the service.
func UseGeneralAuthMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("generalAuth"), middleware)
}

// NewGeneralAuthSecurity creates a generalAuth security definition.
func NewGeneralAuthSecurity() *goa.APIKeySecurity {
	def := goa.APIKeySecurity{
		In:   goa.LocHeader,
		Name: "X-Authorization",
	}
	def.Description = "トークン(general)"
	return &def
}

// handleSecurity creates a handler that runs the auth middleware for the security scheme.
func handleSecurity(schemeName string, h goa.Handler, scopes ...string) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		scheme := ctx.Value(authMiddlewareKey(schemeName))
		am, ok := scheme.(goa.Middleware)
		if !ok {
			return goa.NoAuthMiddleware(schemeName)
		}
		ctx = goa.WithRequiredScopes(ctx, scopes)
		return am(h)(ctx, rw, req)
	}
}
