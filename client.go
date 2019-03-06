package fcm

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	// DefaultEndpoint contains endpoint URL of FCM service.
	DefaultEndpoint = "https://fcm.googleapis.com/fcm/send"

	// DefaultTimeout duration in second
	DefaultTimeout time.Duration = 30 * time.Second
)

var (
	// ErrInvalidAPIKey occurs if API key is not set.
	ErrInvalidAPIKey = errors.New("client API Key is invalid")
)

// Client abstracts the interaction between the application server and the
// FCM server via HTTP protocol. The developer must obtain an API key from the
// Google APIs Console page and pass it to the `Client` so that it can
// perform authorized requests on the application server's behalf.
// To send a message to one or more devices use the Client's Send.
//
// If the `HTTP` field is nil, a zeroed http.Client will be allocated and used
// to send messages.
type Client struct {
	apiKey   string
	client   *fasthttp.Client
	endpoint string
	timeout  time.Duration
}

// NewClient creates new Firebase Cloud Messaging Client based on API key and
// with default endpoint and http client.
func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, ErrInvalidAPIKey
	}

	httpclient := fasthttp.Client{}
	if proxy := os.Getenv("HTTPS_PROXY"); proxy != "" {
		httpclient.Dial = FasthttpHTTPDialer(proxy)
	}

	c := &Client{
		apiKey:   apiKey,
		endpoint: DefaultEndpoint,
		client:   &httpclient,
		timeout:  DefaultTimeout,
	}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// Send sends a message to the FCM server without retrying in case of service
// unavailability. A non-nil error is returned if a non-recoverable error
// occurs (i.e. if the response status is not "200 OK").
func (c *Client) Send(msg *Message) (*Response, []byte, error) {
	// validate
	if err := msg.Validate(); err != nil {
		return nil, nil, err
	}

	// marshal message
	data, err := msg.MarshalJSON()
	if err != nil {
		return nil, nil, err
	}

	return c.send(data)
}

// SendWithRetry sends a message to the FCM server with defined number of
// retrying in case of temporary error.
func (c *Client) SendWithRetry(msg *Message, retryAttempts int) (*Response, []byte, error) {
	// validate
	if err := msg.Validate(); err != nil {
		return nil, nil, err
	}
	// marshal message
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, nil, err
	}

	resp := new(Response)
	var body []byte
	err = retry(func() error {
		var er error
		resp, body, er = c.send(data)
		return er
	}, retryAttempts)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, nil
}

// send sends a request.
func (c *Client) send(data []byte) (*Response, []byte, error) {
	// create request
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(c.endpoint)
	req.Header.SetMethod("POST")
	req.SetBody(data)

	req.Header.Set("Authorization", fmt.Sprintf("key=%s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	// create response
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := c.client.DoTimeout(req, resp, c.timeout)
	if err != nil {
		return nil, resp.Body(), connectionError(err.Error())
	}

	// check response status
	if resp.StatusCode() != fasthttp.StatusOK {
		if resp.StatusCode() >= fasthttp.StatusInternalServerError {
			return nil, resp.Body(), serverError(fmt.Sprintf("%d error: %s", resp.StatusCode(), fasthttp.StatusMessage(resp.StatusCode())))
		}
		return nil, resp.Body(), fmt.Errorf("%d error: %s", resp.StatusCode(), fasthttp.StatusMessage(resp.StatusCode()))
	}

	// build return
	response := new(Response)

	if err := response.UnmarshalJSON(resp.Body()); err != nil {
		return nil, nil, err
	}

	return response, resp.Body(), nil
}
