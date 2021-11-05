package loginmodels

type LoginPayload struct {
	NickName *string `json:"nick" binding:"required"`
	Password *string `json:"password" binding:"required"`
}

// Token login response
type TokenResponse struct {
	Token   string `json:"token,omitempty"`   // Session token
	Expires string `json:"expires,omitempty"` // Expiration timestamp
}
