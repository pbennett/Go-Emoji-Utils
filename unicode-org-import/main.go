package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// This helper program grabs the text description of the Unicode 15.1 character set, grabbing all fully-qualified emojis
// and replacing the Go-Emoji-Utils emojidata.go file for update.

// example lines to parse
// 1F600         ; fully-qualified     # 😀 E1.0 grinning face
// 1F170 FE0F    ; fully-qualified     # 🅰️ E0.6 A button (blood type)
// 1F9D6 1F3FC 200D 2642 FE0F   ; fully-qualified     # 🧖🏼‍♂️ E5.0 man in steamy room: medium-light skin tone
var unicodeRegex = regexp.MustCompile(`(?P<bytes>( ?[0-9A-F]{2,5})+).+fully-qualified.+# (?P<emoji>[^ ]+) E\d+\.\d+ (?P<title>.+)$`)

func main() {
	resp, err := http.Get("https://unicode.org/Public/emoji/15.1/emoji-test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	outFile, err := os.Create("emojidata.go")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	outFile.WriteString(`package emoji

// Emojis - Map of Emoji Runes as Hex keys to their description
var Emojis = map[string]Emoji{
`)

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line := scanner.Text()
		matches := unicodeRegex.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		byteVals := matches[unicodeRegex.SubexpIndex("bytes")]
		keyPart := strings.Join(strings.Split(byteVals, " "), "-")

		outFile.WriteString(fmt.Sprintf(`	"%s": {
		Key:        "%s",
		Value:      "%s",
		Descriptor: "%s",
	},
`, keyPart, keyPart, matches[unicodeRegex.SubexpIndex("emoji")], matches[unicodeRegex.SubexpIndex("title")]))
	}

	outFile.WriteString(`}
`)
}
