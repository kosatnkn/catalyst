package adapters

import "github.com/kosatnkn/log"

// LogAdapterInterface is implemented by all logging adapters.
type LogAdapterInterface interface {
	log.AdapterInterface
}
