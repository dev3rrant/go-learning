package main

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
	t.Run("table tests", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"Struct with one string field",
				struct {
					Name string
				}{"Paul"},
				[]string{"Paul"},
			},
			{
				"Struct with two string field",
				struct {
					Name string
					City string
				}{"Paul", "Arrakis"},
				[]string{"Paul", "Arrakis"},
			},
			{
				"Struct with non string field",
				struct {
					Name string
					Age  int
				}{"Paul", 18},
				[]string{"Paul"},
			},
			{
				"struct with nested fields",
				Person{
					"Paul",
					Profile{18, "Arrakis"},
				},
				[]string{"Paul", "Arrakis"},
			},
			{
				"struct with pointers",
				&Person{
					"Paul",
					Profile{18, "Arrakis"},
				},
				[]string{"Paul", "Arrakis"},
			},
			{
				"slice",
				[]Profile{
					{18, "Arrakis"},
					{17, "Andor"},
				},
				[]string{"Arrakis", "Andor"},
			},
			{
				"array",
				[2]Profile{
					{18, "Arrakis"},
					{17, "Andor"},
				},
				[]string{"Arrakis", "Andor"},
			},
			{
				"maps",
				map[string]string{
					"Dune":          "Herbert",
					"Wheel of Time": "Jordan",
				},
				[]string{"Herbert", "Jordan"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("Got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}

	})

	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Arrakis"}
			aChannel <- Profile{33, "Amber"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Arrakis", "Amber"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("func", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{18, "Arrakis"}, Profile{18, "Andor"}
		}

		var got []string
		want := []string{"Arrakis", "Andor"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
