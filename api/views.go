package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/badge"
)

func ViewCount(w http.ResponseWriter, req *http.Request) {
	badge.UpdateCounter()
	s := badge.ConvertToJson("Views", badge.ChannelStats.ViewCount)
	fmt.Fprint(w, s)
}
