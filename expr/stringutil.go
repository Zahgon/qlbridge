package expr

import (
	"bytes"

	u "github.com/araddon/gou"
)

var _ = u.EMPTY

// IsValidIdentity test the given string to determine if any characters are
// not valid and therefore must be quoted
func IsValidIdentity(identity string) bool { _ = "STUB: not implemented"; return false }

// LeftRight Return left, right values if is of form `table.column` or `schema`.`table`
// also return true/false for if it even has left/right
func LeftRight(val string) (string, string, bool) { _ = "STUB: not implemented"; return "", "", false }

// wat, no idea what this is

// wat, no idea what this is

// IdentityTrim trims the leading/trailing identity quote marks  ` or []
func IdentityTrim(ident string) string { _ = "STUB: not implemented"; return "" }

// IdentityMaybeQuote
func IdentityMaybeQuote(quote byte, ident string) string { _ = "STUB: not implemented"; return "" }

// IdentityMaybeEscape Quote an identity/literal
// if need be (has illegal characters or spaces)
func IdentityMaybeEscapeBuf(buf *bytes.Buffer, quote byte, ident string) {
	_ = "STUB: not implemented"
	return
}

// IdentityMaybeQuoteStrict Quote an identity if need be (has illegal characters or spaces)
//
//	First character MUST be alpha (not numeric or any other character)
func IdentityMaybeQuoteStrictBuf(buf *bytes.Buffer, quote byte, ident string) {
	_ = "STUB: not implemented"
	return
}

// Already escaped??

// IdentityMaybeQuoteStrict Quote an identity if need be (has illegal characters or spaces)
// First character MUST be alpha (not numeric or any other character)
func IdentityMaybeQuoteStrict(quote byte, ident string) string {
	_ = "STUB: not implemented"
	return ""
}

func escapeQuote(buf *bytes.Buffer, quote rune, val string) { _ = "STUB: not implemented"; return }

// LiteralQuoteEscape escape string that may need characters escaped
//
//	LiteralQuoteEscape("'","item's") => 'item''s'
//	LiteralQuoteEscape(`"`,"item's") => "item's"
//	LiteralQuoteEscape(`"`,`item"s`) => "item""s"
func LiteralQuoteEscape(quote rune, literal string) string { _ = "STUB: not implemented"; return "" }

// Already escaped??

// LiteralQuoteEscapeBuf escape string that may need characters escaped
//
//	LiteralQuoteEscapeBuf("'","item's") => 'item''s'
//	LiteralQuoteEscapeBuf(`"`,"item's") => "item's"
//	LiteralQuoteEscapeBuf(`"`,`item"s`) => "item""s"
func LiteralQuoteEscapeBuf(buf *bytes.Buffer, quote rune, literal string) {
	_ = "STUB: not implemented"
	return
}

// Already escaped??

// StringEscape escape string that may need characters escaped
//
//	StringEscape("'","item's") => "item''s"
func StringEscape(quote rune, literal string) string { _ = "STUB: not implemented"; return "" }

// StringUnEscape remove escaping on string that may need characters escaped
//
//	StringUnEscape(`"`,`item"s`) => "item""s", true
func StringUnEscape(quote rune, val string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// A break, is some character such as comma, ;, whitespace
func isBreak(r rune) bool { _ = "STUB: not implemented"; return false }
