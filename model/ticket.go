package model

// Ticket model for app_key table
type Ticket struct {
	ID      int64  `json:"id"`
	Site    string `json:"site"`
	Secret  string `json:"secret"`
	UserID  int64  `json:"userId"`
	Enabled string `json:"enabled"`
}
