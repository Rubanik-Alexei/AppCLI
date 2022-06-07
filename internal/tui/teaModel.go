package tui

import (
	"fmt"
	"strconv"

	"github.com/Rubanik-Alexei/AppCLI/cmd"
	"github.com/Rubanik-Alexei/AppCLI/internal/models"
	tea "github.com/charmbracelet/bubbletea"
)

// type Task struct {
// 	Id          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	CompletedAt int64  `json:"completedAt"`
// }

// type cmd interface {
// 	DoTask([]string)
// }

type model struct {
	tasks    []models.Task
	cursor   int
	selected map[int]int
	//cmd      cmd
}

var _ tea.Model = (*model)(nil)

func NewModel(tasks []models.Task) model {
	return model{
		tasks:    tasks,
		selected: make(map[int]int),
	}
}

func (m model) View() string {
	// The header
	s := ""
	// Iterate over our choices
	for i, task := range m.tasks {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %v, %s \n", cursor, checked, task.Id, task.Name)
	}
	if s != "" {
		s = "List of incompleted tasks:\n\n" + s
	}
	// The footer
	s += "\nPress q to quit.\n"
	s += "\nPress d to complete tasks.\n"

	// Send the UI for rendering
	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "d":
			args := []string{}
			for _, id := range m.selected {
				args = append(args, strconv.Itoa(id))
			}
			cmd.DoTask(args)
			//time.Sleep(3 * time.Second)
			m.UpdateTasks()
			return m, tea.Quit
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = m.tasks[m.cursor].Id
			}
		}
	}
	return m, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) UpdateTasks() {
	m.tasks = cmd.ListTasksTUI()
}

// func (m model) DoTask(args []string) {
// 	m.cmd.DoTask(args)
// }
