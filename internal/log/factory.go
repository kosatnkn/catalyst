package log

import (
	"github.com/kosatnkn/catalyst/v3/app/adapters"
)

func NewAdapter(cfg Config) (adapters.LogAdapterInterface, error) {
	if err := validateCfg(cfg); err != nil {
		return nil, err
	}

	return newJSONLogger(cfg)
}
