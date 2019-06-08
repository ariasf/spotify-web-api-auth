package structures

type TokenResponse struct {
	AccessToken             string `json:"access_token"`
	TokenType               string `json:"token_type"`
	ExpiresIn               int64  `json:"expires_in"`
	RefreshToken            string `json:"refresh_token"`
	Scope                   string `json:"scope"`
	ErrorCode               string `json:"error_code"`
	ExpirationTimeStampInMs int64  `json:"expiration_time_stamp_in_ms"`
}
