package utils

import "github.com/GabrielAchumba/school-mgt-backend/modules/result-module/dtos"

type ResultUtils interface {
	GroupByStudents(results []dtos.ResultResponse) map[string][]dtos.ResultResponse
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
			if v.Id == u {
				skip = true
				break
			}
		}
		if !skip {
			uniqueIds = append(uniqueIds, v.Id)
		}
	}

	grouppedStudentsResults := make(map[string][]dtos.ResultResponse, 0)

	for _, v := range uniqueIds {
		studentResults := make([]dtos.ResultResponse, 0)
		for _, u := range results {
			if v == u.Id {
				studentResults = append(studentResults, u)
			}
		}

		if len(studentResults) > 0 {
			grouppedStudentsResults[v] = studentResults
		}
	}

	return grouppedStudentsResults
}
