package loginmodels

type LoginPayload struct {
	NickName *string `json:"nick" binding:"required"`
	Password *string `json:"password" binding:"required"`
}
