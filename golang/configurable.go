package openconfig

type Configurable interface {
	// Decode parses byte slice as `Encoded Configuration Message` in JSON or proto format and unmarshal it on
	// the Configurable struct. It supports all `OpenConfig Proto Extensions Format 1.0` extenstion
	// (validation, default values etc).
	// Use `proto.Unmarshal` or `protojson.Unmarshal` for decoding without `OpenConfig 1.0` extension support.
	Decode(ecm []byte) error
	// DecodeString parses string `Encoded Configuration Message` in JSON or proto format and unmarshal it on
	// the Configurable struct. It supports all `OpenConfig Proto Extensions Format 1.0` extenstion
	// (validation, default values etc).
	// Use `proto.Unmarshal` or `protojson.Unmarshal` for decoding without `OpenConfig 1.0` extension support.
	DecodeString(ecm string) error
}
