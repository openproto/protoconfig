package openconfig

type Configurable interface {
	Unmarshal([]byte) error
	UnmarshalString(string) error
}
