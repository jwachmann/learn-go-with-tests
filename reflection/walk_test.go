package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"James"},
			[]string{"James"},
		},
		{
			"Struct with multiple fields",
			struct {
				Name    string
				Age     int
				Address string
			}{"name", 1, "address"},
			[]string{"name", "address"},
		},
		{
			"Struct with embedded structs",
			Person{
				"James",
				Profile{
					34,
					"Burlington",
				},
			},
			[]string{"James", "Burlington"},
		},
		{
			"Pointer to a struct",
			&Person{
				"James",
				Profile{
					34,
					"Burlington",
				},
			},
			[]string{"James", "Burlington"},
		},
		{
			"Slices",
			[]Profile{
				{34, "James"},
				{3, "Leah"},
			},
			[]string{"James", "Leah"},
		},
		{
			"Arrays",
			[2]Profile{
				{34, "James"},
				{3, "Leah"},
			},
			[]string{"James", "Leah"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(tt.ExpectedCalls, got) {
				t.Errorf("Want '%v' got '%v'", tt.ExpectedCalls, got)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		input := map[string]string{
			"Name":     "James",
			"Location": "Burlington",
		}

		got := make(map[string]struct{})

		walk(input, func(value string) {
			got[value] = struct{}{}
		})

		assertLength := func(lengthExpected int) {
			lengthGot := len(got)
			if lengthGot != lengthExpected {
				t.Errorf("Expected '%v' values but got '%v'", lengthExpected, lengthGot)
			}
		}
		assertContains := func(value string) {
			if _, ok := got[value]; !ok {
				t.Errorf("Want '%v' but didn't find it", value)
			}
		}

		assertLength(2)
		assertContains("James")
		assertContains("Burlington")
	})
}
