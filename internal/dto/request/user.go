package request

type LoginReq struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
