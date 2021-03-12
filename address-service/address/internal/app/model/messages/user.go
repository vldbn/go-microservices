package messages

type UserMsg struct {
	ID       string `json:"_id"`
	Username string `json:"username,omitempty"`
}
