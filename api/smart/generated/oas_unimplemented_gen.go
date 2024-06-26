// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// GetAllDrivesSmartInfo implements get-all-drives-smart-info operation.
//
// Retrieve all discovered disks information/metrics.
//
// GET /smart/api/v1/drives
func (UnimplementedHandler) GetAllDrivesSmartInfo(ctx context.Context) (r GetAllDrivesSmartInfoRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetDriveSmartInfo implements get-drive-smart-info operation.
//
// Retrieve the disk information/metrics with the matching serial number.
//
// GET /smart/api/v1/drive/{serialNumber}
func (UnimplementedHandler) GetDriveSmartInfo(ctx context.Context, params GetDriveSmartInfoParams) (r GetDriveSmartInfoRes, _ error) {
	return r, ht.ErrNotImplemented
}
