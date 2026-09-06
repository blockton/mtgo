package types

import (
	"testing"

	"github.com/mtgo-labs/mtgo/telegram/params"
	"github.com/mtgo-labs/mtgo/tg"
)

func TestInlineArticleParsesParseMode(t *testing.T) {
	a := &InlineArticle{
		ID:        "1",
		Title:     "T",
		Text:      "<b>bold</b> plain",
		ParseMode: params.ParseModeHTML,
	}
	res, err := a.TL()
	if err != nil {
		t.Fatalf("TL() error: %v", err)
	}
	r, ok := res.(*tg.InputBotInlineResult)
	if !ok {
		t.Fatalf("result type = %T", res)
	}
	if r.ID != "1" || r.Type != "article" {
		t.Fatalf("id/type = %q/%q", r.ID, r.Type)
	}
	if !r.Flags.Has(1) {
		t.Fatalf("title flag 1 not set, flags=%032b", r.Flags)
	}
	msg, ok := r.SendMessage.(*tg.InputBotInlineMessageText)
	if !ok {
		t.Fatalf("send_message type = %T", r.SendMessage)
	}
	msg.SetFlags() // Encode computes flags lazily; mirror it before asserting.
	if msg.Message != "bold plain" {
		t.Fatalf("message = %q, want markup stripped", msg.Message)
	}
	if len(msg.Entities) != 1 {
		t.Fatalf("entities = %v, want one", msg.Entities)
	}
	if _, isBold := msg.Entities[0].(*tg.MessageEntityBold); !isBold {
		t.Fatalf("entity type = %T, want *tg.MessageEntityBold", msg.Entities[0])
	}
	// Entities must be flagged via the generated SetFlags path.
	if !msg.Flags.Has(1) {
		t.Fatalf("entities flag 1 not set on message, flags=%032b", msg.Flags)
	}
	// Regression: a parse mode must not leak into no_webpage (flag 0).
	if msg.Flags.Has(0) {
		t.Fatalf("no_webpage flag 0 unexpectedly set, flags=%032b", msg.Flags)
	}
}

func TestInlineArticlePrebuiltEntitiesWin(t *testing.T) {
	prebuilt := []tg.MessageEntityClass{&tg.MessageEntityItalic{Offset: 0, Length: 2}}
	a := &InlineArticle{
		ID:        "1",
		Text:      "<b>ignored</b>",
		Entities:  prebuilt,
		ParseMode: params.ParseModeHTML,
	}
	res, err := a.TL()
	if err != nil {
		t.Fatalf("TL() error: %v", err)
	}
	r := res.(*tg.InputBotInlineResult)
	msg := r.SendMessage.(*tg.InputBotInlineMessageText)
	if len(msg.Entities) != 1 {
		t.Fatalf("entities = %v, want prebuilt only", msg.Entities)
	}
	if _, isItalic := msg.Entities[0].(*tg.MessageEntityItalic); !isItalic {
		t.Fatalf("entity type = %T, want the prebuilt italic", msg.Entities[0])
	}
	if msg.Message != "<b>ignored</b>" {
		t.Fatalf("message = %q, want raw text with prebuilt entities", msg.Message)
	}
}

func TestInlineArticleInvalidParseMode(t *testing.T) {
	a := &InlineArticle{ID: "1", Text: "x", ParseMode: "bogus"}
	if _, err := a.TL(); err == nil {
		t.Fatal("TL() with unknown parse mode should error")
	}
}

