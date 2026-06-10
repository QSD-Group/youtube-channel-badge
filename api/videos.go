package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/badge"
)

func VideoCount(w http.ResponseWriter, req *http.Request) {
	badge.UpdateCounter()
	s := badge.ConvertToJson("Videos", badge.ChannelStats.VideoCount)
	fmt.Fprint(w, s)
}
