package models

type Weather struct {
	Wind  uint `json:"wind"`
	Water uint `json:"water"`
}

func (w *Weather) GetStatusWind() string {
	switch {
	case w.Wind > 15:
		return "Bahaya"
	case w.Wind <= 15 && w.Wind >= 7:
		return "Siaga"
	case w.Wind < 7:
		return "Aman"

	default:
		return ""
	}
}

func (w *Weather) GetStatusWater() string {
	switch {
	case w.Water > 8:
		return "Bahaya"
	case w.Water <= 8 && w.Water >= 6:
		return "Siaga"
	case w.Water < 6:
		return "Aman"

	default:
		return ""
	}
}
