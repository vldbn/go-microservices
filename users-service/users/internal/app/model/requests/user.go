package requests

type UserReq struct {
	ID       string `json:"_id,omitempty"`
	Username string `json:"username"`
}
