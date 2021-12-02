package filterconfig

import (
	"errors"

	"github.com/GANGAV08/filterset/filterset"
)

type MatchConfig struct {
	Include *MatchProperties `mapstructure:"include"`

	Exclude *MatchProperties `mapstructure:"exclude"`
}

type MatchProperties struct {
	filterset.Config `mapstructure:",squash"`

	Services []string `mapstructure:"services"`

	SpanNames []string `mapstructure:"span_names"`

	LogNames []string `mapstructure:"log_names"`

	Attributes []Attribute `mapstructure:"attributes"`

	Resources []Attribute `mapstructure:"resources"`

	Libraries []InstrumentationLibrary `mapstructure:"libraries"`
}

func (mp *MatchProperties) ValidateForSpans() error {
	if len(mp.LogNames) > 0 {
		return errors.New("log_names should not be specified for trace spans")
	}

	if len(mp.Services) == 0 && len(mp.SpanNames) == 0 && len(mp.Attributes) == 0 &&
		len(mp.Libraries) == 0 && len(mp.Resources) == 0 {
		return errors.New(`at least one of "services", "span_names", "attributes", "libraries" or "resources" field must be specified`)
	}

	return nil
}

// ValidateForLogs validates properties for logs.
func (mp *MatchProperties) ValidateForLogs() error {
	if len(mp.SpanNames) > 0 || len(mp.Services) > 0 {
		return errors.New("neither services nor span_names should be specified for log records")
	}

	if len(mp.LogNames) == 0 && len(mp.Attributes) == 0 && len(mp.Libraries) == 0 && len(mp.Resources) == 0 {
		return errors.New(`at least one of "log_names", "attributes", "libraries" or "resources" field must be specified`)
	}

	return nil
}

type Attribute struct {
	Key string `mapstructure:"key"`

	Value interface{} `mapstructure:"value"`
}

type InstrumentationLibrary struct {
	Name string `mapstructure:"name"`

	Version *string `mapstructure:"version"`
}
