package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/badge"
)

func CommentCount(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := badge.UpdateCounter(); err != nil {
		fmt.Fprint(w, badge.ErrorToJson("Comments", err.Error()))
		return
	}
	fmt.Fprint(w, badge.ConvertToJson("Comments", badge.ChannelStats.CommentCount))
}
