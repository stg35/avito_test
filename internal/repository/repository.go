package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/stg35/avito_test/internal/model"
)

type Segment interface {
	CreateSegment(segment model.Segment) (*model.Segment, error)
	DeleteSegment(id uint64) error
}

type User interface {
	CreateUser(user model.User) (*model.User, error)
	AddSegments(id uint64, segmentsName []string) error
	DeleteSegments(id uint64, segmentsName []string) error
	GetSegments(id uint64) ([]model.Segment, error)
}

type Repository struct {
	Segment
	User
}

func NewRepository(db *pg.DB) (*Repository, error) {
	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	return &Repository{
		Segment: NewSegmentRepository(db),
		User:    NewUserRepository(db),
	}, nil
}

func createSchema(db *pg.DB) error {
	orm.RegisterTable((*model.UserSegment)(nil))

	models := []interface{}{
		(*model.User)(nil),
		(*model.Segment)(nil),
		(*model.UserSegment)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
