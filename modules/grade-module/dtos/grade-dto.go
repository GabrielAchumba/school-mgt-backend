package dtos

import "time"

type CreateGradeRequest struct {
	Type     string  `json:"type" binding:"required"`
	Point    float64 `json:"point"`
	From     float64 `json:"from"`
	To       float64 `json:"to"`
	SchoolId string  `json:"schoolId" binding:"required"`
}

type UpdateGradeRequest struct {
	Type     string  `json:"type" binding:"required"`
	Point    float64 `json:"point"`
	From     float64 `json:"from"`
	To       float64 `json:"to"`
	SchoolId string  `json:"schoolId" binding:"required"`
}

type GradeResponse struct {
	Id        string    `json:"id"  bson:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Type      string    `json:"type" binding:"required"`
	Point     float64   `json:"point"`
	From      float64   `json:"from"`
	To        float64   `json:"to"`
	SchoolId  string    `json:"schoolId" binding:"required"`
}

func GetGradeAndPoint(grades []GradeResponse, score float64) (string, float64) {

	_grade := "A"
	point := 1.0

	for _, grade := range grades {
		if score >= grade.From && score < grade.To {
			_grade = grade.Type
			point = grade.Point
			break
		}
	}

	return _grade, point
}
