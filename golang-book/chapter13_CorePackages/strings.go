package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		"string.Contains(\"string\", \"section\")\n",
		strings.Contains("test", "es"),

		"\n\nstrings.Count(\"string\", \"char\")\n",
		strings.Count("test", "t"),

		"\n\nstrings.HasPrefix(\"string\", \"start\")\n",
		strings.HasPrefix("test", "te"),
		"\n\nstrings.HasSuffix(\"string\", \"ending\")\n",
		strings.HasSuffix("test", "st"),
		"\n\nstrings.Index(\"string\", \"char\")\n",
		strings.Index("test", "e"),

		"\n\nstrings.Join(array, \"delimiter\")\n",
		strings.Join([]string{"a", "b"}, "-"),

		"\n\nstrings.Repeat(\"string\", times)\n",
		strings.Repeat("a", 5),

		"\n\nstrings.Replace(\"string\", \"old\", \"new\")\n",
		strings.Replace("aaaa", "a", "b", 2),

		"\n\nstrings.Split(\"string\", \"delimiter\")\n",
		strings.Split("a-b-c-d-e", "-"),

		"\n\nstrings.ToUpper(\"string\")\n",
		strings.ToUpper("test"),
	)

	arr := []byte("test")
	fmt.Println("\n\nconvert string to a slice of bytes --> []byte(\"test\")")
	fmt.Println(arr)
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println("\nconvert slice of bytes to a string --> string([]byte{'t', 'e', 's', 't'})")
	fmt.Println(str)
}
