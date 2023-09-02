package dto

type SegmentDto struct {
	Name string `json:"name" binding:"required"`
}