func TestInlineArticleOptions(t *testing.T) {
	a := &InlineArticle{
		ID:          "1",
		Title:       "T",
		URL:         "https://example.com",
		Description: "D",
		ThumbURL:    "https://example.com/t.png",
		ThumbMIME:   "image/png",
		ThumbWidth:  64,
		ThumbHeight: 64,
		NoWebpage:   true,
		InvertMedia: true,
		ReplyMarkup: &tg.ReplyInlineMarkup{},
	}
	res, err := a.TL()
	if err != nil {
		t.Fatalf("TL() error: %v", err)
	}
	r := res.(*tg.InputBotInlineResult)
	if !r.Flags.Has(2) || !r.Flags.Has(3) || !r.Flags.Has(4) {
		t.Fatalf("description/url/thumb flags not set, flags=%032b", r.Flags)
	}
	thumb := r.Thumb
	if thumb == nil || thumb.MimeType != "image/png" || len(thumb.Attributes) != 1 {
		t.Fatalf("thumb = %+v", thumb)
	}
	msg := r.SendMessage.(*tg.InputBotInlineMessageText)
	msg.SetFlags()
	if !msg.Flags.Has(0) || !msg.Flags.Has(3) || !msg.Flags.Has(2) {
		t.Fatalf("no_webpage/invert_media/reply_markup flags not set, flags=%032b", msg.Flags)
	}
}

func TestInlinePhotoAndDocumentEntities(t *testing.T) {
	p := &InlinePhoto{ID: "2", PhotoID: 1, AccessHash: 2, Text: "*b*", ParseMode: params.MarkdownV2}
	res, err := p.TL()
	if err != nil {
		t.Fatalf("photo TL() error: %v", err)
	}
	pr := res.(*tg.InputBotInlineResultPhoto)
	pmsg := pr.SendMessage.(*tg.InputBotInlineMessageMediaAuto)
	if pmsg.Message != "b" || len(pmsg.Entities) != 1 {
		t.Fatalf("photo caption not parsed: %#v", pmsg)
	}

	d := &InlineDocument{ID: "3", DocumentID: 1, AccessHash: 2, Type: "gif", Text: "<i>i</i>", ParseMode: params.ParseModeHTML}
	res, err = d.TL()
	if err != nil {
		t.Fatalf("document TL() error: %v", err)
	}
	dr := res.(*tg.InputBotInlineResultDocument)
	if dr.Type != "gif" {
		t.Fatalf("type = %q, want gif", dr.Type)
	}
	dmsg := dr.SendMessage.(*tg.InputBotInlineMessageMediaAuto)
	if dmsg.Message != "i" || len(dmsg.Entities) != 1 {
		t.Fatalf("document caption not parsed: %#v", dmsg)
	}
}

func TestInlineGameReplyMarkup(t *testing.T) {
	g := &InlineGame{ID: "4", ShortName: "game", ReplyMarkup: &tg.ReplyInlineMarkup{}}
	res, err := g.TL()
	if err != nil {
		t.Fatalf("TL() error: %v", err)
	}
	gr := res.(*tg.InputBotInlineResultGame)
	msg := gr.SendMessage.(*tg.InputBotInlineMessageGame)
	msg.SetFlags()
	if msg.ReplyMarkup == nil || !msg.Flags.Has(2) {
		t.Fatalf("game reply markup not set: %#v", msg)
	}
}

func TestInlineLocation(t *testing.T) {
	l := &InlineLocation{
		ID: "5", Title: "HQ", Latitude: 55.75, Longitude: 37.61,
		Heading: 90, Period: 3600, ProximityRadius: 100,
	}
	res, err := l.TL()
	if err != nil {
		t.Fatalf("TL() error: %v", err)
	}
	r := res.(*tg.InputBotInlineResult)
	if r.Type != "geo" {
		t.Fatalf("type = %q, want geo", r.Type)
	}
	msg, ok := r.SendMessage.(*tg.InputBotInlineMessageMediaGeo)
	if !ok {
		t.Fatalf("send_message type = %T", r.SendMessage)
	}
	geo := msg.GeoPoint.(*tg.InputGeoPoint)
	if geo.Lat != 55.75 || geo.Long != 37.61 {
		t.Fatalf("geo = %+v", geo)
	}
	if msg.Heading != 90 || msg.Period != 3600 || msg.ProximityNotificationRadius != 100 {
		t.Fatalf("live location fields = %+v", msg)
	}
	msg.SetFlags()
	if !msg.Flags.Has(0) || !msg.Flags.Has(1) {
		t.Fatalf("heading/period flags not set, flags=%032b", msg.Flags)
	}
}

