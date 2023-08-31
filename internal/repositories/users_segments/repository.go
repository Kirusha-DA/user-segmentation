package userssegments

import (
	"errors"

	"github.com/Kirusha-DA/user-segmentation/internal/entities"
	"github.com/Kirusha-DA/user-segmentation/internal/models"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewReporepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) InsertSegments(segments []models.Segment, clientId int) ([]models.UsersSegments, error) {
	var slugs []string
	for _, segment := range segments {
		slugs = append(slugs, segment.Slug)
	}

	arg := map[string]interface{}{
		"user_id": clientId,
		"slugs":   slugs,
	}

	query, args, _ := sqlx.Named(
		`
			INSERT INTO users_segments(user_id, segment_id)
			SELECT :user_id, id
			FROM segments
			WHERE slug IN (:slugs)
			RETURNING segment_id;
		`,
		arg,
	)

	query, args, _ = sqlx.In(query, args...)
	query = r.db.Rebind(query)

	var res []entities.UsersSegments
	err := r.db.Select(
		&res,
		query,
		args...,
	)
	if err != nil {
		return nil, errors.New("execute query")
	}

	var modelSegments []models.UsersSegments
	for _, value := range res {
		modelSegments = append(modelSegments, models.UsersSegments(value))
	}
	return modelSegments, nil
}
