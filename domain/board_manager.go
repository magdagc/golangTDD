package domain

// BoardManager manages a map of task slices using the status as the key
type BoardManager struct {
	Tasks map[string][]string
}

// NewBoardManager returns a pointer to an empty declared BoardManager
func NewBoardManager() *BoardManager {
	tasks := map[string][]string{
		"TODO": make([]string, 0),
		"WIP":  make([]string, 0),
		"DONE": make([]string, 0),
	}
	return &BoardManager{Tasks: tasks}
}

// NewBoardManagerCap10 returns a pointer to an empty declared BoardManager with capacity for 10 tasks in each status
func NewBoardManagerCap10() *BoardManager {
	tasks := map[string][]string{
		"TODO": make([]string, 0, 10),
		"WIP":  make([]string, 0, 10),
		"DONE": make([]string, 0, 10),
	}
	return &BoardManager{Tasks: tasks}
}

// NewBoardManagerCap100 returns a pointer to an empty declared BoardManager with capacity for 100 tasks in each status
func NewBoardManagerCap100() *BoardManager {
	tasks := map[string][]string{
		"TODO": make([]string, 0, 100),
		"WIP":  make([]string, 0, 100),
		"DONE": make([]string, 0, 100),
	}
	return &BoardManager{Tasks: tasks}
}

// NewBoardManagerCap1000 returns a pointer to an empty declared BoardManager with capacity for 1000 tasks in each status
func NewBoardManagerCap1000() *BoardManager {
	tasks := map[string][]string{
		"TODO": make([]string, 0, 1000),
		"WIP":  make([]string, 0, 1000),
		"DONE": make([]string, 0, 1000),
	}
	return &BoardManager{Tasks: tasks}
}

// NewBoardManagerCap10000 returns a pointer to an empty declared BoardManager with capacity for 10 tasks in each status
func NewBoardManagerCap10000() *BoardManager {
	tasks := map[string][]string{
		"TODO": make([]string, 0, 1000),
		"WIP":  make([]string, 0, 1000),
		"DONE": make([]string, 0, 1000),
	}
	return &BoardManager{Tasks: tasks}
}

// AddTask adds a new task to the board
func (b *BoardManager) AddTask(task, status string) {
	b.Tasks[status] = append(b.Tasks[status], task)
}

// GetTasks returns the slice with the tasks defined for a specific status
func (b *BoardManager) GetTasks(status string) []string {
	return b.Tasks[status]
}

// ChangeStatus modifies the status of a task inside the board
func (b *BoardManager) ChangeStatus(task, originalStatus, newStatus string) {

	tasksOriginalStatus := b.Tasks[originalStatus]

	for i, taskInBoard := range tasksOriginalStatus {
		if taskInBoard == task {
			tasksOriginalStatus = append(tasksOriginalStatus[:i], tasksOriginalStatus[i+1:]...)
		}
	}

	b.Tasks[originalStatus] = tasksOriginalStatus

	b.Tasks[newStatus] = append(b.Tasks[newStatus], task)
}

// RemoveTask removes a task from the board
func (b *BoardManager) RemoveTask(task, status string) {
	tasks := b.Tasks[status]

	for i, taskInBoard := range tasks {
		if taskInBoard == task {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	b.Tasks[status] = tasks
}
