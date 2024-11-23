package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type TaskStatus string

const (
	TODO       TaskStatus = "todo"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
)

type Task struct {
	Id          uint64     `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Tracker struct {
	filename string

	tasks []*Task
}

func NewTaskTracker(filename string) *Tracker {
	var (
		err     error
		tracker = &Tracker{
			tasks:    make([]*Task, 0),
			filename: filename,
		}
		file *os.File
	)

	if file, err = os.OpenFile(filename, os.O_CREATE, os.ModePerm); err != nil {
		panic(err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&tracker.tasks); err != nil {
		if !errors.Is(err, io.EOF) {
			fmt.Println(err)
			panic(err)
		}
	}

	return tracker
}

func (t *Tracker) Save() error {
	var (
		err  error
		file *os.File
	)

	if file, err = os.OpenFile(t.filename, os.O_WRONLY|os.O_TRUNC, os.ModePerm); err != nil {
		panic(err)
	}
	defer file.Close()

	if len(t.tasks) > 0 {
		if err = json.NewEncoder(file).Encode(t.tasks); err != nil {
			return err
		}
	}

	return nil
}

func (t *Tracker) Add(desc string) {
	var maxId uint64
	for _, task := range t.tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	nextId := maxId + 1
	t.tasks = append(t.tasks, &Task{
		Id:          nextId,
		Description: desc,
		Status:      TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	fmt.Printf("Task added successfully (ID: %d)", nextId)

}

func (t *Tracker) Update(id uint64, desc string) {
	for _, task := range t.tasks {
		if task.Id == id {
			task.Description = desc
			task.UpdatedAt = time.Now()
		}
	}
}

func (t *Tracker) Delete(id uint64) {
	var filtered []*Task
	for _, task := range t.tasks {
		if task.Id == id {
			continue
		}
		filtered = append(filtered, task)
	}
	t.tasks = filtered
}

func (t *Tracker) MarkInProgress(id uint64) {
	for _, task := range t.tasks {
		if task.Id == id {
			task.Status = InProgress
			task.UpdatedAt = time.Now()
		}
	}
}

func (t *Tracker) MarkDone(id uint64) {
	for _, task := range t.tasks {
		if task.Id == id {
			task.Status = Done
			task.UpdatedAt = time.Now()
		}
	}
}

func (t *Tracker) List(status TaskStatus) {
	for _, task := range t.tasks {
		if status == "" || task.Status == status {
			if b, err := json.Marshal(task); err == nil {
				fmt.Println(string(b))
			}

		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Hello!")
		return
	}

	tracker := NewTaskTracker("./store.json")
	needWrite := true

	command := strings.TrimSpace(args[0])
	var params []string
	if len(args) > 1 {
		params = args[1:]
	}

	switch command {
	case "add":
		var desc string
		if len(params) > 0 {
			desc = params[0]
		}
		tracker.Add(desc)
	case "update":
		if len(params) != 2 {
			panic(errors.New("update need 2 params"))
		}
		taskId, err := strconv.ParseUint(params[0], 10, 64)
		if err != nil {
			fmt.Print(err.Error())
		}

		tracker.Update(taskId, params[1])
	case "delete":
		if len(params) == 0 {
			panic(errors.New("need task id"))
		}
		taskId, err := strconv.ParseUint(params[0], 10, 64)
		if err != nil {
			fmt.Print(err.Error())
		}

		tracker.Delete(taskId)
	case "mark-in-progress":
		if len(params) == 0 {
			panic(errors.New("need task id"))
		}
		taskId, err := strconv.ParseUint(params[0], 10, 64)
		if err != nil {
			fmt.Print(err.Error())
		}
		tracker.MarkInProgress(taskId)
	case "mark-done":
		if len(params) == 0 {
			panic(errors.New("need task id"))
		}
		taskId, err := strconv.ParseUint(params[0], 10, 64)
		if err != nil {
			fmt.Print(err.Error())
		}
		tracker.MarkDone(taskId)
	case "list":
		var status string
		if len(params) > 0 {
			status = params[0]
		}
		tracker.List(TaskStatus(status))
		needWrite = false
	default:
		fmt.Println("Unknown command")
		needWrite = false
	}

	if needWrite {
		if err := tracker.Save(); err != nil {
			fmt.Println(err)
		}
	}
}
