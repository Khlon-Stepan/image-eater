package storage

import (
	"context"
)

func (s *Storage) CreateTask(ctx context.Context, originalFilename, originalURL string) (string, error) {
	query := `
	INSERT INTO image_tasks (original_filename, original_url)
	VALUES ($1, $2)
	RETURNING id
	`

	var id string
	err := s.pool.QueryRow(ctx, query, originalFilename, originalURL).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, err
}

func (s *Storage) UpdateTask(ctx context.Context, taskID, status, processedURL string) error {
	query := `
	UPDATE image_tasks
	SET status = $1,
	processedURL = $2,
	updated_at = CURRENT_TIMESTAMP
	WHERE id = $3
	`
	_, err := s.pool.Exec(ctx, query, status, processedURL, taskID)
	if err != nil {
		return err
	}
	return nil
}
