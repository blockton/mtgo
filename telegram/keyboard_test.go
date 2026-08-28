package telegram

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/mtgo-labs/mtgo/tg"
)

func TestKeyboardEmpty(t *testing.T) {
	kb := Keyboard()
	if kb.Build() != nil {
		t.Error("empty builder should return nil")
	}
	if kb.BuildReply() != nil {
		t.Error("empty builder should return nil for reply")
	}
}

func TestKeyboardCallback(t *testing.T) {
	markup := Keyboard().
		Callback("Click", "data123").
		Build()

	inner, ok := markup.(*tg.ReplyInlineMarkup)
	if !ok {
		t.Fatal("expected ReplyInlineMarkup")
	}
	if len(inner.Rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(inner.Rows))
	}
	if len(inner.Rows[0].Buttons) != 1 {
		t.Fatalf("expected 1 button, got %d", len(inner.Rows[0].Buttons))
	}
	btn := inner.Rows[0].Buttons[0]
	if btn.Text != "Click" {
		t.Errorf("text = %q, want %q", btn.Text, "Click")
	}
	cb, ok := btn.Type.(*tg.InlineButtonTypeCallback)
	if !ok {
		t.Fatalf("expected InlineButtonTypeCallback, got %T", btn.Type)
	}
	if string(cb.Data) != "data123" {
		t.Errorf("data = %q, want %q", cb.Data, "data123")
	}
}

func TestKeyboardCallbackTruncation(t *testing.T) {
	longData := make([]byte, 128)
	for i := range longData {
		longData[i] = 'x'
	}
	markup := Keyboard().
		Callback("Btn", string(longData)).
		Build()

	inner := markup.(*tg.ReplyInlineMarkup)
	btn := inner.Rows[0].Buttons[0]
	if cb, ok := btn.Type.(*tg.InlineButtonTypeCallback); !ok || len(cb.Data) != 64 {
		t.Errorf("data length = %d, want 64", len(cb.Data))
	}
}

func TestKeyboardURL(t *testing.T) {
	markup := Keyboard().URL("Link", "https://example.com").Build()
	inner := markup.(*tg.ReplyInlineMarkup)
	btn := inner.Rows[0].Buttons[0]
	u, ok := btn.Type.(*tg.InlineButtonTypeURL)
	if !ok {
		t.Fatalf("expected InlineButtonTypeURL, got %T", btn.Type)
	}
	if btn.Text != "Link" || u.URL != "https://example.com" {
		t.Errorf("got text=%q url=%q", btn.Text, u.URL)
	}
}

func TestKeyboardTextReply(t *testing.T) {
	markup := Keyboard().
		Text("A").
		Text("B").
		BuildReply(ReplyOpts{Resize: true, OneTime: true})

	inner, ok := markup.(*tg.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("expected ReplyKeyboardMarkup")
	}
	if !inner.Resize || !inner.SingleUse {
		t.Error("expected Resize=true, SingleUse=true")
	}
	if len(inner.Rows) != 1 || len(inner.Rows[0].Buttons) != 2 {
		t.Fatalf("expected 1 row with 2 buttons, got %d rows", len(inner.Rows))
	}
}

func TestKeyboardNext(t *testing.T) {
	markup := Keyboard().
		Callback("R1A", "a").
		Callback("R1B", "b").
		Next().
		Callback("R2A", "c").
		Build()

	inner := markup.(*tg.ReplyInlineMarkup)
	if len(inner.Rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(inner.Rows))
	}
	if len(inner.Rows[0].Buttons) != 2 {
		t.Errorf("row 0: expected 2 buttons, got %d", len(inner.Rows[0].Buttons))
	}
	if len(inner.Rows[1].Buttons) != 1 {
		t.Errorf("row 1: expected 1 button, got %d", len(inner.Rows[1].Buttons))
	}
}

func TestKeyboardRow(t *testing.T) {
	markup := Keyboard().
		InlineRow(
			&tg.KeyboardInlineButton{Text: "A", Type: &tg.InlineButtonTypeCallback{Data: []byte("a")}},
			&tg.KeyboardInlineButton{Text: "B", Type: &tg.InlineButtonTypeCallback{Data: []byte("b")}},
		).
		InlineRow(
			&tg.KeyboardInlineButton{Text: "C", Type: &tg.InlineButtonTypeCallback{Data: []byte("c")}},
		).
		Build()

	inner := markup.(*tg.ReplyInlineMarkup)
	if len(inner.Rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(inner.Rows))
	}
}

