package composition

type engine struct{
	hp int
}

type wheel struct{
	wheelDiemensions int
}


/* This is how we can achieve inheritence in go basically has a relatioship */
type Car struct{
	CarName string
	Engine engine
	Wheel wheel
}

func NewCar(carName string, hp, wd int) Car{
	return Car{
		CarName: carName,
		Engine: engine {
			hp:hp,
		},
		Wheel : wheel {
			wheelDiemensions: wd,
		},
	}
}