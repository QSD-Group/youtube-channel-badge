package badge

import (
	"encoding/json"
)

type EndpointResponse struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color,omitempty"`
	NamedLogo     string `json:"namedLogo,omitempty"`
	LogoColor     string `json:"logoColor,omitempty"`
	IsError       bool   `json:"isError,omitempty"`
}

func ConvertToJson(l, m string) (s string) {
	er := EndpointResponse{
		SchemaVersion: 1,
		Label:         l,
		Message:       m,
		Color:         "Red",
		NamedLogo:     "Youtube",
		LogoColor:     "Red",
	}

	bs, _ := json.MarshalIndent(er, "", "	")
	s = string(bs)
	return
}

// ErrorToJson renders a shields.io "error" badge carrying msg, so a
// misconfiguration shows up on the badge (and in a direct curl) instead of
// crashing the function with an opaque 500.
func ErrorToJson(l, msg string) (s string) {
	er := EndpointResponse{
		SchemaVersion: 1,
		Label:         l,
		Message:       msg,
		Color:         "red",
		IsError:       true,
	}

	bs, _ := json.MarshalIndent(er, "", "	")
	s = string(bs)
	return
}
