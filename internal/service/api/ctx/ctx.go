package ctx

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/apodeixis/mailjet-svc/internal/mailjet"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota + 1
	mailjetClientCtxKey
)

func SetLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func SetMailjetClient(entry mailjet.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, mailjetClientCtxKey, entry)
	}
}

func MailjetClient(r *http.Request) mailjet.Connector {
	return r.Context().Value(mailjetClientCtxKey).(mailjet.Connector)
}
