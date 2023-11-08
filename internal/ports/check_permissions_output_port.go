package ports

import "context"

type CheckPermissions interface {
	Check(ctx context.Context, p CheckPermissionsRequest) (bool, error)
}

type CheckPermissionsRequest struct {
	ObjectId string
	UserID   string
}
