package ports

import (
	"github.com/go-chi/render"
	"github.com/noodlensk/task-tracker/internal/common/server/httperr"
	"github.com/noodlensk/task-tracker/internal/tasks/app"
	"github.com/noodlensk/task-tracker/internal/tasks/app/command"
	"github.com/noodlensk/task-tracker/internal/tasks/app/query"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/task"
	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
	"net/http"
)

type HTTPServer struct {
	app *app.Application
}

func (h HTTPServer) GetTasks(w http.ResponseWriter, r *http.Request, params GetTasksParams) {
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

func (h HTTPServer) CreateTask(w http.ResponseWriter, r *http.Request) {
	taskFromReq := &Task{}
	if err := render.Decode(r, taskFromReq); err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	var u user.User
	// TODO get user
	t, err := task.NewTask(taskFromReq.Title, taskFromReq.Description, u)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	err = h.app.Commands.CreateTask.Handle(r.Context(), command.CreateTask{Task: *t})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HTTPServer) MarkTaskAsComplete(w http.ResponseWriter, r *http.Request) {
	taskFromReq := &TaskUpdate{}
	if err := render.Decode(r, taskFromReq); err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	var u user.User
	// TODO get user

	for _, uid := range taskFromReq.Uid {
		if err := h.app.Commands.CompleteTask.Handle(r.Context(), command.CompleteTask{TaskUID: uid, User: u}); err != nil {
			httperr.RespondWithSlugError(err, w, r)

			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HTTPServer) ReassignTasks(w http.ResponseWriter, r *http.Request) {
	var u user.User
	// TODO get user

	if err := h.app.Commands.ReAssignAllTasks.Handle(r.Context(), command.ReAssignAllTasks{User: u}); err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}
}

func NewHTTPServer(application *app.Application) HTTPServer {
	return HTTPServer{application}
}
