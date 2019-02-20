package domain

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestBoardManagerAddTaskShouldAddAsTODO(t *testing.T) {
	boardManager := NewBoardManager()

	boardManager.AddTask("task 1", "TODO")

	assert.Equal(t, len(boardManager.Tasks),1)
	assert.Equal(t, len(boardManager.Tasks["TODO"]), 1)
	assert.Equal(t, boardManager.Tasks["TODO"][0],"task 1")
}
