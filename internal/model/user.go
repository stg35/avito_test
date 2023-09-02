package model

type User struct {
	Id       uint64    `pg:",pk"`
	Username string    `pg:",unique"`
	Segments []Segment `pg:"many2many:user_segment"`
}
