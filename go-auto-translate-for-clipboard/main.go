package main

import (
	"fmt"
	"os"
	"time"

	gt "github.com/bas24/googletranslatefree"
	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	lastText := ""
	for {
		time.Sleep(150)
		result := ""
		jpresult := ""
		koresult := ""
		cnresult := ""
		clipText := string(clipboard.Read(clipboard.FmtText))
		if clipText == "" {
			continue
		}
		if clipText == lastText {
			continue
		}
		lastText = clipText

		if len(os.Args) > 1 {
			result, _ = gt.Translate(clipText, "auto", "ko")
		} else {
			result, _ = gt.Translate(clipText, "auto", "en")
			jpresult, _ = gt.Translate(clipText, "auto", "ja")
			koresult, _ = gt.Translate(clipText, "auto", "ko")
			cnresult, _ = gt.Translate(clipText, "auto", "zh_CN")
			result = result + "\n" + jpresult + "\n" + koresult + "\n" + cnresult
		}
		clipboard.Write(clipboard.FmtText, []byte(result))
		lastText = result
		fmt.Println(clipText + " to " + result)
	}
}
