package domain

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBoardManagerAddTaskShouldAddAsTODO(t *testing.T) {
	boardManager := NewBoardManager()

	boardManager.AddTask("task 1", "TODO")

	assert.Equal(t, len(boardManager.Tasks), 1)
	assert.Equal(t, len(boardManager.Tasks["TODO"]), 1)
	assert.Equal(t, boardManager.Tasks["TODO"][0], "task 1")
}

func TestBoardManagerAddTasksShouldAddAsTODOAndWIP(t *testing.T) {
	boardManager := NewBoardManager()

	boardManager.AddTask("task 1", "TODO")
	boardManager.AddTask("task 2", "TODO")
	boardManager.AddTask("task 3", "WIP")

	assert.Equal(t, len(boardManager.Tasks), 2)
	assert.Equal(t, len(boardManager.Tasks["TODO"]), 2)
	assert.Equal(t, len(boardManager.Tasks["WIP"]), 1)
	assert.Equal(t, boardManager.Tasks["TODO"][0], "task 1")
	assert.Equal(t, boardManager.Tasks["TODO"][1], "task 2")
	assert.Equal(t, boardManager.Tasks["WIP"][0], "task 3")
}
