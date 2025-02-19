package bmm

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log/slog"
	"net/http"
	"net/url"
)

type APIClient struct {
	httpClient *resty.Client
	token      *Token
	logger     *slog.Logger
}

// NewApiClient creates a new BMM API client, using the provided token
//
// It sets the language to norwegian. There is currently no way to change this. PRs welcome :).
func NewApiClient(baseURL string, token *Token) *APIClient {
	client := &APIClient{}
	client.httpClient = resty.New()
	client.httpClient.BaseURL = baseURL
	client.httpClient.SetHeader("Accept-Language", "nb")
	client.token = token

	client.logger = slog.Default().With("component", "bmm")

	return client
}

func (c *APIClient) SetLogger(logger *slog.Logger) *APIClient {
	c.logger = logger
	return c
}

func (c *APIClient) SetBaseURL(baseURL string) *APIClient {
	c.httpClient.BaseURL = baseURL
	return c
}

func parseResponse[T any](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}

func (c *APIClient) makeRequest(method, path string, body any) ([]byte, error) {
	token, err := c.token.GetAccessToken()

	if err != nil {
		return nil, err
	}

	req := c.httpClient.R().
		SetAuthToken(token).
		SetBody(body)

	var res *resty.Response

	if method == "GET" {
		res, err = req.Get(path)
	} else if method == "POST" {
		res, err = req.Post(path)
	} else if method == "PUT" {
		res, err = req.Put(path)
	} else if method == "DELETE" {
		res, err = req.Delete(path)
	}

	if err != nil {
		return nil, err
	}

	if res == nil {
		slog.Error("request failed, response is nil")
		return nil, fmt.Errorf("request failed, response is nil")
	}

	if res.StatusCode() != http.StatusOK {
		slog.Error("request failed with status code", "code", res.StatusCode(), "response", res.Body())
		return nil, fmt.Errorf("request failed with status code %d", res.StatusCode())
	}

	return res.Body(), nil
}

func (c *APIClient) GetYears() ([]Year, error) {
	data, err := c.makeRequest("GET", "/facets/album_published/years", nil)

	if err != nil {
		return nil, err
	}

	return parseResponse[[]Year](data)
}

func (c *APIClient) GetAlbums(year int) ([]Item, error) {
	data, err := c.makeRequest("GET", fmt.Sprintf("/album/published/%d/", year), nil)

	if err != nil {
		return nil, err
	}

	return parseResponse[[]Item](data)
}

func (c *APIClient) GetAlbumTracks(albumId string) ([]Item, error) {
	data, err := c.makeRequest("GET", fmt.Sprintf("/album/%s", albumId), nil)

	if err != nil {
		return nil, err
	}

	return parseResponse[[]Item](data)
}

func (c *APIClient) GetPodcastTracks(podcastTag string, limit int) ([]Item, error) {
	data, err := c.makeRequest("GET", fmt.Sprintf("/track?tags=%s&size=%d&unpublished=show", url.QueryEscape(podcastTag), limit), nil)

	if err != nil {
		return nil, err
	}

	return parseResponse[[]Item](data)
}

func (c *APIClient) GetLanguages() ([]Overview, error) {
	data, err := c.makeRequest("GET", "/languages", nil)

	if err != nil {
		return nil, err
	}

	return parseResponse[[]Overview](data)
}
