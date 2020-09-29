package service

// ContextKey is a type used for the keys in the context.
type ContextKey string

const (
	// IDKey is used to store the request ID.
	IDKey ContextKey = "IDKey"
	// TimeInitialKey is used to store the start time of the request.
	TimeInitialKey = "TimeInitialKey"
	// TimeElapsedKey is used to store the duration of the request.
	TimeElapsedKey = "TimeElapsedKey"
	// RequestCounterKey is used to store the number of requests.
	RequestCounterKey = "RequestCounterKey"
)
