package progressbar

import (
	"fmt"
	"strings"
	"time"
)

const ProgressBarWidth = 50

type ProgressBar struct {
	title   string
	min     uint32
	max     uint32
	percent int
	start  time.Time
}

func New(title string, min, max uint32) *ProgressBar {
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

func (bar *ProgressBar) Update(current uint32) {
	percent := int(float64(current-bar.min) / float64(bar.max-bar.min) * 100)
	if bar.percent != percent {
		bar.percent = percent
		width := percent * ProgressBarWidth / 100
		speed := float64(current-bar.min) / float64(time.Since(bar.start).Seconds())
		fmt.Printf("\r%s %s%s %d%% | %.0f ops/secs", bar.title, strings.Repeat("▓", width), strings.Repeat("░", ProgressBarWidth-width), percent, speed)
	}
}

func (bar *ProgressBar) Done() {
	fmt.Println()
}
