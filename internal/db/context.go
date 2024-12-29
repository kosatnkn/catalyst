package db

// Context key type to be used with contexts.
type key string

// TxKey is the key used to bind a transaction to context.
const TxKey key = "tx"
