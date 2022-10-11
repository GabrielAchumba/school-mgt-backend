package utils

import (
	"sort"

	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/dtos"
)

type ResultsSorterByOverallScore []dtos.StudentResults

func (t ResultsSorterByOverallScore) Len() int      { return len(t) }
func (t ResultsSorterByOverallScore) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t ResultsSorterByOverallScore) Less(i, j int) bool {
	return t[i].OverallScore > t[j].OverallScore
}

func SortPositionResults(results []dtos.StudentResults) []dtos.StudentResults {

	sort.Sort(ResultsSorterByOverallScore(results))
	return results
}
