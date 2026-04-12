package user

import (
	"encoding/json"
	"example/web-service-gin/cmd/config"
	"example/web-service-gin/service/auth"
	"example/web-service-gin/types"
	"example/web-service-gin/utils"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	Store types.StoreUser
}

type Error struct {
	Err     string
	Message string
	Code    int
}

func NewHandler(store types.StoreUser) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.loginHandler).Methods("POST")
	router.HandleFunc("/register", h.registerHandler).Methods("POST")
}

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	//login flow
	//.take payload->parse it ->2.check unique with email(response->400) ->3hash the password
	//->check if password matches create a session jwt token
	var payload types.UserLoginPayload
	if err := utils.DecodePayload(r, &payload); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, err)
		return
	}
	//2
	user, err := h.Store.GetUserBYEmail(payload.Email)
	if err != nil {
		err := Error{Err: "", Message: "Wrong Credentials", Code: http.StatusBadRequest}
		errJson, _ := json.Marshal(err)
		http.Error(w, string(errJson), http.StatusBadRequest)
		return
	}

	//3
	if !auth.ComparePasswords(user.Password, []byte(payload.Password)) {
		err := Error{Err: "", Message: "Wrong Credentials", Code: http.StatusBadRequest}
		errJson, _ := json.Marshal(err)
		http.Error(w, string(errJson), http.StatusBadRequest)
		return
	}
	//4 JWT Token
	secret := []byte(config.GetEnv("Secret", "fallb1ckSecretPassw0r4")) //TODO:ADD SECRET TO ENV
	log.Println("UserId:", user)
	token, err := auth.CreateJWT(secret, user.Id.String(), user.Email)
	if err != nil {
		err := Error{Err: "", Message: "Error While lookup!", Code: http.StatusInternalServerError}
		errJson, _ := json.Marshal(err)
		http.Error(w, string(errJson), http.StatusInternalServerError)
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{"Token": token})

}

func (h *Handler) registerHandler(w http.ResponseWriter, r *http.Request) {
	//Registration Flow
	//1.take payload->parse it ->2.check unique with email(response->400) ->3hash the password->store->response OK
	//1
	var payload types.UserRegisterPayload
	if err := utils.DecodePayload(r, &payload); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"err": "Encountered an error"})
		return
	}
	//2
	_, err := h.Store.GetUserBYEmail(payload.Email)
	if err == nil {
		err := Error{Err: "", Message: "User Exists with " + payload.Email, Code: http.StatusConflict}
		errJson, _ := json.Marshal(err)
		http.Error(w, string(errJson), http.StatusConflict)
		return
	}
	//3
	hashedPassword, err := auth.HashFromPassword(payload.Password)
	if err != nil {
		utils.WriteJson(w, http.StatusFailedDependency, map[string]string{"err": "Encountered an error"})
		return
	}

	//4
	var UserInsertion types.User
	UserInsertion.Id = uuid.New()
	UserInsertion.Created_at = time.Now()
	UserInsertion.Name = payload.Name
	UserInsertion.Email = payload.Email
	UserInsertion.Password = hashedPassword

	err = h.Store.CreateUser(&UserInsertion)
	if err != nil {
		utils.WriteJson(w, http.StatusFailedDependency, map[string]string{"err": "Encountered an error"})
		return
	}

	utils.WriteJson(w, http.StatusAccepted, map[string]string{"Message": "User added to DB!"})

}
