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

  world.SetCellAlive(2,3)
  world.SetCellAlive(2,4)

  world.SetCellAlive(3,2)
  world.SetCellAlive(3,5)

  world.SetCellAlive(4,3)
  world.SetCellAlive(4,4)

  newWorld := world.Step()

  if(newWorld.GetCell(1,2)) {
    t.Error("this cell should be dead")
  }
  if(!newWorld.GetCell(2,3)) {
    t.Error("this cell should be alive")
  }
}
