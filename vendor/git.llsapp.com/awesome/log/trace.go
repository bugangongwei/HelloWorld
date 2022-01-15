package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

// TraceIDKey trace ID context key
var TraceIDKey = "TraceID"

// FromContextTraceID load trace ID from context
func FromContextTraceID(ctx context.Context) string {
	id, ok := ctx.Value(TraceIDKey).(string)
	if !ok {
		return ""
	}
	return id
}

// WithTraceID returns a copy of parent in which the value associated with key is trace_id
func WithTraceID(ctx context.Context, id interface{}) context.Context {
	return context.WithValue(ctx, TraceIDKey, id)
}

// Ctx creates an entry from the standard logger and adds a context to it.
// Add a single field(trace_id) to the Entry
func Ctx(ctx context.Context) *Entry {
	return logrus.WithContext(ctx).WithField("trace_id", FromContextTraceID(ctx))
}
