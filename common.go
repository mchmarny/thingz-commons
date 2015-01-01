package commons

import (
	"fmt"
)

// FormatMetricName
func FormatMetricName(src, dim, met string) string {
	return fmt.Sprintf("src.%s.dim.%s.met.%s", src, dim, met)
}
