package openconfig

// Configurator allows to produce `Encoded Configuration Messages` from the `Configuration Proto Definition`.
type Configurator interface {
	// Encode encodes self as `Encoded Configuration Message` in proto format so it can be understood and
	// passed to Configurable struct. It supports all `OpenConfig Proto Extensions Format 1.0` extenstion
	// (validation, default values etc).
	// Use `proto.Marshal` encoding without `OpenConfig 1.0` extension support.
	Encode() ([]byte, error)
	// EncodeJSON encodes self as `Encoded Configuration Message` in JSON format so it can be understood and
	// passed to Configurable struct. It supports all `OpenConfig Proto Extensions Format 1.0` extenstion
	// Use `protojson.Marshal` encoding without `OpenConfig 1.0` extension support.
	EncodeJSON() ([]byte, error)

	// Metadata returns metadata defined in `OpenConfig Proto Extensions Format 1.0`.
	Metadata() Metadata
}
