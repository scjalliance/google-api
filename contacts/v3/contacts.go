package contacts

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"google.golang.org/api/gensupport"
	"google.golang.org/api/googleapi"
)

const (
	// Version is the API version of the Google Contacts API
	Version = "3.0"
	// BasePath is the base URL of the Google Contacts API
	BasePath = "https://www.google.com/m8/feeds/"
)

// OAuth2 scopes used by this API.
const (
	// Manage your contacts
	ContactsScope = "https://www.googleapis.com/auth/contacts"

	// View your contacts
	ContactsReadonlyScope = "https://www.googleapis.com/auth/contacts.readonly"
)

// New returns a new service that will rely upon the provided HTTP client for
// authenticated transport.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("Could not create contacts service because a nil HTTP client was provided.")
	}
	s := &Service{client: client, BasePath: BasePath}
	s.Contacts = NewContactsService(s)
	return s, nil
}

// Service represents the root of the contacts API.
type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Contacts *ContactsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

// NewContactsService creates a new contacts service that operates via the
// given service root.
func NewContactsService(s *Service) *ContactsService {
	cs := &ContactsService{s: s}
	//cs.Connections
	return cs
}

// ContactsService provides access to contact feeds.
type ContactsService struct {
	s *Service
}

// Feed returns a contact feed for the given user.
func (cs *ContactsService) Feed(userID, projection string) *ContactFeedCall {
	c := &ContactFeedCall{s: cs.s, urlParams: make(gensupport.URLParams)}
	c.userID = userID
	c.projection = projection
	return c
}

// ContactFeedCall represents a call to the contact feed API.
type ContactFeedCall struct {
	s           *Service
	userID      string
	projection  string
	urlParams   gensupport.URLParams
	ifNoneMatch string
	ctx         context.Context
	header      http.Header
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ContactFeedCall) Header() http.Header {
	if c.header == nil {
		c.header = make(http.Header)
	}
	return c.header
}

// Query sets the contact search query to be used in the call.
func (c *ContactFeedCall) Query(q string) *ContactFeedCall {
	c.urlParams.Set("q", q)
	return c
}

// Group constrains the resulting contacts to members of the specified group.
func (c *ContactFeedCall) Group(group string) *ContactFeedCall {
	c.urlParams.Set("group", group)
	return c
}

// MaxResults specifies the maximum number of results to retrieve. The default
// is 25.
func (c *ContactFeedCall) MaxResults(max int) *ContactFeedCall {
	c.urlParams.Set("max-results", strconv.Itoa(max))
	return c
}

// Context sets the context to be used in the call. Any pending HTTP request
// created during the call will be aborted if the provided context is canceled.
func (c *ContactFeedCall) Context(ctx context.Context) *ContactFeedCall {
	c.ctx = ctx
	return c
}

func (c *ContactFeedCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	reqHeaders.Set("GData-Version", Version)
	var body io.Reader
	c.urlParams.Set("alt", alt)
	url := googleapi.ResolveRelative(c.s.BasePath, "contacts/{+userID}/{+projection}")
	url += "?" + c.urlParams.Encode()
	req, _ := http.NewRequest("GET", url, body)
	googleapi.Expand(req.URL, map[string]string{
		"userID":     c.userID,
		"projection": c.projection,
	})
	//log.Printf("URL: %v", req.URL)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx, c.s.client, req)
}

// Do performs the API call for the contacts feed and returns the response.
func (c *ContactFeedCall) Do(opts ...googleapi.CallOption) (*ContactFeedResponse, error) {
	gensupport.SetOptions(c.urlParams, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err = googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//log.Printf("Raw response: %v", string(data))
	ret := &ContactFeedResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	//if err := json.NewDecoder(res.Body).Decode(target); err != nil {
	if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}
