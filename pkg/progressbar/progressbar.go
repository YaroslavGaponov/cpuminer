package progressbar

import (
	"fmt"
	"strings"
)

const ProgressBarWidth = 50

type ProgressBar struct {
	title   string
	min     int
	max     int
	percent int
}

func New(title string, min, max int) *ProgressBar {
	return &ProgressBar{
		title:   title,
		min:     min,
		max:     max,
		percent: -1,
	}
}

func (bar *ProgressBar) Begin() {
	fmt.Println()
	bar.Update(bar.min)
}

func (bar *ProgressBar) Update(current int) {
	percent := int(float64(current-bar.min) / float64(bar.max-bar.min) * 100)
	if bar.percent != percent {
		bar.percent = percent
		width := percent * ProgressBarWidth / 100
		fmt.Printf("\r%s %s%s %d%%", bar.title, strings.Repeat("▓", width), strings.Repeat("░", ProgressBarWidth-width), percent)
	}
}

func (bar *ProgressBar) Done() {
	bar.Update(bar.max)
	fmt.Println()
}
