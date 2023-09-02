package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/stg35/avito_test/internal/model"
)

type SegmentRepository struct {
	db *pg.DB
}

func NewSegmentRepository(db *pg.DB) *SegmentRepository {
	return &SegmentRepository{
		db,
	}
}

func (r *SegmentRepository) CreateSegment(segment model.Segment) (*model.Segment, error) {
	_, err := r.db.Model(&segment).Insert()
	if err != nil {
		return nil, err
	}
	return &segment, nil
}

func (r *SegmentRepository) DeleteSegment(id uint64) error {
	userSegment := new(model.UserSegment)
	_, err := r.db.Model(userSegment).Where("segment_id = ?", id).Delete()
	if err != nil {
		return err
	}
	seg := new(model.Segment)
	_, err = r.db.Model(seg).Where("id = ?", id).Delete()
	if err != nil {
		return err
	}
	return nil
}
