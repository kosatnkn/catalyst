package log

import (
	"fmt"

	"github.com/kosatnkn/catalyst/v3/app/adapters"
)

func NewAdapter(cfg Config) (adapters.LogAdapterInterface, error) {
	if err := validateCfg(cfg); err != nil {
		return nil, err
	}

	switch cfg.Flavor {
	case flavourText:
		return newTextLogger(cfg)
	case flavourJSON:
		return newJSONLogger(cfg)
	default:
		return nil, fmt.Errorf("logger: unknown log flavour '%s'", cfg.Flavor)
	}
}
