// GENERATED CODE - DO NOT EDIT
// This file was generated by protoc-gen-fastmarshal

package googlev2

import (
	"fmt"
	"github.com/CrowdStrike/csproto"
)

//------------------------------------------------------------------------------
// Custom Protobuf size/marshal/unmarshal code for EmbeddedEvent

// Size calculates and returns the size, in bytes, required to hold the contents of m using the Protobuf
// binary encoding.
func (m *EmbeddedEvent) Size() int {
	if m == nil {
		return 0
	}
	var sz, l int
	_ = l // avoid unused variable

	// ID (int32,optional)
	sz += csproto.SizeOfTagKey(1) + csproto.SizeOfVarint(uint64(m.ID))
	// Stuff (string,optional)
	if l = len(m.Stuff); l > 0 {
		sz += csproto.SizeOfTagKey(2) + csproto.SizeOfVarint(uint64(l)) + l
	}
	// FavoriteNumbers (int32,repeated,packed)
	if len(m.FavoriteNumbers) > 0 {
		sz += csproto.SizeOfTagKey(3)
		l = 0
		for _, iv := range m.FavoriteNumbers {
			l += csproto.SizeOfVarint(uint64(iv))
		}
		sz += csproto.SizeOfVarint(uint64(l)) + l
	}
	// RandomThings (bytes,repeated)
	for _, bv := range m.RandomThings {
		if l = len(bv); l > 0 {
			sz += csproto.SizeOfTagKey(4) + csproto.SizeOfVarint(uint64(l)) + l
		}
	}
	return sz
}

// Marshal converts the contents of m to the Protobuf binary encoding and returns the result or an error.
func (m *EmbeddedEvent) Marshal() ([]byte, error) {
	siz := m.Size()
	buf := make([]byte, siz)
	err := m.MarshalTo(buf)
	return buf, err
}

// MarshalTo converts the contents of m to the Protobuf binary encoding and writes the result to dest.
func (m *EmbeddedEvent) MarshalTo(dest []byte) error {
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

	// ID (1,int32,optional)
	enc.EncodeInt32(1, m.ID)
	// Stuff (2,string,optional)
	if len(m.Stuff) > 0 {
		enc.EncodeString(2, m.Stuff)
	}
	// FavoriteNumbers (3,int32,repeated,packed)
	if len(m.FavoriteNumbers) > 0 {
		enc.EncodePackedInt32(3, m.FavoriteNumbers)
	}
	// RandomThings (4,bytes,repeated)
	for _, val := range m.RandomThings {
		enc.EncodeBytes(4, val)
	}
	return nil
}

// Unmarshal decodes a binary encoded Protobuf message from p and populates m with the result.
func (m *EmbeddedEvent) Unmarshal(p []byte) error {
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
		case 1: // ID (int32,optional)
			if wt != csproto.WireTypeVarint {
				return fmt.Errorf("incorrect wire type %v for tag field 'ID' (tag=1), expected 0 (varint)", wt)
			}
			if v, err := dec.DecodeInt32(); err != nil {
				return fmt.Errorf("unable to decode int32 value for field 'ID' (tag=1): %w", err)
			} else {
				m.ID = v
			}
		case 2: // Stuff (string,optional)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for field 'stuff' (tag=2), expected 2 (length-delimited)", wt)
			}
			if s, err := dec.DecodeString(); err != nil {
				return fmt.Errorf("unable to decode string value for field 'stuff' (tag=2): %w", err)
			} else {
				m.Stuff = s
			}

		case 3: // FavoriteNumbers (int32,repeated,packed)
			switch wt {
			case csproto.WireTypeVarint:
				if v, err := dec.DecodeInt32(); err != nil {
					return fmt.Errorf("unable to decode int32 value for field 'favoriteNumbers' (tag=3): %w", err)
				} else {
					m.FavoriteNumbers = append(m.FavoriteNumbers, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedInt32(); err != nil {
					return fmt.Errorf("unable to decode packed int32 values for field 'favoriteNumbers' (tag=3): %w", err)
				} else {
					m.FavoriteNumbers = append(m.FavoriteNumbers, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'favoriteNumbers' (tag=3), expected 0 (varint) or 1 (length-delimited)", wt)
			}
		case 4: // RandomThings (bytes,repeated)

			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for field 'randomThings' (tag=4), expected 2 (length-delimited)", wt)
			}
			if b, err := dec.DecodeBytes(); err != nil {
				return fmt.Errorf("unable to decode bytes value for field 'randomThings' (tag=4): %w", err)
			} else {
				m.RandomThings = append(m.RandomThings, b)
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
