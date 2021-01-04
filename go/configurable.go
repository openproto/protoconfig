package protoconfig

type Configurable interface {
	// Decode parses byte slice as `Encoded Configuration Message` in JSON or proto format and unmarshal it on
	// the Configurable struct. It supports all `Proto Config Extensions Format 1.0` extenstion
	// (validation, default values etc).
	// Use `proto.Unmarshal` or `protojson.Unmarshal` for decoding without `ProtoConfig 1.0` extension support.
	Decode(ecm []byte) error
	// DecodeString parses string `Encoded Configuration Message` in JSON or proto format and unmarshal it on
	// the Configurable struct. It supports all `Proto Config Extensions Format 1.0` extenstion
	// (validation, default values etc).
	// Use `proto.Unmarshal` or `protojson.Unmarshal` for decoding without `ProtoConfig 1.0` extension support.
	DecodeString(ecm string) error
}
