package reflectutils

import (
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
)

type MyData struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	HasPet bool   `csv:"has_pet"`
}

func TestCsv(t *testing.T) {
	data := `name,age,has_pet
Jon,"100",true
"Fred"" The Hammer ""Smith",42,false
Martha,37,"true"
`

	r := csv.NewReader(strings.NewReader(data))
	allData, err := r.ReadAll()
	if err != nil {
		t.Error(err)
	}

	entries := make([]MyData, 0)
	Unmarshal(allData, &entries)
	fmt.Println(entries)

	out, err := Marshal(entries)
	if err != nil {
		t.Error(err)
	}
	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	w.WriteAll(out)
	fmt.Println(sb)
}
