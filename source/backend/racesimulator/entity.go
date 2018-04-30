package racesimulator

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"sports-event-timing/source/backend/database"
)

var (
	errEntityFullName              = errors.New("entity full Name is invalid")
	errEntityChipIdentifier        = errors.New("entity chip identifier is invalid")
	errEntityInvalidLocation       = errors.New("entity starting location is invalid")
	errEntitySpeedInvalid          = errors.New("entity speed is invalid")
	errEntitySprintDistanceInvalid = errors.New("entity sprint distance is invalid")
)

func init() {
	// seeding rand
	rand.Seed(time.Now().UTC().UnixNano())
}

// IEntity is the interface that wraps the underlying entity-oriented methods.
// eg. when the dummy entity is a car and not an athlete, the same constraints will
// need to be implemented for it to be considered an IEntity.
// This helps form consistency with derivative structures.
type IEntity interface {
	// EntityName returns the entity's name.
	EntityName() string
	setEntityName(string)

	// StartNumber returns the entity's start number.
	StartNumber() int
	setStartNumber(int)

	// Chip returns the chip that the entity is carrying.
	Chip() IChip
	setChip(IChip)

	// Speed returns the entity's speed.
	Speed() int
	setSpeed(int)

	// Location returns the entity's current location.
	Location() int
	setLocation(int)

	// SprintDistance returns the distance that the entity will race for.
	SprintDistance() PointAtDistance
	setSprintDistance(PointAtDistance)

	// InFinishCorridor returns whether the entity is in the finish corridor or not.
	InFinishCorridor() bool

	// HasFinished returns whether the entity has finished the race or not.
	HasFinished() bool

	// TimeTakenToReachFinishCorridor returns the time in which the entity reached the finish corridor.
	TimeTakenToReachFinishCorridor() time.Duration

	// TimeTakenToFinish returns the time in which the entity reached the finish line.
	TimeTakenToFinish() time.Duration

	// EnsureEntity verifies that the given entity is valid for race conditions.
	EnsureEntity() error

	// Run handles the race simulation for the entities competing. Whenever an entity trips
	// a timepoint, their information is stored in the database.
	Run(*EventState, ...ITimePoint)
}

// randInt is a helper function to generate a random int value between a min and a max.
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// athlete is an implementation struct based on IEntity principles.
type athlete struct {
	fullName                       string
	startNumber                    int
	chip                           IChip
	location                       int
	speed                          int
	sprintTo                       PointAtDistance
	inFinishCorridor               bool
	hasFinished                    bool
	timeStart                      time.Time
	timeTakenToReachFinishCorridor time.Duration
	timeTakenToFinish              time.Duration
}

func (ath *athlete) EntityName() string { return ath.fullName }

func (ath *athlete) setEntityName(name string) { ath.fullName = name }

func (ath *athlete) StartNumber() int { return ath.startNumber }

func (ath *athlete) setStartNumber(number int) { ath.startNumber = number }

func (ath *athlete) Chip() IChip { return ath.chip }

func (ath *athlete) setChip(chip IChip) { ath.chip = chip }

func (ath *athlete) Speed() int { return ath.speed }

func (ath *athlete) setSpeed(speed int) { ath.speed = speed }

func (ath *athlete) Location() int { return ath.location }

func (ath *athlete) setLocation(location int) { ath.location = location }

func (ath *athlete) SprintDistance() PointAtDistance { return ath.sprintTo }

func (ath *athlete) setSprintDistance(sprintTo PointAtDistance) { ath.sprintTo = sprintTo }

func (ath *athlete) InFinishCorridor() bool { return ath.inFinishCorridor }

func (ath *athlete) HasFinished() bool { return ath.hasFinished }

func (ath *athlete) TimeTakenToFinish() time.Duration { return ath.timeTakenToFinish }

func (ath *athlete) TimeTakenToReachFinishCorridor() time.Duration {
	return ath.timeTakenToReachFinishCorridor
}

// EnsureEntity ensures that the entity passed in the parameter is valid for the race.
func (ath *athlete) EnsureEntity() error {
	if strings.TrimSpace(ath.fullName) == "" {
		return errEntityFullName
	}

	if strings.TrimSpace(ath.chip.Identifier()) == "" {
		return errEntityChipIdentifier
	}

	if ath.location < 0 {
		return errEntityInvalidLocation
	}

	if ath.speed <= 0 {
		return errEntitySpeedInvalid
	}

	if ath.sprintTo <= 0 {
		return errEntitySprintDistanceInvalid
	}

	return nil
}

// Run handles the server-based race simulation for the entities competing. Whenever an entity trips
// a timepoint, their information is stored in the database.
func (ath *athlete) Run(state *EventState, timePoints ...ITimePoint) {
	ath.timeStart = time.Now()
	for int(ath.location) <= int(ath.sprintTo) {
		if *state == RaceNotRunning {
			break
		}

		time.Sleep(500 * time.Millisecond)
		ath.location += randInt(10, 16)

		for _, t := range timePoints {
			switch t.Name() {
			default:
				break
			case CorridorTimePoint:
				if int(ath.location) >= int(t.Location()) && !ath.inFinishCorridor {
					ath.inFinishCorridor = true
					ath.timeTakenToReachFinishCorridor = time.Since(ath.timeStart) / time.Millisecond
					database.Operator.Update(ath.Chip().Identifier(), database.AthleteDBModel{
						InFinishCorridor:               ath.inFinishCorridor,
						TimeTakenToReachFinishCorridor: ath.timeTakenToReachFinishCorridor,
						Location:                       ath.location,
					})
					fmt.Println(fmt.Sprintf("%v with ID: %v in finish corridor - took %v ms",
						ath.fullName, ath.chip.Identifier(), float64(ath.timeTakenToReachFinishCorridor)))
				} else {
					database.Operator.Update(ath.Chip().Identifier(), database.AthleteDBModel{
						Location: ath.location,
					})
				}
				break
			case FinishLineTimePoint:
				if int(ath.location) >= int(t.Location()) && !ath.hasFinished {
					ath.hasFinished = true
					ath.timeTakenToFinish = time.Since(ath.timeStart) / time.Millisecond
					database.Operator.Update(ath.Chip().Identifier(), database.AthleteDBModel{
						HasFinished:       ath.hasFinished,
						TimeTakenToFinish: ath.timeTakenToFinish,
						Location:          ath.location,
					})
					fmt.Println(fmt.Sprintf("%v with ID: %v in finish line - took %v ms",
						ath.fullName, ath.chip.Identifier(), float64(ath.timeTakenToFinish)))
				} else {
					database.Operator.Update(ath.Chip().Identifier(), database.AthleteDBModel{
						Location: ath.location,
					})
				}
				break
			}
		}
	}
}

// NewEntity returns a new interface of IEntity.
// It takes the fullname, starting number and sprint distance as arguments
// and are then ensured to be correct before returning a new IEntity.
func NewEntity(fullName string, startNumber int, sprintTo PointAtDistance) IEntity {
	ath := &athlete{
		fullName:                       fullName,
		startNumber:                    startNumber,
		chip:                           NewChip(),
		location:                       0,
		speed:                          1,
		sprintTo:                       sprintTo,
		inFinishCorridor:               false,
		hasFinished:                    false,
		timeTakenToReachFinishCorridor: -1,
		timeTakenToFinish:              -1,
	}

	err := ath.EnsureEntity()
	if err != nil {
		log.Fatal(err)
	}

	return ath
}
