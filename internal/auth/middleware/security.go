package middleware

import "trpc.group/trpc-go/trpc-mcp-go/internal/auth"

type SecurityMiddlewareOption struct {
	verifier auth.TokenVerifier
}
