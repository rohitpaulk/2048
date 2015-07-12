package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// sentinel
var sen = 0
var winningScore = 64

type grid [][]int

func (g grid) MoveRight() {
	for i := range g {
		g.collapseRowRight(i)
		// combine
		for j := len(g[i]) - 1; j > 0; j-- {
			if g[i][j] == g[i][j-1] {
				g[i][j-1] = sen
				g[i][j] *= 2
				j--
			}
		}
		g.collapseRowRight(i)
	}
}

func (g grid) collapseRowRight(ri int) {
	row := g[ri]
	si := -1
	for j := len(row) - 1; j >= 0; j-- {
		if row[j] != sen && si >= 0 {
			// fmt.Println(j, si)
			row[j], row[si] = row[si], row[j]
			si--
		}
		if si == -1 && row[j] == sen {
			si = j
		}
	}
}

func (g grid) MoveLeft() {
	for i := range g {
		g.collapseRowLeft(i)
		// combine
		for j := 0; j < len(g[i])-1; j++ {
			if g[i][j] == g[i][j+1] {
				g[i][j+1] = sen
				g[i][j] *= 2
				j++
			}
		}
		g.collapseRowLeft(i)
	}
}

func (g grid) collapseRowLeft(ri int) {
	row := g[ri]
	si := -1
	for j := 0; j < len(row); j++ {
		if row[j] != sen && si >= 0 {
			row[j], row[si] = row[si], row[j]
			si++
		}
		if si == -1 && row[j] == sen {
			si = j
		}
	}
}

func (g grid) MoveDown() {
	for j := range g[0] {
		g.collapseColDown(j)
		for i := len(g) - 1; i > 0; i-- {
			if g[i][j] == g[i-1][j] {
				g[i][j] *= 2
				g[i-1][j] = sen
				i--
			}
		}
		g.collapseColDown(j)
	}
}

func (g grid) collapseColDown(ci int) {
	si := -1
	for i := len(g) - 1; i >= 0; i-- {
		if g[i][ci] != sen && si >= 0 {
			g[si][ci], g[i][ci] = g[i][ci], g[si][ci]
			si--
		}
		if si == -1 && g[i][ci] == sen {
			si = i
		}
	}
}

func (g grid) MoveUp() {
	for j := range g[0] {
		g.collapseColUp(j)
		for i := 0; i < len(g)-1; i++ {
			if g[i][j] == g[i+1][j] {
				g[i][j] *= 2
				g[i+1][j] = sen
				i++
			}
		}
		g.collapseColUp(j)
	}
}

func (g grid) collapseColUp(ci int) {
	si := -1
	for i := 0; i < len(g); i++ {
		if g[i][ci] != sen && si >= 0 {
			g[si][ci], g[i][ci] = g[i][ci], g[si][ci]
			si++
		}
		if si == -1 && g[i][ci] == sen {
			si = i
		}
	}
}

func (g grid) AddNumber() {
	possibles := make([][2]int, 0, len(g)*len(g[0]))

	for i := range g {
		for j := range g[i] {
			if g[i][j] == sen {
				possibles = append(possibles, [2]int{i, j})
			}
		}
	}

	x := possibles[rand.Intn(len(possibles))]
	g[x[0]][x[1]] = 2
}

func (g grid) Win() bool {
	for i := range g {
		for j := range g[i] {
			if g[i][j] == winningScore {
				return true
			}
		}
	}
	return false
}

func (g grid) Full() bool {
	for i := range g {
		for j := range g[i] {
			if g[i][j] == sen {
				return false
			}
		}
	}
	return true
}

func (g grid) String() string {
	var b bytes.Buffer
	for i := range g {
		b.Write([]byte(fmt.Sprintf("%v\n", g[i])))
	}
	return b.String()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	g := grid{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0}}
	g.AddNumber()
	g.AddNumber()

	for {
		fmt.Print(g)
		fmt.Println("1 - Up, 2 - Down, 3 - Left, 4 - Right")

		// Windows workaround
		t, errread := reader.ReadString('\n')
		i, errconv := strconv.Atoi(t[0:1])
		if errread != nil || errconv != nil || i < 1 || i > 4 {
			fmt.Println("Exiting the game", errread, errconv)
			break
		}
		switch i {
		case 1:
			g.MoveUp()
		case 2:
			g.MoveDown()
		case 3:
			g.MoveLeft()
		case 4:
			g.MoveRight()
		}
		if g.Win() {
			fmt.Printf("You won! You reached %d!\n", winningScore)
			break
		}
		g.AddNumber()
		if g.Full() {
			fmt.Println("Game over")
			break
		}
	}
}
