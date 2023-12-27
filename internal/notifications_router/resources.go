package notifications_router

const (
	registerServiceType = "notificator-service"
)

type (
	registerServiceRequest struct {
		Data registerServiceData `json:"data"`
	}
	registerServiceData struct {
		Type       string                    `json:"type"`
		Attributes registerServiceAttributes `json:"attributes"`
	}
	registerServiceAttributes struct {
		Endpoint string   `json:"endpoint"`
		Channels []string `json:"channels"`
	}
)
