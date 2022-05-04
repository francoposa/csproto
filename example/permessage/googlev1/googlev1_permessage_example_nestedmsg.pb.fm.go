// GENERATED CODE - DO NOT EDIT
// This file was generated by protoc-gen-fastmarshal

package googlev1

import (
	"fmt"
	"github.com/CrowdStrike/csproto"
)

//------------------------------------------------------------------------------
// Custom Protobuf size/marshal/unmarshal code for TestEvent_NestedMsg

// Size calculates and returns the size, in bytes, required to hold the contents of m using the Protobuf
// binary encoding.
func (m *TestEvent_NestedMsg) Size() int {
	if m == nil {
		return 0
	}
	var sz, l int
	_ = l // avoid unused variable

	// Details (string,optional)
	if l = len(m.Details); l > 0 {
		sz += csproto.SizeOfTagKey(1) + csproto.SizeOfVarint(uint64(l)) + l
	}
	return sz
}

// Marshal converts the contents of m to the Protobuf binary encoding and returns the result or an error.
func (m *TestEvent_NestedMsg) Marshal() ([]byte, error) {
	siz := m.Size()
	buf := make([]byte, siz)
	err := m.MarshalTo(buf)
	return buf, err
}

// MarshalTo converts the contents of m to the Protobuf binary encoding and writes the result to dest.
func (m *TestEvent_NestedMsg) MarshalTo(dest []byte) error {
	var (
		enc    = csproto.NewEncoder(dest)
		buf    []byte
		err    error
		extVal interface{}
	)
	// ensure no unused variables
	_ = enc
	_ = buf
	_ = err
	_ = extVal

	// Details (1,string,optional)
	if len(m.Details) > 0 {
		enc.EncodeString(1, m.Details)
	}
	return nil
}

// Unmarshal decodes a binary encoded Protobuf message from p and populates m with the result.
func (m *TestEvent_NestedMsg) Unmarshal(p []byte) error {
	if len(p) == 0 {
		return fmt.Errorf("cannot unmarshal from an empty buffer")
	}
	// clear any existing data
	m.Reset()
	dec := csproto.NewDecoder(p)
	for dec.More() {
		tag, wt, err := dec.DecodeTag()
		if err != nil {
			return err
		}
		switch tag {
		case 1: // Details (string,optional)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for field 'details' (tag=1), expected 2 (length-delimited)", wt)
			}
			if s, err := dec.DecodeString(); err != nil {
				return fmt.Errorf("unable to decode string value for field 'details' (tag=1): %w", err)
			} else {
				m.Details = s
			}

		default:
			if skipped, err := dec.Skip(tag, wt); err != nil {
				return fmt.Errorf("invalid operation skipping tag %v: %w", tag, err)
			} else {
				m.unknownFields = append(m.unknownFields, skipped...)
			}
		}
	}
	return nil
}
