package repository

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/stg35/avito_test/internal/model"
)

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) CreateUser(user model.User) (*model.User, error) {
	_, err := r.db.Model(&user).Insert()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) AddSegments(id uint64, segmentsName []string) error {
	ctx := context.Background()
	err := r.execTx(ctx, func(tx *pg.Tx) error {
		for _, name := range segmentsName {
			seg := new(model.Segment)
			err := tx.Model(seg).Where("name = ?", name).Select()
			if err != nil {
				return err
			}
			user := new(model.User)
			err = tx.Model(user).Where("id = ?", id).Select()
			if err != nil {
				return err
			}
			userSegment := &model.UserSegment{
				UserId:    id,
				SegmentId: seg.Id,
			}
			_, err = tx.Model(userSegment).Insert()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetSegments(id uint64) ([]model.Segment, error) {
	user := new(model.User)
	err := r.db.Model(user).Relation("Segments").Where("id = ?", id).Select()
	if err != nil {
		return []model.Segment{}, err
	}
	return user.Segments, nil
}

func (r *UserRepository) DeleteSegments(id uint64, segmentsName []string) error {
	ctx := context.Background()
	err := r.execTx(ctx, func(tx *pg.Tx) error {
		for _, name := range segmentsName {
			seg := new(model.Segment)
			err := tx.Model(seg).Where("name = ?", name).Select()
			if err != nil {
				return err
			}
			user := new(model.User)
			err = tx.Model(user).Where("id = ?", id).Select()
			if err != nil {
				return err
			}
			userSegment := new(model.UserSegment)
			_, err = tx.Model(userSegment).Where("user_id = ?", id).Where("segment_id = ?", seg.Id).Delete()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) execTx(ctx context.Context, fn func(*pg.Tx) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Close()

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
