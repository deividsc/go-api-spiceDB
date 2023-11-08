package web

import (
	"api-spiceDB/internal/adapters/secondary/mock"
	"api-spiceDB/internal/core/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockGetPostHandler struct {
	post          domain.Post
	responseError error
}

func (h *MockGetPostHandler) GetPost(ctx context.Context, id string) (domain.Post, error) {
	return h.post, h.responseError
}
func TestGetPostController_GetById(t *testing.T) {
	t.Run("Should return a post", func(t *testing.T) {
		service := mock.CheckPermissionsMock{Response: true}
		post := domain.Post{
			Id:   "Id1",
			Body: "Test post",
			Date: time.Now(),
		}
		handler := MockGetPostHandler{post: post}

		controller, _ := NewGetPostController(&handler, &service)

		req := httptest.NewRequest("GET", "/"+post.Id, nil)
		req.Header.Set(HEADER_USER_ID, "testUserID")

		w := httptest.NewRecorder()

		controller.GetById(w, req)

		assert.Equal(t, w.Code, 200)

		encoded, err := json.Marshal(post)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, string(encoded)+"\n", w.Body.String())
	})
}

func TestGetPostController_GetById_Error(t *testing.T) {
	testCases := []struct {
		desc                string
		postId              string
		userId              string
		permissionsResponse bool
		handlerError        error
		statusCode          int
		wantError           error
	}{
		{
			desc:       "Empty Id should return an error",
			postId:     "",
			statusCode: http.StatusBadRequest,
			wantError:  errors.New("id cannot be empty"),
		},
		{
			desc:       "Empty UserId should return an error",
			postId:     "testId",
			statusCode: http.StatusForbidden,
			wantError:  errors.New("Forbidden"),
		},
		{
			desc:                "If check permisssions return false should return an error",
			postId:              "testId",
			userId:              "testUserId",
			permissionsResponse: false,
			statusCode:          http.StatusForbidden,
			wantError:           errors.New("Forbidden"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := mock.CheckPermissionsMock{Response: tC.permissionsResponse}

			handler := MockGetPostHandler{responseError: tC.handlerError}

			controller, _ := NewGetPostController(&handler, &service)

			req := httptest.NewRequest("GET", "/"+tC.postId, nil)

			req.Header.Set(HEADER_USER_ID, tC.userId)

			w := httptest.NewRecorder()

			controller.GetById(w, req)

			assert.Equal(t, tC.statusCode, w.Code)

			assert.Equal(t, fmt.Sprint(tC.wantError), strings.ReplaceAll(w.Body.String(), "\n", ""))
		})

	}
}
