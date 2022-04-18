package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (b *Bot) getAPI(method string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.token, method)
}

func (b *Bot) getFileURL(filePath string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", b.token, filePath)
}

func (b *Bot) postJSON(method string, values interface{}, rb interface{}) (*APIResponse, error) {
	url := b.getAPI(method)
	if b.Debug {
		fmt.Println("POST", url)
	}

	data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("ok")

	result := &APIResponse{Result: rb}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(result); err != nil {
		return result, nil
	}

	return result, nil
}

func parseResponse(body io.ReadCloser, r *APIResponse) error {
	dec := json.NewDecoder(body)
	if err := dec.Decode(r); err != nil {
		return err
	}

	return nil
}
