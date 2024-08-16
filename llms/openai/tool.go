package openai

import (
	"context"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai/internal/openaiclient"
)

type Tool struct {
	openaiclient.Tool
	CallFunc func(ctx context.Context, input string) (string, error)
}

// NewTool creates a new tool.
func NewTool(funcDef llms.FunctionDefinition, callFunc func(ctx context.Context, input string) (string, error)) Tool {
	oaiFuncDef := openaiclient.FunctionDefinition{
		Name:        funcDef.Name,
		Description: funcDef.Description,
		Parameters:  funcDef.Parameters,
		Strict:      funcDef.Strict,
	}
	return Tool{
		Tool: openaiclient.Tool{
			Type:		openaiclient.ToolTypeFunction,
			Function:	oaiFuncDef,
		},
		CallFunc: callFunc,
	}
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

