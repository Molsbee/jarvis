package hacker_news

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://hacker-news.firebaseio.com"

func GetNewStories(storyCount int) ([]NewsItem, error) {
	endpoint := fmt.Sprintf("%s/v0/newstories.json", baseURL)
	resp, err := http.DefaultClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to get response from top stories (error - %s)", err)
	}

	return convertToItem(resp, storyCount)
}

func GetTopStories(storyCount int) ([]NewsItem, error) {
	endpoint := fmt.Sprintf("%s/v0/topstories.json", baseURL)
	resp, err := http.DefaultClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to get response from top stories (error - %s)", err)
	}

	return convertToItem(resp, storyCount)
}

func GetBestStories(storyCount int) ([]NewsItem, error) {
	endpoint := fmt.Sprintf("%s/v0/beststories.json", baseURL)
	resp, err := http.DefaultClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to get response from top stories (error - %s)", err)
	}

	return convertToItem(resp, storyCount)
}

func convertToItem(response *http.Response, storyCount int) ([]NewsItem, error) {
	defer response.Body.Close()

	var ids []int
	if err := json.NewDecoder(response.Body).Decode(&ids); err != nil {
		return nil, fmt.Errorf("failed to read response body of initial list request")
	}

	var items []NewsItem
	for _, id := range ids {
		item, err := getItem(id)
		if err == nil {
			items = append(items, item)
			if len(items) == storyCount {
				break
			}
		}
	}

	return items, nil
}

func getItem(id int) (item NewsItem, err error) {
	endpoint := fmt.Sprintf("%s/v0/item/%d.json", baseURL, id)
	resp, getErr := http.DefaultClient.Get(endpoint)
	if getErr != nil {
		err = fmt.Errorf("failed to get article %d details (error - %s)", id, err)
		return
	}
	defer resp.Body.Close()

	jErr := json.NewDecoder(resp.Body).Decode(&item)
	if jErr != nil {
		err = fmt.Errorf("failed to convert article json (error - %s)", err)
	}

	return item, nil
}
