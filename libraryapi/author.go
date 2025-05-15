package libraryapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func CallAuthorOpenLibraryInformation(input string) (*Book, error) {
	query := url.QueryEscape(input)
	baseURL := fmt.Sprintf("https://openlibrary.org/search.json?author=%s&sort=new", query)

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Could not send get req to open library general api: %v", err)
	}
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Open Library API returned non-200 status: %s", resp.Status)
	}

	var result OpenLibraryGeneralQueryAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Failed to decode JSON response from Open Library API: %w", err)
	}

	if len(result.Docs) == 0 {
		return nil, fmt.Errorf("No results found for query: %q", input)
	}

	return &result.Docs[0], nil
}