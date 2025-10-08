package rest

type accountRequest struct {
	Owner    string  `json:"owner" binding:"required"`
	Currency string  `json:"currency"`
	Balance  float32 `json:"balance"`
}
