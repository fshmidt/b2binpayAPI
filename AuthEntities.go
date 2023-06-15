package b2binpay

type AuthRequest struct {
	Data AuthRequestData `json:"data"`
}

type AuthRequestData struct {
	Type       string                `json:"type"`
	Attributes AuthRequestAttributes `json:"attributes"`
}

type AuthRequestAttributes struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Data AuthResponseData `json:"data"`
	Meta AuthResponseMeta `json:"meta"`
}

type AuthResponseData struct {
	Type       string                 `json:"type"`
	ID         string                 `json:"id"`
	Attributes AuthResponseAttributes `json:"attributes"`
}

type AuthResponseAttributes struct {
	Refresh          string `json:"refresh"`
	Access           string `json:"access"`
	AccessExpiredAt  string `json:"access_expired_at"`
	RefreshExpiredAt string `json:"refresh_expired_at"`
	Is2FAConfirmed   bool   `json:"is_2fa_confirmed"`
}

type AuthResponseMeta struct {
	Time string `json:"time"`
	Sign string `json:"sign"`
}

type RefreshRequest struct {
	Data RefreshRequestData `json:"data"`
}

type RefreshRequestData struct {
	Type       string                   `json:"type"`
	Attributes RefreshRequestAttributes `json:"attributes"`
}

type RefreshRequestAttributes struct {
	Refresh string `json:"refresh"`
}

type RefreshResponse struct {
	Data RefreshResponseData `json:"data"`
}

type RefreshResponseData struct {
	Type       string                    `json:"type"`
	ID         string                    `json:"id"`
	Attributes RefreshResponseAttributes `json:"attributes"`
}

type RefreshResponseAttributes struct {
	Refresh          string `json:"refresh"`
	Access           string `json:"access"`
	AccessExpiredAt  string `json:"access_expired_at"`
	RefreshExpiredAt string `json:"refresh_expired_at"`
	Is2FAConfirmed   bool   `json:"is_2fa_confirmed"`
}
