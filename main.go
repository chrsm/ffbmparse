package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chrsm/ffbmparse/bookmark"
)

func main() {
	fi, err := os.Open("bookmarks.json")
	if err != nil {
		panic(err)
	}

	bookmarks := &bookmark.Item{}

	dec := json.NewDecoder(fi)
	dec.Decode(bookmarks)

	recurse(bookmarks, "")
}

func recurse(bm *bookmark.Item, prefix string) {
	if bm.TypeCode == bookmark.TypeFolder && len(bm.Children) > 0 {
		fmt.Printf("%sFolder: %s\n", prefix, bm.Title)

		for i := range bm.Children {
			recurse(bm.Children[i], prefix+"\t")
		}
	} else if bm.TypeCode == bookmark.TypeBookmark {
		fmt.Printf("%sTitle: %s\n", prefix, bm.Title)
		fmt.Printf("%sURI: %s\n\n", prefix, bm.URI)
	}
}
