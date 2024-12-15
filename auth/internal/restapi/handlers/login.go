package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/impl"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/models"
	"github.com/h4x4d/go_hsse_hotels/auth/internal/restapi/operations"
	"log/slog"
)

func (h *Handler) LoginHandler(api operations.PostLoginParams) middleware.Responder {
	// Tracing
	_, span := h.tracer.Start(context.Background(), "login")
	defer span.End()

	token, err := impl.LoginUser(h.Client, api.Body)
	conflict := int64(operations.PostLoginUnauthorizedCode)
	if err != nil {
		// Logging
		slog.Error(
			"failed login user",
			slog.String("method", "POST"),
			slog.Group("user-properties",
				slog.String("login", *api.Body.Login),
			),
			slog.Int("status_code", operations.PostRegisterConflictCode),
			slog.String("error", err.Error()),
		)
		return new(operations.PostRegisterConflict).WithPayload(&models.Error{
			ErrorMessage:    err.Error(),
			ErrorStatusCode: &conflict,
		})
	}
	// Logging
	slog.Info(
		"user login",
		slog.String("method", "POST"),
		slog.Group("user-properties",
			slog.String("login", *api.Body.Login),
		),
		slog.Int("status_code", operations.PostLoginOKCode),
	)

	result := new(operations.PostLoginOK).WithPayload(&operations.PostLoginOKBody{
		Token: *token,
	})
	return result
}