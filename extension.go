package uuid

import (
	"encoding/hex"
	"io"
)

// LitterDump  dump
// interface for github.com/sanity-io/litter
func (u UUID) LitterDump(w io.Writer) {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	w.Write(buf)
}

// extension for github.com/tinylib/msgp/msgp
// func init() {
// 	// Registering an extension is as simple as matching the
// 	// appropriate type number with a function that initializes
// 	// a freshly-allocated object of that type
// 	msgp.RegisterExtension(99, func() msgp.Extension { return new(RGBA) } )
// }

// RGBA will be the concrete type
// of our new extension

// ExtensionType for uuid
// Here, we'll pick an arbitrary number between
// 0 and 127 that isn't already in use
func (u *UUID) ExtensionType() int8 { return 101 }

// Len We'll always use 4 bytes to encode the data
func (u *UUID) Len() int { return Size }

// MarshalBinaryTo simply copies the value
// of the bytes into 'b'
func (u *UUID) MarshalBinaryTo(b []byte) error {
	copy(b, u.Bytes())
	return nil
}

// UnmarshalBinary copies the value of 'b'
// into the RGBA object. (We might want to add
// a sanity check here later that len(b)==4.)
// func (u UUID) UnmarshalBinary(b []byte) error {
// 	copy(u.Bytes(), b)
// 	return nil
// }
