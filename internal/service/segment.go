package service

import (
	"github.com/stg35/avito_test/internal/handler/dto"
	"github.com/stg35/avito_test/internal/model"
	"github.com/stg35/avito_test/internal/repository"
)

type SegmentService struct {
	repo *repository.Repository
}

func NewSegmentService(repo *repository.Repository) *SegmentService {
	return &SegmentService{
		repo,
	}
}

func (s *SegmentService) CreateSegment(dto dto.SegmentDto) (*model.Segment, error) {
	segment := model.Segment{Name: dto.Name}
	model, err := s.repo.CreateSegment(segment)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *SegmentService) DeleteSegment(id uint64) error {
	err := s.repo.DeleteSegment(id)
	if err != nil {
		return err
	}
	return nil
}
