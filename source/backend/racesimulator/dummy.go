package racesimulator

// dummyEntityData returns a slice of generated dummy IEntity.
func dummyEntityData() []IEntity {
	// Dummy Athletes
	var (
		manish     = NewEntity("Manish Singh", 1, FinishPoint)
		madhushree = NewEntity("Madhushree Singh", 2, FinishPoint)
		siim       = NewEntity("Siim Kaspar Uustalu", 3, FinishPoint)
		eliisabeth = NewEntity("Eliisabeth Käbin", 4, FinishPoint)
		ahti       = NewEntity("Ahti Liin", 5, FinishPoint)
	)

	// Ready the Athletes
	var athletes = append([]IEntity{},
		manish,
		madhushree,
		siim,
		eliisabeth,
		ahti,
	)

	return athletes
}

// dummyTimePointData returns a slice of generated dummy ITimePoint.
func dummyTimePointData() []ITimePoint {
	var timePoints []ITimePoint
	timePoints = append(timePoints,
		NewTimePoint(CorridorTimePoint, CorridorPoint),
		NewTimePoint(FinishLineTimePoint, FinishPoint))

	return timePoints
}
