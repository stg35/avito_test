package dto

type ChangeSegmentDto struct {
	Id       uint64   `json:"id" binding:"required,min=1"`
	Segments []string `json:"segments" binding:"required"`
	TTL      uint64   `json:"ttl"`
}
