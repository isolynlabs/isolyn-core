package execution

// Executor defines the interface for isolated execution components.
//
// Implementations are responsible for placing and managing perpetual
// positions using ephemeral execution accounts.
//
// An Executor MUST NOT:
// - custody user funds
// - access user identity
// - infer capital origin
// - persist execution identifiers across sessions
//
// This interface is intentionally minimal and subject to change.
type Executor interface {

	// OpenPosition places a new perpetual position based on the provided
	// execution context.
	OpenPosition(ctx ExecutionContext) error

	// ClosePosition closes an existing position identified by the context.
	ClosePosition(ctx ExecutionContext) error
}

// ExecutionContext carries execution-scoped information.
//
// This context MUST NOT include user identifiers, wallet addresses,
// or capital attribution data.
type ExecutionContext struct {

	// Market identifies the target perpetual market.
	Market string

	// Side indicates long or short execution intent.
	Side PositionSide

	// Size represents the notional position size.
	Size float64
}

// PositionSide represents the directional intent of a position.
type PositionSide string

const (
	PositionLong  PositionSide = "long"
	PositionShort PositionSide = "short"
)
