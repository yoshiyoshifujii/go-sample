package main

import "fmt"

// Animal ベース構造体
type Animal struct{}

func (a Animal) Speak() {
	fmt.Println("Animal makes a sound")
}

// Dog Animalを継承したDog構造体
type Dog struct {
	Animal
}

// Speak DogがAnimalのSpeakメソッドをオーバーライド
func (d Dog) Speak() {
	fmt.Println("Dog barks")
}

type Speaker interface {
	Speak()
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Cat meows")
}

func performSpeak(s Speaker) {
	s.Speak()
}

func main() {
	animal := Animal{}
	dog := Dog{}
	cat := Cat{}

	performSpeak(animal)
	performSpeak(dog)
	performSpeak(cat)
}
