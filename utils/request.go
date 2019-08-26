package utils

type Request struct {
	method  string
	data    string
	headers string
	url     string
}

// Constructor
func NewRequest(m, d, h, u string) *Request {
	r := new(Request)
	r.method = m
	r.data = d
	r.headers = h
	r.url = u
	return r
}

// Method
func (r *Request) GetMethod() string {
	return r.method
}
func (r *Request) SetMethod(newMethod string) {
	r.method = newMethod
}

// Data
func (r *Request) GetData() string {
	return r.data
}
func (r *Request) SetData(newData string) {
	r.data = newData
}

// Headers
func (r *Request) GetHeaders() string {
	return r.headers
}
func (r *Request) SetHeaders(newHeaders string) {
	r.headers = newHeaders
}

// URL
func (r *Request) GetURL() string {
	return r.url
}
func (r *Request) SetURL(newURL string) {
	r.url = newURL
}
