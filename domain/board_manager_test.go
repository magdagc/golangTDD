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

func TestBoardManagerGetTasksShouldGetTODOTasks(t *testing.T) {
	boardManager := NewBoardManager()

	boardManager.AddTask("task 1", "TODO")
	boardManager.AddTask("task 2", "TODO")
	boardManager.AddTask("task 3", "WIP")

	tasks := boardManager.GetTasks("TODO")

	assert.Equal(t, tasks[0], "task 1")
	assert.Equal(t, tasks[1], "task 2")
}

func TestBoardManagerGetTasksShouldGetTODOAndWIPTasks(t *testing.T) {
	boardManager := NewBoardManager()

	boardManager.AddTask("task 1", "TODO")
	boardManager.AddTask("task 2", "TODO")
	boardManager.AddTask("task 3", "WIP")

	todoTasks := boardManager.GetTasks("TODO")
	wipTasks := boardManager.GetTasks("WIP")

	assert.Equal(t, todoTasks[0], "task 1")
	assert.Equal(t, todoTasks[1], "task 2")
	assert.Equal(t, wipTasks[0], "task 3")
}

func TestBoardManagerGetTasksShouldGetEmptyStringSlice(t *testing.T) {
	boardManager := NewBoardManager()

	boardManager.AddTask("task 1", "TODO")
	boardManager.AddTask("task 2", "TODO")
	boardManager.AddTask("task 3", "WIP")

	doneTasks := boardManager.GetTasks("DONE")

	assert.Equal(t, len(doneTasks), 0)
}

func TestBoardManagerChangeStatusFromTODOToWIP(t *testing.T) {
	boardManager := NewBoardManager()

	boardManager.AddTask("task 1", "TODO")
	boardManager.AddTask("task 2", "TODO")
	boardManager.AddTask("task 3", "WIP")

	boardManager.ChangeStatus("task 1", "TODO", "WIP")

	assert.Equal(t, len(boardManager.Tasks), 2)
	assert.Equal(t, len(boardManager.Tasks["TODO"]), 1)
	assert.Equal(t, len(boardManager.Tasks["WIP"]), 2)
	assert.Equal(t, boardManager.Tasks["TODO"][0], "task 2")
	assert.Equal(t, boardManager.Tasks["WIP"][0], "task 3")
	assert.Equal(t, boardManager.Tasks["WIP"][1], "task 1")
}

func TestBoardManagerChangeStatusFromWIPToDone(t *testing.T) {
	boardManager := NewBoardManager()

	boardManager.AddTask("task 1", "TODO")
	boardManager.AddTask("task 2", "TODO")
	boardManager.AddTask("task 3", "WIP")

	boardManager.ChangeStatus("task 3", "WIP", "DONE")

	assert.Equal(t, len(boardManager.Tasks), 3)
	assert.Equal(t, len(boardManager.Tasks["TODO"]), 2)
	assert.Equal(t, len(boardManager.Tasks["WIP"]), 0)
	assert.Equal(t, len(boardManager.Tasks["DONE"]), 1)
	assert.Equal(t, boardManager.Tasks["TODO"][0], "task 1")
	assert.Equal(t, boardManager.Tasks["TODO"][1], "task 2")
	assert.Equal(t, boardManager.Tasks["DONE"][0], "task 3")
}
