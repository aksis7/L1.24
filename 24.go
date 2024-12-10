package main

import (
	"fmt"
	"math"
)

// Структура Point с инкапсулированными полями x и y
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Метод для вычисления расстояния между текущей точкой и другой точкой
func (p Point) Distance(other Point) float64 {
	// Формула для вычисления Евклидова расстояния
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	// Создаем две точки с помощью конструктора
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисляем и выводим расстояние между точками
	distance := point1.Distance(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
