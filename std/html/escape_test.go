package html

import (
	"fmt"
	"html"
	"testing"
)

func TestEscape(t *testing.T) {

	//  `&`, "&amp;",
	//	`'`, "&#39;", // "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	//	`<`, "&lt;",
	//	`>`, "&gt;",
	//	`"`, "&#34;", // "&#34;" is shorter than "&quot;".
	fmt.Println(html.EscapeString("<h1>hello world</h1>"))

	//&lt;&amp;&#39;&#39;&#34;&gt;
	fmt.Println(html.EscapeString("<&''\">"))

	//&lt;&amp;&#39;&#39;&#34;&gt;
	fmt.Println(html.EscapeString(`<&''">`))

	fmt.Println(html.EscapeString("&lt;&amp;&#39;&#39;&#34;&gt;"))

}

func TestUnescape(t *testing.T) {
	//<&''">
	fmt.Println(html.UnescapeString("&lt;&amp;&#39;&#39;&#34;&gt;"))
}
