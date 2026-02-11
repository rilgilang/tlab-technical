package dto

// Handle http response with Pagination
type (
	Pagination struct {
		Cursor     int64 `json:"cursor"`
		NextCursor int64 `json:"next_cursor"`
		Total      int   `json:"total"`
	}
)
