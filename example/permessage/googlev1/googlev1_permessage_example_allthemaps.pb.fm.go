// GENERATED CODE - DO NOT EDIT
// This file was generated by protoc-gen-fastmarshal

package googlev1

import (
	"fmt"
	"sync/atomic"
	"github.com/CrowdStrike/csproto"
)

//------------------------------------------------------------------------------
// Custom Protobuf size/marshal/unmarshal code for AllTheMaps

// Size calculates and returns the size, in bytes, required to hold the contents of m using the Protobuf
// binary encoding.
func (m *AllTheMaps) Size() int {
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

	// ToInt32 (message,repeated)
	for k, v := range m.ToInt32 {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		valueSize := 1 + csproto.SizeOfVarint(uint64(v))
		sz += csproto.SizeOfTagKey(1) + csproto.SizeOfVarint(uint64(keySize+valueSize)) + keySize + valueSize
	}

	// ToInt64 (message,repeated)
	for k, v := range m.ToInt64 {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		valueSize := 1 + csproto.SizeOfVarint(uint64(v))
		sz += csproto.SizeOfTagKey(2) + csproto.SizeOfVarint(uint64(keySize+valueSize)) + keySize + valueSize
	}

	// ToUInt32 (message,repeated)
	for k, v := range m.ToUInt32 {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		valueSize := 1 + csproto.SizeOfVarint(uint64(v))
		sz += csproto.SizeOfTagKey(3) + csproto.SizeOfVarint(uint64(keySize+valueSize)) + keySize + valueSize
	}

	// ToUInt64 (message,repeated)
	for k, v := range m.ToUInt64 {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		valueSize := 1 + csproto.SizeOfVarint(uint64(v))
		sz += csproto.SizeOfTagKey(4) + csproto.SizeOfVarint(uint64(keySize+valueSize)) + keySize + valueSize
	}

	// ToString (message,repeated)
	for k, v := range m.ToString {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		l = len(v)
		valueSize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		sz += csproto.SizeOfTagKey(5) + csproto.SizeOfVarint(uint64(keySize+valueSize)) + keySize + valueSize
	}

	// ToBytes (message,repeated)
	for k, v := range m.ToBytes {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		l = len(v)
		valueSize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		sz += csproto.SizeOfTagKey(6) + csproto.SizeOfVarint(uint64(keySize+valueSize)) + keySize + valueSize
	}

	// ToSInt32 (message,repeated)
	for k, v := range m.ToSInt32 {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		valueSize := 1 + csproto.SizeOfZigZag(uint64(v))
		sz += csproto.SizeOfTagKey(7) + csproto.SizeOfVarint(uint64(keySize+valueSize)) + keySize + valueSize
	}

	// ToSInt64 (message,repeated)
	for k, v := range m.ToSInt64 {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		valueSize := 1 + csproto.SizeOfZigZag(uint64(v))
		sz += csproto.SizeOfTagKey(8) + csproto.SizeOfVarint(uint64(keySize+valueSize)) + keySize + valueSize
	}

	// ToFixed32 (message,repeated)
	for k, v := range m.ToFixed32 {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		_ = v
		sz += csproto.SizeOfTagKey(9) + csproto.SizeOfVarint(uint64(keySize+5)) + keySize + 5
	}

	// ToSFixed32 (message,repeated)
	for k, v := range m.ToSFixed32 {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		_ = v
		sz += csproto.SizeOfTagKey(10) + csproto.SizeOfVarint(uint64(keySize+5)) + keySize + 5
	}

	// ToFixed64 (message,repeated)
	for k, v := range m.ToFixed64 {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		_ = v
		sz += csproto.SizeOfTagKey(11) + csproto.SizeOfVarint(uint64(keySize+9)) + keySize + 9
	}

	// ToSFixed64 (message,repeated)
	for k, v := range m.ToSFixed64 {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		_ = v
		sz += csproto.SizeOfTagKey(12) + csproto.SizeOfVarint(uint64(keySize+9)) + keySize + 9
	}

	// ToFloat (message,repeated)
	for k, v := range m.ToFloat {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		_ = v
		sz += csproto.SizeOfTagKey(13) + csproto.SizeOfVarint(uint64(keySize+5)) + keySize + 5
	}

	// ToDouble (message,repeated)
	for k, v := range m.ToDouble {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		_ = v
		sz += csproto.SizeOfTagKey(14) + csproto.SizeOfVarint(uint64(keySize+9)) + keySize + 9
	}

	// ToMessage (message,repeated)
	for k, v := range m.ToMessage {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		if v != nil {
			l = csproto.Size(v)
			valueSize := 1 + csproto.SizeOfVarint(uint64(l)) + l
			sz += csproto.SizeOfTagKey(15) + csproto.SizeOfVarint(uint64(keySize+valueSize)) + keySize + valueSize
		}
	}

	// ToEnum (message,repeated)
	for k, v := range m.ToEnum {
		// size of key (always has an internal tag of 1)
		l = len(k)
		keySize := 1 + csproto.SizeOfVarint(uint64(l)) + l
		// size of value (always has an internal tag of 2)
		valueSize := 1 + csproto.SizeOfVarint(uint64(v))
		sz += csproto.SizeOfTagKey(16) + csproto.SizeOfVarint(uint64(keySize+valueSize)) + keySize + valueSize
	}

	// cache the size so it can be re-used in Marshal()/MarshalTo()
	atomic.StoreInt32(&m.sizeCache, int32(sz))
	return sz
}

