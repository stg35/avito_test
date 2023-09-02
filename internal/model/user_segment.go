package model

type UserSegment struct {
	tableName struct{} `pg:"user_segment"`
	UserId    uint64   `pg:",pk"`
	SegmentId uint64   `pg:",pk"`
}
