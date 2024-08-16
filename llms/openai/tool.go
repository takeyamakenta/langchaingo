package openai

import (
	"context"
	"github.com/tmc/langchaingo/llms/openai/internal/openaiclient"
)

type Tool struct {
	openaiclient.Tool
	CallFunc func(ctx context.Context, input string) (string, error)
}

func (t Tool) Name() string {
	return t.Function.Name
}

func (t Tool) Description() string {
	return t.Function.Description
}

func (t Tool) Call(ctx context.Context, input string) (string, error) {
	return t.CallFunc(context.Background(), input)
}

