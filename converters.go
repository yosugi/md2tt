package main

import (
	"bytes"
	"fmt"
	"regexp"

   "github.com/yosugi/md2tt/lib/freeze"
)

// string -> Freeze
func link(line string) freeze.Freeze {
    pattern := `\[(.*?)\]\((.*?)\)`
    re := regexp.MustCompile(pattern)

	replacedLine := re.ReplaceAllString(line, `"$1":$2`)

	return freeze.Wrap(replacedLine)
}

// string -> Freeze
func heading(line string) freeze.Freeze {
	pattern := "^#*"
	re := regexp.MustCompile(pattern)

	matched := re.FindString(line)
	matchedLen := len(matched)
	if matchedLen == 0 {
		return freeze.Wrap(line)
	}

	times := matchedLen
	headingStr := fmt.Sprintf("h%d.", times)
	replacedLine := re.ReplaceAllString(line, headingStr)

	return freeze.WrapFrozen(replacedLine)
}

// int -> string -> string -> (string -> Freeze)
func list(indent int, listPattern, replaceStr string) func(string) freeze.Freeze {
	pattern := fmt.Sprintf(`^(( {%d})*)%s`, indent, listPattern)
	re := regexp.MustCompile(pattern)

	return func(line string) freeze.Freeze {
		matched := re.FindStringSubmatch(line)
		if matched == nil {
			return freeze.Wrap(line)
		}

		indentLen := len(matched[1])
		times := indentLen/indent + 1
		str := []byte(replaceStr)
		repeatedStr := string(bytes.Repeat(str, times))
		replacedLine := re.ReplaceAllString(line, repeatedStr)

		return freeze.WrapFrozen(replacedLine)
	}
}

// int -> (string -> Freeze)
func bulletList(indent int) func(string) freeze.Freeze {
	return list(indent, `[-*+]`, `*`)
}

// int -> (string -> Freeze)
func numberedList(indent int) func(string) freeze.Freeze {
	return list(indent, `\d+\.`, `#`)
}
