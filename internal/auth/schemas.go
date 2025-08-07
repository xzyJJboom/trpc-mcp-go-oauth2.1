package auth

// OAuthClientMetadata 定义OAuth客户端元数据的结构，基于OAuth 2.1规范。
// Defines the structure for OAuth client metadata, based on OAuth 2.1 specification.
type OAuthClientMetadata struct {
	// RedirectURIs 是客户端注册的重定向URI列表，必须包含有效的URL。
	// 验证规则：每个URI必须可以通过 net/url.Parse 解析为有效URL。
	// RedirectURIs is a list of redirect URIs for the client, must contain valid URLs.
	// Validation rule: Each URI must be parseable by net/url.Parse.
	RedirectURIs []string `json:"redirect_uris"`

	// TokenEndpointAuthMethod 是令牌端点认证方法，可选。
	// TokenEndpointAuthMethod is the authentication method for the token endpoint, optional.
	TokenEndpointAuthMethod *string `json:"token_endpoint_auth_method,omitempty"`

	// GrantTypes 是支持的授权类型列表，可选。
	// GrantTypes is a list of supported grant types, optional.
	GrantTypes []string `json:"grant_types,omitempty"`

	// ResponseTypes 是支持的响应类型列表，可选。
	// ResponseTypes is a list of supported response types, optional.
	ResponseTypes []string `json:"response_types,omitempty"`

	// ClientName 是客户端的名称，可选。
	// ClientName is the name of the client, optional.
	ClientName *string `json:"client_name,omitempty"`

	// ClientURI 是客户端的URI，可选。
	// ClientURI is the URI of the client, optional.
	ClientURI *string `json:"client_uri,omitempty"`

	// LogoURI 是客户端的Logo URI，可选。
	// LogoURI is the URI of the client's logo, optional.
	LogoURI *string `json:"logo_uri,omitempty"`

	// Scope 是客户端请求的权限范围（以空格分隔的字符串），可选。
	// Scope is the requested scope for the client (space-separated string), optional.
	Scope *string `json:"scope,omitempty"`

	// Contacts 是客户端的联系人列表（通常为邮箱地址），可选。
	// Contacts is a list of contact information for the client (typically email addresses), optional.
	Contacts []string `json:"contacts,omitempty"`

	// TosURI 是客户端的服务条款URI，可选。
	// TosURI is the URI of the client's terms of service, optional.
	TosURI *string `json:"tos_uri,omitempty"`

	// PolicyURI 是客户端的隐私策略URI，可选。
	// PolicyURI is the URI of the client's privacy policy, optional.
	PolicyURI *string `json:"policy_uri,omitempty"`

	// JwksURI 是客户端的JSON Web Key Set URI，可选。
	// JwksURI is the URI of the client's JSON Web Key Set, optional.
	JwksURI *string `json:"jwks_uri,omitempty"`

	// Jwks 是客户端的JSON Web Key Set，任意类型，可选。
	// Jwks is the client's JSON Web Key Set, any type, optional.
	Jwks interface{} `json:"jwks,omitempty"`

	// SoftwareID 是客户端的软件ID，可选。
	// SoftwareID is the software ID of the client, optional.
	SoftwareID *string `json:"software_id,omitempty"`

	// SoftwareVersion 是客户端的软件版本，可选。
	// SoftwareVersion is the software version of the client, optional.
	SoftwareVersion *string `json:"software_version,omitempty"`

	// SoftwareStatement 是客户端的软件声明，可选。
	// SoftwareStatement is the software statement of the client, optional.
	SoftwareStatement *string `json:"software_statement,omitempty"`
}

// OAuthClientInformation 定义RFC 7591 OAuth 2.0动态客户端注册的客户端信息。(OAuth 2.1沿用)
// Defines RFC 7591 OAuth 2.0 Dynamic Client Registration client information.(OAuth 2.1 follows)
type OAuthClientInformation struct {
	// ClientID 是客户端的唯一标识符，必须提供。
	// 验证规则：必须是非空字符串。
	// ClientID is the unique identifier for the client, required.
	// Validation rule: Must be a non-empty string.
	ClientID string `json:"client_id"`

	// ClientSecret 是客户端的密钥，可选。
	// 验证规则：如果提供，必须是非空字符串。
	// ClientSecret is the client's secret, optional.
	// Validation rule: If provided, must be a non-empty string.
	ClientSecret *string `json:"client_secret,omitempty"`

	// ClientIDIssuedAt 是客户端ID的颁发时间（自Unix纪元以来的秒数），可选。
	// 验证规则：如果提供，必须是有效的Unix时间戳（数字）。
	// ClientIDIssuedAt is the issuance time of the client ID (seconds since Unix epoch), optional.
	// Validation rule: If provided, must be a valid Unix timestamp (number).
	ClientIDIssuedAt *int64 `json:"client_id_issued_at,omitempty"`

	// ClientSecretExpiresAt 是客户端密钥的过期时间（自Unix纪元以来的秒数），可选。
	// 验证规则：如果提供，必须是有效的Unix时间戳（数字）。
	// ClientSecretExpiresAt is the expiration time of the client secret (seconds since Unix epoch), optional.
	// Validation rule: If provided, must be a valid Unix timestamp (number).
	ClientSecretExpiresAt *int64 `json:"client_secret_expires_at,omitempty"`
}

type OAuthClientInformationFull struct {
	OAuthClientMetadata
	OAuthClientInformation
}

// OAuthTokens 定义OAuth 2.1令牌响应的结构。
// Defines the OAuth 2.1 token response.
type OAuthTokens struct {
	// AccessToken 是访问令牌，必须提供。
	// 验证规则：必须是非空字符串。
	// AccessToken is the access token, required.
	// Validation rule: Must be a non-empty string.
	AccessToken string `json:"access_token"`

	// IDToken 是OpenID Connect的ID令牌，可选。
	// 在OAuth 2.1中可选，但在OpenID Connect中必需。
	// 验证规则：如果提供，必须是非空字符串。
	// IDToken is the OpenID Connect ID token, optional.
	// Optional for OAuth 2.1, but necessary in OpenID Connect.
	// Validation rule: If provided, must be a non-empty string.
	IDToken *string `json:"id_token,omitempty"`

	// TokenType 是令牌类型（例如"Bearer"），必须提供。
	// 验证规则：必须是非空字符串。
	// TokenType is the token type (e.g., "Bearer"), required.
	// Validation rule: Must be a non-empty string.
	TokenType string `json:"token_type"`

	// ExpiresIn 是访问令牌的过期时间（以秒为单位），可选。
	// 验证规则：如果提供，必须是正整数。
	// ExpiresIn is the expiration time of the access token (in seconds), optional.
	// Validation rule: If provided, must be a positive integer.
	ExpiresIn *int64 `json:"expires_in,omitempty"`

	// Scope 是授权的权限范围（以空格分隔的字符串），可选。
	// 验证规则：如果提供，必须是非空字符串。
	// Scope is the authorized scope (space-separated string), optional.
	// Validation rule: If provided, must be a non-empty string.
	Scope *string `json:"scope,omitempty"`

	// RefreshToken 是刷新令牌，可选。
	// 验证规则：如果提供，必须是非空字符串。
	// RefreshToken is the refresh token, optional.
	// Validation rule: If provided, must be a non-empty string.
	RefreshToken *string `json:"refresh_token,omitempty"`
}
