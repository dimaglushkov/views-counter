package app

import (
	"bytes"
	"github.com/ajstarks/svgo"
	"strconv"
)

func buildBadge(label string, viewsCnt int64) *bytes.Buffer {
	buf := bytes.Buffer{}
	viewsCntString := beautify(viewsCnt)

	labelFieldWidth := len(label)*6 + 15
	viewsFieldWidth := len(viewsCntString)*6 + 15

	canvas := svg.New(&buf)
	canvas.Start(labelFieldWidth+viewsFieldWidth, 20)
	canvas.Rect(0, 0, labelFieldWidth, 20, `fill="#555"`)
	canvas.Rect(labelFieldWidth, 0, viewsFieldWidth, 20, `fill="#08C"`)
	canvas.Text(labelFieldWidth/2, 14, label, `fill="#fff"`, `font-family="Verdana,DejaVu Sans,sans-serif"`, `font-size="12"`, `text-anchor="middle"`)
	canvas.Text(labelFieldWidth+viewsFieldWidth/2, 14, viewsCntString, `fill="#fff"`, `font-family="Verdana,DejaVu Sans,sans-serif"`, `font-size="12"`, `text-anchor="middle"`)
	canvas.End()

	return &buf
}

func beautify(num int64) string {
	switch {
	case num == 0:
		return "none"
	case num < 1000:
		return strconv.FormatInt(num, 10)
	case num < 10000:
		return strconv.FormatFloat(float64(num)/float64(1000), 'f', 1, 64) + "K"
	case num < 1000000:
		return strconv.FormatInt(num/1000, 10) + "K"
	case num < 10000000:
		return strconv.FormatFloat(float64(num)/float64(1000000), 'f', 1, 64) + "M"
	case num < 1000000000:
		return strconv.FormatInt(num/1000000, 10) + "M"
	default:
		return "∞"
	}
}
