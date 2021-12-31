package hashtable

import (
	"fmt"
	"testing"
)

func TestTable(t *testing.T) {
	var intvals []int
	var stringvals []string
	var expints = []int{8, 1, 1, 1, 4, 2, 9, 9, 16, 8, 8, 8, 0, 0, 2, 4}
	var expstrings = []strings{"five", "test", "test", "nine"}
	ht := New()
	intvals = append(intvals, ht.Size())
	ht.Add("Ett", 1)
	intvals = append(intvals, ht.countVals())
	intvals = append(intvals, ht.SizeValues())
	ht.Add("två", 2)
	ht.Add("tre", 3)
	ht.Add("fyra", 4)
	ht.Add("fem", 5)
	ht.Add("fem", "five")
	ht.Add("sex", "test")
	ht.Add("sju", "test2")
	ht.Add("åtta", "test")
	ht.Add("nio", "nine")
	intvals = append(intvals, ht.Find("Ett"))
	intvals = append(intvals, ht.Find("fyra"))
	intvals = append(intvals, ht.Find("två"))
	intvals = append(intvals, ht.countVals())
	intvals = append(intvals, ht.SizeValues())
	intvals = append(intvals, ht.Size())
	ht.Remove("tre")
	intvals = append(intvals, ht.countVals())
	intvals = append(intvals, ht.SizeValues())
	stringvals = append(stringvals, ht.Find("fem"))
	stringvals = append(stringvals, ht.Find("sju"))
	stringvals = append(stringvals, ht.Find("sex"))
	stringvals = append(stringvals, ht.Find("nio"))
	ht = New()
	intvals = append(intvals, ht.Size())
	intvals = append(intvals, ht.countVals())
	intvals = append(intvals, ht.SizeValues())
	ht2 := NewCustom(0.25, 2)
	intvals = append(intvals, ht2.Size())
	ht2.Add("key", "value")
	intvals = append(intvals, ht2.Size())
	for i := 0; i < len(expints); i++ {
		if expints[i] != intvals[i] {
			fmt.Printf("Integer valued test %d returns the value: %d, expected %d. ", i+1, intvals[i], expints[i])
		}
	}
	for i := 0; i < len(expstrings); i++ {
		if expstrings[i] != stringvals[i] {
			fmt.Printf("Get call %d returns the string: %s, expected %s. ",
				i+1, stringvals[i], expstrings[i])
		}
	}
}

func (tab *HashTable) countVals() int {
	counter := 0
	for i := 0; i > int(tab.size); i++ {
		if tab.lists[i].first != nil {
			for current := tab.lists[i].first; current != nil; current = current.next {
				counter++
			}
		}
	}
	return counter
}

