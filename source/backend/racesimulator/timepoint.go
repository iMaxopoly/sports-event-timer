package racesimulator

type ITimePoint interface {
	// Name is the name of the timepoint which denotes where the timepoint resides. TimePointName is a simple
	// wrapper over string type.
	Name() TimePointName
	setName(TimePointName)

	// Chip returns the valid chip information specific to the current timepoint. It contains a unique identifier.
	Chip() IChip
	setChip(IChip)

	// Location returns the location of the current timepoint. PointAtDistance being a wrapper over int type.
	Location() PointAtDistance
	setLocation(PointAtDistance)
}

// TimePointName is a helper type that masks a string type to help with better segregation of constant variables.
// This is currently used to declare TimePoint names.
type TimePointName string

const (
	CorridorTimePoint   TimePointName = "Corridor Timepoint"
	FinishLineTimePoint TimePointName = "Finish Line Timepoint"
)

type timePoint struct {
	name     TimePointName
	chip     IChip
	location PointAtDistance
}

// Name is the name of the timepoint which denotes where the timepoint resides. TimePointName is a simple
// wrapper over string type.
func (tp *timePoint) Name() TimePointName { return tp.name }

func (tp *timePoint) setName(name TimePointName) { tp.name = name }

// Chip returns the valid chip information specific to the current timepoint. It contains a unique identifier.
func (tp *timePoint) Chip() IChip { return tp.chip }

func (tp *timePoint) setChip(chip IChip) { tp.chip = chip }

// Location returns the location of the current timepoint. PointAtDistance being a wrapper over int type.
func (tp *timePoint) Location() PointAtDistance { return tp.location }

func (tp *timePoint) setLocation(location PointAtDistance) { tp.location = location }

// NewTimePoint returns a new interface of ITimePoint.
// It takes in a name for the timepoint, and a location of the timepoint.
// A new chip is then embedded into the resultant interface of timepoint with a unique identifier.
func NewTimePoint(name TimePointName, location PointAtDistance) ITimePoint {
	return &timePoint{chip: NewChip(), name: name, location: location}
}
