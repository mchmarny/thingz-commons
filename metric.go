package commons

import (
	"encoding/json"
	"fmt"
	"time"
)

// NewMetric is a factory method for minimal valid Metric
func NewMetric(source, dimension, metric string, value interface{}) *Metric {
	return &Metric{
		Source:    source,
		Timestamp: time.Now(),
		Dimension: dimension,
		Metric:    metric,
		Value:     value,
	}
}

// Metric represents a generic sample of a metric in time
type Metric struct {

	// Source this collection represents
	Source string `json:"src"`

	// Timestamp of when the metric was captured
	Timestamp time.Time `json:"ts"`

	// Dimension of metrics this collection represents
	Dimension string `json:"dim"`

	// Dimension this metric represents
	Metric string `json:"met"`

	// Value of this metric
	Value interface{} `json:"val"`

	// Unit of this metric
	Unit string `json:"unit,omitempty"`

	// Context data for this metric
	Context map[string]string `json:"c,omitempty"`
}

// AddContext adds context to this metric
// and return itself to allow for chaining
func (m *Metric) AddContext(key, val string) *Metric {

	if m.Context == nil {
		m.Context = make(map[string]string)
	}

	m.Context[key] = val
	return m
}

// FormatFQName
func (m *Metric) FormatFQName() string {
	return fmt.Sprintf("src.%s.dim.%s.met.%s",
		m.Source, m.Dimension, m.Metric)
}

// String
func (m *Metric) String() string {
	return fmt.Sprintf(
		"Metric: [ Timestamp:%v, FQName:%s, Value:%v, Unit:%s, Context:%v ]",
		m.Timestamp, m.FormatFQName(), m.Value, m.Unit, m.Context,
	)
}

// ToBytes converts content of the current message into byte array
func (m *Metric) ToBytes() ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("unable to marshal: %v", err.Error())
		return nil, err
	}
	return b, nil
}

// ParseMetricCollection converts array of bytes into MetricCollection pointer
func ParseMetric(data []byte) (*Metric, error) {
	col := &Metric{}
	if err := json.Unmarshal(data, &col); err != nil {
		fmt.Printf("unable to unmarshal: %s", string(data))
		return nil, err
	}
	return col, nil
}
