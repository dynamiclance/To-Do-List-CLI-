// CLI Based To Do List

// 1. Add Task
// 2. List Task
// 3. Complete Task as Complete
// 4. Delete Task
// 5. Exit Task

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func displayMessage() {
	fmt.Print(`1. Add Task
2. List Task
3. Update Task as Complete
4. Delete Task
5. Exit Task
`)
}

func getUserInput(reader *bufio.Reader) (int, error) {

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return strconv.Atoi(input)
}

func getTaskByUser(reader *bufio.Reader) string {
	task, _ := reader.ReadString('\n')
	return strings.TrimSpace(task)
}

func taskStatusUpdate(tasks []string, taskId int) []string {
	if taskId > 0 && taskId <= len(tasks) {
		tasks[taskId-1] = "[+] " + tasks[taskId-1]
		fmt.Println("Task Status has been Updated")

		return tasks
	}

	fmt.Println("Invalid Task Id")
	return tasks

}

func deletedTask(tasks []string, taskId int) []string {
	if taskId > 0 && taskId <= len(tasks) {

		tasks = append(tasks[:taskId-1], tasks[taskId:]...)
		fmt.Println("TaskId has been deleted successfully")

		return tasks
	}

	fmt.Println("Invalid Task Id")
	return tasks

}

func main() {

	var tasks []string
	reader := bufio.NewReader(os.Stdin)

	displayMessage()

	for {

		fmt.Print("Enter Your Choice: ")

		choice, error := getUserInput(reader)

		if choice > 5 || error != nil {
			fmt.Println("Invalid Choice. Please Try Again")
		}
		switch choice {
		case 1:
			// add task
			fmt.Printf("Enter Your Task: ")

			// add this task to task list
			tasks = append(tasks, getTaskByUser(reader))
			fmt.Println("Task Added")
		case 2:

			// List of Task
			if len(tasks) == 0 {
				fmt.Println("No Task is Available Now")
			} else {
				fmt.Println("Your Task List: ")
				for index, task := range tasks {
					fmt.Printf("%d. %s\n", index+1, task)
				}
			}
		case 3:
			// Task Status Update

			if len(tasks) == 0 {
				fmt.Println("No Task is Available Now")
			} else {
				fmt.Printf("Enter Your Task ID to Update Task Status: ")

				taskId, error := getUserInput(reader)

				if error == nil {

					taskStatusUpdate(tasks, taskId)

				} else {
					fmt.Println("You Put Wrong Task Id, Please try agin")
				}

			}

		case 4:
			// Delete Task

			if len(tasks) == 0 {
				fmt.Println("No Task is Available Now")
			} else {
				fmt.Printf("Enter Your Task ID to Delete Task: ")

				taskId, error := getUserInput(reader)

				if error == nil {

					tasks = deletedTask(tasks, taskId)

				} else {
					fmt.Println("You Put Wrong Task Id, Please try agin")
				}

			}

		case 5:
			// Exit Task
			fmt.Println("Goodbye")
			os.Exit(1)

		default:
			fmt.Println("Inavalid Input Choice")

		}

	}

}
