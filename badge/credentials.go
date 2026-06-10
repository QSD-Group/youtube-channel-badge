package badge

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Info Env

type Env struct {
	CacheTime int
	ApiKey    string
	ChannelID string
}

func init() {
	Info.CacheTime = 300
	if s := strings.TrimSpace(os.Getenv("CACHE_TIME")); s != "" {
		if n, err := strconv.Atoi(s); err == nil && n > 0 {
			Info.CacheTime = n
		}
	}

	// Trim whitespace so a stray space/newline pasted into a dashboard env
	// var doesn't silently break the API key or channel-id match.
	Info.ApiKey = strings.TrimSpace(os.Getenv("API_KEY"))
	Info.ChannelID = strings.TrimSpace(os.Getenv("CHANNEL_ID"))
}

// Validate reports any missing required configuration. It returns an error
// instead of panicking so the handlers can surface the reason on the badge
// rather than crashing the function with an opaque 500.
func (e Env) Validate() error {
	var missing []string
	if e.ApiKey == "" {
		missing = append(missing, "API_KEY")
	}
	if e.ChannelID == "" {
		missing = append(missing, "CHANNEL_ID")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing env var(s): %s", strings.Join(missing, ", "))
	}
	return nil
}
