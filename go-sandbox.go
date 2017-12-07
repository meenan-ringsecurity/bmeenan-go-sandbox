package main


import (
	"fmt"
	"time"
	"reflect"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("hi")
	dateThat()
	passThatFunction()
	callThat()
	configThat()
}

func dateThat() {
	var start = time.Now()
	fmt.Println(start.Month())
	fmt.Println(start.Day())
	fmt.Println(start.Year())
}

func callThat() {
	fmt.Println("rest call to some endpoint goes here")
}

type Foo int

func passThatFunction() {
	methodExpression()
	methodValues()
	methodReflection()
}

func methodExpression() {
	var f Foo
	bar := func(m func(f Foo)) {
		m(f)
	}
	bar(Foo.A)
	bar(Foo.B)
	bar(Foo.C)
}

func methodValues() {
	var f Foo
	bar := func(m func()) {
		m()
	}
	bar(f.A)
	bar(f.B)
	bar(f.C)
}

func methodReflection() {
	var f Foo
	bar := func(name string) {
		reflect.ValueOf(f).MethodByName(name).Call(nil)
	}
	bar("A")
	bar("B")
	bar("C")
}

func (f Foo) A() {
	fmt.Println("A")
}
func (f Foo) B() {
	fmt.Println("B")
}
func (f Foo) C() {
	fmt.Println("C")
}

func configThat() {
	var key, value = "yo", "yo yo"
	x := viper.New()
	x.SetDefault(key, value)
	y := x.GetString(key)
	fmt.Println(y)
}