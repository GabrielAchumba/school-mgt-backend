package utils

import (
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module/dtos"
)

type CompetitionResultUtils interface {
	GroupByStudents(CompetitionResults []dtos.CompetitionResultResponse) map[string][]dtos.CompetitionResultResponse
	GroupByStudentsAndMonthYear(CompetitionResults []dtos.CompetitionResultResponse,
		monthYears []dtos.MonthYear) map[string]dtos.MonthYear
	GroupByStudentsAndYear(CompetitionResults []dtos.CompetitionResultResponse,
		monthYears []dtos.MonthYear) map[string]dtos.MonthYear
}

type CompetitionResultUtilsImpl struct {
}

func New() CompetitionResultUtils {
	return &CompetitionResultUtilsImpl{}
}

func (impl CompetitionResultUtilsImpl) GroupByStudents(CompetitionResults []dtos.CompetitionResultResponse) map[string][]dtos.CompetitionResultResponse {

	var uniqueIds []string

	for _, v := range CompetitionResults {
		skip := false
		for _, u := range uniqueIds {
			if v.StudentId == u {
				skip = true
				break
			}
		}
		if !skip {
			uniqueIds = append(uniqueIds, v.StudentId)
		}
	}

	grouppedStudentsCompetitionResults := make(map[string][]dtos.CompetitionResultResponse, 0)

	for _, v := range uniqueIds {
		studentCompetitionResults := make([]dtos.CompetitionResultResponse, 0)
		for _, u := range CompetitionResults {
			if v == u.StudentId {
				studentCompetitionResults = append(studentCompetitionResults, u)
			}
		}

		if len(studentCompetitionResults) > 0 {
			grouppedStudentsCompetitionResults[v] = studentCompetitionResults
		}
	}

	return grouppedStudentsCompetitionResults
}

func (impl CompetitionResultUtilsImpl) GroupByStudentsAndMonthYear(CompetitionResults []dtos.CompetitionResultResponse,
	monthYears []dtos.MonthYear) map[string]dtos.MonthYear {

	var uniqueIds []string

	for _, v := range CompetitionResults {
		skip := false
		for _, u := range uniqueIds {
			if v.Id == u {
				skip = true
				break
			}
		}
		if !skip {
			uniqueIds = append(uniqueIds, v.Id)
		}
	}

	groupedMonthYears := make(map[string]dtos.MonthYear)
	for _, monthYear := range monthYears {

		grouppedStudentsCompetitionResults := make(map[string][]dtos.CompetitionResultResponse, 0)

		for _, v := range uniqueIds {
			studentCompetitionResults := make([]dtos.CompetitionResultResponse, 0)
			for _, u := range CompetitionResults {
				if v == u.Id &&
					monthYear.Month == int(u.CreatedAt.Month()) &&
					monthYear.Year == u.CreatedAt.Year() {
					studentCompetitionResults = append(studentCompetitionResults, u)
				}
			}

			if len(studentCompetitionResults) > 0 {
				grouppedStudentsCompetitionResults[v] = studentCompetitionResults
			}
		}

		key := strconv.Itoa(monthYear.Month) + "_" + strconv.Itoa(monthYear.Year)
		groupedMonthYears[key] = dtos.MonthYear{
			Month:    monthYear.Month,
			Year:     monthYear.Year,
			Students: grouppedStudentsCompetitionResults,
		}
	}

	return groupedMonthYears
}

func (impl CompetitionResultUtilsImpl) GroupByStudentsAndYear(CompetitionResults []dtos.CompetitionResultResponse,
	monthYears []dtos.MonthYear) map[string]dtos.MonthYear {

	var uniqueIds []string

	for _, v := range CompetitionResults {
		skip := false
		for _, u := range uniqueIds {
			if v.Id == u {
				skip = true
				break
			}
		}
		if !skip {
			uniqueIds = append(uniqueIds, v.Id)
		}
	}

	groupedMonthYears := make(map[string]dtos.MonthYear)
	for _, monthYear := range monthYears {

		grouppedStudentsCompetitionResults := make(map[string][]dtos.CompetitionResultResponse, 0)

		for _, v := range uniqueIds {
			studentCompetitionResults := make([]dtos.CompetitionResultResponse, 0)
			for _, u := range CompetitionResults {
				if v == u.Id &&
					monthYear.Year == u.CreatedAt.Year() {
					studentCompetitionResults = append(studentCompetitionResults, u)
				}
			}

			if len(studentCompetitionResults) > 0 {
				grouppedStudentsCompetitionResults[v] = studentCompetitionResults
			}
		}

		key := strconv.Itoa(monthYear.Year)
		groupedMonthYears[key] = dtos.MonthYear{
			Month:    monthYear.Month,
			Year:     monthYear.Year,
			Students: grouppedStudentsCompetitionResults,
		}
	}

	return groupedMonthYears
}