func TestInlineVenueAndContact(t *testing.T) {
	v := &InlineVenue{
		ID: "6", Title: "Coffee", Latitude: 1, Longitude: 2,
		Address: "Main St", Provider: "foursquare", VenueID: "v1", VenueType: "cafe",
	}
	res, err := v.TL()
	if err != nil {
		t.Fatalf("venue TL() error: %v", err)
	}
	vr := res.(*tg.InputBotInlineResult)
	if vr.Type != "venue" {
		t.Fatalf("type = %q, want venue", vr.Type)
	}
	vmsg := vr.SendMessage.(*tg.InputBotInlineMessageMediaVenue)
	if vmsg.Title != "Coffee" || vmsg.Address != "Main St" || vmsg.VenueID != "v1" {
		t.Fatalf("venue message = %#v", vmsg)
	}

	ct := &InlineContact{ID: "7", FirstName: "Ada", LastName: "L", PhoneNumber: "+15550100"}
	res, err = ct.TL()
	if err != nil {
		t.Fatalf("contact TL() error: %v", err)
	}
	cr := res.(*tg.InputBotInlineResult)
	if cr.Type != "contact" {
		t.Fatalf("type = %q, want contact", cr.Type)
	}
	if cr.Title != "Ada L" {
		t.Fatalf("title = %q, want %q", cr.Title, "Ada L")
	}
	cmsg := cr.SendMessage.(*tg.InputBotInlineMessageMediaContact)
	if cmsg.PhoneNumber != "+15550100" || cmsg.FirstName != "Ada" || cmsg.LastName != "L" {
		t.Fatalf("contact message = %#v", cmsg)
	}
}

func TestInlineRich(t *testing.T) {
	r := &InlineRich{
		ID:          "8",
		Title:       "Doc",
		RichMessage: &tg.InputRichMessageMarkdown{Markdown: "**x**"},
		ReplyMarkup: &tg.ReplyInlineMarkup{},
	}
	res, err := r.TL()
	if err != nil {
		t.Fatalf("TL() error: %v", err)
	}
	rr := res.(*tg.InputBotInlineResult)
	if rr.Type != "rich" {
		t.Fatalf("type = %q, want rich", rr.Type)
	}
	msg, ok := rr.SendMessage.(*tg.InputBotInlineMessageRichMessage)
	if !ok {
		t.Fatalf("send_message type = %T", rr.SendMessage)
	}
	if _, ok := msg.RichMessage.(*tg.InputRichMessageMarkdown); !ok {
		t.Fatalf("rich message type = %T", msg.RichMessage)
	}
	msg.SetFlags()
	if !msg.Flags.Has(2) {
		t.Fatalf("reply_markup flag not set, flags=%032b", msg.Flags)
	}

	if _, err := (&InlineRich{ID: "9"}).TL(); err == nil {
		t.Fatal("InlineRich without a rich message should error")
	}
}

func TestBuildInlineResultsStopsOnError(t *testing.T) {
	good := &InlineArticle{ID: "1", Text: "x"}
	bad := &InlineArticle{ID: "2", Text: "x", ParseMode: "bogus"}
	if _, err := BuildInlineResults([]InlineResultBuilder{good, bad}); err == nil {
		t.Fatal("BuildInlineResults should propagate builder errors")
	}
	if _, err := BuildInlineResults([]InlineResultBuilder{good}); err != nil {
		t.Fatalf("BuildInlineResults(good) error: %v", err)
	}
	if _, err := BuildInlineResults([]InlineResultBuilder{nil}); err == nil {
		t.Fatal("BuildInlineResults(nil builder) should error")
	}
}
