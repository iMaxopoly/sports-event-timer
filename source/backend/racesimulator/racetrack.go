package racesimulator

// IRaceTrack is the interface that wraps the underlying racing-track-oriented methods.
// eg. when the race track is a swimming and not foot-race, the same constraints will
// need to be implemented for it to be considered an IRaceTrack.
// This helps form consistency with derivative structures.
type IRaceTrack interface {
	Distance() PointAtDistance
	setDistance(PointAtDistance)

	Athletes() []IEntity
	setAthletes([]IEntity)

	TimePoints() []ITimePoint
	setTimePoints([]ITimePoint)

	Race(*EventState)
}

// PointAtDistance is a helper type that masks an int type to help with better segregation of constant variables.
// This is currently used to declare TimePoint positions and the sprint distance for IEntity.
type PointAtDistance int

const (
	// FinishPoint is a value flag of type PointAtDistance that denotes a linear point or distance
	// where an IEntity will finish the race or where an ITimePoint resides to mark the finish line.
	FinishPoint PointAtDistance = 1000

	// CorridorPoint is a value flag of type PointAtDistance that denotes a linear point or distance
	// where an IEntity will enter the finish corridor of the race
	// or where an ITimePoint resides to mark the entrance of the finish corridor.
	CorridorPoint PointAtDistance = 800
)

type raceTrack struct {
	distance         PointAtDistance
	timePoints       []ITimePoint
	athletes         []IEntity
	inFinishCorridor []IEntity
	pastFinishLine   []IEntity
}

// Distance returns the length of the the racetrack where PointAtDistance is a wrapper over int type.
func (rt *raceTrack) Distance() PointAtDistance { return rt.distance }

func (rt *raceTrack) setDistance(distance PointAtDistance) { rt.distance = distance }

// Atheletes returns a collection of IEntity that are expected to compete in the racetrack.
func (rt *raceTrack) Athletes() []IEntity { return rt.athletes }

func (rt *raceTrack) setAthletes(athletes []IEntity) { rt.athletes = athletes }

// TimePoints returns a collection of timepoint locations where the chips register the position of competitors.
func (rt *raceTrack) TimePoints() []ITimePoint { return rt.timePoints }

func (rt *raceTrack) setTimePoints(timePoints []ITimePoint) { rt.timePoints = timePoints }

// Race is a server-simulated behavior specific function where competitors are simulated to race eachother.
func (rt *raceTrack) Race(state *EventState) {
	done := make(chan bool)
	totalEntities := len(rt.Athletes())

	for _, entity := range rt.Athletes() {
		ent := entity
		go func(ent IEntity) {
			ent.Run(state, rt.timePoints...)
			done <- true
		}(ent)
	}

	for i := 0; i < totalEntities; i++ {
		<-done
	}
}

// NewRaceTrack returns a new interface of IRaceTrack.
// It takes the distance, slice of IEntity, and a variadic input of ITimePoint.
func NewRaceTrack(distance PointAtDistance, athletes *[]IEntity, timepoints *[]ITimePoint) IRaceTrack {
	return &raceTrack{distance: distance, athletes: *athletes, timePoints: *timepoints}
}
