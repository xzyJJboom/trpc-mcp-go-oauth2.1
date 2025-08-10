package client

import (
	"net/http"
	"net/url"
	"trpc.group/trpc-go/trpc-mcp-go/internal/auth"
)

// OAuthClientProvider 定义一个完整的OAuth客户端接口，用于与一个MCP服务器交互。
// 该客户端依赖于“会话（session）”的概念，其具体含义由应用程序定义。
// 令牌、授权码和代码验证器不应跨不同会话使用。
// Implements an end-to-end OAuth client to be used with one MCP server.
// This client relies upon a concept of an authorized "session," the exact
// meaning of which is application-defined. Tokens, authorization codes, and
// code verifiers should not cross different sessions.
type OAuthClientProvider interface {
	// RedirectURL getter,返回重定向用户代理的URL。
	// The URL to redirect the user agent to after authorization.
	RedirectURL() (*url.URL, error)

	// ClientMetadata getter,返回此OAuth客户端的元数据。
	// Metadata about this OAuth client.
	ClientMetadata() (auth.OAuthClientMetadata, error)

	// State 返回一个OAuth2状态参数，可选。
	// 如果未实现，返回空字符串表示未提供状态。
	// Returns a OAuth2 state parameter, optional.
	// If unimplemented, returns an empty string to indicate no state provided.
	// 可选方法 / Optional method
	//fixme 可选方法转配置选项
	State() (string, error)

	// ClientInformation 加载已注册到服务器的OAuth客户端信息。
	// 如果客户端未注册到服务器，返回nil。
	// Loads information about this OAuth client, as registered already with the
	// server, or returns `nil` if the client is not registered with the server.
	ClientInformation() (*auth.OAuthClientInformation, error)

	// SaveClientInformation 动态将客户端注册到服务器，可选。
	// 通过此方法保存的客户端信息应通过 ClientInformation() 读取。
	// 如果客户端信息是静态已知的（例如预注册），则无需实现此方法。
	// If implemented, this permits the OAuth client to dynamically register with
	// the server. Client information saved this way should later be read via
	// `ClientInformation()`.
	// This method is not required to be implemented if client information is
	// statically known (e.g., pre-registered).
	// 可选方法 / Optional method
	//fixme 可选方法转配置选项
	SaveClientInformation(clientInformation auth.OAuthClientInformationFull) error

	// Tokens 加载当前会话的现有OAuth令牌。
	// 如果没有保存的令牌，返回nil。
	// Loads any existing OAuth tokens for the current session, or returns
	// `nil` if there are no saved tokens.
	Tokens() (*auth.OAuthTokens, error)

	// SaveTokens 保存当前会话的新OAuth令牌，在授权成功后调用。
	// Stores new OAuth tokens for the current session, after a successful
	// authorization.
	SaveTokens(tokens auth.OAuthTokens) error

	// RedirectToAuthorization 将用户代理重定向到给定的URL以开始授权流程。
	// Invoked to redirect the user agent to the given URL to begin the authorization flow.
	// fixme url结构体与ts SDK不一致
	RedirectToAuthorization(authorizationUrl *url.URL) error

	// SaveCodeVerifier 保存当前会话的PKCE代码验证器，在重定向到授权流程之前调用。
	// Saves a PKCE code verifier for the current session, before redirecting to
	// the authorization flow.
	SaveCodeVerifier(codeVerifier string) error

	// CodeVerifier 加载当前会话的PKCE代码验证器，用于验证授权结果。
	// Loads the PKCE code verifier for the current session, necessary to validate
	// the authorization result.
	CodeVerifier() (string, error)

	// AddClientAuthentication 为OAuth令牌请求添加自定义客户端认证，可选。
	// 允许实现自定义客户端凭据在令牌交换和刷新请求中的包含方式。
	// 如果提供此方法，将替代默认认证逻辑，完全控制认证机制。
	// 常见用例包括：
	// - 支持超出标准OAuth 2.0的认证方法
	// - 添加专有认证方案的自定义头
	// - 实现基于客户端断言的认证（例如JWT bearer令牌）
	// Adds custom client authentication to OAuth token requests.
	// This optional method allows implementations to customize how client credentials
	// are included in token exchange and refresh requests. When provided, this method
	// is called instead of the default authentication logic, giving full control over
	// the authentication mechanism.
	// Common use cases include:
	// - Supporting authentication methods beyond the standard OAuth 2.0 methods
	// - Adding custom headers for proprietary authentication schemes
	// - Implementing client assertion-based authentication (e.g., JWT bearer tokens)
	// 可选方法 / Optional method
	//fixme 可选方法，metadata为可选
	AddClientAuthentication(headers http.Header, params url.Values, urlStr string, metadata *auth.AuthorizationServerMetadata) error

	// ValidateResourceURL 覆盖RFC 8707资源指示器的选择和验证，可选。
	// 如果未实现，将使用默认验证行为。
	// 实现必须验证返回的资源与MCP服务器匹配。
	// If defined, overrides the selection and validation of the
	// RFC 8707 Resource Indicator. If left undefined, default
	// validation behavior will be used.
	// Implementations must verify the returned resource matches the MCP server.
	// 可选方法 / Optional method
	//fixme 可选方法，resource为可选
	ValidateResourceURL(serverUrl string, resource string) (*url.URL, error)

	// InvalidateCredentials 使指定的凭据失效（例如删除），在服务器指示凭据不再有效时调用，可选。
	// 避免用户手动干预。
	// scope: 'all' | 'client' | 'tokens' | 'verifier'
	// If implemented, provides a way for the client to invalidate (e.g. delete) the specified
	// credentials, in the case where the server has indicated that they are no longer valid.
	// This avoids requiring the user to intervene manually.
	// 可选方法 / Optional method
	//fixme 可选方法
	InvalidateCredentials(scope string) error
}
