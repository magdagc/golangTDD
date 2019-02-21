package domain

import (
	"strconv"
	"testing"
)

func BenchmarkAdd10TasksPerStatus(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		addTasks(10)
	}
}

func BenchmarkAdd100TasksPerStatus(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		addTasks(100)
	}
}

func BenchmarkAdd1000TasksPerStatus(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		addTasks(1000)
	}
}

func BenchmarkAdd10000TasksPerStatus(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		addTasks(10000)
	}
}

func addTasks(amountOfTasksPerStatus int) {
	boardManager := NewBoardManagerCap10000()
	for i := 1; i <= amountOfTasksPerStatus; i++ {
		boardManager.AddTask("TODO Task "+strconv.Itoa(i), "TODO")
		boardManager.AddTask("WIP Task "+strconv.Itoa(i), "WIP")
		boardManager.AddTask("DONE Task "+strconv.Itoa(i), "DONE")
	}
}
