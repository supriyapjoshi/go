package game_of_life

import "testing"


func TestCellAlive(t *testing.T) {
  cell_step := func (neighbour int) bool {
      cell := NewCell(true)
      cell.Step(neighbour)
      return cell.alive
  }
  if(cell_step(1)) {
      t.Error("only 1 neighbour alive, cell must die")
  }
  if(!cell_step(2)) {
      t.Error("only 2 neighbours alive, cell must live on")
  }
  if(!cell_step(3)) {
      t.Error("3 live neighbours, cell must live on")
  }
  if(cell_step(4)) {
      t.Error("4 live neighbours, cell must die")
  }

}
func TestDeadCellBehaviour(t *testing.T) {
  cell_step := func(neighbour int) bool {
      cell := NewCell(false)
      cell.Step(neighbour)
      return cell.alive
  }
  if(cell_step(1)) {
    t.Error("only 1 neighbour alive, cell must be dead")
  }
  if(cell_step(2)) {
    t.Error("only 2 neighbours alive, cell must be dead")
  }
  if(!cell_step(3)) {
    t.Error("3 live neighbours, cell should become alive")
  }
  if(cell_step(4)) {
    t.Error("cell should continue to die with 4 live neighbours")
  }
 }
