package html

import (
	"fmt"
	"html"
)

func ExampleEscapeString() {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Println(html.EscapeString(s))
	// Output: &#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;
}

func ExampleUnescapeString() {
	const s = `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
	fmt.Println(html.UnescapeString(s))
	// Output: "Fran & Freddie's Diner" <tasty@example.com>
}
