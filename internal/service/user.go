package service

import (
	"github.com/stg35/avito_test/internal/handler/dto"
	"github.com/stg35/avito_test/internal/model"
	"github.com/stg35/avito_test/internal/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repo,
	}
}

func (s *UserService) CreateUser(dto dto.UserDto) (*model.User, error) {
	user := model.User{Username: dto.Username}
	model, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *UserService) AddSegments(dto dto.ChangeSegmentDto) error {
	err := s.repo.AddSegments(dto.Id, dto.Segments)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteSegments(dto dto.ChangeSegmentDto) error {
	err := s.repo.DeleteSegments(dto.Id, dto.Segments)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetSegments(id uint64) ([]model.Segment, error) {
	segments, err := s.repo.GetSegments(id)
	if err != nil {
		return []model.Segment{}, err
	}
	return segments, nil
}
