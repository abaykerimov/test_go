package algo_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCall1(t *testing.T) {
	a1 := A1{}
	got, _ := call(a1)
	want := []string{"A", "B", "C"}

	assert.Equal(t, got, want, "The two vars should be the same.")
}

func TestCall2(t *testing.T) {
	a2 := A2{}
	got, err := call(a2)
	want := []string{"B", "C"}

	var expectedError = errors.New("error in A")
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}

	assert.Equal(t, got, want, "The two vars should be the same.")
}

func TestCall3(t *testing.T) {
	a3 := A3{}
	got, err := call(a3)
	want := []string{"B", "C"}

	var expectedError = errors.New("runtime error: integer divide by zero")
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}

	assert.Equal(t, got, want, "The two vars should be the same.")
}

func TestCall4(t *testing.T) {
	a4 := A4{}
	got, _ := call(a4)
	want := []string{"B", "C"}

	assert.Equal(t, got, want, "The two vars should be the same.")
}

type A1 struct{}
type A2 struct{}
type A3 struct{}
type A4 struct{}

type funcA interface {
	A(a chan string, p chan interface{}) (string, error)
}

func call(i funcA) (d []string, err error) {
	a := make(chan string)
	b := make(chan string)
	c := make(chan string)
	p := make(chan interface{})

	go i.A(a, p)
	go B(b)
	go C(c)

	switch v := i.(type) {
	case A1:
		d = append(d, <-a, <-b, <-c)
	case A2:
		select {
		case <-a:
			d = append(d, <-a, <-b, <-c)
			return
		case errA := <-p:
			err = fmt.Errorf("%s", errA)
			d = append(d, <-b, <-c)
			return
		}
	case A3:
		select {
		case <-a:
			d = append(d, <-a, <-b, <-c)
			return
		case rec := <-p:
			err = fmt.Errorf("%s", rec)
			d = append(d, <-b, <-c)
			return
		}
	case A4:
		select {
		case <-time.After(2 * time.Second):
			d = append(d, <-b, <-c)
			return
		}
	default:
		fmt.Printf("I don't know about type %T!\n", v)
		return
	}

	return
}

func (an A1) A(a chan string, p chan interface{}) (string, error) {
	a <- "A"
	return "", nil
}

func (an A2) A(a chan string, p chan interface{}) (string, error) {
	err := errors.New("error in A")
	p <- err
	a <- "A"
	return "", err
}

func (an A3) A(a chan string, p chan interface{}) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			p <- r
			fmt.Println("recover err: ", r)
		}
	}()
	var one, two = 100, 0
	div := one / two
	fmt.Println(div)
	a <- "A"
	return "", nil
}

func (an A4) A(a chan string, p chan interface{}) (string, error) {
	time.Sleep(10 * time.Second)
	a <- "A"
	return "", nil
}

func B(b chan string) (string, error) {
	b <- "B"
	return "", nil
}

func C(c chan string) (string, error) {
	c <- "C"
	return "", nil
}
