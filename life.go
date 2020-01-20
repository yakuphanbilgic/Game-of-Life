package main

import(
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math/rand"
	"log"
)

const scale = 10
const screenW = 10
const screenH = 15

var cellColor color.RGBA = color.RGBA{255,204,128,255}
var backgroundColor color.RGBA = color.RGBA{0,137,123,255}

var grid [screenW][screenH] uint8 = [screenW][screenH] uint8{}
var buffer [screenW][screenH] uint8 = [screenW][screenH] uint8{}

var counter int = 0

//game logic is here
func update() error {
	for x := 1; x < screenW - 1; x++ {
		for y := 1; y < screenH - 1; y++ {
			buffer[x][y] = 0

			aliveCount := grid[x][y]

			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					row := (x + i + screenW) % screenW
					col := (y + j + screenH) % screenH
					aliveCount = aliveCount + grid[row][col]
				}   
			}

			aliveCount = aliveCount - 2 * grid[x][y]

			// Rule 2:
			// Any dead cell with exactly three live neighbours becomes alive cell.
			if grid[x][y] == 0 && aliveCount == 3 {
				buffer[x][y] = 1
			} else if aliveCount > 3 || aliveCount < 2 {
				// Rule 1:
				// Any alive cell with less than 2 or more than 3 neighbours dies.
				buffer[x][y] = 0
			} else {
				// Rule 3:
				// Any alive cell with 2 or 3 neighbours lives through next generation.
				buffer[x][y]=grid[x][y]
			}
		}
	}

	temp := buffer
	buffer = grid
	grid = temp
	return nil
}

func render(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	for x:=0; x<screenW; x++ {
		for y:=0; y<screenH; y++ {
			if grid[x][y] > 0 {
				for x1:=0; x1<scale; x1++ {
					for y1:=0; y1<scale; y1++ {
						screen.Set((x*scale)+x1, (y*scale)+y1, cellColor)
					}
				}
			}
		}
	}
}

func frame(screen *ebiten.Image) error {
	counter++
	var err error = nil
	if counter == 60 {
		err = update()
		counter = 0
	}
	if !ebiten.IsDrawingSkipped(){
    	render(screen)
  	}
  	return err
}

func main() {
	//set inital cells
	for x := 1; x < screenW-1; x++ {
		for y := 1; y < screenH-1; y++ {
			if(rand.Float32() < 0.5){
				grid[x][y] = 1
			}
		}
	}

	if err := ebiten.Run(frame, screenW*scale, screenH*scale, 2, "Conway's Game of Life");

	err != nil{
    	log.Fatal(err)
	}
}




