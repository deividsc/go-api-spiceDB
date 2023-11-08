package mock

import (
	"api-spiceDB/internal/ports"
	"context"
)

type CheckPermissionsMock struct {
	Response bool
}

func (c *CheckPermissionsMock) Check(_ context.Context, _ ports.CheckPermissionsRequest) (bool, error) {
	return c.Response, nil
}
