package healthcheck

// PingResponseContract is healthcheck response contract.
//
// @tag.name PingResponseContract
// @tag.description Ping response API contract.
type PingResponseContract struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}