// Marshal converts the contents of m to the Protobuf binary encoding and returns the result or an error.
func (m *AllTheMaps) Marshal() ([]byte, error) {
	siz := m.Size()
	if siz == 0 {
		return []byte{}, nil
	}
	buf := make([]byte, siz)
	err := m.MarshalTo(buf)
	return buf, err
}

// MarshalTo converts the contents of m to the Protobuf binary encoding and writes the result to dest.
func (m *AllTheMaps) MarshalTo(dest []byte) error {
	// nil message == no-op
	if m == nil {
		return nil
	}
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

	// ToInt32 (1,map)
	for k, v := range m.ToInt32 {
		itemSize := 1 + csproto.SizeOfVarint(uint64(v))
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(1, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeInt32(2, v)
	}

	// ToInt64 (2,map)
	for k, v := range m.ToInt64 {
		itemSize := 1 + csproto.SizeOfVarint(uint64(v))
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(2, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeInt64(2, v)
	}

	// ToUInt32 (3,map)
	for k, v := range m.ToUInt32 {
		itemSize := 1 + csproto.SizeOfVarint(uint64(v))
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(3, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeUInt32(2, v)
	}

	// ToUInt64 (4,map)
	for k, v := range m.ToUInt64 {
		itemSize := 1 + csproto.SizeOfVarint(uint64(v))
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(4, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeUInt64(2, v)
	}

	// ToString (5,map)
	for k, v := range m.ToString {
		valueSize := len(v)
		itemSize := 1 + csproto.SizeOfVarint(uint64(valueSize)) + valueSize
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(5, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeString(2, v)
	}

	// ToBytes (6,map)
	for k, v := range m.ToBytes {
		valueSize := len(v)
		itemSize := 1 + csproto.SizeOfVarint(uint64(valueSize)) + valueSize
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(6, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeBytes(2, v)
	}

	// ToSInt32 (7,map)
	for k, v := range m.ToSInt32 {
		itemSize := 1 + csproto.SizeOfZigZag(uint64(v))
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(7, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeSInt32(2, v)
	}

	// ToSInt64 (8,map)
	for k, v := range m.ToSInt64 {
		itemSize := 1 + csproto.SizeOfZigZag(uint64(v))
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(8, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeSInt64(2, v)
	}

	// ToFixed32 (9,map)
	for k, v := range m.ToFixed32 {
		itemSize := 5
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(9, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeFixed32(2, uint32(v))
	}

	// ToSFixed32 (10,map)
	for k, v := range m.ToSFixed32 {
		itemSize := 5
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(10, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeFixed32(2, uint32(v))
	}

	// ToFixed64 (11,map)
	for k, v := range m.ToFixed64 {
		itemSize := 9
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(11, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeFixed64(2, uint64(v))
	}

	// ToSFixed64 (12,map)
	for k, v := range m.ToSFixed64 {
		itemSize := 9
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(12, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeFixed64(2, uint64(v))
	}

	// ToFloat (13,map)
	for k, v := range m.ToFloat {
		itemSize := 5
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(13, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeFloat32(2, v)
	}

	// ToDouble (14,map)
	for k, v := range m.ToDouble {
		itemSize := 9
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(14, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeFloat64(2, v)
	}

	// ToMessage (15,map)
	for k, v := range m.ToMessage {
		if v == nil {
			continue
		}
		valueSize := csproto.Size(v)
		itemSize := 1 + csproto.SizeOfVarint(uint64(valueSize)) + valueSize
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(15, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeNested(2, v)
	}

	// ToEnum (16,map)
	for k, v := range m.ToEnum {
		itemSize := 1 + csproto.SizeOfVarint(uint64(v))
		keySize := len(k)
		itemSize += 1 + csproto.SizeOfVarint(uint64(keySize)) + keySize
		enc.EncodeMapEntryHeader(16, itemSize)
		enc.EncodeString(1, k)
		enc.EncodeInt32(2, int32(v))
	}

	return nil
}

// Unmarshal decodes a binary encoded Protobuf message from p and populates m with the result.
func (m *AllTheMaps) Unmarshal(p []byte) error {
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
		case 1: // ToInt32 (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toInt32' (tag=1), expected 2 (length-delimited)", wt)
			}

			if m.ToInt32 == nil {
				m.ToInt32 = make(map[string]int32)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue int32
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toInt32' (tag=1), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeVarint {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toInt32' (tag=1), expected 0 (varint)", ewt)
					}
					if entryValue, err = dec.DecodeInt32(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToInt32[entryKey] = entryValue
		case 2: // ToInt64 (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toInt64' (tag=2), expected 2 (length-delimited)", wt)
			}

			if m.ToInt64 == nil {
				m.ToInt64 = make(map[string]int64)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue int64
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toInt64' (tag=2), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeVarint {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toInt64' (tag=2), expected 0 (varint)", ewt)
					}
					if entryValue, err = dec.DecodeInt64(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToInt64[entryKey] = entryValue
		case 3: // ToUInt32 (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toUInt32' (tag=3), expected 2 (length-delimited)", wt)
			}

			if m.ToUInt32 == nil {
				m.ToUInt32 = make(map[string]uint32)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue uint32
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toUInt32' (tag=3), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeVarint {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toUInt32' (tag=3), expected 0 (varint)", ewt)
					}
					if entryValue, err = dec.DecodeUInt32(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToUInt32[entryKey] = entryValue
		case 4: // ToUInt64 (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toUInt64' (tag=4), expected 2 (length-delimited)", wt)
			}

			if m.ToUInt64 == nil {
				m.ToUInt64 = make(map[string]uint64)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue uint64
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toUInt64' (tag=4), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeVarint {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toUInt64' (tag=4), expected 0 (varint)", ewt)
					}
					if entryValue, err = dec.DecodeUInt64(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToUInt64[entryKey] = entryValue
		case 5: // ToString (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toString' (tag=5), expected 2 (length-delimited)", wt)
			}

			if m.ToString == nil {
				m.ToString = make(map[string]string)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue string
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toString' (tag=5), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toString' (tag=5), expected 2 (length-delimited)", ewt)
					}
					if entryValue, err = dec.DecodeString(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToString[entryKey] = entryValue
		case 6: // ToBytes (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toBytes' (tag=6), expected 2 (length-delimited)", wt)
			}

			if m.ToBytes == nil {
				m.ToBytes = make(map[string][]byte)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue []byte
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toBytes' (tag=6), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toBytes' (tag=6), expected 2 (length-delimited)", ewt)
					}
					if entryValue, err = dec.DecodeBytes(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToBytes[entryKey] = entryValue
		case 7: // ToSInt32 (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toSInt32' (tag=7), expected 2 (length-delimited)", wt)
			}

			if m.ToSInt32 == nil {
				m.ToSInt32 = make(map[string]int32)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue int32
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toSInt32' (tag=7), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeVarint {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toSInt32' (tag=7), expected 0 (varint)", ewt)
					}
					if entryValue, err = dec.DecodeSInt32(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToSInt32[entryKey] = entryValue
		case 8: // ToSInt64 (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toSInt64' (tag=8), expected 2 (length-delimited)", wt)
			}

			if m.ToSInt64 == nil {
				m.ToSInt64 = make(map[string]int64)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue int64
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toSInt64' (tag=8), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeVarint {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toSInt64' (tag=8), expected 0 (varint)", ewt)
					}
					if entryValue, err = dec.DecodeSInt64(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToSInt64[entryKey] = entryValue
		case 9: // ToFixed32 (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toFixed32' (tag=9), expected 2 (length-delimited)", wt)
			}

			if m.ToFixed32 == nil {
				m.ToFixed32 = make(map[string]uint32)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue uint32
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toFixed32' (tag=9), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeFixed32 {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toFixed32' (tag=9), expected 5 (fixed32)", ewt)
					}
					if entryValue, err = dec.DecodeFixed32(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToFixed32[entryKey] = entryValue
		case 10: // ToSFixed32 (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toSFixed32' (tag=10), expected 2 (length-delimited)", wt)
			}

			if m.ToSFixed32 == nil {
				m.ToSFixed32 = make(map[string]int32)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue int32
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toSFixed32' (tag=10), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeFixed32 {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toSFixed32' (tag=10), expected 5 (fixed32)", ewt)
					}
					if sv, err := dec.DecodeFixed32(); err != nil {
						return err
					} else {
						entryValue = int32(sv)
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToSFixed32[entryKey] = entryValue
		case 11: // ToFixed64 (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toFixed64' (tag=11), expected 2 (length-delimited)", wt)
			}

			if m.ToFixed64 == nil {
				m.ToFixed64 = make(map[string]uint64)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue uint64
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toFixed64' (tag=11), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeFixed64 {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toFixed64' (tag=11), expected 1 (fixed64)", ewt)
					}
					if entryValue, err = dec.DecodeFixed64(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToFixed64[entryKey] = entryValue
		case 12: // ToSFixed64 (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toSFixed64' (tag=12), expected 2 (length-delimited)", wt)
			}

			if m.ToSFixed64 == nil {
				m.ToSFixed64 = make(map[string]int64)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue int64
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toSFixed64' (tag=12), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeFixed64 {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toSFixed64' (tag=12), expected 1 (fixed64)", ewt)
					}
					if sv, err := dec.DecodeFixed64(); err != nil {
						return err
					} else {
						entryValue = int64(sv)
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToSFixed64[entryKey] = entryValue
		case 13: // ToFloat (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toFloat' (tag=13), expected 2 (length-delimited)", wt)
			}

			if m.ToFloat == nil {
				m.ToFloat = make(map[string]float32)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue float32
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toFloat' (tag=13), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeFixed32 {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toFloat' (tag=13), expected 5 (fixed32)", ewt)
					}
					if entryValue, err = dec.DecodeFloat32(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToFloat[entryKey] = entryValue
		case 14: // ToDouble (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toDouble' (tag=14), expected 2 (length-delimited)", wt)
			}

			if m.ToDouble == nil {
				m.ToDouble = make(map[string]float64)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue float64
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toDouble' (tag=14), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeFixed64 {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toDouble' (tag=14), expected 1 (fixed64)", ewt)
					}
					if entryValue, err = dec.DecodeFloat64(); err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToDouble[entryKey] = entryValue
		case 15: // ToMessage (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toMessage' (tag=15), expected 2 (length-delimited)", wt)
			}

			if m.ToMessage == nil {
				m.ToMessage = make(map[string]*EmbeddedEvent)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue *EmbeddedEvent
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toMessage' (tag=15), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toMessage' (tag=15), expected 2 (length-delimited)", ewt)
					}
					var v EmbeddedEvent
					if err = dec.DecodeNested(&v); err != nil {
						return err
					} else {
						entryValue = &v
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToMessage[entryKey] = entryValue
		case 16: // ToEnum (map)
			if wt != csproto.WireTypeLengthDelimited {
				return fmt.Errorf("incorrect wire type %v for map field 'toEnum' (tag=16), expected 2 (length-delimited)", wt)
			}

			if m.ToEnum == nil {
				m.ToEnum = make(map[string]EventType)
			}
			// consume the map entry size
			// TODO - should we validate this?
			if _, err = dec.DecodeInt32(); err != nil {
				return err
			}
			// always 2 values
			var (
				entryKey   string
				entryValue EventType
			)
			for i := 0; i < 2; i++ {
				etag, ewt, err := dec.DecodeTag()
				if err != nil {
					return err
				}
				switch etag {
				case 1: // key
					if ewt != csproto.WireTypeLengthDelimited {
						return fmt.Errorf("incorrect wire type %v for map key for field 'toEnum' (tag=16), expected 2 (length-delimited)", ewt)
					}
					if entryKey, err = dec.DecodeString(); err != nil {
						return err
					}
				case 2: // value
					if ewt != csproto.WireTypeVarint {
						return fmt.Errorf("incorrect wire type %v for map value for field 'toEnum' (tag=16), expected 0 (varint)", ewt)
					}
					if v, err := dec.DecodeInt32(); err != nil {
						return err
					} else {
						entryValue = EventType(v)
					}
				default:
					return fmt.Errorf("invalid map entry field tag %d, expected 1 or 2", etag)
				}
			}
			m.ToEnum[entryKey] = entryValue

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
