package http

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/noodlensk/task-tracker/internal/common/auth"
	"github.com/noodlensk/task-tracker/internal/common/server/httperr"
	"github.com/noodlensk/task-tracker/internal/tasks/app"
	"github.com/noodlensk/task-tracker/internal/tasks/app/command"
	"github.com/noodlensk/task-tracker/internal/tasks/app/query"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type Server struct {
	app *app.Application
}

func (h Server) GetTasks(w http.ResponseWriter, r *http.Request, params GetTasksParams) {
	tasks, err := h.app.Queries.AllTasks.Handle(r.Context(), query.AllTasks{
		Limit:  params.Limit,
		Offset: params.Offset,
	})

	var respTasks []Task

	for _, t := range tasks {
		assignedTo := t.AssignedTo().UID
		createdAt := t.CreatedAt()
		createdBy := t.CreatedBy().UID
		uid := t.UID()
		modifiedAt := t.UpdatedAt()

		status := NEW

		if t.Status() == task.StatusDone {
			status = DONE
		}

		respTasks = append(respTasks, Task{
			AssignedTo:  &assignedTo,
			CreatedAt:   &createdAt,
			CreatedBy:   &createdBy,
			Description: t.Description(),
			Uid:         &uid,
			ModifiedAt:  &modifiedAt,
			Status:      &status,
			Title:       t.Title(),
		})
	}

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	render.Respond(w, r, Tasks{Tasks: respTasks})
}

func (h Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	taskFromReq := &Task{}
	if err := render.Decode(r, taskFromReq); err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	u, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	t, err := task.NewTask(taskFromReq.Title, taskFromReq.Description, userToDomain(u))
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	err = h.app.Commands.CreateTask.Handle(r.Context(), command.CreateTask{Task: *t})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h Server) MarkTaskAsComplete(w http.ResponseWriter, r *http.Request) {
	taskFromReq := &TaskUpdate{}
	if err := render.Decode(r, taskFromReq); err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	u, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	for _, uid := range taskFromReq.Uid {
		if err := h.app.Commands.CompleteTask.Handle(r.Context(), command.CompleteTask{TaskUID: uid, User: userToDomain(u)}); err != nil {
			httperr.RespondWithSlugError(err, w, r)

			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Server) ReassignTasks(w http.ResponseWriter, r *http.Request) {
	u, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	if u.Role != "admin" && u.Role != "manager" {
		httperr.Unauthorized("invalid-role", nil, w, r)
		return
	}

	if err := h.app.Commands.ReAssignAllTasks.Handle(r.Context(), command.ReAssignAllTasks{User: userToDomain(u)}); err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}
}

func NewHTTPServer(application *app.Application) Server {
	return Server{application}
}

func userToDomain(u auth.User) user.User {
	return user.User{
		UID:   u.UUID,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role, // TODO: make it more secure
	}
}
