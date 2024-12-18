package main

import (
	"github.com/h4x4d/go_hsse_hotels/notification/internal/handlers"
	"github.com/h4x4d/go_hsse_hotels/notification/internal/server"
	"log"
	"log/slog"
	"os"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")
	groupID := os.Getenv("KAFKA_GROUP_ID")

	notifyHandlers := map[string]func([]byte) error{
		"send_notification": handlers.SendNotificationHandler,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	notificationServer := server.NewNotificationServer(&[]string{broker}, &topic, &groupID, notifyHandlers)

	if err := notificationServer.Serve(); err != nil {
		log.Fatalln(err)
	}
}
