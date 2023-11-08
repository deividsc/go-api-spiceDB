package web

import (
	"api-spiceDB/internal/get_posts/application"
	"api-spiceDB/internal/ports"
	"encoding/json"
	"net/http"
)

const HEADER_USER_ID = "User-Id"

type GetPostController struct {
	handler             application.GetPostByIdHandler
	checkPermissionsSrv ports.CheckPermissions
}

func NewGetPostController(h application.GetPostByIdHandler, c ports.CheckPermissions) (*GetPostController, error) {
	return &GetPostController{
		handler:             h,
		checkPermissionsSrv: c,
	}, nil
}

func (c *GetPostController) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	if id == "" {
		http.Error(w, "id cannot be empty", http.StatusBadRequest)
		return
	}
	userId := r.Header.Get(HEADER_USER_ID)
	if userId == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	post, err := c.handler.GetPost(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	valid, err := c.checkPermissionsSrv.Check(r.Context(), ports.CheckPermissionsRequest{
		ObjectId: id,
		UserID:   userId,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !valid {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(post)
}
