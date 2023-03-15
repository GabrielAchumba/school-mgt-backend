package dtos

type Schedule struct {
	TypeOfSchedule string    `json:"typeOfSchedule"`
	CumulativeTime []float64 `json:"cumulativeTime"`
}
