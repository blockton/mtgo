package tg

import (
	"bytes"
	"testing"
)

func TestKeyboardButtonRequestPeerSchema(t *testing.T) {
	button := &KeyboardButton{
		Text: "Select channel",
		Type: &InputButtonTypeRequestPeer{
			UsernameRequested: true,
			ButtonID:          1,
			PeerType:          &RequestPeerTypeBroadcast{Creator: true},
			MaxQuantity:       1,
		},
	}

	var buf bytes.Buffer
	if err := button.Encode(&buf); err != nil {
		t.Fatalf("encode: %v", err)
	}

	r := NewReader(buf.Bytes())
	defer ReleaseReader(r)

	id, err := r.ReadUint32()
	if err != nil {
		t.Fatalf("read constructor: %v", err)
	}
	if id != KeyboardButtonTypeID {
		t.Fatalf("constructor id = %#08x, want %#08x", id, KeyboardButtonTypeID)
	}

	flags, err := r.ReadUint32()
	if err != nil {
		t.Fatalf("read flags: %v", err)
	}
	if got := Fields(flags).Has(10); got {
		t.Fatalf("style flag = %v, want false", got)
	}

	text, err := r.ReadString()
	if err != nil {
		t.Fatalf("read text: %v", err)
	}
	if text != "Select channel" {
		t.Fatalf("text = %q", text)
	}

	typeID, err := r.ReadUint32()
	if err != nil {
		t.Fatalf("read type constructor: %v", err)
	}
	if typeID != InputButtonTypeRequestPeerTypeID {
		t.Fatalf("type constructor = %#08x, want %#08x", typeID, InputButtonTypeRequestPeerTypeID)
	}

	typeFlags, err := r.ReadUint32()
	if err != nil {
		t.Fatalf("read type flags: %v", err)
	}
	if got, want := Fields(typeFlags).Has(1), true; got != want {
		t.Fatalf("username_requested flag = %v, want %v", got, want)
	}
}
