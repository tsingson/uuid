package uuid

import (
	"io"
)

// LitterDump  dump
// interface for github.com/sanity-io/litter
func (u UUID) LitterDump(w io.Writer) {
	w.Write([]byte(u.String()))
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
func (u UUID) ExtensionType() int8 { return 99 }

// Len We'll always use 4 bytes to encode the data
func (u UUID) Len() int { return Size }

// MarshalBinaryTo simply copies the value
// of the bytes into 'b'
func (u UUID) MarshalBinaryTo(b []byte) error {
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
