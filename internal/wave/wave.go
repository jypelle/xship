package wave

var waves = []Wave{{}}

type Wave struct {
	tickCount int
}

func (w *Wave) Move() {
	w.tickCount++
}
