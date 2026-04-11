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
	USERID		string `json:"userid"`
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
	Email     string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ----------------------------------------------------------
//Project types

type StoreProject interface {
	GetProjectBYId(name string) (*Project, error)
	GetProjectBYName(name string,id string) (*Project, error) 
	GetProjectBYOwner(id string) (*[]Project, error)
	CreateProject(*Project) error
	UpdateProject(projectId string,payload UpdateProjectPayload)(*Project,error)
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
