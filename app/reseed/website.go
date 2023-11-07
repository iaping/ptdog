package reseed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	headers = map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	httpc = &http.Client{
		Timeout: time.Minute,
	}
)

type Website struct {
	name     string
	domain   string
	api      string
	passkey  string
	download string
	limit    int
}

func (w *Website) String() string {
	return fmt.Sprintf("%s (%s)", w.name, w.domain)
}

func (w *Website) Query(hashes []string) (map[string]int, error) {
	data, err := w.do(hashes)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data map[string]int `json:"data"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		result.Data = nil
	}

	if result.Data != nil {
		return result.Data, nil
	}
	return map[string]int{}, nil
}

func (w *Website) do(hashes []string) ([]byte, error) {
	params, err := w.params(hashes)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", w.api, bytes.NewReader(params))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := httpc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func (w *Website) params(hashes []string) ([]byte, error) {
	params := map[string]interface{}{
		"passkey":     w.passkey,
		"pieces_hash": hashes,
	}
	return json.Marshal(params)
}

func (w *Website) FormatDownload(id int) string {
	if w.download != "" {
		url := strings.Replace(w.download, "{id}", strconv.Itoa(id), 1)
		return strings.Replace(url, "{passkey}", w.passkey, 1)
	}
	return fmt.Sprintf("%s/download.php?id=%d&passkey=%s", w.domain, id, w.passkey)
}
