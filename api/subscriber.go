package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/badge"
)

func SubscriberCount(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := badge.UpdateCounter(); err != nil {
		fmt.Fprint(w, badge.ErrorToJson("Subscribe", err.Error()))
		return
	}
	fmt.Fprint(w, badge.ConvertToJson("Subscribe", badge.ChannelStats.SubscriberCount))
}
