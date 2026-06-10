package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/badge"
)

func CommentCount(w http.ResponseWriter, req *http.Request) {
	badge.UpdateCounter()
	s := badge.ConvertToJson("Comments", badge.ChannelStats.CommentCount)
	fmt.Fprint(w, s)
}
