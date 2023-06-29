package response

type LoginRes struct {
	Name   string `json:"name"`
	Token  string `json:"token"`
	Status uint   `json:"status"`
}
