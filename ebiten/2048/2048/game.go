// Copyright 2016 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package twenty48

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"fmt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 340
	ScreenHeight = 340
	boardSize    = 4
)

// Game represents a game state.
type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image
	activePlayer int
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{
		input: NewInput(),
		activePlayer: 0,
	}
	var err error
	g.board, err = NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.input.Update()
	dir, ok := g.input.Dir()
	fmt.Println(getBlock(g.input.mouseInitPosX), getBlock(g.input.mouseInitPosY), g.activePlayer, dir, ok, g.input.mouseState)
	if g.activePlayer == 1 {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.activePlayer = 0
			return addSpecificTile(g.board.tiles, getBlock(g.input.mouseInitPosX), getBlock(g.input.mouseInitPosY), g.board.size)
		}
		return nil
	}
	if _, ok := g.input.Dir(); ok {
		g.activePlayer = 1
	}
	if err := g.board.Update(g.input); err != nil {
		return err
	}
	return nil
}

func getBlock(pos int) int {
	return pos/84
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := g.board.Size()
		g.boardImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}
