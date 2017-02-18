package resource

// Response is a general HTTP response holder
type Response struct {
	Res  interface{}
	Code int
}

// Metadata is used to set extra object metadata (like pagination)
func (r Response) Metadata() map[string]interface{} {
	return map[string]interface{}{}
}

// Result is used to get the final resulting JSON
func (r Response) Result() interface{} {
	return r.Res
}

// StatusCode is used to get the final resulting HTTP status code
func (r Response) StatusCode() int {
	return r.Code
}
