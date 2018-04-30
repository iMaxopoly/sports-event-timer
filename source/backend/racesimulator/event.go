package racesimulator

import (
	"sort"
	"sync"
	"time"

	"sports-event-timing/source/backend/database"
)

// EventState is a helper type that masks an int type to help with better segregation of constant variables
type EventState int

const (
	// RaceRunning is an event flag of type EventState that signifies that a simulation is underway
	RaceRunning EventState = iota
	// RaceNotRunning is an event flag of type EventState that signifies that a simulation is not underway
	RaceNotRunning
)

// IEvent is the interface that wraps the underlying event-oriented methods
// as a set of guidelines for an event-type function to follow.
// This helps form consistency with derivative structures.
type IEvent interface {
	StartServerSimulatedRace()
	StartClientSimulatedRace()
	StopSimulatedRace()
	Entities() []IEntity
	SetEntities([]IEntity)
	UpdateEntitiesFromTimePoint(IChip, IChip, time.Duration)
	EntityLocation(IChip) int
	EntityPosition(IChip) int
	TimePoints() []ITimePoint
	TimeTaken() time.Duration
	EventState() EventState
	SetEventState(EventState)
	ResetPlatform()
}

// RaceEvent is an implementation struct based on IEvent principles.
type RaceEvent struct {
	eventState          EventState
	raceTrack           IRaceTrack
	timeStarted         time.Time
	timeTaken           time.Duration
	raceEventStateMutex sync.Mutex
	entitiesSetMutex    sync.Mutex
}

func (re *RaceEvent) StartServerSimulatedRace() {
	// Ensuring singular instance
	if re.EventState() == RaceRunning {
		return
	}

	// Flag Start of the Race
	re.SetEventState(RaceRunning)

	// Loading up the TimePoints
	dummyTimePoints := helperFuncTimePointDBSliceToITimePointSlice(database.Operator.TimePoints())

	// Loading up the Athletes
	dummyAthletes := helperFuncAthleteDBSliceToIEntitySlice(database.Operator.GetDummies())

	// Read the racetrack
	re.raceTrack = NewRaceTrack(FinishPoint, dummyAthletes, dummyTimePoints...)

	re.timeStarted = time.Now()
	re.raceTrack.Race(&re.eventState)
	re.timeTaken = time.Since(re.timeStarted)

	// Flag Stop of the Race
	re.SetEventState(RaceNotRunning)
}

func (re *RaceEvent) StartClientSimulatedRace() {
	// Flag Start of the Race
	re.SetEventState(RaceRunning)
}

func (re *RaceEvent) StopSimulatedRace() { re.SetEventState(RaceNotRunning) }

func (re *RaceEvent) Entities() []IEntity {
	return helperFuncAthleteDBSliceToIEntitySlice(database.Operator.Entities())
}

func (re *RaceEvent) SetEntities(entities []IEntity) {
	defer re.entitiesSetMutex.Unlock()
	re.entitiesSetMutex.Lock()

	database.Operator.SetEntities(helperFuncIEntitySliceToAthleteDBSlice(entities))
}

func (re *RaceEvent) UpdateEntitiesFromTimePoint(entityIdentifier IChip, timePointIdentifier IChip, timeElapsed time.Duration) {
	defer re.entitiesSetMutex.Unlock()
	re.entitiesSetMutex.Lock()

	for _, timePoint := range re.raceTrack.TimePoints() {
		if timePoint.Chip().Identifier() == timePointIdentifier.Identifier() {
			switch timePoint.Name() {
			case CorridorTimePoint:
				database.Operator.Update(entityIdentifier.Identifier(), database.AthleteDBModel{
					InFinishCorridor:               true,
					TimeTakenToReachFinishCorridor: timeElapsed,
				})
				break
			case FinishLineTimePoint:
				database.Operator.Update(entityIdentifier.Identifier(), database.AthleteDBModel{
					HasFinished:       true,
					TimeTakenToFinish: timeElapsed,
				})
				break
			}
			return
		}
	}
}

func (re *RaceEvent) EntityLocation(chip IChip) int {
	for _, athlete := range re.Entities() {
		if athlete.Chip().Identifier() == chip.Identifier() {
			return athlete.Location()
		}
	}

	return -1
}

func (re *RaceEvent) EntityPosition(chip IChip) int {
	entities := re.Entities()

	sort.Slice(entities, func(i, j int) bool {
		return entities[i].Location() > entities[i].Location()
	})

	for i, athlete := range entities {
		if athlete.Chip().Identifier() == chip.Identifier() {
			return i
		}
	}

	return -1
}

func (re *RaceEvent) TimePoints() []ITimePoint {
	return helperFuncTimePointDBSliceToITimePointSlice(database.Operator.TimePoints())
}

func (re *RaceEvent) TimeTaken() time.Duration { return re.timeTaken }

func (re *RaceEvent) EventState() EventState {
	defer re.raceEventStateMutex.Unlock()
	re.raceEventStateMutex.Lock()

	return re.eventState
}

func (re *RaceEvent) SetEventState(state EventState) {
	defer re.raceEventStateMutex.Unlock()
	re.raceEventStateMutex.Lock()

	re.eventState = state
}

func (re *RaceEvent) ResetPlatform() {
	database.Operator.ResetDB(helperFuncIEntitySliceToAthleteDBSlice(dummyEntityData()),
		helperFuncITimePointSliceToTimePointDBSlice(dummyTimePointData()))

	// Loading up the TimePoints
	dummyTimePoints := helperFuncTimePointDBSliceToITimePointSlice(database.Operator.TimePoints())

	// Loading up the Athletes
	dummyAthletes := helperFuncAthleteDBSliceToIEntitySlice(database.Operator.GetDummies())

	// Read the racetrack
	re.raceTrack = NewRaceTrack(FinishPoint, dummyAthletes, dummyTimePoints...)
}
