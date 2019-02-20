package main

import (
	"github.com/magdagc/golangTDD/domain"
	"gopkg.in/abiosoft/ishell.v2"
)

// PossibleStatuses are the three statuses defined in the board
var PossibleStatuses = []string{"TODO", "WIP", "DONE"}

func main() {

	boardManager := domain.NewBoardManager()

	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("       _                        \r\n       \\`*-.                    \r\n        )  _`-.                 \r\n       .  : `. .                \r\n       : _   '  \\               \r\n       ; *` _.   `*-._          \r\n       `-.-'          `-.       \r\n         ;       `       `.     \r\n         :.       .        \\    \r\n         . \\  .   :   .-'   .   \r\n         '  `+.;  ;  '      :   \r\n         :  '  |    ;       ;-. \r\n         ; '   : :`-:     _.`* ;\r\n      .*' /  .*' ; .*`- +'  `*' \r\n      `*-*   `*-*  `*-*'")
	shell.Println("\n ____  ____  _  _   __   __ _   __                \r\n/ ___)(  __)( \\/ ) / _\\ (  ( \\ / _\\               \r\n\\___ \\ ) _) / \\/ \\/    \\/    //    \\              \r\n(____/(____)\\_)(_/\\_/\\_/\\_)__)\\_/\\_/              \r\n ____  ____    __     __                          \r\n(    \\(  __)  (  )   / _\\                         \r\n ) D ( ) _)   / (_/\\/    \\                        \r\n(____/(____)  \\____/\\_/\\_/                        \r\n  __    ___  __  __    __  ____   __   ____       \r\n / _\\  / __)(  )(  )  (  )(    \\ / _\\ (    \\      \r\n/    \\( (_ \\ )( / (_/\\ )(  ) D (/    \\ ) D (      \r\n\\_/\\_/ \\___/(__)\\____/(__)(____/\\_/\\_/(____/ \n")

	shell.AddCmd(&ishell.Cmd{
		Name: "new-todo-task",
		Help: "Agrega una tarea en la columna TODO",
		Func: func(c *ishell.Context) {
			c.Println("Ingresar nombre de la tarea a agregar")
			taskName := c.ReadLine()
			boardManager.AddTask(taskName, "TODO")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "change-task-status",
		Help: "Cambia el estado de una tarea",
		Func: func(c *ishell.Context) {
			statuses := make([]string, 0)
			statuses = append(statuses, PossibleStatuses...)
			statusIndex := c.MultiChoice(statuses, "Status original de la tarea")
			originalStatus := statuses[statusIndex]
			tasks := boardManager.GetTasks(originalStatus)
			statuses = append(statuses[:statusIndex], statuses[statusIndex+1:]...)
			taskIndex := c.MultiChoice(tasks, "Tarea a modificar")
			statusIndex = c.MultiChoice(statuses, "Nuevo status")
			newStatus := statuses[statusIndex]
			boardManager.ChangeStatus(tasks[taskIndex], originalStatus, newStatus)
			c.Println("¡Tarea modificada con éxito!")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "setup-sample-tasks",
		Help: "Configura tareas de ejemplo",
		Func: func(c *ishell.Context) {
			setupSampleTasks(boardManager)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "show-board",
		Help: "Muestra el tablero con las tareas por columna",
		Func: func(c *ishell.Context) {
			// 30 caracteres por columna
			c.Println("TODO                          WIP                           DONE")
			c.Println("Task 1                        Task 3")
			c.Println("Task 2")
		},
	})

	// run shell
	shell.Run()
}

func setupSampleTasks(boardManager *domain.BoardManager) {
	boardManager = domain.NewBoardManager()
	/*	boardManager.AddTask("Task 1", "TODO")
		boardManager.AddTask("Task 2", "TODO")
		boardManager.AddTask("Task 3", "WIP")*/
}
