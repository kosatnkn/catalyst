package log

// Context key type to be used with contexts.
type key string

// ID is the universally unique identifier key to be used with context.
const ID key = "id"

// TraceKey is the key to add an additional trace values to the context.
const TraceKey key = "trace"
