package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Record struct {
	ID          uint64    `json:"id"`
	Description string    `json:"description"`
	Amount      uint64    `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}

type Tracker struct {
	filename string

	records []*Record
}

func NewTracker(filename string) *Tracker {
	var (
		err     error
		tracker = &Tracker{
			records:  make([]*Record, 0),
			filename: filename,
		}
		file *os.File
	)

	if file, err = os.OpenFile(filename, os.O_CREATE, os.ModePerm); err != nil {
		panic(err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&tracker.records); err != nil {
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

	if len(t.records) > 0 {
		if err = json.NewEncoder(file).Encode(t.records); err != nil {
			return err
		}
	}

	return nil
}

func (t *Tracker) Add(desc string, amount uint64) {
	var maxId uint64
	for _, record := range t.records {
		if record.ID > maxId {
			maxId = record.ID
		}
	}

	nextID := maxId + 1

	t.records = append(t.records, &Record{
		ID:          nextID,
		Description: desc,
		Amount:      amount,
		CreatedAt:   time.Now(),
	})

	fmt.Printf("Expense added successfully (ID: %d)\n", nextID)
}

func (t *Tracker) List() {
	fmt.Println("ID\tDate\tDescription\tAmount")
	for _, record := range t.records {
		fmt.Printf("%d\t%s\t%s\t$%d\n", record.ID, record.CreatedAt.Format("2006-01-02"), record.Description, record.Amount)
	}
}

func (t *Tracker) Summary(month uint32) {
	if month > 12 {
		fmt.Printf("Invalid month: %d\n", month)
		return
	}

	var total uint64
	for _, record := range t.records {
		if month == 0 || record.CreatedAt.Month() == time.Month(month) {
			total += record.Amount
		}
	}
	if month == 0 {
		fmt.Printf("Total expenses: $%d", total)
	} else {
		fmt.Printf("Total expenses for %s: $%d", time.Month(month), total)
	}

}

func (t *Tracker) Delete(id uint64) {
	filtered := make([]*Record, 0)
	for _, record := range t.records {
		if record.ID == id {
			fmt.Println("Expense deleted successfully")
			continue
		}
		filtered = append(filtered, record)
	}
	t.records = filtered
}

func main() {
	if len(os.Args) == 0 {
		fmt.Println("No command provided")
		return
	}
	command := strings.TrimSpace(os.Args[1])

	var (
		description string
		amount      uint64
		id          uint64
		month       uint64
	)

	if len(os.Args) > 2 {
		fs := flag.NewFlagSet("expense-tracker", flag.ExitOnError)
		fs.StringVar(&description, "description", "", "Description of the expense")
		fs.Uint64Var(&amount, "amount", 0, "Amount of the expense")
		fs.Uint64Var(&id, "id", 0, "ID of the expense")
		fs.Uint64Var(&month, "month", 0, "Month to filter expenses")
		fs.Parse(os.Args[2:])
	}

	tracker := NewTracker("./store.json")
	needWrite := true

	switch command {
	case "add":
		tracker.Add(description, amount)
	case "list":
		tracker.List()
		needWrite = false
	case "summary":
		tracker.Summary(uint32(month))
		needWrite = false
	case "delete":
		tracker.Delete(id)
	default:
		needWrite = false
	}

	if needWrite {
		if err := tracker.Save(); err != nil {
			fmt.Println(err)
		}
	}
}
