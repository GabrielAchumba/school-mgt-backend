package helpers

import (
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cycle-module/models"
)

type CyclesUtils struct {
}

func NewCyclesUtils() CyclesUtils {

	return CyclesUtils{}
}

func (impl CyclesUtils) GetListOfStream() []models.Stream {

	StreamNames := Streams
	LevelCounts := LevelCounts
	var streams []models.Stream
	var stream models.Stream
	for i := 0; i < 1; i++ {

		stream = models.Stream{}
		stream.Id = i + 1
		stream.Label = StreamNames[i]

		var children []models.PsuedoLevel
		psuedoLevel := models.PsuedoLevel{}
		LevelCount := LevelCounts[i]

		for j := 0; j < LevelCount; j++ {
			psuedoLevel = models.PsuedoLevel{
				Id:    strconv.Itoa(j+1) + "-" + StreamNames[i],
				Label: "Level-" + strconv.Itoa(j+1),
			}

			children = append(children, psuedoLevel)
		}

		stream.Children = children

		streams = append(streams, stream)
	}

	return streams
}
