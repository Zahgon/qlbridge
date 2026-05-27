package builtins

import (
	"regexp"

	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

var _ = u.EMPTY

// email a string, parses email and makes sure it is valid
//
//	email("Bob <bob@bob.com>")  =>  bob@bob.com, true
//	email("Bob <bob>")          =>  "", false
type Email struct{}

// Type string
func (m *Email) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Email) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func emailEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// emailname a string, parses email
//
//	emailname("Bob <bob@bob.com>") =>  Bob
type EmailName struct{}

// Type string
func (m *EmailName) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *EmailName) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func emailNameEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// emaildomain parses email and returns domain
//
//	emaildomain("Bob <bob@bob.com>") =>  bob.com
type EmailDomain struct{}

// Type string
func (m *EmailDomain) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *EmailDomain) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func emailDomainEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Domains Extract Domains from a Value, or Values (must be urlish), doesn't do much/any validation
//
//	domains("http://www.lytics.io/index.html") =>  []string{"lytics.io"}
type Domains struct{}

// Type strings
func (m *Domains) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Domains) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func domainsEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Since its empty, we will just re-use it

// Now convert to domains

// May not have an http prefix, if not assume it

// Extract Domain from a Value, or Values (must be urlish), doesn't do much/any validation.
// if input is a list of strings, only first is evaluated, for plural see domains()
//
//	domain("http://www.lytics.io/index.html") =>  "lytics.io"
type Domain struct{}

// Type string
func (m *Domain) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Domain) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func domainEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Extract host from a String (must be urlish), doesn't do much/any validation
// In the event the value contains more than one input url, will ONLY evaluate first
//
//	host("http://www.lytics.io/index.html") =>  www.lytics.io
type Host struct{}

// Type string
func (m *Host) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Host) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func HostEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

//u.Infof("url.parse: %#v", urlParsed)

// Extract hosts from a Strings (must be urlish), doesn't do much/any validation
//
//	hosts("http://www.lytics.io", "http://www.activate.lytics.io") => www.lytics.io, www.activate.lytics.io
type Hosts struct{}

// Type strings
func (m *Hosts) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Hosts) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func HostsEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

//u.Infof("url.parse: %#v", urlParsed)

// url decode a string
//
//	urldecode("http://www.lytics.io/index.html") =>  http://www.lytics.io
//
// In the event the value contains more than one input url, will ONLY evaluate first
type UrlDecode struct{}

// Type string
func (m *UrlDecode) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *UrlDecode) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func urlDecodeEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// UrlPath Extract url path from a String (must be urlish), doesn't do much/any validation
//
//	path("http://www.lytics.io/blog/index.html") =>  blog/index.html
//
// In the event the value contains more than one input url, will ONLY evaluate first
type UrlPath struct{}

// Type string
func (m *UrlPath) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *UrlPath) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func urlPathEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Qs Extract qs param from a string (must be url valid)
//
//	qs("http://www.lytics.io/?utm_source=google","utm_source")  => "google", true
type Qs struct{}

// Type string
func (m *Qs) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Qs) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func qsEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Qs Extract qs param from a string (must be url valid)
//
//	qs("http://www.lytics.io/?utm_source=google","utm_source")  => "google", true
type QsDeprecate struct{}

// Type string
func (m *QsDeprecate) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *QsDeprecate) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func qsDeprecateEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// UrlMain remove the querystring and scheme from url
//
//	urlmain("http://www.lytics.io/?utm_source=google")  => "www.lytics.io/", true
type UrlMain struct{}

// Type string
func (m *UrlMain) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *UrlMain) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func urlMainEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// UrlMinusQs removes a specific query parameter and its value from a url
//
//	urlminusqs("http://www.lytics.io/?q1=google&q2=123", "q1") => "http://www.lytics.io/?q2=123", true
type UrlMinusQs struct{}

// Type string
func (m *UrlMinusQs) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *UrlMinusQs) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func urlMinusQsEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// UrlWithQueryFunc strips a url and retains only url parameters that match
// the supplied regular expressions.
//
//	url.matchqs(url, re1, re2, ...)  => url_withoutqs
type UrlWithQuery struct{}

// Type string
func (m *UrlWithQuery) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (*UrlWithQuery) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

// Memoize these compiled reg-expressions

// UrlWithQueryEval pass reg-expressions to match qs args.
// Must match one regexp or else the qs param is dropped.
func UrlWithQueryEval(include []*regexp.Regexp) expr.EvaluatorFunc {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc)
}

// include fields specified as arguments

// UserAgent Extract user agent features
//
//	useragent(user_agent_field,"mobile")  => "true", true
type UserAgent struct{}

// Type string
func (m *UserAgent) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *UserAgent) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func userAgentEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

/*
   fmt.Printf("%v\n", ua.Mobile())   // => false
   fmt.Printf("%v\n", ua.Bot())      // => false
   fmt.Printf("%v\n", ua.Mozilla())  // => "5.0"

   fmt.Printf("%v\n", ua.Platform()) // => "X11"
   fmt.Printf("%v\n", ua.OS())       // => "Linux x86_64"

   name, version := ua.Engine()
   fmt.Printf("%v\n", name)          // => "AppleWebKit"
   fmt.Printf("%v\n", version)       // => "537.11"

   name, version = ua.Browser()
   fmt.Printf("%v\n", name)          // => "Chrome"
   fmt.Printf("%v\n", version)       // => "23.0.1271.97"

   // Let's see an example with a bot.

   ua.Parse("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

   fmt.Printf("%v\n", ua.Bot())      // => true

   name, version = ua.Browser()
   fmt.Printf("%v\n", name)          // => Googlebot
   fmt.Printf("%v\n", version)       // => 2.1
*/

// UserAgentMap Extract user agent features
//
//	useragent.map(user_agent_field)  => {"mobile": "false","platform":"X11"}, true
type UserAgentMap struct{}

// Type MapString
func (m *UserAgentMap) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}

func (m *UserAgentMap) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func userAgentMapEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

/*
   fmt.Printf("%v\n", ua.Mobile())   // => false
   fmt.Printf("%v\n", ua.Bot())      // => false
   fmt.Printf("%v\n", ua.Mozilla())  // => "5.0"

   fmt.Printf("%v\n", ua.Platform()) // => "X11"
   fmt.Printf("%v\n", ua.OS())       // => "Linux x86_64"

   name, version := ua.Engine()
   fmt.Printf("%v\n", name)          // => "AppleWebKit"
   fmt.Printf("%v\n", version)       // => "537.11"

   name, version = ua.Browser()
   fmt.Printf("%v\n", name)          // => "Chrome"
   fmt.Printf("%v\n", version)       // => "23.0.1271.97"

   // Let's see an example with a bot.

   ua.Parse("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

   fmt.Printf("%v\n", ua.Bot())      // => true

   name, version = ua.Browser()
   fmt.Printf("%v\n", name)          // => Googlebot
   fmt.Printf("%v\n", version)       // => 2.1
*/
