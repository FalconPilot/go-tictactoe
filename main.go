package main

import (
  "fmt"
)

// Grid struct
type Grid struct {
  grid [][]int
}

// Grid interface (G)
type G interface {
  Reset()
  Show()
}

// Blank grid template
func blankGrid() [][]int {
  return [][]int{
    []int{0, 0, 0},
    []int{0, 0, 0},
    []int{0, 0, 0},
  }
}

// Show a grid in the terminal
func (g Grid) Show() {
  for n, y := range g.grid {
    for i, x := range y {
      switch {
      case x == 1:
        fmt.Printf("X")
      case x == 2:
        fmt.Printf("O")
      default:
        fmt.Printf(" ")
      }
      if i < 2 {
        fmt.Printf("|")
      } else {
        fmt.Printf("\n")
      }
    }
    if n < 2 {
      fmt.Println("-----")
    }
  }
}

// Reset grid to blank state
func (g Grid) Reset() {
  g.grid = blankGrid()
}

// Main function
func main() {
  grid := Grid{blankGrid()}
  grid.Show()
}
