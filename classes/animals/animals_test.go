package animals

import (
	"testing"
)

var (
	animal *Animal
)

// to run once before all tests for initializing data
func setupSuite(t testing.TB){
	animal = NewAnimal("eats", "moves", "speaks")
}

func TestEat(t *testing.T)	{
	setupSuite(t)
	actual := animal.Eat()
	expected := "eats"

	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestMove(t *testing.T)	{
	actual := animal.Move()
	expected := "moves"

	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}
func TestSpeak(t *testing.T)	{
	actual := animal.Speak()
	expected := "speaks"

	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}
