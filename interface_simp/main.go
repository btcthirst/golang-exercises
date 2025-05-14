package main

type Animal interface {
	Speak() string
}
type Dog struct {
	name string
	age  int
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	name string
	age  int
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Bird struct {
	name string
	age  int
}

func (b Bird) Speak() string {
	return "Tweet!"
}

func main() {
	doggy := Dog{
		name: "Buddy",
		age:  3,
	}
	teddy := Dog{
		name: "Max",
		age:  5,
	}
	catty := Cat{
		name: "Whiskers",
		age:  2,
	}
	kitty := Cat{
		name: "Mittens",
		age:  4,
	}
	birdy := Bird{
		name: "Tweety",
		age:  1,
	}
	animals := []Animal{doggy, catty, teddy, birdy, kitty}
	for _, animal := range animals {
		println(animal.Speak())
	}

}
