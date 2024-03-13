// Code generated by ogen, DO NOT EDIT.

package api

// GetAllDrivesSmartInfoBadRequest is response for GetAllDrivesSmartInfo operation.
type GetAllDrivesSmartInfoBadRequest struct{}

func (*GetAllDrivesSmartInfoBadRequest) getAllDrivesSmartInfoRes() {}

// GetAllDrivesSmartInfoInternalServerError is response for GetAllDrivesSmartInfo operation.
type GetAllDrivesSmartInfoInternalServerError struct{}

func (*GetAllDrivesSmartInfoInternalServerError) getAllDrivesSmartInfoRes() {}

// GetAllDrivesSmartInfoNotFound is response for GetAllDrivesSmartInfo operation.
type GetAllDrivesSmartInfoNotFound struct{}

func (*GetAllDrivesSmartInfoNotFound) getAllDrivesSmartInfoRes() {}

// GetDriveSmartInfoBadRequest is response for GetDriveSmartInfo operation.
type GetDriveSmartInfoBadRequest struct{}

func (*GetDriveSmartInfoBadRequest) getDriveSmartInfoRes() {}

// GetDriveSmartInfoInternalServerError is response for GetDriveSmartInfo operation.
type GetDriveSmartInfoInternalServerError struct{}

func (*GetDriveSmartInfoInternalServerError) getDriveSmartInfoRes() {}

// GetDriveSmartInfoNotFound is response for GetDriveSmartInfo operation.
type GetDriveSmartInfoNotFound struct{}

func (*GetDriveSmartInfoNotFound) getDriveSmartInfoRes() {}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/SmartMetrics
type SmartMetrics struct {
	// Smart info Json.
	SmartInfo OptString `json:"smartInfo"`
}

// GetSmartInfo returns the value of SmartInfo.
func (s *SmartMetrics) GetSmartInfo() OptString {
	return s.SmartInfo
}

// SetSmartInfo sets the value of SmartInfo.
func (s *SmartMetrics) SetSmartInfo(val OptString) {
	s.SmartInfo = val
}

func (*SmartMetrics) getAllDrivesSmartInfoRes() {}
func (*SmartMetrics) getDriveSmartInfoRes()     {}