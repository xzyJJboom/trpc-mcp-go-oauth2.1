package errors

// OAuthErrorResponse OAuth 2.1草案标准错误响应。
type OAuthErrorResponse struct {
	Error            string  `json:"error"`
	ErrorDescription *string `json:"error_description,omitempty"`
	ErrorURI         *string `json:"error_uri,omitempty"`
}

// ToResponseObject 转换为OAuthErrorResponse。
func (e OAuthError) ToResponseObject() OAuthErrorResponse {
	resp := OAuthErrorResponse{
		Error:            e.Code,
		ErrorDescription: &e.Message,
	}
	if e.URI != nil {
		resp.ErrorURI = e.URI
	}
	return resp
}
