package main

/*
Suppose we have a class:

	public class Foo {
	  public void first() { print("first"); }
	  public void second() { print("second"); }
	  public void third() { print("third"); }
	}

The same instance of Foo will be passed to three different threads.
Thread A will call first(), thread B will call second(), and thread C will call third().
Design a mechanism and modify the program to ensure that second() is executed after first(), and third()
is executed after second().

Note:

We do not know how the threads will be scheduled in the operating system,
even though the numbers in the input seem to imply the ordering.
The input format you see is mainly to ensure our tests' comprehensiveness.
*/
func mainasdjfijio2ih1h1h1h1h() {
}

type Foo struct {
	first  chan bool
	second chan bool
}

func NewFoo() *Foo {
	return &Foo{
		first:  make(chan bool, 1),
		second: make(chan bool, 1),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.first <- true
}

func (f *Foo) Second(printSecond func()) {
	select {
	case <-f.first:
		/// Do not change this line
		printSecond()
		f.second <- true
	}
}

func (f *Foo) Third(printThird func()) {
	select {
	case <-f.second:
		// Do not change this line
		printThird()
	}
}
