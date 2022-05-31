package colors

import "fmt"

// Generator - рандомный генератор кода цвета консоли
type Generator interface {
	NextColor() string
}

type generator struct {
	colorIndex int
}

// NewGenerator - конструктор генератора кодов цветов
func NewGenerator() Generator {
	return &generator{}
}

func (g *generator) NextColor() string {
	color := g.colorIndex%14 + 31
	if color > 37 {
		color = color + 90 - 37
	}

	g.colorIndex++

	return fmt.Sprintf("%dm", color)
}