func TestKeyboardReplyPlaceholder(t *testing.T) {
	markup := Keyboard().
		Text("OK").
		BuildReply(ReplyOpts{Placeholder: "Pick one"})

	inner := markup.(*tg.ReplyKeyboardMarkup)
	if inner.Placeholder != "Pick one" {
		t.Error("expected placeholder to be set")
	}
}

func TestKeyboardMultipleTypes(t *testing.T) {
	markup := Keyboard().
		Callback("Cb", "cb").
		URL("Link", "https://example.com").
		Copy("Copy", "text").
		Game("Play").
		Build()

	inner := markup.(*tg.ReplyInlineMarkup)
	if len(inner.Rows[0].Buttons) != 4 {
		t.Fatalf("expected 4 buttons, got %d", len(inner.Rows[0].Buttons))
	}
	if _, ok := inner.Rows[0].Buttons[0].Type.(*tg.InlineButtonTypeCallback); !ok {
		t.Error("button 0 should be Callback")
	}
	if _, ok := inner.Rows[0].Buttons[1].Type.(*tg.InlineButtonTypeURL); !ok {
		t.Error("button 1 should be URL")
	}
	if _, ok := inner.Rows[0].Buttons[2].Type.(*tg.InlineButtonTypeCopy); !ok {
		t.Error("button 2 should be Copy")
	}
	if _, ok := inner.Rows[0].Buttons[3].Type.(*tg.InlineButtonTypeGame); !ok {
		t.Error("button 3 should be Game")
	}
}

func TestKeyboardReplyButtons(t *testing.T) {
	markup := Keyboard().
		Text("A").
		RequestPhone("Phone").
		RequestGeo("Location").
		RequestPoll("Poll", false).
		BuildReply()

	inner := markup.(*tg.ReplyKeyboardMarkup)
	btns := inner.Rows[0].Buttons
	if _, ok := btns[0].Type.(*tg.ButtonTypeDefault); !ok {
		t.Error("button 0 should be Text")
	}
	if _, ok := btns[1].Type.(*tg.ButtonTypeRequestPhone); !ok {
		t.Error("button 1 should be RequestPhone")
	}
	if _, ok := btns[2].Type.(*tg.ButtonTypeRequestGeoLocation); !ok {
		t.Error("button 2 should be RequestGeo")
	}
	if _, ok := btns[3].Type.(*tg.ButtonTypeRequestPoll); !ok {
		t.Error("button 3 should be RequestPoll")
	}
}

