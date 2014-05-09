package game_of_life
import "testing"

func TestWorld(t *testing.T) {
  world := NewBoardGame()
  world.SetCellAlive(1,2)
  if(!world.GetCell(1,2)) {
    t.Error("this cell cannot be dead")
  }
}
func TestWorldStep(t *testing.T) {
  world := NewBoardGame()
  world.SetCellAlive(1,2)
  newWorld := world.Step()
  if(newWorld.GetCell(1,2)) {
    t.Error("this cell should be dead")
  }
}
