package internal

import "./img"

type BadGuyStatus struct {
	State     BadGuyState
	img.Position
	TickCount int
}

type BadGuyState int8

const (
	BadGuyDisabledState BadGuyState = iota
	BadGuyExplodingState
	BadGuyEnabledState
)