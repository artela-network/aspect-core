package types

// Gas measured by the SDK
type Gas = uint64

// GasMeter interface to track gas consumption
type CosmosGasMeter interface {
	GasConsumed() Gas
	GasConsumedToLimit() Gas
	GasRemaining() Gas
	Limit() Gas
}
