package notifications_router

const (
	registerServiceType = "notificator-service"
)

type (
	registerServiceRequest struct {
		Data registerServiceData
	}
	registerServiceData struct {
		Type       string
		Attributes registerServiceAttributes
	}
	registerServiceAttributes struct {
		Endpoint string   `json:"endpoint"`
		Channels []string `json:"channels"`
	}
)
