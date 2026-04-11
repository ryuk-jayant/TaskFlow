package project

import (
	"database/sql"
	"example/web-service-gin/types"
	"fmt"
	"log"
	"strings"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProjectBYId(id string) (*types.Project, error) {
	query := `
		SELECT CONVERT(VARCHAR(36), id) AS id, name, description, created_at
		FROM Project
		WHERE id = @p1
	`

	row := s.db.QueryRow(query, id)

	u := new(types.Project)
	err := row.Scan(
		&u.Id,
		&u.Name,
		&u.Description,
		&u.Created_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product not found")
		}
		return nil, err
	}

	return u, nil
}

func (s *Store) GetProjectBYName(name string, id string) (*types.Project, error) {
	query := `
		SELECT CONVERT(VARCHAR(36), id) AS id, name, description, created_at
		FROM Project
		WHERE name = @p1 AND owner_id =@p2
	`

	row := s.db.QueryRow(query, name, id)

	u := new(types.Project)
	err := row.Scan(
		&u.Id,
		&u.Name,
		&u.Description,
		&u.Created_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product not found")
		}
		return nil, err
	}

	return u, nil
}

func (s *Store) GetProjectBYOwner(id string) (*[]types.Project, error) {
	query := `
		SELECT CONVERT(VARCHAR(36), id) AS id, name, description, created_at
		FROM Project
		WHERE owner_id = @p1
	`
	// 	SELECT CONVERT(VARCHAR(36), your_uuid_column) AS uuid
	// FROM your_table

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []types.Project

	for rows.Next() {
		var p types.Project

		err := rows.Scan(
			&p.Id,
			&p.Name,
			&p.Description,
			&p.Created_at,
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, p)
	}

	// check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// optional: handle empty result
	if len(projects) == 0 {
		return nil, fmt.Errorf("no projects found for owner")
	}

	return &projects, nil
}

func (s *Store) CreateProject(p *types.Project) error {
	query := `
		INSERT INTO Project (id, name, description, created_at,owner_id)
		VALUES (@p1, @p2, @p3, @p4, @p5)
	`

	_, err := s.db.Exec(
		query,
		p.Id,
		p.Name,
		p.Description,
		p.Created_at,
		p.Owner_id,
	)

	return err
}

func (s *Store) UpdateProject(projectId string, payload types.UpdateProjectPayload) (*types.Project, error) {
	query := "UPDATE Project SET "
	args := []interface{}{}
	i := 1
	// log.Println("Payload", *payload.Name, payload.Description)
	if payload.Name != nil {
		query += fmt.Sprintf(" name =@p%d", i)
		args = append(args, *payload.Name)
		i++
	}

	if payload.Description != nil {
		query += fmt.Sprintf(" description =@p%d", i)
		args = append(args, *payload.Description)
		i++
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("Nothing to Update!")
	}
	query = strings.TrimSuffix(query, ", ")
	query += " OUTPUT INSERTED.id, INSERTED.name, INSERTED.description, INSERTED.created_at"
	query += fmt.Sprintf(" WHERE id= @p%d", i)
	args = append(args, projectId)
	log.Println("Query:",query)
	log.Println("args:",args)
	row := s.db.QueryRow(query, args...)
	p := new(types.Project)
	err := row.Scan(
		&p.Id,
		&p.Name,
		&p.Description,
		&p.Created_at,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}
