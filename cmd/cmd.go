package cmd

type Command struct {
}

func (c *Command) Execute(command string) {
	switch command {
	// category methods
	case "create-category":
	case "edit-category":
	case "list-category":

	// category methods
	case "create-task":
	case "list-task":
	case "list-task-today":
	case "list-task-bydate":
	case "edit-task":
	case "toggle-task":

	// user methods
	case "register":
	case "login":

	}
}
