package logging

import (
	"github.com/prometheus/client_golang/prometheus"
)

// ActionResponseTime represents the Telemetry counter used to classify each executed action by response time
var ActionResponseTime = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "bdjuno_action_response_time",
		Help:    "Time it has taken to execute an action",
		Buckets: []float64{0.5, 1, 2, 3, 4, 5},
	}, []string{"path"})

// ActionCounter represents the Telemetry counter used to track the total number of actions executed
var ActionCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "bdjuno_actions_total_count",
		Help: "Total number of actions executed.",
	}, []string{"path", "http_status_code"})

// ActionErrorCounter represents the Telemetry counter used to track the number of action's errors emitted
var ActionErrorCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "bdjuno_actions_error_count",
		Help: "Total number of errors emitted.",
	}, []string{"path", "http_status_code"},
)

// BlockTimeGauge represents the Telemetry gauge used to track chain block time
var BlockTimeGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "bdjuno_block_time",
		Help: "The current bdjuno block time.",
	}, []string{
		"period",
	},
)

// ProposalSummary represents the Telemetry summary used to track proposals
var ProposalSummary = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "bdjuno_proposal",
		Help: "Counts successful proposals.",
	}, []string{
		"validator",
	},
)

// ValidatorBlockMismatchCounter represents the Telemetry counter used to track cases when height in processed block
// differs from the one in returned validator set.
var ValidatorBlockMismatchCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "bdjuno_validator_block_mismatch",
		Help: "Total number of mismatched block heights in validator set.",
	},
)

func init() {
	for _, c := range []prometheus.Collector{
		ActionResponseTime,
		ActionCounter,
		ActionErrorCounter,
		BlockTimeGauge,
		ProposalSummary,
		ValidatorBlockMismatchCounter,
	} {
		if err := prometheus.Register(c); err != nil {
			panic(err)
		}
	}
}
