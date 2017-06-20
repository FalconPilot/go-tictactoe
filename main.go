package main

import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "runtime"
  "strconv"
)

// Clearing shell
var clear map[string]string

func clearShell() {
  cmd, ok := clear[runtime.GOOS]
  if ok {
    cmd := exec.Command(cmd)
    cmd.Stdout = os.Stdout
    cmd.Run()
  } else {
    panic("Unsupported platform !")
  }
}

func init() {
  clear = make(map[string]string)
  clear["linux"] = "clear"
  clear["darwin"] = "clear"
  clear["windows"] = "cls"
}

// Grid struct
type Grid struct {
  grid [][]int
}

// Grid interface (G)
type G interface {
  CheckVictory() int
  Full() bool
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

// Slice contains int ?

func contains(s []int, n int) bool {
  for _, x := range s {
    if x == n {
      return true
    }
  }
  return false
}

// Check if someone has won
func (g Grid) CheckVictory() int {

  // Check rows
  for _, y := range g.grid {
    if y[0] == y[1] && y[1] == y[2] {
      return y[0]
    }
  }

  // Set grid alias
  x := g.grid

  // Check columns
  for i := 0; i < 3; i++ {
    if x[0][i] == x[1][i] && x[1][i] == x[2][i] {
      return x[0][i]
    }
  }

  // Check diagonals
  if x[0][0] == x[1][1] && x[1][1] == x[2][2] || x[0][2] == x[1][1] && x[1][1] == x[2][0] {
    return x[1][1]
  }

  // No victory
  return 0
}

// Check if grid is full
func (g Grid) Full() bool {
  for _, y := range g.grid {
    if contains(y, 0) {
      return false
    }
  }
  return true
}

// Show a grid in the terminal
func (g Grid) Show() {
  clearShell()
  for n, y := range g.grid {
    for i, x := range y {
      switch {
      case x == 1:
        fmt.Print("X")
      case x == 2:
        fmt.Print("O")
      default:
        fmt.Printf("%v", (n * 3 + i + 1))
      }
      if i < 2 {
        fmt.Print("|")
      } else {
        fmt.Print("\n")
      }
    }
    if n < 2 {
      fmt.Println("-----")
    }
  }
}

// Inject a value in the grid
func (g Grid) Inject(char string, index int) {
  value := 0
  if char == "X" {
    value = 1
  } else if char == "O" {
    value = 2
  }
  i := 0
  for ;index > 2; i++ {
    index -= 3
  }
  g.grid[i][index] = value
}

// Reset grid to blank state
func (g Grid) Reset() {
  g.grid = blankGrid()
}

// Main function
func main() {
  buf := bufio.NewReader(os.Stdin)
  grid := Grid{blankGrid()}
  turn := "X"

  for {
    // Grid display
    grid.Show()
    fmt.Printf("Place a %v\n> ", turn)
    input, err := buf.ReadBytes('\n')
    if err != nil {
      fmt.Println(err)
    } else {
      // No error, regular program behavior
      fmt.Println(string(input[0]))
      value, err := strconv.ParseInt(string(input[0]), 10, 0)
      if len(input) != 2 || err != nil || value <= 0 {
        fmt.Println("Please enter a valid number [1-9]")
      } else {

        // Inject new value
        grid.Inject(turn, int(value) - 1)

        // Check victory conditions
        if v := grid.CheckVictory(); v != 0 {
          grid.Show()
          fmt.Printf("Victoire de %v!\n", turn)
          break

        // Check if grid is full
        } else if grid.Full() {
          grid.Show()
          fmt.Println("Match nul !")
          break
        }

        // Switch turn
        if turn == "X" {
          turn = "O"
        } else {
          turn = "X"
        }
      }
    }
  }
}
