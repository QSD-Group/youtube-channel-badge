package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/badge"
)

func ViewCount(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := badge.UpdateCounter(); err != nil {
		fmt.Fprint(w, badge.ErrorToJson("Views", err.Error()))
		return
	}
	fmt.Fprint(w, badge.ConvertToJson("Views", badge.ChannelStats.ViewCount))
}
