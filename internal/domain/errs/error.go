package errs

// Err

type Err string

func (e Err) Error() string {
	return string(e)
}

// ErrWithDesc

type ErrWithDesc struct {
	Err  Err
	Desc string
}

func (e ErrWithDesc) Error() string {
	return e.Err.Error() + ", desc:" + e.Desc
}

// errors

const (
	BadJson          = Err("bad_json")
	BadQueryParams   = Err("bad_query_params")
	ServiceNA        = Err("service_not_available")
	NotImplemented   = Err("not_implemented")
	NotAuthorized    = Err("not_authorized")
	PermissionDenied = Err("permission_denied")
	ObjectNotFound   = Err("object_not_found")
	BadStatusCode    = Err("bad_status_code")
	OrgBinRequired   = Err("org_bin_required")
)
