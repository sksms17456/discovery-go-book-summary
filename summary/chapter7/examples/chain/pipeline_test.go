package concurrency

import (
	"reflect"
	"testing"
)

func TestChain(t *testing.T) {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()

	var got[]int
	want := []int{36, 16, 81}

	for num := range Chain(PlusOne, Square)(c) {
		got = append(got, num)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFanIn(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	sendInts := func(c chan<- int, begin, end int) {
		defer close(c)
		for i := begin; i < end; i++ {
			c <- i
		}
	}
	go sendInts(c1, 11, 14)
	go sendInts(c2, 21, 23)
	go sendInts(c3, 31, 35)

	var got []int
	var want = []int{11, 12, 13, 21, 22, 31, 32, 33, 34}

	for i := range FanIn(c1, c2, c3) {
		got = append(got, i)
	}
	sort.Ints(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
