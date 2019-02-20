package domain

// BoardManager manages a map of task slices using the status as the key
type BoardManager struct {
	Tasks map[string][]string
}

// NewBoardManager returns a pointer to an empty declared BoardManager
func NewBoardManager() *BoardManager {
	tasks := make(map[string][]string, 0)
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
}
