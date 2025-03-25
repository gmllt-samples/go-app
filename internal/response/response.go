package response

type JSON struct {
	Status   int    `json:"status"`
	URI      string `json:"uri"`
	Duration string `json:"duration"`
	Body     string `json:"body,omitempty"`
}
