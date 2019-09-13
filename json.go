package uuid

import (
	"fmt"
)

/**
MarshalToString(v interface{}) (string, error)
Marshal(v interface{}) ([]byte, error)
MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
UnmarshalFromString(str string, v interface{}) error
Unmarshal(data []byte, v interface{}) error


*/

// MarshalJSON marshals the NullUUID as null or the nested UUID
func (u *UUID) MarshalJSON() ([]byte, error) {
	return []byte(u.String()), nil
}

// UnmarshalJSON unmarshals a NullUUID
func (u *UUID) UnmarshalJSON(text []byte) error {
	switch len(text) {
	case 32:
		return u.decodeHashLike(text)
	case 34, 38:
		return u.decodeBraced(text)
	case 36:
		return u.decodeCanonical(text)
	case 41, 45:
		return u.decodeURN(text)
	default:
		return fmt.Errorf("uuid: incorrect UUID length: %s", text)
	}
}
