package secondary

import (
	"api-spiceDB/internal/ports"
	"context"
)

type CheckPermissionsSpiceDB struct {
}

func (db *CheckPermissionsSpiceDB) Check(ctx context.Context, p ports.CheckPermissionsRequest) (bool, error) {
	return true, nil
}
