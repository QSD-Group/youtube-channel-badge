package badge

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ChannelListResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID         string     `json:"id"`
	Statistics Statistics `json:"statistics"`
}

type Statistics struct {
	ViewCount       string `json:"viewCount"`
	CommentCount    string `json:"commentCount"`
	SubscriberCount string `json:"subscriberCount"`
	VideoCount      string `json:"videoCount"`
	Time            time.Time
}

var ChannelStats Statistics

// UpdateCounter refreshes ChannelStats from the YouTube API when the cached
// value is stale. It returns an error (rather than panicking) so callers can
// render the reason on the badge.
func UpdateCounter() error {
	if err := Info.Validate(); err != nil {
		return err
	}

	now := time.Now()
	if !now.After(ChannelStats.Time.Add(time.Second * time.Duration(Info.CacheTime))) {
		return nil // cached value is still fresh
	}

	reqString := fmt.Sprintf("https://www.googleapis.com/youtube/v3/channels?part=statistics&id=%s&key=%s", Info.ChannelID, Info.ApiKey)
	resp, err := http.Get(reqString)
	if err != nil {
		return fmt.Errorf("YouTube API request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading YouTube API response failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		// Surface Google's own error message (e.g. referrer-restricted key,
		// API not enabled, quota exceeded) instead of hiding it.
		return fmt.Errorf("YouTube API %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}

	var r ChannelListResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return fmt.Errorf("parsing YouTube API response failed: %w", err)
	}

	var channel Item
	for _, v := range r.Items {
		if v.ID == Info.ChannelID {
			channel = v
			break
		}
	}
	if channel.ID == "" {
		return fmt.Errorf("no channel found with id %q", Info.ChannelID)
	}

	channel.Statistics.Time = time.Now()
	ChannelStats = channel.Statistics
	return nil
}
