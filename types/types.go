package types

import (
	"time"

	"github.com/google/uuid"
)

// ----------------------------------------------------
//User types

type StoreUser interface {
	GetUserBYEmail(email string) (*User, error)
	GetUserBYID(id string) (*User, error)
	CreateUser(*User) error
}

type User struct {
	Id         uuid.UUID `json:"id"`
	USERID     string    `json:"userid"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Created_at time.Time `json:"created_at"`
}
type UserRegisterPayload struct {
	Name     string `json:"Name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=4,max=12"`
}
type UserLoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ----------------------------------------------------------
//Project types

type StoreProject interface {
	GetProjectBYId(name string) (*Project, error)
	GetProjectBYName(name string, id string) (*Project, error)
	GetProjectBYOwner(id string) (*[]Project, error)
	CreateProject(*Project) error
	UpdateProject(projectId string, payload UpdateProjectPayload) (*Project, error)
	DeleteProject(projectId string, userId string) error
}

type Project struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Owner_id    string    `json:"-"`
	Created_at  time.Time `json:"created_at"`
}
type NewProjectPayload struct {
	Name        string `json:"Name" validate:"required"`
	Description string `json:"Description"`
}
type UpdateProjectPayload struct {
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
}

// ---------------------------------------------------------------------
// Task
type StoreTask interface {
	GetTaskBYId(name string) (*Task, error)
	GetTaskBYName(name string, id string) (*Task, error)
	GetTaskBYOwner(id string) (*[]Task, error)
	CreateTask(*Task) error
	UpdateTask(projectId string, payload any) (*Task, error)
	DeleteTask(projectId string, userId string) error
}
type Task struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	ProjectId   string    `json:"project_id"`
	AssigneeId  string    `json:"assignee_id"`
	DueDate     time.Time `json:"due_date"`
	Updated_at  time.Time `json:"updated_at"`
	Created_at  time.Time `json:"created_at"`
}
