package service

import (
	"github.com/stg35/avito_test/internal/handler/dto"
	"github.com/stg35/avito_test/internal/model"
	"github.com/stg35/avito_test/internal/repository"
)

type Segment interface {
	CreateSegment(dto dto.SegmentDto) (*model.Segment, error)
	DeleteSegment(id uint64) error
}

type User interface {
	CreateUser(dto dto.UserDto) (*model.User, error)
	AddSegments(dto dto.ChangeSegmentDto) error
	DeleteSegments(dto dto.ChangeSegmentDto) error
	GetSegments(id uint64) ([]model.Segment, error)
}

type Service struct {
	User
	Segment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repo),
		Segment: NewSegmentService(repo),
	}
}
