// Code generated from specification version 7.4.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

func newSnapshotCleanupRepositoryFunc(t Transport) SnapshotCleanupRepository {
	return func(repository string, o ...func(*SnapshotCleanupRepositoryRequest)) (*Response, error) {
		var r = SnapshotCleanupRepositoryRequest{Repository: repository}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SnapshotCleanupRepository removes stale data from repository.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html.
//
type SnapshotCleanupRepository func(repository string, o ...func(*SnapshotCleanupRepositoryRequest)) (*Response, error)

// SnapshotCleanupRepositoryRequest configures the Snapshot Cleanup Repository API request.
//
type SnapshotCleanupRepositoryRequest struct {
	Body io.Reader

	Repository string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SnapshotCleanupRepositoryRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_snapshot") + 1 + len(r.Repository) + 1 + len("_cleanup"))
	path.WriteString("/")
	path.WriteString("_snapshot")
	path.WriteString("/")
	path.WriteString(r.Repository)
	path.WriteString("/")
	path.WriteString("_cleanup")

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), r.Body)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f SnapshotCleanupRepository) WithContext(v context.Context) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.ctx = v
	}
}

// WithBody - .
//
func (f SnapshotCleanupRepository) WithBody(v io.Reader) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.Body = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
//
func (f SnapshotCleanupRepository) WithMasterTimeout(v time.Duration) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
//
func (f SnapshotCleanupRepository) WithTimeout(v time.Duration) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SnapshotCleanupRepository) WithPretty() func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SnapshotCleanupRepository) WithHuman() func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SnapshotCleanupRepository) WithErrorTrace() func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SnapshotCleanupRepository) WithFilterPath(v ...string) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f SnapshotCleanupRepository) WithHeader(h map[string]string) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}
