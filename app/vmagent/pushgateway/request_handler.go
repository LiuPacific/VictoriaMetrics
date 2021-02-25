package pushgateway

import (
	parser "github.com/VictoriaMetrics/VictoriaMetrics/app/vmagent/pushgateway/parser"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/prompbmarshal"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/writeconcurrencylimiter"
	"github.com/VictoriaMetrics/metrics"
	"net/http"
)

var (
	rowsInserted  = metrics.NewCounter(`vmagent_rows_inserted_total{type="pushgateway"}`)
	rowsPerInsert = metrics.NewHistogram(`vmagent_rows_per_insert{type="pushgateway"}`)
)

// InsertHandler processes `/api/v1/pushgateway` request.
//
// See https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6
func InsertHandler(req *http.Request) error {
	return writeconcurrencylimiter.Do(func() error {
		return parser.ParseStream(req, func(labels map[string]string) error {
			return insertRows(labels, nil)
		})
	})
}

func insertRows(labels map[string]string, extraLabels []prompbmarshal.Label) error {
	return nil
}
