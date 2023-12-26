package requests

import (
	"encoding/json"
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/apodeixis/mailjet-svc/resources"

	"github.com/pkg/errors"
)

func NewCreateNotification(r *http.Request) (*resources.CreateNotificationRequest, error) {
	request := new(resources.CreateNotificationRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to decode request body"),
		}
	}
	return request, nil
}
