/*
A Cookie represents an HTTP cookie as sent in the Set-Cookie Header of an HTTP Response.
OR, The Cookie Header of the HTTP Request

More Detailed:
Cookies are small key-value pairs that the server sends to the client in a Set-Cookie header.
Then, on subsequent requests, the client includes them in the Cookie header.


*/

package cookies

import (
	"fmt"
	"strings"
	"time"
)

/*
SameSite allows a server to define a cookie attribute making it impossible
for the browser to send this cookie along with cross-site requests. The main
goal is to mitigate the risk of cross-origin information leakage, and provide
some protection against cross-site request forgery attacks.
*/
type SameSite int

const (
	SameSiteDefaultMode SameSite = iota + 1
	SameSiteLaxMode              // 2
	SameSiteStrictMode           // 3
	SameSiteNoneMode             // 4
)

type Cookie struct {
	Name   string // name of the cookie
	Value  string // cookie value
	Quoted bool   // indicates whether the Value was initially Quoted or not

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // optional

	// MaxAge = 0; means no 'MaxAge' attributes set
	// MaxAge < 0; means delete cookie now, equivalently MaxAge = 0
	// MaxAge > 0; means Max-Age attribute present and available in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

// serialize the cookie struct into string
// to set in the Set-Cookie header
func FormatSetCookie(c *Cookie) string {

	const extraCookieLength = 110

	// A Builder is used to efficiently build a string using [Builder.Write] methods.
	// It minimizes memory copying. The zero value is ready to use.
	// Do not copy a non-zero Builder.
	var b strings.Builder

	// Grow grows b's capacity, if necessary, to guarantee space for another n bytes.
	// now, atleast n bytes can be written to b without re-allocation
	b.Grow(len(c.Name) + len(c.Path) + len(c.Value) + len(c.Domain) + extraCookieLength)

	// Name and value is required
	fmt.Fprintf(&b, "%s=%s", c.Name, c.Value)

	// optional fields
	if c.Path != "" {
		fmt.Fprintf(&b, "; Path=%s", c.Path)
	}

	if c.Domain != "" {
		fmt.Fprintf(&b, "; Domain=%s", c.Domain)
	}

	if !c.Expires.IsZero() {
		fmt.Fprintf(&b, "; Expires=%s", c.Expires.UTC().Format(time.RFC1123))
	}

	if c.MaxAge > 0 {
		fmt.Fprintf(&b, "; Max-Age=%d", c.MaxAge)
	} else if c.MaxAge < 0 {
		fmt.Fprintf(&b, "; Max-Age=0") // for deletion of cookie
	}

	if c.Secure {
		b.WriteString("; secure")
	}

	if c.HttpOnly {
		b.WriteString("; HttpOnly") // prevents the use of javascript
	}

	switch c.SameSite {
	case SameSiteLaxMode:
		b.WriteString("; SameSite=Lax")

	case SameSiteStrictMode:
		b.WriteString("; SameSite=Strict")
	case SameSiteNoneMode:
		b.WriteString("; SameSite=None")
	}

	return b.String() // .String() returns the accumulated string of Builder Type String
}
