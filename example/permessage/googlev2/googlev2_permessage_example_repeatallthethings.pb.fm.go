// GENERATED CODE - DO NOT EDIT
// This file was generated by protoc-gen-fastmarshal

package googlev2

import (
	"fmt"
	"sync/atomic"
	"github.com/CrowdStrike/csproto"
)

//------------------------------------------------------------------------------
// Custom Protobuf size/marshal/unmarshal code for RepeatAllTheThings

// Size calculates and returns the size, in bytes, required to hold the contents of m using the Protobuf
// binary encoding.
func (m *RepeatAllTheThings) Size() int {
	// nil message is always 0 bytes
	if m == nil {
		return 0
	}
	// return cached size, if present
	if csz := int(atomic.LoadInt32(&m.sizeCache)); csz > 0 {
		return csz
	}
	// calculate and cache
	var sz, l int
	_ = l // avoid unused variable

	// ID (int32,optional)
	if m.ID != 0 {
		sz += csproto.SizeOfTagKey(1) + csproto.SizeOfVarint(uint64(m.ID))
	}
	// TheStrings (string,repeated)
	for _, sv := range m.TheStrings {
		l = len(sv)
		sz += csproto.SizeOfTagKey(2) + csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheBools (bool,repeated,packed)
	if l = len(m.TheBools); l > 0 {
		sz += csproto.SizeOfTagKey(3) + csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheInt32S (int32,repeated,packed)
	if len(m.TheInt32S) > 0 {
		sz += csproto.SizeOfTagKey(4)
		l = 0
		for _, iv := range m.TheInt32S {
			l += csproto.SizeOfVarint(uint64(iv))
		}
		sz += csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheInt64S (int64,repeated,packed)
	if len(m.TheInt64S) > 0 {
		sz += csproto.SizeOfTagKey(5)
		l = 0
		for _, iv := range m.TheInt64S {
			l += csproto.SizeOfVarint(uint64(iv))
		}
		sz += csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheUInt32S (uint32,repeated,packed)
	if len(m.TheUInt32S) > 0 {
		sz += csproto.SizeOfTagKey(6)
		l = 0
		for _, iv := range m.TheUInt32S {
			l += csproto.SizeOfVarint(uint64(iv))
		}
		sz += csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheUInt64S (uint64,repeated,packed)
	if len(m.TheUInt64S) > 0 {
		sz += csproto.SizeOfTagKey(7)
		l = 0
		for _, iv := range m.TheUInt64S {
			l += csproto.SizeOfVarint(uint64(iv))
		}
		sz += csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheSInt32S (sint32,repeated,packed)
	if len(m.TheSInt32S) > 0 {
		sz += csproto.SizeOfTagKey(8)
		l = 0
		for _, iv := range m.TheSInt32S {
			l += csproto.SizeOfZigZag(uint64(iv))
		}
		sz += csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheSInt64S (sint64,repeated,packed)
	if len(m.TheSInt64S) > 0 {
		sz += csproto.SizeOfTagKey(9)
		l = 0
		for _, iv := range m.TheSInt64S {
			l += csproto.SizeOfZigZag(uint64(iv))
		}
		sz += csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheFixed32S (fixed32,repeated,packed)
	if l = len(m.TheFixed32S); l > 0 {
		sz += csproto.SizeOfTagKey(10) + csproto.SizeOfVarint(uint64(l)) + (l * 4)
	}
	// TheFixed64S (fixed64,repeated,packed)
	if l = len(m.TheFixed64S); l > 0 {
		sz += csproto.SizeOfTagKey(11) + csproto.SizeOfVarint(uint64(l)) + (l * 8)
	}
	// TheSFixed32S (sfixed32,repeated,packed)
	if l = len(m.TheSFixed32S); l > 0 {
		sz += csproto.SizeOfTagKey(12) + csproto.SizeOfVarint(uint64(l)) + (l * 4)
	}
	// TheSFixed64S (sfixed64,repeated,packed)
	if l = len(m.TheSFixed64S); l > 0 {
		sz += csproto.SizeOfTagKey(13) + csproto.SizeOfVarint(uint64(l)) + (l * 8)
	}
	// TheFloats (float,repeated,packed)
	if l = len(m.TheFloats); l > 0 {
		sz += csproto.SizeOfTagKey(14) + csproto.SizeOfVarint(uint64(l)) + (l * 4)
	}
	// TheDoubles (double,repeated,packed)
	if l = len(m.TheDoubles); l > 0 {
		sz += csproto.SizeOfTagKey(15) + csproto.SizeOfVarint(uint64(l)) + (l * 8)
	}
	// TheEventTypes (enum,repeated,packed)
	if len(m.TheEventTypes) > 0 {
		sz += csproto.SizeOfTagKey(16)
		l = 0
		for _, iv := range m.TheEventTypes {
			l += csproto.SizeOfVarint(uint64(iv))
		}
		sz += csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheBytes (bytes,repeated)
	for _, bv := range m.TheBytes {
		l = len(bv)
		sz += csproto.SizeOfTagKey(17) + csproto.SizeOfVarint(uint64(l)) + l
	}
	// TheMessages (message,repeated)
	for _, val := range m.TheMessages {
		if l = csproto.Size(val); l > 0 {
			sz += csproto.SizeOfTagKey(18) + csproto.SizeOfVarint(uint64(l)) + l
		}
	}
	// cache the size so it can be re-used in Marshal()/MarshalTo()
	atomic.StoreInt32(&m.sizeCache, int32(sz))
	return sz
}

// Marshal converts the contents of m to the Protobuf binary encoding and returns the result or an error.
func (m *RepeatAllTheThings) Marshal() ([]byte, error) {
	siz := m.Size()
	buf := make([]byte, siz)
	err := m.MarshalTo(buf)
	return buf, err
}

// MarshalTo converts the contents of m to the Protobuf binary encoding and writes the result to dest.
func (m *RepeatAllTheThings) MarshalTo(dest []byte) error {
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
	if m.ID != 0 {
		enc.EncodeInt32(1, m.ID)
	}
	// TheStrings (2,string,repeated)
	for _, val := range m.TheStrings {
		enc.EncodeString(2, val)
	}
	// TheBools (3,bool,repeated,packed)
	if len(m.TheBools) > 0 {
		enc.EncodePackedBool(3, m.TheBools)
	}
	// TheInt32S (4,int32,repeated,packed)
	if len(m.TheInt32S) > 0 {
		enc.EncodePackedInt32(4, m.TheInt32S)
	}
	// TheInt64S (5,int64,repeated,packed)
	if len(m.TheInt64S) > 0 {
		enc.EncodePackedInt64(5, m.TheInt64S)
	}
	// TheUInt32S (6,uint32,repeated,packed)
	if len(m.TheUInt32S) > 0 {
		enc.EncodePackedUInt32(6, m.TheUInt32S)
	}
	// TheUInt64S (7,uint64,repeated,packed)
	if len(m.TheUInt64S) > 0 {
		enc.EncodePackedUInt64(7, m.TheUInt64S)
	}
	// TheSInt32S (8,sint32,repeated,packed)
	if len(m.TheSInt32S) > 0 {
		enc.EncodePackedSInt32(8, m.TheSInt32S)
	}
	// TheSInt64S (9,sint64,repeated,packed)
	if len(m.TheSInt64S) > 0 {
		enc.EncodePackedSInt64(9, m.TheSInt64S)
	}
	// TheFixed32S (10,fixed32,repeated,packed)
	if len(m.TheFixed32S) > 0 {
		enc.EncodePackedFixed32(10, m.TheFixed32S)
	}
	// TheFixed64S (11,fixed64,repeated,packed)
	if len(m.TheFixed64S) > 0 {
		enc.EncodePackedFixed64(11, m.TheFixed64S)
	}
	// TheSFixed32S (12,sfixed32,repeated,packed)
	if len(m.TheSFixed32S) > 0 {
		enc.EncodePackedSFixed32(12, m.TheSFixed32S)
	}
	// TheSFixed64S (13,sfixed64,repeated,packed)
	if len(m.TheSFixed64S) > 0 {
		enc.EncodePackedSFixed64(13, m.TheSFixed64S)
	}
	// TheFloats (14,float,repeated,packed)
	if len(m.TheFloats) > 0 {
		enc.EncodePackedFloat32(14, m.TheFloats)
	}
	// TheDoubles (15,double,repeated,packed)
	if len(m.TheDoubles) > 0 {
		enc.EncodePackedFloat64(15, m.TheDoubles)
	}
	// TheEventTypes (16,enum,repeated,packed)
	if l := len(m.TheEventTypes); l > 0 {
		ivs := make([]int32, l)
		for i, v := range m.TheEventTypes {
			ivs[i] = int32(v)
		}
		enc.EncodePackedInt32(16, ivs)
	}
	// TheBytes (17,bytes,repeated)
	for _, val := range m.TheBytes {
		enc.EncodeBytes(17, val)
	}
	// TheMessages (18,message,repeated)
	for _, mm := range m.TheMessages {
		if err = enc.EncodeNested(18, mm); err != nil {
			return fmt.Errorf("unable to encode message data for field 'theMessages' (tag=18): %w", err)
		}
	}
	return nil
}

// Unmarshal decodes a binary encoded Protobuf message from p and populates m with the result.
func (m *RepeatAllTheThings) Unmarshal(p []byte) error {
	m.Reset()
	if len(p) == 0 {
		return nil
	}
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
		case 2: // TheStrings (string,repeated)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for field 'theStrings' (tag=2), expected 2 (length-delimited)", wt)
			}
			if s, err := dec.DecodeString(); err != nil {
				return fmt.Errorf("unable to decode string value for field 'theStrings' (tag=2): %w", err)
			} else {
				m.TheStrings = append(m.TheStrings, s)
			}

		case 3: // TheBools (bool,repeated,packed)
			switch wt {
			case csproto.WireTypeVarint:
				if v, err := dec.DecodeBool(); err != nil {
					return fmt.Errorf("unable to decode boolean value for field 'theBools' (tag=3): %w", err)
				} else {
					m.TheBools = append(m.TheBools, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedBool(); err != nil {
					return fmt.Errorf("unable to decode packed boolean values for field 'theBools' (tag=3): %w", err)
				} else {
					m.TheBools = append(m.TheBools, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theBools' (tag=3), expected 0 (varint) or 1 (length-delimited)", wt)
			}
		case 4: // TheInt32S (int32,repeated,packed)
			switch wt {
			case csproto.WireTypeVarint:
				if v, err := dec.DecodeInt32(); err != nil {
					return fmt.Errorf("unable to decode int32 value for field 'theInt32s' (tag=4): %w", err)
				} else {
					m.TheInt32S = append(m.TheInt32S, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedInt32(); err != nil {
					return fmt.Errorf("unable to decode packed int32 values for field 'theInt32s' (tag=4): %w", err)
				} else {
					m.TheInt32S = append(m.TheInt32S, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theInt32s' (tag=4), expected 0 (varint) or 1 (length-delimited)", wt)
			}
		case 5: // TheInt64S (int64,repeated,packed)
			switch wt {
			case csproto.WireTypeVarint:
				if v, err := dec.DecodeInt64(); err != nil {
					return fmt.Errorf("unable to decode int64 value for field 'theInt64s' (tag=5): %w", err)
				} else {
					m.TheInt64S = append(m.TheInt64S, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedInt64(); err != nil {
					return fmt.Errorf("unable to decode packed int64 values for field 'theInt64s' (tag=5): %w", err)
				} else {
					m.TheInt64S = append(m.TheInt64S, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theInt64s' (tag=5), expected 0 (varint) or 1 (length-delimited)", wt)
			}
		case 6: // TheUInt32S (uint32,repeated,packed)
			switch wt {
			case csproto.WireTypeVarint:
				if v, err := dec.DecodeUInt32(); err != nil {
					return fmt.Errorf("unable to decode uint32 value for field 'theUInt32s' (tag=6): %w", err)
				} else {
					m.TheUInt32S = append(m.TheUInt32S, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedUint32(); err != nil {
					return fmt.Errorf("unable to decode packed uint32 values for field 'theUInt32s' (tag=6): %w", err)
				} else {
					m.TheUInt32S = append(m.TheUInt32S, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theUInt32s' (tag=6), expected 0 (varint) or 1 (length-delimited)", wt)
			}
		case 7: // TheUInt64S (uint64,repeated,packed)
			switch wt {
			case csproto.WireTypeVarint:
				if v, err := dec.DecodeUInt64(); err != nil {
					return fmt.Errorf("unable to decode uint64 value for field 'theUInt64s' (tag=7): %w", err)
				} else {
					m.TheUInt64S = append(m.TheUInt64S, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedUint64(); err != nil {
					return fmt.Errorf("unable to decode packed uint64 values for field 'theUInt64s' (tag=7): %w", err)
				} else {
					m.TheUInt64S = append(m.TheUInt64S, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theUInt64s' (tag=7), expected 0 (varint) or 1 (length-delimited)", wt)
			}
		case 8: // TheSInt32S (sint32,repeated,packed)
			switch wt {
			case csproto.WireTypeVarint:
				if v, err := dec.DecodeSInt32(); err != nil {
					return fmt.Errorf("unable to decode sint32 value for field 'theSInt32s' (tag=8): %w", err)
				} else {
					m.TheSInt32S = append(m.TheSInt32S, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedSint32(); err != nil {
					return fmt.Errorf("unable to decode packed sint32 values for field 'theSInt32s' (tag=8): %w", err)
				} else {
					m.TheSInt32S = append(m.TheSInt32S, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theSInt32s' (tag=8), expected 0 (varint) or 1 (length-delimited)", wt)
			}
		case 9: // TheSInt64S (sint64,repeated,packed)
			switch wt {
			case csproto.WireTypeVarint:
				if v, err := dec.DecodeSInt64(); err != nil {
					return fmt.Errorf("unable to decode sint64 value for field 'theSInt64s' (tag=9): %w", err)
				} else {
					m.TheSInt64S = append(m.TheSInt64S, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedSint64(); err != nil {
					return fmt.Errorf("unable to decode packed sint64 values for field 'theSInt64s' (tag=9): %w", err)
				} else {
					m.TheSInt64S = append(m.TheSInt64S, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theSInt64s' (tag=9), expected 0 (varint) or 1 (length-delimited)", wt)
			}
		case 10: // TheFixed32S (fixed32,repeated,packed)
			switch wt {
			case csproto.WireTypeFixed32:
				if v, err := dec.DecodeFixed32(); err != nil {
					return fmt.Errorf("unable to decode uint32 value for field 'theFixed32s' (tag=10): %w", err)
				} else {
					m.TheFixed32S = append(m.TheFixed32S, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedFixed32(); err != nil {
					return fmt.Errorf("unable to decode packed uint32 values for field 'theFixed32s' (tag=10): %w", err)
				} else {
					m.TheFixed32S = append(m.TheFixed32S, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theFixed32s' (tag=10), expected 5 (32-bit) or 1 (length-delimited)", wt)
			}
		case 11: // TheFixed64S (fixed64,repeated,packed)
			switch wt {
			case csproto.WireTypeFixed64:
				if v, err := dec.DecodeFixed64(); err != nil {
					return fmt.Errorf("unable to decode uint64 value for field 'theFixed64s' (tag=11): %w", err)
				} else {
					m.TheFixed64S = append(m.TheFixed64S, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedFixed64(); err != nil {
					return fmt.Errorf("unable to decode packed uint64 values for field 'theFixed64s' (tag=11): %w", err)
				} else {
					m.TheFixed64S = append(m.TheFixed64S, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theFixed64s' (tag=11), expected 1 (64-bit) or 1 (length-delimited)", wt)
			}
		case 12: // TheSFixed32S (sfixed32,repeated,packed)
			if wt != csproto.WireTypeFixed32 {
				return fmt.Errorf("incorrect wire type %v for field 'theSFixed32s' (tag=12), expected 5 (32-bit)", wt)
			}
			if v, err := dec.DecodeFixed32(); err != nil {
				return fmt.Errorf("unable to decode sfixed32 value for field 'theSFixed32s' (tag=12): %w", err)
			} else {
				m.TheSFixed32S = append(m.TheSFixed32S, int32(v))
			}

		case 13: // TheSFixed64S (sfixed64,repeated,packed)
			if wt != csproto.WireTypeFixed64 {
				return fmt.Errorf("incorrect wire type %v for field 'theSFixed64s' (tag=13), expected 1 (64-bit)", wt)
			}
			if v, err := dec.DecodeFixed64(); err != nil {
				return fmt.Errorf("unable to decode sfixed64 value for field 'theSFixed64s' (tag=13): %w", err)
			} else {
				m.TheSFixed64S = append(m.TheSFixed64S, int64(v))
			}

		case 14: // TheFloats (float,repeated,packed)
			switch wt {
			case csproto.WireTypeFixed32:
				if v, err := dec.DecodeFloat32(); err != nil {
					return fmt.Errorf("unable to decode float value for field 'theFloats' (tag=14): %w", err)
				} else {
					m.TheFloats = append(m.TheFloats, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedFloat32(); err != nil {
					return fmt.Errorf("unable to decode packed float values for field 'theFloats' (tag=14): %w", err)
				} else {
					m.TheFloats = append(m.TheFloats, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theFloats' (tag=14), expected 5 (32-bit) or 1 (length-delimited)", wt)
			}
		case 15: // TheDoubles (double,repeated,packed)
			switch wt {
			case csproto.WireTypeFixed64:
				if v, err := dec.DecodeFloat64(); err != nil {
					return fmt.Errorf("unable to decode double value for field 'theDoubles' (tag=15): %w", err)
				} else {
					m.TheDoubles = append(m.TheDoubles, v)
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedFloat64(); err != nil {
					return fmt.Errorf("unable to decode packed double values for field 'theDoubles' (tag=15): %w", err)
				} else {
					m.TheDoubles = append(m.TheDoubles, v...)
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theDoubles' (tag=15), expected 1 (64-bit) or 1 (length-delimited)", wt)
			}
		case 16: // TheEventTypes (enum,repeated,packed)
			switch wt {
			case csproto.WireTypeVarint:
				if v, err := dec.DecodeInt32(); err != nil {
					return fmt.Errorf("unable to decode int32 enum value for field 'theEventTypes' (tag=16): %w", err)
				} else {
					m.TheEventTypes = append(m.TheEventTypes, EventType(v))
				}
			case csproto.WireTypeLengthDelimited:
				if v, err := dec.DecodePackedInt32(); err != nil {
					return fmt.Errorf("unable to decode packed int32 enum values for field 'theEventTypes' (tag=16): %w", err)
				} else {
					for _, elem := range v {
						m.TheEventTypes = append(m.TheEventTypes, EventType(elem))
					}
				}
			default:
				return fmt.Errorf("incorrect wire type %v for repeated field 'theEventTypes' (tag=16), expected 0 (varint) or 1 (length-delimited)", wt)
			}
		case 17: // TheBytes (bytes,repeated)

			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for field 'theBytes' (tag=17), expected 2 (length-delimited)", wt)
			}
			if b, err := dec.DecodeBytes(); err != nil {
				return fmt.Errorf("unable to decode bytes value for field 'theBytes' (tag=17): %w", err)
			} else {
				m.TheBytes = append(m.TheBytes, b)
			}

		case 18: // TheMessages (message,repeated)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for field 'theMessages' (tag=18), expected 2 (length-delimited)", wt)
			}
			var mm EmbeddedEvent
			if err = dec.DecodeNested(&mm); err != nil {
				return fmt.Errorf("unable to decode message value for field 'theMessages' (tag=18): %w", err)
			}
			m.TheMessages = append(m.TheMessages, &mm)

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
