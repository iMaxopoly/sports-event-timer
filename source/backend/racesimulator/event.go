package racesimulator

import (
	"sort"
	"sync"
	"time"

	"sports-event-timer/source/backend/database"
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
	Entities() *[]IEntity
	SetEntities(*[]IEntity)
	UpdateEntitiesFromTimePoint(IChip, IChip, time.Duration)
	EntityLocation(IChip) int
	EntityPosition(IChip) int
	TimePoints() *[]ITimePoint
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

// StartServerSimulatedRace starts a server simulated race using goroutines allowing
// the client to periodically fetch live data which gets stored into the database
// by the server.
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
	re.raceTrack = NewRaceTrack(FinishPoint, dummyAthletes, dummyTimePoints)

	re.timeStarted = time.Now()
	re.raceTrack.Race(&re.eventState)
	re.timeTaken = time.Since(re.timeStarted)

	// Flag Stop of the Race
	re.SetEventState(RaceNotRunning)
}

// StartClientSimulatedRace readies the server to begin client-simulation
// transactions where the client supplies event data for the server to store
// into the database.
func (re *RaceEvent) StartClientSimulatedRace() {
	// Flag Start of the Race
	re.SetEventState(RaceRunning)
}

// StopSimulatedRace sets the event state to RaceNotRunning which causes a re-initialization
// of the platform when a new race is requested to be started.
func (re *RaceEvent) StopSimulatedRace() { re.SetEventState(RaceNotRunning) }

// Entities returns a pointer to the collection of IEntities presently stored in the database.
func (re *RaceEvent) Entities() *[]IEntity {
	return helperFuncAthleteDBSliceToIEntitySlice(database.Operator.Entities())
}

// SetEntities sets up the entities provided and stores it in there database.
func (re *RaceEvent) SetEntities(entities *[]IEntity) {
	defer re.entitiesSetMutex.Unlock()
	re.entitiesSetMutex.Lock()

	database.Operator.SetEntities(helperFuncIEntitySliceToAthleteDBSlice(entities))
}

// UpdateEntitiesFromTimePoint takes in the chip indentifiers of given entitiy and timepoint
// that the entity tripped. It also takes in the time elapsed since entity started racing
// as parameter. The function then decides whether it reached within the corridor
// or has finished the race, and stores valid information into the database as such.
func (re *RaceEvent) UpdateEntitiesFromTimePoint(entityIdentifier IChip, timePointIdentifier IChip, timeElapsed time.Duration) {
	defer re.entitiesSetMutex.Unlock()
	re.entitiesSetMutex.Lock()

	for _, timePoint := range re.raceTrack.TimePoints() {
		if timePoint.Chip().Identifier() == timePointIdentifier.Identifier() {
			switch timePoint.Name() {
			case CorridorTimePoint:
				database.Operator.Update(entityIdentifier.Identifier(), &database.AthleteDBModel{
					InFinishCorridor:               true,
					TimeTakenToReachFinishCorridor: timeElapsed,
				})
				break
			case FinishLineTimePoint:
				database.Operator.Update(entityIdentifier.Identifier(), &database.AthleteDBModel{
					HasFinished:       true,
					TimeTakenToFinish: timeElapsed,
				})
				break
			}
			return
		}
	}
}

// EntityLocation takes in the chip identifier of a given entity in order
// to return its present location in the race.
func (re *RaceEvent) EntityLocation(chip IChip) int {
	for _, athlete := range *re.Entities() {
		if athlete.Chip().Identifier() == chip.Identifier() {
			return athlete.Location()
		}
	}

	return -1
}

// EntityPosition returns the position of an entity with respect to other entities
// competing in the race. It, basically, returns the relative position with respect to
// who is ahead.
func (re *RaceEvent) EntityPosition(chip IChip) int {
	entities := *re.Entities()

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

// TimePoints returns a pointer to the collection of TimePoints that are set for the
// current race.
func (re *RaceEvent) TimePoints() *[]ITimePoint {
	return helperFuncTimePointDBSliceToITimePointSlice(database.Operator.TimePoints())
}

// TimeTaken returns the time taken for the entirety of the race.
func (re *RaceEvent) TimeTaken() time.Duration { return re.timeTaken }

// EventState returns the current event state of the race;
// i.e. RaceIsRunning or RaceIsNotRunning.
func (re *RaceEvent) EventState() EventState {
	defer re.raceEventStateMutex.Unlock()
	re.raceEventStateMutex.Lock()

	return re.eventState
}

// SetEventState sets the event state of the race;
// i.e. RaceIsRunning or RaceIsNotRunning.
func (re *RaceEvent) SetEventState(state EventState) {
	defer re.raceEventStateMutex.Unlock()
	re.raceEventStateMutex.Lock()

	re.eventState = state
}

// ResetPlatform resets the platform's conditions and variables to
// re-initialize a new race by wiping the current database,
// setting up valid tables, and loading dummies into it.
func (re *RaceEvent) ResetPlatform() {
	database.Operator.ResetDB(helperFuncIEntitySliceToAthleteDBSlice(dummyEntityData()),
		helperFuncITimePointSliceToTimePointDBSlice(dummyTimePointData()))

	// Loading up the TimePoints
	dummyTimePoints := helperFuncTimePointDBSliceToITimePointSlice(database.Operator.TimePoints())

	// Loading up the Athletes
	dummyAthletes := helperFuncAthleteDBSliceToIEntitySlice(database.Operator.GetDummies())

	// Read the racetrack
	re.raceTrack = NewRaceTrack(FinishPoint, dummyAthletes, dummyTimePoints)
}
