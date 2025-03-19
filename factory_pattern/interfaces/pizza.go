package interfaces

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}
