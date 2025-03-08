package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Notification struct {
	UpdatedAt  string `json:"updated_at,omitempty"`
	Reason     string `json:"reason"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	Subject struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	} `json:"subject"`
}

func getNotifications(ctx context.Context, token string, since *time.Time) ([]Notification, error) {
	client := &http.Client{
		Timeout: time.Second * 4,
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		getEndpoint(since),
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("non-success response code: %d", res.StatusCode)
	}

	raw, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	response := []Notification{}
	err = json.Unmarshal(raw, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return response, nil
}

func getEndpoint(since *time.Time) string {
	endpoint := &url.URL{
		Scheme: "https",
		Host:   "api.github.com",
		Path:   "/notifications",
	}

	query := endpoint.Query()

	if since != nil {
		query.Add("since", since.Format(time.RFC3339))
	}

	query.Add("all", "false")

	endpoint.RawQuery = query.Encode()

	return endpoint.String()
}
