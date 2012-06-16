package ini_test

import (
	"ini"
	"testing"
)

var (
	dict ini.Dict
	err  error
)

func init() {
	dict, err = ini.Load("example.ini")
}

func TestLoad(t *testing.T) {
	if err != nil {
		t.Error("Example: load error:", err)
	}
}

func TestGetBool(t *testing.T) {
	b, found := dict.GetBool("pizza", "ham")
	if !found || !b {
		t.Error("Example: parse error for key ham of section pizza.")
	}
	b, found = dict.GetBool("pizza", "mushrooms")
	if !found || !b {
		t.Error("Example: parse error for key mushrooms of section pizza.")
	}
	b, found = dict.GetBool("pizza", "capres")
	if !found || b {
		t.Error("Example: parse error for key capres of section pizza.")
	}
	b, found = dict.GetBool("pizza", "cheese")
	if !found || b {
		t.Error("Example: parse error for key cheese of section pizza.")
	}
	// now check for case insensitvity
	b, found = dict.GetBool("Pizza", "Mushrooms")
	if !found || !b {
		t.Error("Example: parse error for key Mushrooms of section Pizza.")
	}
	b, found = dict.GetBool("Pizza", "capres")
	if !found || b {
		t.Error("Example: parse error for key capres of section Pizza.")
	}
	b, found = dict.GetBool("pizza", "Cheese")
	if !found || b {
		t.Error("Example: parse error for key Cheese of section pizza.")
	}
}

func TestGetStringIntAndDouble(t *testing.T) {
	str, found := dict.GetString("wine", "grape")
	if !found || str != "Cabernet Sauvignon" {
		t.Error("Example: parse error for key grape of section wine.")
	}
	i, found := dict.GetInt("wine", "year")
	if !found || i != 1989 {
		t.Error("Example: parse error for key year of section wine.")
	}
	str, found = dict.GetString("wine", "country")
	if !found || str != "Spain" {
		t.Error("Example: parse error for key country of section wine.")
	}
	d, found := dict.GetDouble("wine", "alcohol")
	if !found || d != 12.5 {
		t.Error("Example: parse error for key wine of section wine.")
	}
	// now check for case insensitvity
	str, found = dict.GetString("Wine", "Grape")
	if !found || str != "Cabernet Sauvignon" {
		t.Error("Example: parse error for key Grape of section Wine.")
	}
	i, found = dict.GetInt("Wine", "year")
	if !found || i != 1989 {
		t.Error("Example: parse error for key year of section Wine.")
	}
	str, found = dict.GetString("wine", "Country")
	if !found || str != "Spain" {
		t.Error("Example: parse error for key Country of section wine.")
	}
	d, found = dict.GetDouble("Wine", "Alcohol")
	if !found || d != 12.5 {
		t.Error("Example: parse error for key Alcohol of section Wine.")
	}
}

func TestGetNotExist(t *testing.T) {
	_, found := dict.GetString("not", "exist")
	if found {
		t.Error("There is no key exist of section not.")
	}
}

func TestGetSections(t *testing.T) {
	sections := dict.GetSections()
	if len(sections) != 4 {
		t.Error("The number of sections is wrong:", len(sections))
	}
	for _, section := range sections {
		if section != "" && section != "pizza" && section != "wine" &&
			section != "monkeys" {
			t.Errorf("Section '%s' should not exist.", section)
		}
	}
}

func TestGetStringArray(t *testing.T) {
	keys, found := dict.SequencedKeyValues("Monkeys", "Monkey")
	if !found {
		t.Error("Monkeys sections should be found")
	}
	if len(keys) != 3 {
		t.Error("The number of sections is wrong:", len(keys))
	}
	for _, key := range keys {
		if key != "" && key != "Chimpanzee" && key != "Gorilla" && key != "Orangutan" {
			t.Errorf("Monkey '%s' should not exist.", key)
		}
	}

}
