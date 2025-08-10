package middleware

import (
	"trpc.group/trpc-go/trpc-mcp-go/internal/auth/server"
)

type SecurityMiddlewareOption struct {
	verifier server.TokenVerifier
}
