// Copyright 2012 Paul Bellamy. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pipe

import (
	"testing"
)

func TestZipPipe(t *testing.T) {
	other := make(chan interface{}, 5)
	in := make(chan interface{}, 5)
	out := Zip(in, other)

	in <- 5
	in <- 10
	in <- 20
	other <- 6
	other <- 11

	for i := 1; i <= 2; i++ {
		result := <-out
		expected := []int{i * 5, (i * 5) + 1}
		if len(result.([]interface{})) != len(expected) {
			t.Fatal("expected channel output to match", expected, "but got", result.([]int))
		}

		for j := 0; j < len(result.([]interface{})); j++ {
			if result.([]interface{})[j].(int) != expected[j] {
				t.Fatal("expected channel output to match", expected, "but got", result.([]interface{}))
			}
		}
	}
}

func TestMultiZipPipe(t *testing.T) {
	other1 := make(chan interface{}, 5)
	other2 := make(chan interface{}, 5)
	in := make(chan interface{}, 5)
	out := Zip(in, other1, other2)

	in <- 5
	in <- 10
	in <- 20
	other1 <- 6
	other1 <- 11
	other2 <- 7
	other2 <- 12

	for i := 1; i <= 2; i++ {
		result := <-out
		expected := []int{i * 5, (i * 5) + 1, (i * 5) + 2}
		if len(result.([]interface{})) != len(expected) {
			t.Fatal("expected channel output to match", expected, "but got", result.([]int))
		}

		for j := 0; j < len(result.([]interface{})); j++ {
			if result.([]interface{})[j].(int) != expected[j] {
				t.Fatal("expected channel output to match", expected, "but got", result.([]interface{}))
			}
		}
	}

	close(other1)
	if _, ok := <-out; ok {
		t.Fatal("expected channel to be closed")
	}
}

func TestZipChainedConstructor(t *testing.T) {
	other := make(chan interface{}, 5)
	in := make(chan interface{}, 5)
	out := NewPipe(in).Zip(other).Output

	in <- 5
	in <- 10
	in <- 20
	other <- 6
	other <- 11

	for i := 1; i <= 2; i++ {
		result := <-out
		expected := []int{i * 5, (i * 5) + 1}
		if len(result.([]interface{})) != len(expected) {
			t.Fatal("expected channel output to match", expected, "but got", result.([]int))
		}

		for j := 0; j < len(result.([]interface{})); j++ {
			if result.([]interface{})[j].(int) != expected[j] {
				t.Fatal("expected channel output to match", expected, "but got", result.([]interface{}))
			}
		}
	}
}

func TestMultiZipChainedConstructor(t *testing.T) {
	other1 := make(chan interface{}, 5)
	other2 := make(chan interface{}, 5)
	in := make(chan interface{}, 5)
	out := NewPipe(in).Zip(other1, other2).Output

	in <- 5
	in <- 10
	in <- 20
	other1 <- 6
	other1 <- 11
	other2 <- 7
	other2 <- 12

	for i := 1; i <= 2; i++ {
		result := <-out
		expected := []int{i * 5, (i * 5) + 1, (i * 5) + 2}
		if len(result.([]interface{})) != len(expected) {
			t.Fatal("expected channel output to match", expected, "but got", result.([]int))
		}

		for j := 0; j < len(result.([]interface{})); j++ {
			if result.([]interface{})[j].(int) != expected[j] {
				t.Fatal("expected channel output to match", expected, "but got", result.([]interface{}))
			}
		}
	}

	close(other1)
	if _, ok := <-out; ok {
		t.Fatal("expected channel to be closed")
	}
}
