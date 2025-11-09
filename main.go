package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

func main() {
	flag.Parse()
	tasks, err := LoadTasksFromFile("storage.json")
	checkError(err)
	switch flag.Arg(0) {
	case "add":
		desc := flag.Arg(1)
		if desc == "" {
			err = fmt.Errorf("usage: %s add \"description of the new task.\"", os.Args[0])
			checkError(err)
			break
		}
		newId := tasks.AddFromDescription(desc)
		fmt.Printf("Added task ID: %d\n", newId)
	case "update":
		id, err := strconv.Atoi(flag.Arg(1))
		desc := flag.Arg(2)
		if err != nil || desc == "" {
			err = fmt.Errorf("usage: %s update <id> \"description of the new task.\"", os.Args[0])
			checkError(err)
			break
		}

		err = tasks.UpdateFromDescription(id, desc)
		checkError(err)
		fmt.Println("Updated task.")
	case "delete":
		id, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			err = fmt.Errorf("usage: %s delete <id>", os.Args[0])
			checkError(err)
			break
		}

		err = tasks.Delete(id)
		checkError(err)
		fmt.Println("Deleted task.")
	case "mark":
		id, err := strconv.Atoi(flag.Arg(1))
		s := flag.Arg(2)
		isValid := CheckValidStatus(s)
		if err != nil || !isValid {
			err = fmt.Errorf("usage: %s mark <id> [todo,in-progress,done]", os.Args[0])
			checkError(err)
			break
		}

		status, err := ParseFromString(s)
		checkError(err)
		tasks.MarkAs(id, status)

		fmt.Println("Marked task.")
	case "list":
		s := flag.Arg(1)
		var filteredTasks TaskRepository
		if s == "" {
			filteredTasks = tasks.ListAll()
		} else {
			isValid := CheckValidStatus(s)
			if !isValid {
				err = fmt.Errorf("usage: %s list <none>|[todo,in-progress,done]", os.Args[0])
				checkError(err)
				break
			}
			status, err := ParseFromString(s)
			checkError(err)
			filteredTasks = tasks.listBy(status)
		}

		if len(filteredTasks) == 0 {
			fmt.Println("No tasks")
			break
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "ID\tStatus\tDescription\tCreated At\tUpdated At")
		// When displaying Unicode characters in the terminal,
		// the table cells don’t align properly with the headers.
		// I hope someone can help me figure out this issue.
		for _, task := range filteredTasks {
			var status_c string
			switch task.Status {
			case TODO:
				status_c = "📝"
			case IN_PROGRESS:
				status_c = "⚙️"
			case DONE:
				status_c = "✅"
			}
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", task.Id, status_c, task.Description, task.CreatedAt.ToString(), task.UpdatedAt.ToString())
		}
		w.Flush()
	default:
		err = fmt.Errorf("usage: %s [add,update,delete,mark,list]", os.Args[0])
		checkError(err)
	}
	err = SaveTasksToFile(tasks, "storage.json")
	checkError(err)
}
