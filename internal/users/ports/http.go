package ports

import (
	"github.com/go-chi/render"
	"github.com/noodlensk/task-tracker/internal/common/server/httperr"
	"github.com/noodlensk/task-tracker/internal/users/app"
	"github.com/noodlensk/task-tracker/internal/users/app/command"
	"github.com/noodlensk/task-tracker/internal/users/app/query"
	"github.com/noodlensk/task-tracker/internal/users/domain/user"
	"net/http"
)

type HTTPServer struct {
	app *app.Application
}

func (h HTTPServer) AuthLogin(w http.ResponseWriter, r *http.Request) {
	loginReq := AuthLoginRequest{}
	if err := render.Decode(r, &loginReq); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	res, err := h.app.Queries.AuthLogin.Handle(r.Context(), query.AuthLogin{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	})

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	authResp := AuthLoginResult{Token: res.Token}

	render.Respond(w, r, authResp)
}

func (h HTTPServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	userToCreate := CreateUserRequest{}
	if err := render.Decode(r, &userToCreate); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	var userRole user.UserRole
	switch userToCreate.Role {
	case CreateUserRequestRoleBasic:
		userRole = user.RoleBasic
	case CreateUserRequestRoleManager:
		userRole = user.RoleManager
	case CreateUserRequestRoleAdmin:
		userRole = user.RoleAdmin
	}

	u, err := user.NewUser(userToCreate.Name, userToCreate.Email, userRole, userToCreate.Password)
	if err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	if err := h.app.Commands.CreateUser.Handle(r.Context(), command.CreateUser{User: *u}); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
}

func (h HTTPServer) GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams) {
	users, err := h.app.Queries.AllUsers.Handle(r.Context(), query.AllUsers{
		Limit:  params.Limit,
		Offset: params.Offset,
	})

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	var respUsers []User

	for _, u := range users {
		name := u.Name()
		email := u.Email()

		var role UserRole

		switch u.Role() {
		case user.RoleAdmin:
			role = UserRoleAdmin
		case user.RoleManager:
			role = UserRoleManager
		case user.RoleBasic:
			role = UserRoleBasic
		}

		respUsers = append(respUsers, User{
			Id:    u.UID(),
			Name:  &name,
			Role:  role,
			Email: &email,
		})
	}

	render.Respond(w, r, Users{Users: respUsers})
}

func NewHTTPServer(application *app.Application) HTTPServer {
	return HTTPServer{application}
}
