package game_of_life

type Cell struct {
  alive bool
}
func NewCell(alive bool) *Cell {
  cell := new(Cell)
  cell.alive = alive
  return cell
}
func (cell *Cell) Step(neighbours int) *Cell {

    if(neighbours <2 || neighbours > 3) {
      cell.alive = false
    } else if (!cell.alive && neighbours == 3) {
      cell.alive = true
    }
    return cell
}
