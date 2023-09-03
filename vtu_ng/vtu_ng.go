package vtung

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/mitchellh/mapstructure"
)

const (
	// library version
	version = "0.1.0"

	// defaultHTTPTimeout is the default timeout on the http client
	defaultHTTPTimeout = 60 * time.Second

	// base URL for all Paystack API requests
	baseURL = " https://vtu.ng" // ""

	// User agent used when communicating with the Paystack API.
	// userAgent = "paystack-go/" + version
	userAgent = "Mozilla/5.0 (Unknown; Linux) AppleWebKit/538.1 (KHTML, like Gecko) Chrome/v1.0.0 Safari/538.1"
)

type service struct {
	client *Client
}

type Client struct {
	common service      // Reuse a single struct instead of allocating one for each service on the heap.
	client *http.Client // HTTP client used to communicate with the API.

	baseURL *url.URL
	logger  Logger

	Airtime              *AirtimeService
	Balance              *BalanceService
	CableTV              *CableTVService
	CustomerVerification *CustomerVerificationService
	DataBundle           *DataBundleService
	Electricity          *ElectricityService
	LoggingEnabled       bool
	Log                  Logger
	Username             string
	Password             string
}

// Logger interface for custom loggers
type Logger interface {
	Printf(format string, v ...interface{})
}

// Metadata is an key-value pairs added to Paystack API requests
type Metadata map[string]interface{}

// Response represents arbitrary response data
type Response map[string]interface{}

// RequestValues aliased to url.Values as a workaround
type RequestValues url.Values

// MarshalJSON to handle custom JSON decoding for RequestValues
func (v RequestValues) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 3)
	for k, val := range v {
		m[k] = val[0]
	}
	return json.Marshal(m)
}

// ListMeta is pagination metadata for paginated responses from the Paystack API
type ListMeta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

// NewClient creates a new Paystack API client with the given API key
// and HTTP client, allowing overriding of the HTTP client to use.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultHTTPTimeout}
	}

	u, _ := url.Parse(baseURL)
	c := &Client{
		client:         httpClient,
		baseURL:        u,
		LoggingEnabled: true,
		Log:            log.New(os.Stderr, "", log.LstdFlags),
		Username:       "GabrielAchumba",
		Password:       "gab*012023",
	}

	c.common.client = c

	c.Airtime = (*AirtimeService)(&c.common)
	c.Balance = (*BalanceService)(&c.common)
	c.CableTV = (*CableTVService)(&c.common)
	c.CustomerVerification = (*CustomerVerificationService)(&c.common)
	c.DataBundle = (*DataBundleService)(&c.common)
	c.Electricity = (*ElectricityService)(&c.common)

	return c
}

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(data)
	return err
}

// Call actually does the HTTP request to Paystack API
func (c *Client) Call(method, path string, body interface{}) (interface{}, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	u, _ := c.baseURL.Parse(path)
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		if c.LoggingEnabled {
			c.Log.Printf("Cannot create Paystack request: %v\n", err)
		}
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	//req.Header.Set("Authorization", "Bearer "+c.key)
	//req.Header.Set("User-Agent", userAgent)

	req.Header.Set("api-key", "25393d574b0e6a0f9ee1ae3210ff3d9c")
	req.Header.Set("public-key", "PK_8107679f265e03550e1832566403c38758c089af353")
	req.Header.Set("secret-key", "SK_4588ef499f516f135e0e9b3c2af6b35ad6a5fc15ebe")

	if c.LoggingEnabled {
		c.Log.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
		c.Log.Printf("POST request data %v\n", buf)
	}

	start := time.Now()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.LoggingEnabled {
		c.Log.Printf("Completed in %v\n", time.Since(start))
	}

	defer resp.Body.Close()
	return c.decodeResponse(resp)
}

// decodeResponse decodes the JSON response from the Twitter API.
// The actual response will be written to the `v` parameter
func (c *Client) decodeResponse(httpResp *http.Response) (interface{}, error) {
	var resp Response
	respBody, _ := ioutil.ReadAll(httpResp.Body)
	json.Unmarshal(respBody, &resp)
	//return mapstruct(resp, v)
	//c.Log.Printf("HTTP Response:", resp)

	/* if status, _ := resp["status"].(bool); !status || httpResp.StatusCode >= 400 {
		if c.LoggingEnabled {
			c.Log.Printf("Paystack error: %+v", err)
			c.Log.Printf("HTTP Response: %+v", resp)
		}
		return newAPIError(httpResp)
	}

	if c.LoggingEnabled {
		c.Log.Printf("Paystack response: %v\n", resp)
	}


	if data, ok := resp["data"]; ok {
		switch t := resp["data"].(type) {
		case map[string]interface{}:
			return mapstruct(data, v)
		default:
			_ = t
			return mapstruct(resp, v)
		}
	} */
	// if response data does not contain data key, map entire response to v
	//return mapstruct(resp, v)
	//v = resp
	return resp, nil
}
