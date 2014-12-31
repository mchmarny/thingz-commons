package types

import (
	"fmt"
	"time"
)

// NewMetricSample is a factory method for Metric
func NewMetricSample(metric string, value interface{}) *MetricSample {
	return &MetricSample{
		Timestamp: time.Now(),
		Metric:    metric,
		Value:     value,
	}
}

// MetricSample represents a generic sample of a metric in time
type MetricSample struct {

	// Timestamp of when the metric was captured
	Timestamp time.Time `json:"t"`

	// Dimension this metric represents
	Metric string `json:"m"`

	// Value of this metric
	Value interface{} `json:"v"`

	// Unit of this metric
	Unit string `json:"u,omitempty"`

	// Context data for this metric
	Context map[string]string `json:"c,omitempty"`
}

// SetUnit sets metric unit
// and return itself to allow for chaining
func (m *MetricSample) SetUnit(unit string) *MetricSample {
	m.Unit = unit
	return m
}

// AddContext adds context to this metric
// and return itself to allow for chaining
func (m *MetricSample) AddContext(key, val string) *MetricSample {

	if m.Context == nil {
		m.Context = make(map[string]string)
	}

	m.Context[key] = val
	return m
}

func (m *MetricSample) String() string {
	return fmt.Sprintf(
		"Sample: [ Timestamp:%v, Metric:%s, Value:%v, Unit:%s, Context:%v ]",
		m.Timestamp, m.Metric, m.Value, m.Unit, m.Context,
	)
}
