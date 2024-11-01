package domain

type Card struct {
	Id    string
	Color Color
	Hand  int
}
type Color string

const (
	Red   Color = "RED"
	Green Color = "GREEN"
	Blue  Color = "BLUE"
)

type Set struct {
	Cards []Card
}
