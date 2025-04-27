package jsonapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type jsonApiClient struct {
	BaseURL string
	Client  *http.Client
}

func NewJsonApiClient(baseURL string) *jsonApiClient {
	return &jsonApiClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *jsonApiClient) doRequest(ctx context.Context, method, path string, body interface{}, out interface{}) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)

	// リクエストボディを作成
	var reqBody io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reqBody = bytes.NewBuffer(jsonBytes)
	}

	// リクエストを作成
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	slog.Info("jsonapi: request", "method", method, "baseURL", c.BaseURL, "path", path, "body", body)

	// リクエスト実行
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	slog.Info("jsonapi: response", "method", method, "baseURL", c.BaseURL, "path", path, "status", resp.StatusCode)

	if resp.StatusCode >= 400 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(bodyBytes),
		}
	}

	if out == nil {
		return nil
	}

	// 読み取ったbodyからデコード
	if err := json.Unmarshal(bodyBytes, out); err != nil {
		return err
	}

	return nil
}
