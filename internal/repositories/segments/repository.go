package segments

import (
	"github.com/Kirusha-DA/user-segmentation/internal/entities"
	"github.com/Kirusha-DA/user-segmentation/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repository struct {
	db *sqlx.DB
}

func NewReporepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(segment models.Segment) (*models.Segment, error) {
	var segmentID int
	err := r.db.Get(
		&segmentID,
		`
			INSERT INTO segments (slug)
			VALUES ($1)
			RETURNING id
		`,
		segment.Slug,
	)
	if err != nil {
		return nil, errors.New("slug already exists")
	}
	segment.Id = segmentID
	return &segment, nil
}

func (r *repository) Delete(slug string) (bool, error) {
	res, err := r.db.Exec(
		`
			DELETE FROM segments
			WHERE slug = $1
		`,
		slug,
	)
	if err != nil {
		return false, errors.New("delete failed")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "rows affected")
	}
	if rowsAffected == 0 {
		return false, errors.New("no such slug")
	}

	return true, nil
}

func (r *repository) GetSlugssById(ids []int) []models.Segment {

	query, args, _ := sqlx.In(
		`
			SELECT slug
			FROM segments
			WHERE id IN (?);
		`,
		ids,
	)
	query = r.db.Rebind(query)
	var res []entities.Segment
	r.db.Select(
		&res,
		query,
		args...,
	)
	var modelSegments []models.Segment
	for _, value := range res {
		modelSegments = append(modelSegments, models.Segment(value))
	}

	return modelSegments
}

func (r *repository) ReadClients(userId int) ([]models.Segment, error) {
	segments := make([]entities.Segment, 0)
	err := r.db.Select(
		&segments,
		`
			SELECT s.id as id, s.slug as slug
			FROM users_segments as u_s 
				INNER JOIN segments as s
				ON u_s.segment_id = s.id
			WHERE u_s.user_id = $1
		`,
		userId,
	)
	if err != nil {
		return nil, errors.Wrap(err, "user's segments")
	}

	var modelSegments []models.Segment
	for _, value := range segments {
		modelSegments = append(modelSegments, models.Segment(value))
	}

	return modelSegments, nil
}

func (r *repository) DeleteClients(segments []models.Segment, clientId int) []models.Segment {
	var slugs []string
	for _, value := range segments {
		slugs = append(slugs, value.Slug)
	}

	arg := map[string]interface{}{
		"user_id": clientId,
		"slugs":   slugs,
	}

	query, args, _ := sqlx.Named(
		`
			DELETE FROM users_segments
			WHERE 
				user_id = :user_id 
			AND 
				segment_id IN 
				(
					SELECT id 
					FROM segments
					WHERE slug IN (:slugs)
				)
			RETURNING segment_id;
		`,
		arg,
	)
	query, args, _ = sqlx.In(query, args...)
	query = r.db.Rebind(query)
	var res []entities.UsersSegments
	r.db.Select(
		&res,
		query,
		args...,
	)
	var modelSegments []models.Segment
	for _, value := range res {
		modelSegments = append(modelSegments, models.Segment{
			Id: value.Segment_id,
		})
	}
	return modelSegments
}
