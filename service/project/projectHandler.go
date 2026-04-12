package project

import (
	"encoding/json"
	"example/web-service-gin/cmd/config"
	"example/web-service-gin/service/middleware"
	"example/web-service-gin/types"
	"example/web-service-gin/utils"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	PStore types.StoreProject
	SStore types.StoreTask
}

type Error struct {
	Err     string
	Message string
	Code    int
}

func NewHandler(projectstore types.StoreProject, taskstore types.StoreTask) *Handler {
	return &Handler{PStore: projectstore, SStore: taskstore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	secret := []byte(config.GetEnv("Secret", "fallb1ckSecretPassw0r4")) //TODO:ADD SECRET TO ENV
	protected := router.PathPrefix("/projects").Subrouter()
	protected.Use(middleware.JWTMiddleware(secret))
	protected.HandleFunc("/", h.addProject).Methods("POST")          //add project
	protected.HandleFunc("/", h.getAllProjects).Methods("GET")       //get all projects of user
	protected.HandleFunc("/{id}", h.getProject).Methods("GET")       //get all project by id
	protected.HandleFunc("/{id}", h.editProject).Methods("PUT")      //modify a project by id
	protected.HandleFunc("/{id}", h.deleteProject).Methods("DELETE") //delete a project and all the tasks in it
	// -------------------------------------------------------------------------
	//Task APIS
	protected.HandleFunc("/:id/tasks", h.listTasks).Methods("GET")    //list tasks
	protected.HandleFunc("/:id/tasks", h.addTask).Methods("POST")     //Create a task
	protected.HandleFunc("/:id/tasks", h.listTasks).Methods("PUT")    //Update title, description, status, priority, assignee, due_date
	protected.HandleFunc("/:id/tasks", h.listTasks).Methods("DELETE") //Delete task
}

func (h *Handler) addProject(w http.ResponseWriter, r *http.Request) {
	//0 context extraction
	userId := r.Context().Value("userId").(string)
	// email := r.Context().Value("email").(string)
	// utils.WriteJson(w, http.StatusAccepted, map[string]string{"user":userId,"email":email})
	//1
	var payload types.NewProjectPayload
	if err := utils.DecodePayload(r, &payload); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"err": "Encountered an error"})
		return
	}
	//2 check if project is non unique
	_, err := h.PStore.GetProjectBYName(payload.Name, userId)
	if err == nil {
		err := Error{Err: "", Message: "Project with " + payload.Name + " Name Exists", Code: http.StatusConflict}
		errJson, _ := json.Marshal(err)
		http.Error(w, string(errJson), http.StatusConflict)
		return
	}
	//3
	var ProjectInsertion types.Project
	ProjectInsertion.Id = uuid.New()
	ProjectInsertion.Created_at = time.Now()
	ProjectInsertion.Name = payload.Name
	ProjectInsertion.Description = payload.Description
	ProjectInsertion.Owner_id = userId
	//4
	log.Println("Creating project for ", userId)
	err = h.PStore.CreateProject(&ProjectInsertion)
	if err != nil {
		utils.WriteJson(w, http.StatusFailedDependency, map[string]string{"error": "Error Creating Project!"})
		return
	}

	utils.WriteJson(w, http.StatusAccepted, map[string]string{"Message": "Project added to DB!"})
}

func (h *Handler) getAllProjects(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, map[string]string{
			"error": "unauthorized",
		})
		return
	}

	projects, err := h.PStore.GetProjectBYOwner(userId)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch projects",
		})
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"projects": projects,
		"message":  "Projects for " + userId,
	})

}

func (h *Handler) getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["id"]
	if projectId == "" {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{
			"error": "No valid ID Found!",
		})
		return
	}

	project, err := h.PStore.GetProjectBYId(projectId)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{
			"error": "failed to fetch projects",
		})
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"project": project,
		"message": "Project Fetched!",
	})

}

func (h *Handler) editProject(w http.ResponseWriter, r *http.Request) {
	//1.get id from url
	vars := mux.Vars(r)
	projectId := vars["id"]
	if projectId == "" {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{
			"error": "No valid ID Found!",
		})
		return
	}
	//2 get payload
	var payload types.UpdateProjectPayload
	if err := utils.DecodePayload(r, &payload); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"err": "Encountered an error in Validation"})
		return
	}
	//3 callDB
	project, err := h.PStore.UpdateProject(projectId, payload)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError,
			map[string]string{
				"error": "Error While Updating Project",
			})
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"project": project,
		"message": "Project Updated!",
	})
}

func (h *Handler) deleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["id"]
	if projectId == "" {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{
			"error": "No valid ID Found!",
		})
		return
	}
	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, map[string]string{
			"error": "unauthorized",
		})
		return
	}

	err := h.PStore.DeleteProject(projectId, userId)
	if err != nil {
		utils.WriteJson(w, http.StatusFailedDependency, map[string]string{"Error": "Error Deleting Project!"})
		return
	}

	utils.WriteJson(w, http.StatusAccepted, map[string]string{"Message": "Project Deleted !"})
}

func (h *Handler) addTask(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) listTasks(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {

}
