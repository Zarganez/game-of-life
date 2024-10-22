package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Field представляет двумерное поле клеток.
type Field struct {
	cells [][]bool
	width, height int
}

// NewField создаёт новое поле заданной ширины и высоты.
func NewField(width, height int) *Field {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &Field{cells: cells, width: width, height: height}
}

// Set устанавливает состояние клетки в заданной позиции.
func (f *Field) Set(x, y int, alive bool) {
	f.cells[y][x] = alive
}

// Alive проверяет, жива ли клетка с учётом тороидальных границ.
func (f *Field) Alive(x, y int) bool {
	x = (x + f.width) % f.width
	y = (y + f.height) % f.height
	return f.cells[y][x]
}

// Next возвращает состояние клетки на следующем шаге.
func (f *Field) Next(x, y int) bool {
	aliveNeighbors := f.countAliveNeighbors(x, y)
	if f.Alive(x, y) {
		return aliveNeighbors == 2 || aliveNeighbors == 3
	}
	return aliveNeighbors == 3
}

// countAliveNeighbors подсчитывает количество живых соседей.
func (f *Field) countAliveNeighbors(x, y int) int {
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, dir := range directions {
		if f.Alive(x+dir[0], y+dir[1]) {
			count++
		}
	}
	return count
}

// Update обновляет состояние поля на следующий шаг.
func (f *Field) Update() {
	newField := NewField(f.width, f.height)
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			newField.Set(x, y, f.Next(x, y))
		}
	}
	f.cells = newField.cells
}

// Print выводит текущее состояние поля.
func (f *Field) Print() {
	for _, row := range f.cells {
		for _, cell := range row {
			if cell {
				fmt.Print("O ") // Живая клетка
			} else {
				fmt.Print(". ") // Мертвая клетка
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Randomize заполняет поле случайными живыми клетками.
func (f *Field) Randomize(density float64) {
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			if rand.Float64() < density {
				f.Set(x, y, true)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	field := NewField(20, 10)
	field.Randomize(0.3) // 30% клеток будут живыми

	fmt.Println("Начальное поле:")
	field.Print()

	for i := 0; i < 10; i++ { // Игровой цикл, например, 10 итераций
		field.Update()
		fmt.Printf("Поколение %d:\n", i+1)
		field.Print()
		time.Sleep(time.Second / 2) // Задержка для визуализации
	}
}