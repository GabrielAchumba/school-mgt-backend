package utils

import (
	"sort"

	"github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module/dtos"
)

type CompetitionResultsSorterByOverallScore []dtos.StudentCompetitionResults

func (t CompetitionResultsSorterByOverallScore) Len() int      { return len(t) }
func (t CompetitionResultsSorterByOverallScore) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t CompetitionResultsSorterByOverallScore) Less(i, j int) bool {
	return t[i].OverallScore > t[j].OverallScore
}

func SortPositionCompetitionResults(CompetitionResults []dtos.StudentCompetitionResults) []dtos.StudentCompetitionResults {

	sort.Sort(CompetitionResultsSorterByOverallScore(CompetitionResults))
	return CompetitionResults
}
