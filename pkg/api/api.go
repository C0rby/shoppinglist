package api

import (
	"errors"
	"net/http"

	"github.com/c0rby/shoppinglist/pkg/model"
	"github.com/c0rby/shoppinglist/pkg/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

const (
	_paramListId  = "listId"
	_paramEntryId = "entryId"
)

type api struct {
	service service.Service
}

func Handler(service service.Service) http.Handler {
	api := api{service: service}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.URLFormat)
	r.Route("/v1", func(r chi.Router) {
		r.Route("/shoppinglists", func(r chi.Router) {
			r.Get("/", api.ListShoppingLists)
			r.Post("/", api.CreateShoppingList)
			r.Route("/{listId}", func(r chi.Router) {
				r.Get("/", api.GetShoppingList)
				r.Delete("/", api.DeleteShoppingList)
				r.Route("/entries", func(r chi.Router) {
					r.Get("/", api.ListShoppingListEntries)
					r.Post("/", api.CreateListEntry)
					r.Put("/{entryId}", api.UpdateListEntry)
				})
			})
		})
	})

	return r
}

func (a api) ListShoppingLists(w http.ResponseWriter, r *http.Request) {
	lists, _ := a.service.GetShoppingLists()
	if err := render.RenderList(w, r, NewShoppingListsResponse(lists)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (a api) GetShoppingList(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, _paramListId)
	list, _ := a.service.GetShoppingList(id)
	if err := render.Render(w, r, NewShoppingListResponse(list)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (a api) ListShoppingListEntries(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, _paramListId)
	entries, _ := a.service.GetShoppingListEntries(id)

	if err := render.RenderList(w, r, NewEntryListResponse(entries)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (a api) CreateShoppingList(w http.ResponseWriter, r *http.Request) {
	data := &ShoppingListRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	shoppingList := data.ShoppingList

	created, err := a.service.CreateShoppingList(model.ShoppingList{Name: shoppingList.Name})
	if err != nil {
		render.Render(w, r, ErrInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewShoppingListResponse(created))
}

func (a api) DeleteShoppingList(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, _paramListId)
	if err := a.service.DeleteShoppingList(id); err != nil {
		render.Render(w, r, ErrInternalServerError(err))
		return
	}
	render.NoContent(w, r)
}

func (a api) CreateListEntry(w http.ResponseWriter, r *http.Request) {
	listID := chi.URLParam(r, _paramListId)
	data := &ShoppingListEntryRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	entry := data.ShoppingListEntry
	created, err := a.service.CreateShoppingListEntry(listID, model.Entry{Name: entry.Name})
	if err != nil {
		render.Render(w, r, ErrInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewEntryResponse(created))
}

func (a api) UpdateListEntry(w http.ResponseWriter, r *http.Request) {
	entryID := chi.URLParam(r, _paramEntryId)

	data := &ShoppingListEntryRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	entry := data.ShoppingListEntry
	updated, err := a.service.UpdateShoppingListEntry(model.Entry{ID: entryID, Name: entry.Name, Buy: entry.Buy})
	if err != nil {
		render.Render(w, r, ErrInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewEntryResponse(updated))
}

type ShoppingList struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ShoppingListRequest struct {
	*ShoppingList
}

func (s *ShoppingListRequest) Bind(r *http.Request) error {
	if s.ShoppingList == nil {
		return errors.New("missing required ShoppingList fields")
	}

	// The requests shouldn't contain Ids
	s.ShoppingList.ID = ""
	return nil
}

type ShoppingListEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Buy  bool   `json:"buy"`
}

type ShoppingListResponse struct {
	*ShoppingList
}

func (sr *ShoppingListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ShoppingListEntryRequest struct {
	*ShoppingListEntry
}

func (s *ShoppingListEntryRequest) Bind(r *http.Request) error {
	if s.ShoppingListEntry == nil {
		return errors.New("missing required ShoppingListEntry fields")
	}

	// The requests shouldn't contain Ids
	s.ShoppingListEntry.ID = ""
	return nil
}

type ShoppingListEntryResponse struct {
	*ShoppingListEntry
}

func (sr *ShoppingListEntryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewEntryResponse(entry model.Entry) *ShoppingListEntryResponse {
	return &ShoppingListEntryResponse{
		ShoppingListEntry: &ShoppingListEntry{
			ID:   entry.ID,
			Name: entry.Name,
			Buy:  entry.Buy,
		},
	}
}

func NewEntryListResponse(entries []model.Entry) []render.Renderer {
	response := make([]render.Renderer, 0, len(entries))
	for _, e := range entries {
		response = append(response, NewEntryResponse(e))
	}
	return response
}

func NewShoppingListResponse(list model.ShoppingList) *ShoppingListResponse {
	return &ShoppingListResponse{
		ShoppingList: &ShoppingList{
			ID:   list.ID,
			Name: list.Name,
		},
	}
}

func NewShoppingListsResponse(lists []model.ShoppingList) []render.Renderer {
	response := make([]render.Renderer, 0, len(lists))
	for _, l := range lists {
		response = append(response, NewShoppingListResponse(l))
	}
	return response
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrInternalServerError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal Server error.",
		ErrorText:      err.Error(),
	}
}
