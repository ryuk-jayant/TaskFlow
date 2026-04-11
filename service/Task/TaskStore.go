package Task

import (
	"database/sql"
	"example/web-service-gin/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetTaskBYId(name string) (*types.Task, error) {
	return nil, nil
}
func (s *Store) GetTaskBYName(name string, id string) (*types.Task, error) {
	return nil, nil
}
func (s *Store) GetTaskBYOwner(id string) (*[]types.Task, error) {
	return nil, nil
}
func (s *Store) CreateTask(*types.Task) error {
	return nil
}
func (s *Store) UpdateTask(projectId string, payload any) (*types.Task, error) {
	return nil, nil
}
func (s *Store) DeleteTask(projectId string, userId string) error {
	return nil
}
