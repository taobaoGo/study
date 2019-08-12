package typego

type BagOperation interface {
	AddItem()
	DeleteItem()
	ChangeColour()
	Dispose()
}
