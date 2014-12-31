package types

import (
	"encoding/json"
	"fmt"
	"log"
)

// NewMetricCollection is a factory method for MetricCollection
func NewMetricCollection(src, dim string) *MetricCollection {
	return &MetricCollection{
		Source:    src,
		Dimension: dim,
		Metrics:   make([]*MetricSample, 0),
	}
}

// MetricCollection represents a generic metric collection event
type MetricCollection struct {

	// Source this collection represents
	Source string `json:"src"`

	// Dimension of metrics this collection represents
	Dimension string `json:"dim"`

	// Metrics represents a collection of metric items
	Metrics []*MetricSample `json:"list"`
}

// Add adds metric to collection
// and return itself to allow for chaining
func (m *MetricCollection) Add(item *MetricSample) {
	m.Metrics = append(m.Metrics, item)
}

// String
func (m *MetricCollection) String() string {
	return fmt.Sprintf(
		"MetricCollection: [ Source:%s, Dimension:%s, Metrics:%v ]",
		m.Source, m.Dimension, m.Metrics,
	)
}

// ToBytes converts content of the current message into byte array
func (m *MetricCollection) ToBytes() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		log.Printf("unable to marshal: %v", err.Error())
	}
	return b
}

// ParseMetricCollection converts array of bytes into MetricCollection pointer
func ParseMetricCollection(data []byte) (*MetricCollection, error) {
	col := &MetricCollection{}
	if err := json.Unmarshal(data, &col); err != nil {
		log.Printf("unable to unmarshal: %s", string(data))
		return nil, err
	}
	return col, nil
}
