package json

import (
	"encoding/json"
)

type Container[T any] struct {
	RawValue T
}

// encoding.BinaryMarshaler
func (c Container[T]) MarshalBinary() ([]byte, error) {
	return json.Marshal(c.RawValue)
}

// encoding.BinaryUnmarshaler
func (c *Container[T]) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &c.RawValue)
}
