package proto2_test

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/proto"

	"github.com/CrowdStrike/csproto"
	"github.com/CrowdStrike/csproto/example/proto2/googlev2"
	"github.com/CrowdStrike/csproto/lazyproto"
)

func BenchmarkEncodeGoogleV2(b *testing.B) {
	var (
		evt = createGoogleV2Event(b)
		buf []byte
	)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buf, _ = proto.Marshal(evt)
		_ = buf
	}
	_ = buf
}

func BenchmarkCustomEncodeGoogleV2(b *testing.B) {
	var (
		evt = createGoogleV2Event(b)
		buf []byte
	)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buf, _ = csproto.Marshal(evt)
		_ = buf
	}
	_ = buf
}

func BenchmarkDecodeGoogleV2(b *testing.B) {
	var (
		evt  = createGoogleV2Event(b)
		evt2 googlev2.BaseEvent
	)
	data, _ := proto.Marshal(evt)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		evt2.Reset()
		_ = proto.Unmarshal(data, &evt2)
	}
}

func BenchmarkCustomDecodeGoogleV2(b *testing.B) {
	var (
		evt  = createGoogleV2Event(b)
		evt2 googlev2.BaseEvent
	)
	data, _ := proto.Marshal(evt)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		evt2.Reset()
		_ = csproto.Unmarshal(data, &evt2)
	}
}

func BenchmarkLazyDecodeGoogleV2(b *testing.B) {
	var (
		evt = createGoogleV2Event(b)
		def = lazyproto.NewDef(1, 2, 3, 4)
	)
	_ = def.NestedTag(100, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	data, _ := proto.Marshal(evt)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r, _ := lazyproto.Decode(data, def) //nolint: staticcheck // benchmarking deprecated function to demonstrate the difference
		_ = r.Close()
	}
}

func BenchmarkLazyDecoder(b *testing.B) {
	var (
		evt = createGoogleV2Event(b)
		def = lazyproto.NewDef(1, 2, 3, 4)
	)
	_ = def.NestedTag(100, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	data, _ := proto.Marshal(evt)
	b.Run("safe", func(b *testing.B) {
		dec, _ := lazyproto.NewDecoder(def, lazyproto.WithMaxBufferSize(3))
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			r, _ := dec.Decode(data)
			nest, _ := r.NestedResult(100)
			discardStrings, _ = nest.StringValues(4)
			_ = r.Close()
		}
	})
	b.Run("unsafe", func(b *testing.B) {
		dec, _ := lazyproto.NewDecoder(def, lazyproto.WithMode(csproto.DecoderModeFast), lazyproto.WithMaxBufferSize(3))
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			r, _ := dec.Decode(data)
			nest, _ := r.NestedResult(100)
			discardStrings, _ = nest.StringValues(4)
			_ = r.Close()
		}
	})
}

var discardStrings []string

func createGoogleV2Event(t interface{ Errorf(string, ...interface{}) }) *googlev2.BaseEvent {
	eventType := googlev2.EventType_EVENT_TYPE_ONE
	baseEvent := googlev2.BaseEvent{
		EventID:   csproto.String(uuid.Must(uuid.NewV4()).String()),
		SourceID:  csproto.String(uuid.Must(uuid.NewV4()).String()),
		Timestamp: csproto.Uint64(uint64(time.Now().UTC().Unix())),
		EventType: &eventType,
	}
	extEvent := googlev2.TestEvent{
		Name:      csproto.String("test-event"),
		Info:      csproto.String("blah blah blah"),
		IsAwesome: csproto.Bool(true),
		Labels:    []string{"one", "two", "three"},
		Embedded: &googlev2.EmbeddedEvent{
			ID:              csproto.Int32(1),
			FavoriteNumbers: []int32{42, 1138},
			RandomThings: [][]byte{
				{0x0, 0x1, 0x2, 0x3, 0x4, 0x5},
				{0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf},
			},
		},
		Path: &googlev2.TestEvent_Jedi{
			Jedi: true,
		},
	}
	if err := csproto.SetExtension(&baseEvent, googlev2.E_TestEvent_EventExt, &extEvent); err != nil {
		t.Errorf("unable to set proto2 extension: %v", err)
	}
	return &baseEvent
}
