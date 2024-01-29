package main

import (
	_ "embed"
	"image/color"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"
)

const (
	WIDTH  = 320
	HEIGHT = 240
)

var (
	pressed uint8
)

func Badge() {
	ledColors[0] = color.RGBA{0, 0, 100, 255}
	ledColors[1] = color.RGBA{60, 0, 0, 255}
	leds.WriteColors(ledColors)

	for {
		myNameIs()
		time.Sleep(time.Second * 5)
		zwindlerDenis()
		time.Sleep(time.Second * 5)
		comeToMyTalk()
		time.Sleep(time.Second * 5)
	}
}

func myNameIs() {
	display.FillScreen(colors[WHITE])

	display.FillRectangle(0, 0, 80, HEIGHT, colors[BLUE])
	display.FillRectangle(WIDTH-80, 0, 80, HEIGHT, colors[RED])

	w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, "HELLO")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 34, "HELLO", colors[WHITE])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique12pt7b, "my name is")
	tinyfont.WriteLine(&display, &freesans.Oblique12pt7b, (WIDTH-int16(w32))/2, 54, "my name is", colors[WHITE])

	// gophers
	tinyfont.WriteLine(&display, &gophers.Regular58pt, WIDTH-200, 235, "BE", colors[RED])

	w32, _ = getFontWidthSize("Denis")
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32))/2, 140, "Denis", colors[BLACK])
}

func zwindlerDenis() {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := getFontWidthSize("@zwindler")
	w32bottom, _ := getFontWidthSize("Denis Germain")
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 90, "@zwindler", colors[BLUE])

	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32bottom))/2, 180, "Denis Germain", colors[BLACK])
}

func comeToMyTalk() {
	display.FillScreen(colors[WHITE])

	w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, "Come to my talk")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 35, "Come to my talk", colors[BLACK])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique18pt7b, "Putting an end to")
	tinyfont.WriteLine(&display, &freesans.Oblique18pt7b, (WIDTH-int16(w32))/2, 90, "Putting an end to", colors[BLUE])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique18pt7b, "Makefiles in go proj.")
	tinyfont.WriteLine(&display, &freesans.Oblique18pt7b, (WIDTH-int16(w32))/2, 130, "Makefiles in go proj.", colors[BLUE])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique18pt7b, "with GoReleaser")
	tinyfont.WriteLine(&display, &freesans.Oblique18pt7b, (WIDTH-int16(w32))/2, 170, "with GoReleaser", colors[BLUE])

	w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, "ud2218a Sat. 15:00")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 225, "ud2218a Sat. 15:00", colors[BLACK])

}

func getFontWidthSize(text string) (w32 uint32, size byte) {
	w32, _ = tinyfont.LineWidth(&freesans.Bold24pt7b, text)
	size = 24
	if w32 < 300 {
		size = 24
	} else {
		w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, text)
		if w32 < 300 {
			size = 18
		} else {
			w32, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, text)
			if w32 < 300 {
				size = 12
			} else {
				w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, text)
				size = 9
			}
		}
	}
	return
}

func reduceLedIntensity() {
	start := true
	for i := 0; i < 4999; i++ {
		if start {
			ledColors[0] = color.RGBA{0, 0, 100, 255}
			ledColors[1] = color.RGBA{60, 0, 0, 255}
			time.Sleep(time.Microsecond * 300)
			start = false
		} else {
			ledColors[0] = color.RGBA{0, 0, 0, 255}
			ledColors[1] = color.RGBA{0, 0, 0, 255}
			time.Sleep(time.Microsecond * 1800)
			start = true
		}
		leds.WriteColors(ledColors)
	}
}
