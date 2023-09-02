package model

type Segment struct {
	Id   uint64 `pg:",pk"`
	Name string `pg:",unique"`
}
