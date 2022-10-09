package utils

import (
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/dtos"
)

type ResultUtils interface {
	GroupByStudents(results []dtos.ResultResponse) map[string][]dtos.ResultResponse
	GroupByStudentsAndMonthYear(results []dtos.ResultResponse,
		monthYears []dtos.MonthYear) map[string]dtos.MonthYear
	GroupByStudentsAndYear(results []dtos.ResultResponse,
		monthYears []dtos.MonthYear) map[string]dtos.MonthYear
}

type ResultUtilsImpl struct {
}

func New() ResultUtils {
	return &ResultUtilsImpl{}
}

func (impl ResultUtilsImpl) GroupByStudents(results []dtos.ResultResponse) map[string][]dtos.ResultResponse {

	var uniqueIds []string

	for _, v := range results {
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

	grouppedStudentsResults := make(map[string][]dtos.ResultResponse, 0)

	for _, v := range uniqueIds {
		studentResults := make([]dtos.ResultResponse, 0)
		for _, u := range results {
			if v == u.StudentId {
				studentResults = append(studentResults, u)
			}
		}

		if len(studentResults) > 0 {
			grouppedStudentsResults[v] = studentResults
		}
	}

	return grouppedStudentsResults
}

func (impl ResultUtilsImpl) GroupByStudentsAndMonthYear(results []dtos.ResultResponse,
	monthYears []dtos.MonthYear) map[string]dtos.MonthYear {

	var uniqueIds []string

	for _, v := range results {
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

		grouppedStudentsResults := make(map[string][]dtos.ResultResponse, 0)

		for _, v := range uniqueIds {
			studentResults := make([]dtos.ResultResponse, 0)
			for _, u := range results {
				if v == u.Id &&
					monthYear.Month == int(u.CreatedAt.Month()) &&
					monthYear.Year == u.CreatedAt.Year() {
					studentResults = append(studentResults, u)
				}
			}

			if len(studentResults) > 0 {
				grouppedStudentsResults[v] = studentResults
			}
		}

		key := strconv.Itoa(monthYear.Month) + "_" + strconv.Itoa(monthYear.Year)
		groupedMonthYears[key] = dtos.MonthYear{
			Month:    monthYear.Month,
			Year:     monthYear.Year,
			Students: grouppedStudentsResults,
		}
	}

	return groupedMonthYears
}

func (impl ResultUtilsImpl) GroupByStudentsAndYear(results []dtos.ResultResponse,
	monthYears []dtos.MonthYear) map[string]dtos.MonthYear {

	var uniqueIds []string

	for _, v := range results {
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

		grouppedStudentsResults := make(map[string][]dtos.ResultResponse, 0)

		for _, v := range uniqueIds {
			studentResults := make([]dtos.ResultResponse, 0)
			for _, u := range results {
				if v == u.Id &&
					monthYear.Year == u.CreatedAt.Year() {
					studentResults = append(studentResults, u)
				}
			}

			if len(studentResults) > 0 {
				grouppedStudentsResults[v] = studentResults
			}
		}

		key := strconv.Itoa(monthYear.Year)
		groupedMonthYears[key] = dtos.MonthYear{
			Month:    monthYear.Month,
			Year:     monthYear.Year,
			Students: grouppedStudentsResults,
		}
	}

	return groupedMonthYears
}
