package chains

import (
	"context"

	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/llms"
)

// ChainCallOption is a function that can be used to modify the behavior of the Call function.
type ChainCallOption = llms.CallOption

// For issue #626, each field here has a boolean "set" flag so we can
// distinguish between the case where the option was actually set explicitly
// on chainCallOption, or asked to remain default. The reason we need this is
// that in translating options from ChainCallOption to llms.CallOption, the
// notion of "default value the user didn't explicitly ask to change" is
// violated.
// These flags are hopefully a temporary backwards-compatible solution, until
// we find a more fundamental solution for #626.
type chainCallOption struct {
	// StreamingFunc is a function to be called for each chunk of a streaming response.
	// Return an error to stop streaming early.
	StreamingFunc func(ctx context.Context, chunk []byte) error

	// CallbackHandler is the callback handler for Chain
	CallbackHandler callbacks.Handler
}

// WithModel is an option for LLM.Call.
func WithModel(model string) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*llms.CallOptions); ok {
			a.Model = model
		}
	}
}

// WithMaxTokens is an option for LLM.Call.
func WithMaxTokens(maxTokens int) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*llms.CallOptions); ok {
			a.MaxTokens = maxTokens
		}
	}
}

// WithTemperature is an option for LLM.Call.
func WithTemperature(temperature float64) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*llms.CallOptions); ok {
			a.Temperature = temperature
		}
	}
}

// WithStreamingFunc is an option for LLM.Call that allows streaming responses.
func WithStreamingFunc(streamingFunc func(ctx context.Context, chunk []byte) error) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*chainCallOption); ok {
			a.StreamingFunc = streamingFunc
		}
	}
}

// WithTopK will add an option to use top-k sampling for LLM.Call.
func WithTopK(topK int) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*llms.CallOptions); ok {
			a.TopK = topK
		}
	}
}

// WithTopP	will add an option to use top-p sampling for LLM.Call.
func WithTopP(topP float64) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*llms.CallOptions); ok {
			a.TopP = topP
		}
	}
}

// WithSeed will add an option to use deterministic sampling for LLM.Call.
func WithSeed(seed int) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*llms.CallOptions); ok {
			a.Seed = seed
		}
	}
}

// WithMinLength will add an option to set the minimum length of the generated text for LLM.Call.
func WithMinLength(minLength int) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*llms.CallOptions); ok {
			a.MinLength = minLength
		}
	}
}

// WithMaxLength will add an option to set the maximum length of the generated text for LLM.Call.
func WithMaxLength(maxLength int) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*llms.CallOptions); ok {
			a.MaxLength = maxLength
		}
	}
}

// WithRepetitionPenalty will add an option to set the repetition penalty for sampling.
func WithRepetitionPenalty(repetitionPenalty float64) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*llms.CallOptions); ok {
			a.RepetitionPenalty = repetitionPenalty
		}
	}
}

// WithStopWords is an option for setting the stop words for LLM.Call.
func WithStopWords(stopWords []string) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*llms.CallOptions); ok {
			a.StopWords = stopWords
		}
	}
}

// WithCallback allows setting a custom Callback Handler.
func WithCallback(callbackHandler callbacks.Handler) ChainCallOption {
	return func(o llms.ICallOptions) {
		if a, ok := o.(*chainCallOption); ok {
			a.CallbackHandler = callbackHandler
		}
	}
}

func getLLMCallOptions(options ...ChainCallOption) []llms.CallOption { //nolint:cyclop
	opts := &chainCallOption{}
	for _, option := range options {
		option(opts)
	}
	if opts.StreamingFunc == nil && opts.CallbackHandler != nil {
		opts.StreamingFunc = func(ctx context.Context, chunk []byte) error {
			opts.CallbackHandler.HandleStreamingFunc(ctx, chunk)
			return nil
		}
	}
	options = append(options, llms.WithStreamingFunc(opts.StreamingFunc))

	return options
}
