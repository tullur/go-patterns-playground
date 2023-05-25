package creational

type RequestBuilder struct {
	method      string
	headers     []string
	queryParams []string
	url         string
}

type Request struct {
	method      string
	headers     []string
	queryParams []string
	url         string
}

func (rb *RequestBuilder) SetHTTPMethod(method string) {
	rb.method = method
}

func (rb *RequestBuilder) SetHeaders(headers []string) {
	rb.headers = headers
}

func (rb *RequestBuilder) SetAuthorization() {
	rb.headers = append(rb.headers, "Authorization: Bearer mF_9.B5f-4.1JqM")
}

func (rb *RequestBuilder) SetQueryParams(queryParams []string) {
	rb.queryParams = queryParams
}

func (rb *RequestBuilder) SetURL(url string) {
	rb.url = url
}

func (rb *RequestBuilder) Build() *Request {
	return &Request{
		method:      rb.method,
		headers:     rb.headers,
		queryParams: rb.queryParams,
		url:         rb.url,
	}
}
