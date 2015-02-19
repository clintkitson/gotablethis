package gotablethis

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrintColumn(*testing.T) {
	testing := []string{"test1"}

	table := Table{
		Header:  []string{"testheader1"},
		RowData: reflect.ValueOf(&testing).Elem(),
	}

	fmt.Printf("\n<- gotablethis.PrintColumn ->\n\n")
	table.PrintColumn()
}

func TestPrintKeyValueTable(*testing.T) {
	testing := teststruct{
		Test1: "first",
		Test2: 1,
		Test3: true,
	}

	table := Table{
		RowData: reflect.ValueOf(&testing).Elem(),
	}

	fmt.Printf("\n<- gotablethis.PrintKeyValueTable ->\n\n")
	table.PrintKeyValueTable()
}

type teststruct struct {
	Test1 string
	Test2 int
	Test3 bool
}

type teststructs []teststruct

func TestPrintTable(*testing.T) {
	testing := teststructs{
		teststruct{
			Test1: "first",
			Test2: 1,
			Test3: true,
		},
		teststruct{
			Test1: "second",
			Test2: 2,
			Test3: false,
		},
	}

	table := Table{
		Header:  []string{"test1", "test2", "test3"},
		Columns: []string{"Test1", "Test2", "Test3"},
		RowData: reflect.ValueOf(&testing).Elem(),
	}

	fmt.Printf("\n<- gotablethis.PrintTable ->\n\n")
	table.PrintTable()
}