func TestForceReplyMarkup(t *testing.T) {
	m := ForceReplyMarkup()
	if m == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRemoveKeyboard(t *testing.T) {
	m := RemoveKeyboard()
	if m == nil {
		t.Fatal("expected non-nil")
	}
}

func TestKeyboardSwitch(t *testing.T) {
	markup := Keyboard().Switch("Share", false, "query").Build()
	inner := markup.(*tg.ReplyInlineMarkup)
	btn := inner.Rows[0].Buttons[0]
	sw, ok := btn.Type.(*tg.InlineButtonTypeSwitchInline)
	if !ok {
		t.Fatalf("expected InlineButtonTypeSwitchInline, got %T", btn.Type)
	}
	if btn.Text != "Share" || sw.Query != "query" || sw.SamePeer {
		t.Errorf("unexpected switch button: %+v", sw)
	}
}

func TestKeyboardWebApp(t *testing.T) {
	markup := Keyboard().WebApp("Open", "https://app.com").Build()
	inner := markup.(*tg.ReplyInlineMarkup)
	btn := inner.Rows[0].Buttons[0]
	wv, ok := btn.Type.(*tg.InlineButtonTypeWebView)
	if !ok {
		t.Fatalf("expected InlineButtonTypeWebView, got %T", btn.Type)
	}
	if btn.Text != "Open" || wv.URL != "https://app.com" {
		t.Errorf("unexpected webapp button: %+v", wv)
	}
}

func TestKeyboardBuy(t *testing.T) {
	markup := Keyboard().Buy("Pay").Build()
	inner := markup.(*tg.ReplyInlineMarkup)
	if _, ok := inner.Rows[0].Buttons[0].Type.(*tg.InlineButtonTypeBuy); !ok {
		t.Error("expected Buy button")
	}
}

func TestKeyboardNextNoopOnEmpty(t *testing.T) {
	markup := Keyboard().Next().Next().Callback("A", "a").Build()
	inner := markup.(*tg.ReplyInlineMarkup)
	if len(inner.Rows) != 1 || len(inner.Rows[0].Buttons) != 1 {
		t.Error("Next() on empty row should be noop")
	}
}

func TestKeyboardRowEmpty(t *testing.T) {
	markup := Keyboard().InlineRow().Build()
	if markup != nil {
		t.Error("empty InlineRow() should produce nil Build()")
	}
}

func TestKeyboardRequestPeer(t *testing.T) {
	markup := Keyboard().
		RequestPeer("Channel", 1, &tg.RequestPeerTypeBroadcast{}, 1).
		RequestPeer("User", 2, &tg.RequestPeerTypeUser{}, 1).
		BuildReply(ReplyOpts{Resize: true})

	inner := markup.(*tg.ReplyKeyboardMarkup)
	btns := inner.Rows[0].Buttons
	if len(btns) != 2 {
		t.Fatalf("expected 2 buttons, got %d", len(btns))
	}
	ch, ok := btns[0].Type.(*tg.InputButtonTypeRequestPeer)
	if !ok {
		t.Fatalf("button 0 should be InputButtonTypeRequestPeer, got %T", btns[0].Type)
	}
	if btns[0].Text != "Channel" || ch.ButtonID != 1 {
		t.Errorf("got text=%q buttonID=%d", btns[0].Text, ch.ButtonID)
	}
	usr := btns[1].Type.(*tg.InputButtonTypeRequestPeer)
	if usr.ButtonID != 2 {
		t.Errorf("got buttonID=%d", usr.ButtonID)
	}
}

func TestKeyboardRequestPeerEncodesInputButtonAndBoolFilter(t *testing.T) {
	markup := Keyboard().
		RequestPeer("Bot", 4, &tg.RequestPeerTypeUser{Bot: true}, 1).
		BuildReply()

	btn := markup.(*tg.ReplyKeyboardMarkup).Rows[0].Buttons[0]
	var buf bytes.Buffer
	if err := btn.Encode(&buf); err != nil {
		t.Fatalf("Encode() error: %v", err)
	}

	data := buf.Bytes()
	if got := binary.LittleEndian.Uint32(data[:4]); got != tg.KeyboardButtonTypeID {
		t.Fatalf("constructor = 0x%08x, want 0x%08x", got, tg.KeyboardButtonTypeID)
	}

	inputButton := make([]byte, 4)
	binary.LittleEndian.PutUint32(inputButton, tg.InputButtonTypeRequestPeerTypeID)
	if !bytes.Contains(data, inputButton) {
		t.Fatal("expected nested InputButtonTypeRequestPeer constructor")
	}

	boolTrue := make([]byte, 4)
	binary.LittleEndian.PutUint32(boolTrue, tg.BoolTrueID)
	if !bytes.Contains(data, boolTrue) {
		t.Fatal("expected requestPeerTypeUser Bot filter to encode boolTrue")
	}
}

func TestKeyboardStyleOnNilStyle(t *testing.T) {
	markup := Keyboard().
		Callback("Yes", "yes").
		Success().
		Build()

	inner := markup.(*tg.ReplyInlineMarkup)
	btn := inner.Rows[0].Buttons[0]
	if btn.Style == nil {
		t.Fatal("Style should be initialized after Success()")
	}
	if !btn.Style.BgSuccess {
		t.Error("BgSuccess should be true")
	}
}

func TestKeyboardRequestUser(t *testing.T) {
	markup := Keyboard().
		RequestUser("Pick User", 1, 5).
		BuildReply()

	btn := markup.(*tg.ReplyKeyboardMarkup).Rows[0].Buttons[0]
	peer := btn.Type.(*tg.InputButtonTypeRequestPeer)
	pt, ok := peer.PeerType.(*tg.RequestPeerTypeUser)
	if !ok {
		t.Fatal("expected RequestPeerTypeUser")
	}
	if pt.Bot {
		t.Error("RequestUser should not set Bot=true by default")
	}
	if peer.MaxQuantity != 5 {
		t.Errorf("MaxQuantity = %d, want 5", peer.MaxQuantity)
	}
}

func TestKeyboardRequestUserBot(t *testing.T) {
	markup := Keyboard().
		RequestUser("Pick Bot", 2, 1, PeerUserOpts{Bot: true, Premium: true}).
		BuildReply()

	btn := markup.(*tg.ReplyKeyboardMarkup).Rows[0].Buttons[0]
	peer := btn.Type.(*tg.InputButtonTypeRequestPeer)
	pt, ok := peer.PeerType.(*tg.RequestPeerTypeUser)
	if !ok {
		t.Fatal("expected RequestPeerTypeUser")
	}
	if !pt.Bot {
		t.Error("expected Bot=true")
	}
	if !pt.Premium {
		t.Error("expected Premium=true")
	}
}

func TestKeyboardRequestGroup(t *testing.T) {
	markup := Keyboard().
		RequestGroup("Pick Group", 3).
		BuildReply()

	btn := markup.(*tg.ReplyKeyboardMarkup).Rows[0].Buttons[0]
	peer := btn.Type.(*tg.InputButtonTypeRequestPeer)
	if _, ok := peer.PeerType.(*tg.RequestPeerTypeChat); !ok {
		t.Fatal("expected RequestPeerTypeChat")
	}
}

func TestKeyboardRequestGroupWithOptions(t *testing.T) {
	markup := Keyboard().
		RequestGroup("Pick Forum", 3, PeerGroupOpts{Creator: true, Forum: true, HasUsername: true}).
		BuildReply()

	btn := markup.(*tg.ReplyKeyboardMarkup).Rows[0].Buttons[0]
	peer := btn.Type.(*tg.InputButtonTypeRequestPeer)
	pt, ok := peer.PeerType.(*tg.RequestPeerTypeChat)
	if !ok {
		t.Fatal("expected RequestPeerTypeChat")
	}
	if !pt.Creator || !pt.Forum || !pt.HasUsername {
		t.Errorf("got creator=%v forum=%v hasUsername=%v", pt.Creator, pt.Forum, pt.HasUsername)
	}
}

func TestKeyboardRequestChannel(t *testing.T) {
	markup := Keyboard().
		RequestChannel("Pick Channel", 4).
		BuildReply()

	btn := markup.(*tg.ReplyKeyboardMarkup).Rows[0].Buttons[0]
	peer := btn.Type.(*tg.InputButtonTypeRequestPeer)
	if _, ok := peer.PeerType.(*tg.RequestPeerTypeBroadcast); !ok {
		t.Fatal("expected RequestPeerTypeBroadcast")
	}
}

func TestKeyboardRequestChannelWithOptions(t *testing.T) {
	markup := Keyboard().
		RequestChannel("Pick Channel", 4, PeerChannelOpts{Creator: true}).
		BuildReply()

	btn := markup.(*tg.ReplyKeyboardMarkup).Rows[0].Buttons[0]
	peer := btn.Type.(*tg.InputButtonTypeRequestPeer)
	pt, ok := peer.PeerType.(*tg.RequestPeerTypeBroadcast)
	if !ok {
		t.Fatal("expected RequestPeerTypeBroadcast")
	}
	if !pt.Creator {
		t.Error("expected Creator=true")
	}
}

func TestKeyboardMixedBuildDegradesButtons(t *testing.T) {
	inline := Keyboard().
		Text("Reply-only").
		Build().(*tg.ReplyInlineMarkup)
	btn := inline.Rows[0].Buttons[0]
	if _, ok := btn.Type.(*tg.InlineButtonTypeDisabled); !ok {
		t.Errorf("reply-only button in inline build should degrade to Disabled, got %T", btn.Type)
	}
	if btn.Text != "Reply-only" {
		t.Errorf("text = %q, want %q", btn.Text, "Reply-only")
	}

	reply := Keyboard().
		Callback("Inline-only", "x").
		BuildReply().(*tg.ReplyKeyboardMarkup)
	rbtn := reply.Rows[0].Buttons[0]
	if _, ok := rbtn.Type.(*tg.ButtonTypeDefault); !ok {
		t.Errorf("inline-only button in reply build should degrade to Default, got %T", rbtn.Type)
	}
	if rbtn.Text != "Inline-only" {
		t.Errorf("text = %q, want %q", rbtn.Text, "Inline-only")
	}
}
