package inheri

type (
	Engine struct {
		HorsePower int
	}
	Wheels struct {
		Count int
	}
	Car struct {
		Engine
		Wheels
	}
)
