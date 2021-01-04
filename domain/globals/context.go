package globals

import "context"

// Context key type to be used with contexts.
type contextKey string

// UUIDKey is the universally unique identifier key to be used with context.
const UUIDKey contextKey = "UUID"

// PrefixKey is the key to add an additional prefix value to the context.
const PrefixKey contextKey = "Prefix"

// TxKey is the key to attach a database transaction to the context.
const TxKey contextKey = "Tx"

// AddTrace appends the given prefix string to the globals.PrefixKey.
func AddTrace(ctx context.Context, prefix string) context.Context {

	pfx := ctx.Value(PrefixKey)

	if pfx == nil {
		return context.WithValue(ctx, PrefixKey, prefix)
	}

	return context.WithValue(ctx, PrefixKey, pfx.(string)+"."+prefix)
}
