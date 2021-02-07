package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Productiva")
	w.Resize(fyne.NewSize(400, 500))
	w.CenterOnScreen()

	duration := 30 * 60
	timerLabel := canvas.NewText(screenDuration(duration), color.White)
	timerLabel.TextSize= 72
	timerLabel.Alignment = fyne.TextAlignCenter

	hButtons := container.NewHBox(
		widget.NewButton("Pomodoro", func() {
			duration := 30 * 60
			timerLabel.Text = screenDuration(duration)
		}),
		layout.NewSpacer(),
		widget.NewButton("Short Break", func() {
			duration := 5 * 60
			timerLabel.Text = screenDuration(duration)
		}),
		layout.NewSpacer(),
		widget.NewButton("Long Break", func() {
			duration := 15 * 60
			timerLabel.Text = screenDuration(duration)
		}),
	)

	w.SetContent(container.NewVBox(
		hButtons,
		container.NewCenter(timerLabel),
		widget.NewButton("Start", func() {
			makeTimer(duration, timerLabel)
		}),
	))

	w.ShowAndRun()

}

func screenDuration(inSeconds int) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("%02d:%02d", minutes, seconds)
	return str
}

func makeTimer(duration int, label *canvas.Text) {
	ticker1 := time.NewTicker(1 * time.Second)
	i := duration
	for c := range ticker1.C {
		i--
		label.Text = screenDuration(i)
		if i == 0 {
			ticker1.Stop()
			c.Second()
			break
		}
	}
}
