package processors

import (
	"soloanvill_backend/internals/cfg"
)

type HealthProcessor struct {
	cfg *cfg.Cfg
}

func NewHealthProccessor(cfg *cfg.Cfg) *HealthProcessor {
	processor := new(HealthProcessor)
	processor.cfg = cfg
	return processor
}
