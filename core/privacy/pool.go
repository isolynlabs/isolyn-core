package privacy

// Pool defines the interface for privacy-preserving capital abstraction.
//
// Implementations are responsible for pooling deposits, abstracting
// attribution, and enabling private settlement.
//
// A Pool MUST NOT expose address linkage or reuse persistent identifiers.
type Pool interface {

	// Deposit accepts capital into the privacy pool.
	Deposit(ctx DepositContext) error

	// Withdraw releases capital from the pool.
	Withdraw(ctx WithdrawalContext) error
}

// DepositContext carries deposit-scoped information.
//
// This context MUST NOT include user identity.
type DepositContext struct {

	// Asset identifies the deposited asset.
	Asset string

	// Amount represents the deposited amount.
	Amount float64
}

// WithdrawalContext carries withdrawal-scoped information.
//
// This context MUST NOT include execution linkage.
type WithdrawalContext struct {

	// Asset identifies the withdrawn asset.
	Asset string

	// Amount represents the withdrawal amount.
	Amount float64
}
