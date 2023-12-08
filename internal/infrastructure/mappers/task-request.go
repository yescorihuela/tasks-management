package mappers

type TaskRequest struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ExpiresAt   string `json:"expires_at,omitempty"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}
