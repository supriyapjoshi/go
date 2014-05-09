package game_of_life
type Coord struct {
  x,y int
}
type BoardGame struct {
  game map[Coord]bool
}
func NewBoardGame() *BoardGame {
  board := new(BoardGame)
  board.game = make(map[Coord]bool)
  return board
}
func (board *BoardGame) SetCellAlive(x,y int) {
  coord := Coord{x,y}
  board.game[coord] = true

}
func (board *BoardGame) GetCell(x, y int) bool{
  cell, ok := board.game[Coord{x,y}]
  if(ok) {
    return cell
  } else {
    return false
  }
}
func (board *BoardGame) Step() *BoardGame{
  next_gen := NewBoardGame()

  for key, value := range board.game {
    coord := key
    neighbours := GetLiveNeighbours(coord, board)
    cell := NewCell(value)
    cell.Step(neighbours)
    if(cell.alive) {
      next_gen.SetCellAlive(coord.x, coord.y)
    }
  }
  return next_gen
}

func GetLiveNeighbours(coord Coord, board *BoardGame) int {
  neighbours := 0
  if(board.GetCell(coord.x-1, coord.y)) {neighbours++}
  if(board.GetCell(coord.x,coord.y-1)) {neighbours++;}
  if(board.GetCell(coord.x+1, coord.y)) {neighbours++;}
  if(board.GetCell(coord.x, coord.y+1)) {neighbours++;}
  if(board.GetCell(coord.x-1, coord.y-1)) {neighbours++;}
  if(board.GetCell(coord.x+1, coord.y+1)) {neighbours++;}
  if(board.GetCell(coord.x-1, coord.y+1)) {neighbours++;}
  if(board.GetCell(coord.x+1, coord.y-1)) {neighbours++;}
  return neighbours
}
