package entities

import "fmt"

type TileContent uint8

const (
	Void TileContent = iota
	X
	O
)

type Board struct {
	FreeTiles uint8
	Tiles     [3][3]TileContent
}

func (b *Board) Init() {
	b.FreeTiles = 9
}

func (b Board) Print() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch b.Tiles[i][j] {
			case X:
				fmt.Print(" X")
			case O:
				fmt.Print(" O")
			default:
				fmt.Print("  ")
			}
			if j != 2 {
				fmt.Print(" |")
			}
		}
		if i != 2 {
			fmt.Print("\n-----------\n")
		}
	}
}
