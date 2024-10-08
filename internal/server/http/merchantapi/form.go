package merchantapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mituan8/pay/internal/bus"
	"github.com/mituan8/pay/internal/server/http/common"
	"github.com/mituan8/pay/internal/server/http/middleware"
	"github.com/mituan8/pay/pkg/api-dashboard/v1/model"
	"github.com/pkg/errors"
)

func (h *Handler) CreateFormSubmission(c echo.Context) error {
	var req model.FormRequest
	if !common.BindAndValidateRequest(c, &req) {
		return nil
	}

	u := middleware.ResolveUser(c)
	if u == nil {
		return errors.New("unable to resolve user")
	}

	event := bus.FormSubmittedEvent{
		RequestType: *req.Topic,
		Message:     *req.Message,
		UserID:      u.ID,
	}

	if err := h.publisher.Publish(bus.TopicFormSubmissions, event); err != nil {
		return errors.Wrap(err, "unable to publish event")
	}

	return c.NoContent(http.StatusCreated)
}
