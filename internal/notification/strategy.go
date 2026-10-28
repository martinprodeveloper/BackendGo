// Package notification strategy
package notification

import (
	"context"
	"log/slog"
)

// Notifier interface
type Notifier interface {
	Notify(ctx context.Context, userID string, message string) error
}

// LogNotifier struct
type LogNotifier struct {
	logger *slog.Logger
}

// NewLogNotifier func
func NewLogNotifier(logger *slog.Logger) *LogNotifier {
	return &LogNotifier{logger: logger}
}

// Notify func
func (notifier *LogNotifier) Notify(ctx context.Context, userID string, message string) error {
	notifier.logger.InfoContext(ctx, "notificación enviada",
		slog.String("user_id", userID),
		slog.String("message", message),
	)
	return nil
}

// NoOpNotifier struct
type NoOpNotifier struct{}

// Notify func
func (notifier *NoOpNotifier) Notify(_ context.Context, _ string, _ string) error {
	return nil
}
