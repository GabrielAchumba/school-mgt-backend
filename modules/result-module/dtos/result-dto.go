package dtos

import "time"

type CreateResultRequest struct {
	Score         float64 `json:"score" binding:"required"`
	ScoreMax      float64 `json:"scoreMax" binding:"required"`
	SubjectId     string  `json:"subjectId" binding:"required"`
	StudentId     string  `json:"studentId" binding:"required"`
	TeacherId     string  `json:"teacherId" binding:"required"`
	ClassRoomId   string  `json:"classRoomId" binding:"required"`
	AssessmentId  string  `json:"assessmentId" binding:"required"`
	DesignationId string  `json:"designationId" binding:"required"`
	LevelId       string  `json:"levelId"`
	CreatedAt     string  `json:"createdAt" binding:"required"`
	SchoolId      string  `json:"schoolId" binding:"required"`
}

type UpdateResultRequest struct {
	Score         float64 `json:"score" binding:"required"`
	ScoreMax      float64 `json:"scoreMax" binding:"required"`
	SubjectId     string  `json:"subjectId" binding:"required"`
	StudentId     string  `json:"studentId" binding:"required"`
	TeacherId     string  `json:"teacherId" binding:"required"`
	ClassRoomId   string  `json:"classRoomId" binding:"required"`
	LevelId       string  `json:"levelId"`
	AssessmentId  string  `json:"assessmentId" binding:"required"`
	DesignationId string  `json:"designationId" binding:"required"`
	UpdatedAt     string  `json:"updatedAt" binding:"required"`
	SchoolId      string  `json:"schoolId" binding:"required"`
}

type ResultResponse struct {
	Id                  string    `json:"id"  bson:"_id"`
	CreatedAt           time.Time `json:"createdAt"`
	CreatedBy           string    `json:"createdBy"`
	Score               float64   `json:"score" binding:"required"`
	ScoreMax            float64   `json:"scoreMax" binding:"required"`
	SubjectId           string    `json:"subjectId" binding:"required"`
	StudentId           string    `json:"studentId" binding:"required"`
	TeacherId           string    `json:"teacherId" binding:"required"`
	ClassRoomId         string    `json:"classRoomId" binding:"required"`
	LevelId             string    `json:"levelId"`
	AssessmentId        string    `json:"assessmentId" binding:"required"`
	DesignationId       string    `json:"designationId" binding:"required"`
	SubjectFullName     string    `json:"subjectFullName" binding:"required"`
	StudentFullName     string    `json:"studentFullName" binding:"required"`
	TeacherFullName     string    `json:"teacherFullName" binding:"required"`
	ClassRoomFullName   string    `json:"classRoomFullName" binding:"required"`
	AssessmentFullName  string    `json:"assessmentFullName" binding:"required"`
	LevelFullName       string    `json:"levelFullName"`
	DesignationFullName string    `json:"designationFullName" binding:"required"`
	SchoolId            string    `json:"schoolId" binding:"required"`
}

type RangeOfScore struct {
	From             float64 `json:"from"`
	To               float64 `json:"to"`
	NumberOfStudents int     `json:"numberOfStudents"`
}

type GetResultsRequest struct {
	StartDate     string         `json:"startDate"`
	EndDate       string         `json:"endDate"`
	Score         float64        `json:"score"`
	ScoreMax      float64        `json:"scoreMax"`
	SubjectIds    []string       `json:"subjectIds"`
	SubjectId     string         `json:"subjectId"`
	StudentId     string         `json:"studentId"`
	StudentIds    []string       `json:"studentIds"`
	TeacherId     string         `json:"teacherId"`
	TeacherIds    []string       `json:"teacherIds"`
	ClassRoomId   string         `json:"classRoomId"`
	ClassRoomIds  string         `json:"classRoomIds"`
	SessionId     string         `json:"sessionId"`
	SessionIds    string         `json:"sessionIds"`
	LevelId       string         `json:"levelId"`
	LevelsIds     string         `json:"levelIds"`
	AssessmentId  string         `json:"assessmentId"`
	DesignationId string         `json:"designationId"`
	RangeOfScores []RangeOfScore `json:"rangeOfScores"`
	MonthYears    []MonthYear    `json:"monthYears"`
	IsMonthly     bool           `json:"isMonthly"`
	SchoolId      string         `json:"schoolId" binding:"required"`
}

type AssesmentGroup struct {
	AssessmentScore float64 `json:"assessmentScore"`
	ScoreMax        float64 `json:"scoreMax"`
}

type SubJectResult struct {
	Assessments  map[string]AssesmentGroup `json:"assessments"`
	SubjectScore float64                   `json:"subjectScore"`
	Grade        string                    `json:"grade"`
	Point        float64                   `json:"point"`
}

type StudentResults struct {
	StudentId       string                   `json:"studentId"`
	FullName        string                   `json:"fullName"`
	OverallScore    float64                  `json:"overallScore"`
	OverallScoreMax float64                  `json:"overallScoreMax"`
	CGPA            float64                  `json:"cGPA"`
	CGPAMax         float64                  `json:"cGPAMax"`
	Subjects        map[string]SubJectResult `json:"subjects"`
}

type MonthYear struct {
	Month         int `json:"month"`
	Year          int `json:"year"`
	Students      map[string][]ResultResponse
	RangeOfScores []RangeOfScore `json:"rangeOfScores"`
}
