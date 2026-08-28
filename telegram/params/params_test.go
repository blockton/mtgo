package params

import (
	"testing"

	tl "github.com/mtgo-labs/mtgo/tg"
)

// --- flatToSendMsg propagation tests ---------------------------------------
// Each media type implements getFlatSendFields + ToSendMsg. Verify that all 12
// shared fields propagate correctly through flatToSendMsg.

func makeFlatArgs() (bool, bool, bool, bool, bool, int32, tl.InputReplyToClass, tl.ReplyMarkupClass, *int32, *int64, tl.InputPeerClass, int32) {
	return true, true, true, true, true,
		42,
		&tl.InputReplyToMessage{ReplyToMsgID: 42},
		&tl.ReplyKeyboardMarkup{Rows: []*tl.KeyboardButtonRow{{Buttons: []*tl.KeyboardButton{{Text: "test"}}}}},
		new(int32(100)),
		new(int64(200)),
		&tl.InputPeerUser{UserID: 123, AccessHash: 456},
		789
}

func assertFlatFields(t *testing.T, got *SendMessage) {
	t.Helper()

	if !got.DisableNotification {
		t.Error("DisableNotification not propagated")
	}
	if !got.Silent {
		t.Error("Silent not propagated")
	}
	if !got.Background {
		t.Error("Background not propagated")
	}
	if !got.ClearDraft {
		t.Error("ClearDraft not propagated")
	}
	if !got.NoForwards {
		t.Error("NoForwards not propagated")
	}
	if got.ReplyToMessageID != 42 {
		t.Errorf("ReplyToMessageID = %d, want 42", got.ReplyToMessageID)
	}
	if got.ReplyTo == nil {
		t.Error("ReplyTo not propagated")
	}
	if got.ReplyMarkup == nil {
		t.Error("ReplyMarkup not propagated")
	}
	if got.ScheduleDate == nil || *got.ScheduleDate != 100 {
		t.Errorf("ScheduleDate = %v, want 100", got.ScheduleDate)
	}
	if got.EffectID == nil || *got.EffectID != 200 {
		t.Errorf("EffectID = %v, want 200", got.EffectID)
	}
	if got.SendAs == nil {
		t.Error("SendAs not propagated")
	}
	if got.MessageThreadID != 789 {
		t.Errorf("MessageThreadID = %d, want 789", got.MessageThreadID)
	}
}

