package main

import (
	"bufio"
	"os"
	"testing"
	"io/ioutil"
)
func TestWhatElfHasTheMostCalories(t *testing.T) {
	tests := []struct {
		name string
		input string
		want int
	}{
		{
			name: "all elves have the same number of calories",
			input: "5\n\n5\n\n5\n",
			want: 0,
		},
		{
			name: "first elf has the most calories",
			input: "5\n\n5\n\n1\n",
			want: 0,
		},
		{
			name: "second elf has the most calories",
			input: "1\n\n5\n\n5\n",
			want: 1,
		},
		{
			name: "last elf has the most calories",
			input: "1\n\n1\n\n5\n",
			want: 2,
		},
		{
			name: "empty input file",
			input: "",
			want: 0,
		},
		{
			name: "input file with only one line",
			input: "5",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := writeToTempFile(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(f.Name())

			if got := WhatElfHasTheMostCalories(f.Name()); got != tt.want {
				t.Errorf("WhatElfHasTheMostCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func writeToTempFile(input string) (*os.File, error) {
	f, err := ioutil.TempFile("", "input")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(input)
	if err != nil {
		return nil, err
	}
	err = w.Flush()
	if err != nil {
		return nil, err
	}
	return f, nil
}
