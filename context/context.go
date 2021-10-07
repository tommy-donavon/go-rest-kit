package context

import (
	"context"
	"fmt"
)

type (
	ContextHandler struct{}

	ContextValues struct {
		values map[string]interface{}
	}

	contextKey string
)

const (
	ck contextKey = "contextKey"
)

func (c *ContextHandler) Get(ctx context.Context, key string) (interface{}, error) {
	ctxMap, ok := ctx.Value(ck).(ContextValues)
	if !ok {
		return nil, fmt.Errorf("unable to retrieve request context")
	}
	if value, ok := ctxMap.values[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("error retrieving item at key %s", key)
}

func (c *ContextHandler) Add(ctx context.Context, key string, value interface{}) context.Context {

	ctxMap, ok := ctx.Value(ck).(ContextValues)
	if !ok {
		ctxMap = ContextValues{
			values: make(map[string]interface{}),
		}
	}
	ctxMap.values[key] = value
	return context.WithValue(ctx, ck, ctxMap)

}