func TestFlatToSendMsg_SendPoll(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendPoll{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendVenue(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendVenue{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendContact(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendContact{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendLocation(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendLocation{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendDice(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendDice{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendGame(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendGame{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendInlineBotResult(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendInlineBotResult{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendChecklist(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendChecklist{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendAudio(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendAudio{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendVideo(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendVideo{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendDocument(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendDocument{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendPhoto(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendPhoto{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendAnimation(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendAnimation{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendVoice(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendVoice{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendVideoNote(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendVideoNote{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

func TestFlatToSendMsg_SendSticker(t *testing.T) {
	dn, s, bg, cd, nf, rtm, rt, rm, sd, eid, sa, mtid := makeFlatArgs()
	v := &SendSticker{
		DisableNotification: dn, Silent: s, Background: bg, ClearDraft: cd, NoForwards: nf,
		ReplyToMessageID: rtm, ReplyTo: rt, ReplyMarkup: rm, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: mtid,
	}
	assertFlatFields(t, v.ToSendMsg())
}

// --- flatToSendMsg parse mode propagation ----------------------------------

func TestFlatToSendMsg_ParseModePropagation(t *testing.T) {
	tests := []struct {
		name         string
		construct    func() *SendMessage
		wantMode     ParseMode
		wantEntities bool // true = non-nil
	}{
		{
			"SendPhoto",
			func() *SendMessage {
				return (&SendPhoto{
					ParseMode:       ParseModeHTML,
					CaptionEntities: []tl.MessageEntityClass{&tl.MessageEntityBold{Offset: 0, Length: 1}},
				}).ToSendMsg()
			},
			ParseModeHTML, true,
		},
		{
			"SendVideo",
			func() *SendMessage {
				return (&SendVideo{
					ParseMode:       ParseModeMarkdown,
					CaptionEntities: []tl.MessageEntityClass{&tl.MessageEntityItalic{Offset: 0, Length: 1}},
				}).ToSendMsg()
			},
			ParseModeMarkdown, true,
		},
		{
			"SendAudio",
			func() *SendMessage {
				return (&SendAudio{
					ParseMode:       ParseModeHTML,
					CaptionEntities: nil,
				}).ToSendMsg()
			},
			ParseModeHTML, false,
		},
		{
			"SendDocument",
			func() *SendMessage {
				return (&SendDocument{
					ParseMode:       ParseModeMarkdown,
					CaptionEntities: []tl.MessageEntityClass{&tl.MessageEntityCode{Offset: 0, Length: 1}},
				}).ToSendMsg()
			},
			ParseModeMarkdown, true,
		},
		{
			"SendAnimation",
			func() *SendMessage {
				return (&SendAnimation{
					ParseMode:       ParseModeDisabled,
					CaptionEntities: nil,
				}).ToSendMsg()
			},
			ParseModeDisabled, false,
		},
		{
			"SendVoice",
			func() *SendMessage {
				return (&SendVoice{
					ParseMode:       ParseModeHTML,
					CaptionEntities: []tl.MessageEntityClass{&tl.MessageEntityTextURL{Offset: 0, Length: 1, URL: "https://example.com"}},
				}).ToSendMsg()
			},
			ParseModeHTML, true,
		},
		{
			"SendSticker",
			func() *SendMessage {
				return (&SendSticker{
					ParseMode:       ParseModeMarkdown,
					CaptionEntities: nil,
				}).ToSendMsg()
			},
			ParseModeMarkdown, false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.construct()
			if got.ParseMode != tt.wantMode {
				t.Errorf("ParseMode = %q, want %q", got.ParseMode, tt.wantMode)
			}
			if tt.wantEntities && got.Entities == nil {
				t.Error("Entities is nil, want non-nil")
			}
			if !tt.wantEntities && got.Entities != nil {
				t.Errorf("Entities = %v, want nil", got.Entities)
			}
		})
	}
}

func TestFlatToSendMsg_ZeroParseModePropagated(t *testing.T) {
	// When ParseMode is empty string and CaptionEntities nil, those zero values
	// should be propagated (empty string, nil slice) — not silently dropped.
	v := &SendPhoto{}
	got := v.ToSendMsg()
	if got.ParseMode != "" {
		t.Errorf("ParseMode = %q, want empty", got.ParseMode)
	}
	if got.Entities != nil {
		t.Errorf("Entities = %v, want nil", got.Entities)
	}
}

func TestFlatToSendMsg_NonMediaTypesPreserveDefaults(t *testing.T) {
	// Non-media types (SendPoll, SendVenue, etc.) don't implement parseModeProvider.
	// Their ToSendMsg should still work (ParseMode stays empty string).
	v := &SendPoll{
		DisableNotification: true,
		Silent:              true,
		ReplyToMessageID:    99,
	}
	got := v.ToSendMsg()
	if got.ParseMode != "" {
		t.Errorf("non-media type ParseMode = %q, want empty", got.ParseMode)
	}
	if got.Entities != nil {
		t.Errorf("non-media type Entities = %v, want nil", got.Entities)
	}
	if got.ReplyToMessageID != 99 {
		t.Errorf("ReplyToMessageID = %d, want 99", got.ReplyToMessageID)
	}
}

// --- SendMediaGroup special case: drops ReplyMarkup ------------------------

func TestFlatToSendMsg_SendMediaGroupDropsReplyMarkup(t *testing.T) {
	// SendMediaGroup.getFlatSendFields hardcodes nil for ReplyMarkup even
	// though the struct has the field — the Telegram API doesn't support
	// reply markup on media groups.
	sd := new(int32(100))
	eid := new(int64(200))
	sa := &tl.InputPeerUser{UserID: 123}
	v := &SendMediaGroup{
		DisableNotification: true, Silent: true, Background: true, ClearDraft: true, NoForwards: true,
		ReplyToMessageID: 42, ScheduleDate: sd, EffectID: eid, SendAs: sa, MessageThreadID: 789,
	}
	got := v.ToSendMsg()

	if !got.DisableNotification {
		t.Error("DisableNotification not propagated")
	}
	if got.ReplyToMessageID != 42 {
		t.Errorf("ReplyToMessageID = %d, want 42", got.ReplyToMessageID)
	}
	if got.MessageThreadID != 789 {
		t.Errorf("MessageThreadID = %d, want 789", got.MessageThreadID)
	}
	if got.ReplyMarkup != nil {
		t.Errorf("SendMediaGroup should drop ReplyMarkup, got %T", got.ReplyMarkup)
	}
}

// --- Zero-value defaults ----------------------------------------------------

func TestFlatToSendMsg_ZeroValues(t *testing.T) {
	v := &SendPhoto{}
	got := v.ToSendMsg()

	if got.DisableNotification || got.Silent || got.Background || got.ClearDraft || got.NoForwards {
		t.Error("bool fields should be false by default")
	}
	if got.ReplyToMessageID != 0 {
		t.Errorf("ReplyToMessageID = %d, want 0", got.ReplyToMessageID)
	}
	if got.ReplyTo != nil {
		t.Error("ReplyTo should be nil by default")
	}
	if got.ReplyMarkup != nil {
		t.Error("ReplyMarkup should be nil by default")
	}
	if got.ScheduleDate != nil {
		t.Error("ScheduleDate should be nil by default")
	}
	if got.EffectID != nil {
		t.Error("EffectID should be nil by default")
	}
	if got.SendAs != nil {
		t.Error("SendAs should be nil by default")
	}
	if got.MessageThreadID != 0 {
		t.Errorf("MessageThreadID = %d, want 0", got.MessageThreadID)
	}
}

// --- Keyboard helper tests --------------------------------------------------

func TestReplyKeyboardStructure(t *testing.T) {
	markup := ReplyKeyboard(
		[]KeyboardButton{Button("A"), Button("B")},
		[]KeyboardButton{Button("C")},
	)

	rkm, ok := markup.(*tl.ReplyKeyboardMarkup)
	if !ok {
		t.Fatalf("expected *ReplyKeyboardMarkup, got %T", markup)
	}
	if len(rkm.Rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(rkm.Rows))
	}

	row0 := rkm.Rows[0].Buttons
	if len(row0) != 2 {
		t.Fatalf("row 0: expected 2 buttons, got %d", len(row0))
	}
	assertButtonText(t, row0[0], "A")
	assertButtonText(t, row0[1], "B")

	row1 := rkm.Rows[1].Buttons
	if len(row1) != 1 {
		t.Fatalf("row 1: expected 1 button, got %d", len(row1))
	}
	assertButtonText(t, row1[0], "C")
}

func TestReplyKeyboardEmpty(t *testing.T) {
	markup := ReplyKeyboard()
	rkm, ok := markup.(*tl.ReplyKeyboardMarkup)
	if !ok {
		t.Fatalf("expected *ReplyKeyboardMarkup, got %T", markup)
	}
	if len(rkm.Rows) != 0 {
		t.Fatalf("expected 0 rows, got %d", len(rkm.Rows))
	}
}

func TestInlineKeyboardStructure(t *testing.T) {
	markup := InlineKeyboard(
		[]KeyboardButton{ButtonURL("Open", "https://example.com"), ButtonCB("Delete", "del")},
	)

	rim, ok := markup.(*tl.ReplyInlineMarkup)
	if !ok {
		t.Fatalf("expected *ReplyInlineMarkup, got %T", markup)
	}
	if len(rim.Rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rim.Rows))
	}
	buttons := rim.Rows[0].Buttons
	if len(buttons) != 2 {
		t.Fatalf("expected 2 buttons, got %d", len(buttons))
	}

	urlBtn := buttons[0]
	urlType, ok := urlBtn.Type.(*tl.InlineButtonTypeURL)
	if !ok {
		t.Fatalf("button 0: expected *InlineButtonTypeURL, got %T", urlBtn.Type)
	}
	if urlBtn.Text != "Open" || urlType.URL != "https://example.com" {
		t.Errorf("URL button: text=%q url=%q", urlBtn.Text, urlType.URL)
	}

	cbBtn := buttons[1]
	cbType, ok := cbBtn.Type.(*tl.InlineButtonTypeCallback)
	if !ok {
		t.Fatalf("button 1: expected *InlineButtonTypeCallback, got %T", cbBtn.Type)
	}
	if cbBtn.Text != "Delete" {
		t.Errorf("callback button text=%q, want Delete", cbBtn.Text)
	}
	if string(cbType.Data) != "del" {
		t.Errorf("callback button data=%q, want del", string(cbType.Data))
	}
}

func TestRemoveKeyboard(t *testing.T) {
	markup := RemoveKeyboard()
	if _, ok := markup.(*tl.ReplyKeyboardHide); !ok {
		t.Fatalf("expected *ReplyKeyboardHide, got %T", markup)
	}
}

func TestForceReplyKeyboard(t *testing.T) {
	markup := ForceReplyKeyboard()
	if _, ok := markup.(*tl.ReplyKeyboardForceReply); !ok {
		t.Fatalf("expected *ReplyKeyboardForceReply, got %T", markup)
	}
}

// --- Button type tests ------------------------------------------------------

func TestButtonVariants(t *testing.T) {
	tests := []struct {
		name string
		btn  KeyboardButton
		want tl.InlineButtonTypeClass
	}{
		{"plain", Button("Click"), &tl.InlineButtonTypeDisabled{}},
		{"url", ButtonURL("Go", "https://x.com"), &tl.InlineButtonTypeURL{URL: "https://x.com"}},
		{"callback", ButtonCB("OK", "ok"), &tl.InlineButtonTypeCallback{Data: []byte("ok")}},
		{"switch", ButtonSwitch("Share", "q"), &tl.InlineButtonTypeSwitchInline{Query: "q"}},
		{"game", ButtonGame("Play"), &tl.InlineButtonTypeGame{}},
		{"pay", ButtonPay("Buy"), &tl.InlineButtonTypeBuy{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.btn.toInlineTL()
			if got.Type.ConstructorID() != tc.want.ConstructorID() {
				t.Errorf("constructor: got 0x%x, want 0x%x", got.Type.ConstructorID(), tc.want.ConstructorID())
			}
		})
	}
}

func TestButtonFieldValues(t *testing.T) {
	b := Button("Hello")
	if b.Text != "Hello" || b.URL != "" || b.Data != nil || b.Switch != "" || b.Game || b.Pay {
		t.Errorf("Button fields incorrect: %+v", b)
	}

	bu := ButtonURL("Link", "https://example.com")
	if bu.Text != "Link" || bu.URL != "https://example.com" {
		t.Errorf("ButtonURL fields incorrect: %+v", bu)
	}

	bc := ButtonCB("Press", "data")
	if bc.Text != "Press" || string(bc.Data) != "data" {
		t.Errorf("ButtonCB fields incorrect: %+v", bc)
	}

	bs := ButtonSwitch("Share", "query")
	if bs.Text != "Share" || bs.Switch != "query" {
		t.Errorf("ButtonSwitch fields incorrect: %+v", bs)
	}

	bg := ButtonGame("Play")
	if bg.Text != "Play" || !bg.Game {
		t.Errorf("ButtonGame fields incorrect: %+v", bg)
	}

	bp := ButtonPay("Pay $5")
	if bp.Text != "Pay $5" || !bp.Pay {
		t.Errorf("ButtonPay fields incorrect: %+v", bp)
	}
}

// --- GetOptDef tests --------------------------------------------------------

func TestGetOptDef_Empty(t *testing.T) {
	got := GetOptDef("default")
	if got != "default" {
		t.Errorf("GetOptDef() with no opts = %q, want default", got)
	}
}

func TestGetOptDef_ValidOpt(t *testing.T) {
	got := GetOptDef("default", "custom")
	if got != "custom" {
		t.Errorf("GetOptDef() = %q, want custom", got)
	}
}

func TestGetOptDef_ZeroValueOpt(t *testing.T) {
	got := GetOptDef("default", "")
	if got != "default" {
		t.Errorf("GetOptDef() with zero-value opt = %q, want default", got)
	}
}

func TestGetOptDef_TooManyOptsPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic with >1 opts")
		}
	}()
	_ = GetOptDef("def", "a", "b")
}

func TestGetOptDef_StructType(t *testing.T) {
	type opts struct {
		val int
	}
	custom := opts{val: 42}
	got := GetOptDef(opts{}, custom)
	if got.val != 42 {
		t.Errorf("GetOptDef() with struct = %v, want val=42", got.val)
	}
}

// --- ParseMode tests --------------------------------------------------------

func TestParseModeString(t *testing.T) {
	tests := []struct {
		mode ParseMode
		want string
	}{
		{ParseModeHTML, "html"},
		{ParseModeMarkdown, "markdown"},
		{ParseModeDisabled, "disabled"},
		{ParseModeDefault, "default"},
		{MarkdownV2, "MarkdownV2"},
	}
	for _, tc := range tests {
		if got := tc.mode.String(); got != tc.want {
			t.Errorf("%v.String() = %q, want %q", tc.mode, got, tc.want)
		}
	}
}

// --- helpers ---------------------------------------------------------------

func assertButtonText(t *testing.T, btn *tl.KeyboardButton, want string) {
	t.Helper()
	if btn.Text != want {
		t.Errorf("button text = %q, want %q", btn.Text, want)
	}
}
