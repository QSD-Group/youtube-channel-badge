package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/badge"
)

func SubscriberCount(w http.ResponseWriter, req *http.Request) {
	badge.UpdateCounter()
	s := badge.ConvertToJson("Subscribe", badge.ChannelStats.SubscriberCount)
	fmt.Fprint(w, s)
}
