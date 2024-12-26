package main

import (
	"errors"
	"fmt"
	"math"
)

// ----------------------- Задание 1: Массивы и срезы -----------------------

func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func listEven(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("левая граница больше правой")
	}

	var evens []int
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return evens, nil
}

// ----------------------- Задание 2: Карты -----------------------

func countCharacters(s string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	return charCount
}

// ----------------------- Задание 3: Структуры, методы и интерфейсы -----------------------

type Point struct {
	X, Y float64
}

type LineSegment struct {
	Start, End Point
}

func (ls LineSegment) Length() float64 {
	dx := ls.End.X - ls.Start.X
	dy := ls.End.Y - ls.Start.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// ----------------------- Задание 3: Треугольник и Круг -----------------------

type Triangle struct {
	A, B, C Point
}

type Circle struct {
	Center Point
	Radius float64
}

func (t Triangle) Area() float64 {
	a := distance(t.A, t.B)
	b := distance(t.B, t.C)
	c := distance(t.C, t.A)
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt((p2.X-p1.X)*(p2.X-p1.X) + (p2.Y-p1.Y)*(p2.Y-p1.Y))
}

type Shape interface {
	Area() float64
}

func printArea(s Shape, figureName string) {
	result := s.Area()
	fmt.Printf("Площадь %s: %.2f\n", figureName, result)
}

func main() {
	// ----------------------- Задание 1: Массивы и срезы -----------------------

	ip := [4]byte{192, 168, 0, 1}
	fmt.Println("IP-адрес:", formatIP(ip))
	evens, err := listEven(3, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Четные числа в диапазоне 3-10:", evens)
	}

	evens, err = listEven(10, 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Четные числа:", evens)
	}

	// ----------------------- Задание 2: Карты -----------------------

	str := "hello world"
	charCount := countCharacters(str)
	fmt.Println("Подсчёт символов в строке 'hello world':")
	for char, count := range charCount {
		fmt.Printf("'%c' встречается %d раз(а)\n", char, count)
	}

	// ----------------------- Задание 3: Структуры, методы и интерфейсы -----------------------

	startPoint := Point{X: 0, Y: 0}
	endPoint := Point{X: 3, Y: 4}

	line := LineSegment{Start: startPoint, End: endPoint}

	// Вычисляем длину отрезка
	length := line.Length()
	fmt.Printf("Длина отрезка с началом в (%.1f, %.1f) и концом в (%.1f, %.1f) равна %.2f\n",
		line.Start.X, line.Start.Y, line.End.X, line.End.Y, length) // вывод: 5.00

	// ----------------------- Задание 3: Треугольник и Круг -----------------------

	triangle := Triangle{
		A: Point{X: 0, Y: 0},
		B: Point{X: 3, Y: 0},
		C: Point{X: 0, Y: 4},
	}

	circle := Circle{
		Center: Point{X: 0, Y: 0},
		Radius: 5,
	}

	printArea(triangle, "треугольника")

	printArea(circle, "круга")

	// ----------------------- Задание 4: Функциональное программирование -----------------------

	Map := func(slice []float64, fn func(float64) float64) []float64 {
		result := make([]float64, len(slice))
		for i, v := range slice {
			result[i] = fn(v)
		}
		return result
	}

	square := func(x float64) float64 {
		return x * x
	}

	values := []float64{1, 2, 3, 4, 5}

	fmt.Println("Исходный срез:", values)

	squaredValues := Map(values, square)

	fmt.Println("Срез после возведения в квадрат:", squaredValues)
}
