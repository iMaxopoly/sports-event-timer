package database

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	// verify connection
	Operator.connect()
	defer Operator.db.Close()
}

// Operations encapsulates a pointer to GORM DB that is used to interface with database activities.
type Operations struct {
	db *gorm.DB
}

var Operator Operations

// AthleteDBModel is a gorm model struct to storage athletes.
type AthleteDBModel struct {
	gorm.Model
	FullName                       string
	StartNumber                    int
	ChipIdentifier                 string
	InFinishCorridor               bool
	HasFinished                    bool
	Location                       int
	TimeTakenToReachFinishCorridor time.Duration
	TimeTakenToFinish              time.Duration
}

// TimepointDBModel is a gorm model struct to storage timepoints in the race.
type TimepointDBModel struct {
	gorm.Model
	Name           string
	Location       int
	ChipIdentifier string
}

// connect creates a database connection that must be deferred separately.
func (dbo *Operations) connect() {
	var err error
	dbo.db, err = gorm.Open("sqlite3", "RaceEventDB.db")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	dbo.db.LogMode(false)
}

// ResetPlatform drops the existing database tables and automigrates schema.
// Thereby, it also populates the underlying tables with dummy data.
func (dbo *Operations) ResetDB(dummyAthletes []AthleteDBModel, dummyTimePoints []TimepointDBModel) {
	Operator.connect()
	defer Operator.db.Close()

	Operator.db.DropTableIfExists(&AthleteDBModel{}, &TimepointDBModel{})
	Operator.db.AutoMigrate(&AthleteDBModel{}, &TimepointDBModel{})
	Operator.SetupDummies(dummyAthletes)
	Operator.SetupTimePoints(dummyTimePoints)
}

// Update updates the a single athlete row by comparing the supplied chip identifier string.
// All fields that are changed, get updated.
func (dbo *Operations) Update(chipIdentifier string, athlete AthleteDBModel) {
	dbo.connect()
	defer dbo.db.Close()

	var val AthleteDBModel
	dbo.db.Model(&val).Where(&AthleteDBModel{ChipIdentifier: chipIdentifier}).Update(athlete)
}

// Entities returns a slice of IEntity which are actually all the athletes currently stored.
func (dbo *Operations) Entities() []AthleteDBModel {
	dbo.connect()
	defer dbo.db.Close()

	var valDB []AthleteDBModel
	dbo.db.Find(&valDB)

	return valDB
}

// SetEntities takes in a slice of IEntity and updates the currently stored athletes
// based on observable changes.
func (dbo Operations) SetEntities(entities []AthleteDBModel) {
	dbo.connect()
	defer dbo.db.Close()

	for _, valDB := range entities {
		var val AthleteDBModel
		dbo.db.Model(&val).Where(&AthleteDBModel{ChipIdentifier: valDB.ChipIdentifier}).Update(valDB)
	}
}

// TimePoints returns a slice of ITimePoint which are all the stored TimepointDBModel values.
// This helps track the location of the timepoint as well as the respective chip indentifiers.
func (dbo *Operations) TimePoints() []TimepointDBModel {
	dbo.connect()
	defer dbo.db.Close()

	var timePointsDB []TimepointDBModel
	dbo.db.Find(&timePointsDB)

	return timePointsDB
}

// SetupDummies sets up dummy values for the athletes and stores it in the database
func (dbo *Operations) SetupDummies(dummyAthletes []AthleteDBModel) {
	dbo.connect()
	defer dbo.db.Close()

	for _, val := range dummyAthletes {
		dbo.db.Create(&val)
	}
}

// GetDummies gets a slice of IEntity which are actually all the dummy values currently stored in the database.
// It also modifies the speed to give it a dynamic random value for proper simulation.
func (dbo *Operations) GetDummies() []AthleteDBModel {
	dbo.connect()
	defer dbo.db.Close()

	var valDB []AthleteDBModel
	dbo.db.Find(&valDB)

	return valDB
}

// SetupTimePoints sets up dummy values for the timepoints and stores it in the database.
func (dbo *Operations) SetupTimePoints(dummyTimePoints []TimepointDBModel) {
	dbo.connect()
	defer dbo.db.Close()

	for _, val := range dummyTimePoints {
		dbo.db.Create(&val)
	}
}
