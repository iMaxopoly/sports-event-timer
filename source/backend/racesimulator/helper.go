package racesimulator

import (
	"sports-event-timing/source/backend/database"
)

// helperFuncAthleteDBSliceToIEntitySlice helps in converting from:
// []database.AthleteDBModel to []IEntity type, helping with database value extractions
func helperFuncAthleteDBSliceToIEntitySlice(athletesDBSlice []database.AthleteDBModel) []IEntity {
	var entitySlice []IEntity

	for _, athleteDB := range athletesDBSlice {
		var entity athlete
		entity.setEntityName(athleteDB.FullName)
		entity.setStartNumber(athleteDB.StartNumber)
		var newChip Chip
		newChip.SetIdentifier(athleteDB.ChipIdentifier)
		entity.setChip(&newChip)
		entity.setSprintDistance(FinishPoint)
		entity.setSpeed(1)
		entity.timeTakenToReachFinishCorridor = athleteDB.TimeTakenToReachFinishCorridor
		entity.timeTakenToFinish = athleteDB.TimeTakenToFinish
		entity.inFinishCorridor = athleteDB.InFinishCorridor
		entity.hasFinished = athleteDB.HasFinished
		entity.location = athleteDB.Location

		entitySlice = append(entitySlice, &entity)
	}

	return entitySlice
}

// helperFuncIEntitySliceToAthleteDBSlice helps in converting from:
// []IEntity to []database.AthleteDBModel type, helping with database value extractions
func helperFuncIEntitySliceToAthleteDBSlice(entitiesSlice []IEntity) []database.AthleteDBModel {
	var athleteDBModelSlice []database.AthleteDBModel

	for _, entity := range entitiesSlice {
		var ath = database.AthleteDBModel{
			FullName:                       entity.EntityName(),
			StartNumber:                    entity.StartNumber(),
			ChipIdentifier:                 entity.Chip().Identifier(),
			InFinishCorridor:               entity.InFinishCorridor(),
			HasFinished:                    entity.HasFinished(),
			TimeTakenToReachFinishCorridor: entity.TimeTakenToReachFinishCorridor(),
			TimeTakenToFinish:              entity.TimeTakenToFinish(),
			Location:                       entity.Location(),
		}

		athleteDBModelSlice = append(athleteDBModelSlice, ath)
	}

	return athleteDBModelSlice
}

// helperFuncTimePointDBSliceToITimePointSlice helps in converting from:
// []database.TimepointDBModel to []ITimePoint type, helping with database value extractions
func helperFuncTimePointDBSliceToITimePointSlice(timePointsDBSlice []database.TimepointDBModel) []ITimePoint {
	var timePointsSlice []ITimePoint

	for _, t := range timePointsDBSlice {
		var tp timePoint

		tp.setName(TimePointName(t.Name))
		tp.setLocation(PointAtDistance(t.Location))
		var newChip Chip
		newChip.SetIdentifier(t.ChipIdentifier)
		tp.setChip(&newChip)

		timePointsSlice = append(timePointsSlice, &tp)
	}

	return timePointsSlice
}

// helperFuncITimePointSliceToTimePointDBSlice helps in converting from:
// []ITimePoint to []database.TimepointDBModel type, helping with database value extractions
func helperFuncITimePointSliceToTimePointDBSlice(timePointsSlice []ITimePoint) []database.TimepointDBModel {
	var timePointDBModelSlice []database.TimepointDBModel

	for _, timePoint := range timePointsSlice {
		var tp = database.TimepointDBModel{
			Name:           string(timePoint.Name()),
			Location:       int(timePoint.Location()),
			ChipIdentifier: timePoint.Chip().Identifier(),
		}

		timePointDBModelSlice = append(timePointDBModelSlice, tp)
	}

	return timePointDBModelSlice
}
