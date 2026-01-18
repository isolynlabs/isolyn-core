package intelligence

// Engine defines the interface for market intelligence components.
//
// Implementations provide probabilistic analysis and directional signals.
// This layer MUST NOT custody capital, execute transactions, or access
// user identity or wallet data.
type Engine interface {

	// Evaluate analyzes market state and returns an assessment.
	Evaluate(ctx EvaluationContext) (Signal, error)
}

// EvaluationContext carries analysis-scoped information.
//
// This context MUST NOT include user identifiers or capital attribution.
type EvaluationContext struct {

	// Market identifies the target market.
	Market string

	// Timeframe represents the analysis timeframe.
	Timeframe string
}

// Signal represents a directional assessment.
type Signal struct {

	// Side indicates long or short bias.
	Side SignalSide

	// Confidence represents relative confidence (0.0 - 1.0).
	Confidence float64
}

// SignalSide represents directional bias.
type SignalSide string

const (
	SignalLong  SignalSide = "long"
	SignalShort SignalSide = "short"
)
