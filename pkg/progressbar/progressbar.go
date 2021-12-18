package progressbar

import (
	"fmt"
	"strings"
	"time"
)

const ProgressBarWidth = 50

type ProgressBar struct {
	title   string
	min     int
	max     int
	percent int
	start  time.Time
}

func New(title string, min, max int) *ProgressBar {
	return &ProgressBar{
		title:   title,
		min:     min,
		max:     max,
		percent: -1,
		start: time.Now(),
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
		speed := int(float64(current-bar.min) / time.Now().Sub(bar.start).Seconds())
		fmt.Printf("\r%s %s%s %d%% | %d ops/secs", bar.title, strings.Repeat("▓", width), strings.Repeat("░", ProgressBarWidth-width), percent, speed)
	}
}

func (bar *ProgressBar) Done() {
	fmt.Println()
}
