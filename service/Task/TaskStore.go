package Task

import (
	"database/sql"
	"example/web-service-gin/types"
	"fmt"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetTaskBYId(id string) (*types.Task, error) {
	query := `
		SELECT CONVERT(VARCHAR(36), id) AS id, title, description,status,priority,project_id,assignee_id,due_date,updated_at, created_at
		FROM Task
		WHERE id = @p1
	`

	row := s.db.QueryRow(query, id)

	u := new(types.Task)
	err := row.Scan(
		&u.Id,
		&u.Title,
		&u.Description,
		&u.ProjectId,
		&u.Status,
		&u.Priority,
		&u.AssigneeId,
		&u.DueDate,
		&u.Created_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Task not found")
		}
		return nil, err
	}

	return u, nil
}
func (s *Store) GetTaskBYName(name string, projectId string) (*types.Task, error) {
	query := `
		SELECT CONVERT(VARCHAR(36), id) AS id,title
		FROM Task
		WHERE title = @p1 AND project_id=@p2
	`

	row := s.db.QueryRow(query, name, projectId)

	u := new(types.Task)
	err := row.Scan(
		&u.Id,
		&u.Title,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Task not found")
		}
		return nil, err
	}

	return u, nil
}
func (s *Store) GetTaskBYProject(id string, payload types.GetTaskQueryFilters) (*[]types.Task, error) {
	query := "SELECT CONVERT(VARCHAR(36), id) AS id, title, description,status,priority,CONVERT(VARCHAR(36), project_id) AS project_id,CONVERT(VARCHAR(36), assignee_id) AS assignee_id,due_date,updated_at, created_at FROM Task WHERE "
	args := []interface{}{}
	i := 1
	//project_id = @p1
	log.Println("Payload:", payload)
	if payload.Status != nil {
		query += fmt.Sprintf(" status =@p%d AND", i)
		args = append(args, *payload.Status)
		i++
	}

	if payload.AssigneeId != nil {
		query += fmt.Sprintf(" assignee_id =@p%d AND", i)
		args = append(args, *payload.AssigneeId)
		i++
	}

	query += fmt.Sprintf(" project_id=@p%d", i)
	args = append(args, id)

	log.Println("Query:", query)
	log.Println("Args:", args)
	rows, err := s.db.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []types.Task

	for rows.Next() {
		var u types.Task

		err := rows.Scan(
			&u.Id,
			&u.Title,
			&u.Description,
			&u.Status,
			&u.Priority,
			&u.ProjectId,
			&u.AssigneeId,
			&u.DueDate,
			&u.Updated_at,
			&u.Created_at,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, u)
	}

	// check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	//handle empty result
	if len(tasks) == 0 {
		return nil, fmt.Errorf("no projects found for owner")
	}

	return &tasks, nil
}
func (s *Store) CreateTask(t *types.Task) error {
	query := `
		INSERT INTO Task (id, title, description,status,priority,project_id,assignee_id,due_date,updated_at, created_at)
		VALUES (@p1, @p2, @p3, @p4, @p5,@p6,@p7,@p8,@p9,@p10)
	`

	_, err := s.db.Exec(
		query,
		t.Id,
		t.Title,
		t.Description,
		t.Status,
		t.Priority,
		t.ProjectId,
		t.AssigneeId,
		t.DueDate,
		t.Updated_at,
		t.Created_at,
	)

	return err
}
func (s *Store) UpdateTask(projectId string, payload any) (*types.Task, error) {
	return nil, nil
}
func (s *Store) DeleteTask(projectId string, userId string) error {
	return nil
}
