package common

import (
	"context"
	"errors"
	"fmt"
)

var (
	ExampleSentinelError = errors.New("sentinel-error-example")
	ContextKey1          = "context-key-1"
	ContextKey2          = contextKey{}
)

type StructError struct {
	Message string
}

func (e *StructError) Error() string {
	return fmt.Sprintf("struct error: message=%s", e.Message)
}

type contextKey struct{}

func SetContextOne(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, ContextKey1, value)
}

func GetContextOne(ctx context.Context) string {
	v, _ := ctx.Value(ContextKey1).(string)
	return v
}

func SetContextTwo(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, ContextKey2, value)
}

func GetContextTwo(ctx context.Context) string {
	v, _ := ctx.Value(ContextKey2).(string)
	return v
}
