package algo_test

import (
	"fmt"
	"testing"
	"time"
)

func TestCall1(t *testing.T) {
	got, err := call1()
	want := []string{"A", "B", "C"}

	if err == nil {
		t.Skip("no error")
	} else {
		t.Error("err: ", err)
	}

	if !Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCall2(t *testing.T) {
	got, err := call2()
	want := []string{"B", "C"}

	if err != nil {
		t.Error("err: ", err)
	}

	if !Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCall3(t *testing.T) {
	got, err := call3()
	want := []string{"B", "C"}

	if err != nil {
		t.Error("err: ", err)
	}

	if !Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCall4(t *testing.T) {
	got, err := call4()
	want := []string{"B", "C"}

	if err != nil {
		t.Error("err: ", err)
	}

	if !Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

type aResponse struct {
	a   string
	div int
	err error
	r   interface{}
}

func call1() ([]string, error) {
	a := make(chan string)
	b := make(chan string)
	c := make(chan string)
	d := make([]string, 0)

	go A1(a)
	go B(b)
	go C(c)

	d = append(d, <-a, <-b, <-c)

	return d, nil
}

func call2() (d []string, err error) {
	a := make(chan aResponse)
	b := make(chan string)
	c := make(chan string)

	go A2(a)
	go B(b)
	go C(c)

	d = append(d, <-b, <-c)

	result := <-a
	if result.err != nil {
		return
	}
	d = append(d, result.a)

	return
}

func call3() (d []string, err error) {
	a := make(chan aResponse)
	b := make(chan string)
	c := make(chan string)
	p := make(chan interface{})

	go A3(100, 0, a, p)
	go B(b)
	go C(c)

	select {
	case <-a:
		result := <-a
		if result.err != nil {
			return
		}
		d = append(d, result.a, <-b, <-c)
		return
	case <-p:
		d = append(d, <-b, <-c)
		return
	}
}

func call4() (d []string, err error) {
	a := make(chan string)
	b := make(chan string)
	c := make(chan string)

	go A4(a)
	go B(b)
	go C(c)

	select {
	case <-time.After(2 * time.Second):
		d = append(d, <-b, <-c)
		return
	}
}

func A1(a chan string) (response string, err error) {
	response = "A"
	a <- response
	return
}

func A2(a chan aResponse) (response aResponse, err error) {
	response.err = fmt.Errorf("error in A")
	a <- response
	return
}

func A3(one, two int, a chan aResponse, p chan interface{}) (response aResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			p <- r
			fmt.Println("recover err: ", r)
		}
	}()
	response.a = "A"
	response.div = one / two
	a <- response
	return
}

func A4(a chan string) (response string, err error) {
	response = "A"
	time.Sleep(10 * time.Second)
	a <- response
	return
}

func B(b chan string) (response string, err error) {
	response = "B"
	b <- response
	return
}

func C(c chan string) (response string, err error) {
	response = "C"
	c <- response
	return
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
