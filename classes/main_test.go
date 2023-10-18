package main

import (
	animals "classes/animals"
	"testing"
)

var (cow, bird, snake *animals.Animal)

func setupSuite(t testing.TB)	{
	cow, bird, snake = GenData()
}

func TestCowEat(t *testing.T)	{
	setupSuite(t)
	expected := "grass"
	actual := cow.Eat()

	if expected != actual {
		t.Errorf("Expected %s, received %s for cow eat", expected, actual)
	}
}

func TestCowMove(t *testing.T)	{
	expected := "walk"
	actual := cow.Move()

	if expected != actual {
		t.Errorf("Expected %s, received %s for cow move", expected, actual)
	}
}

func TestCowSpeak(t *testing.T)	{
	expected := "moo"
	actual := cow.Speak()

	if expected != actual {
		t.Errorf("Expected %s, received %s for cow speak", expected, actual)
	}
}

func TestBirdEat(t *testing.T)	{
	setupSuite(t)
	expected := "worms"
	actual := bird.Eat()

	if expected != actual {
		t.Errorf("Expected %s, received %s for bird eat", expected, actual)
	}
}

func TestBirdMove(t *testing.T)	{
	expected := "fly"
	actual := bird.Move()

	if expected != actual {
		t.Errorf("Expected %s, received %s for bird move", expected, actual)
	}
}

func TestBirdSpeak(t *testing.T)	{
	expected := "peep"
	actual := bird.Speak()

	if expected != actual {
		t.Errorf("Expected %s, received %s for bird speak", expected, actual)
	}
}

func TestSnakeEat(t *testing.T)	{
	setupSuite(t)
	expected := "mice"
	actual := snake.Eat()

	if expected != actual {
		t.Errorf("Expected %s, received %s for snake eat", expected, actual)
	}
}

func TestSnakeMove(t *testing.T)	{
	expected := "slither"
	actual := snake.Move()

	if expected != actual {
		t.Errorf("Expected %s, received %s for snake move", expected, actual)
	}
}

func TestSnakeSpeak(t *testing.T)	{
	expected := "hsss"
	actual := snake.Speak()

	if expected != actual {
		t.Errorf("Expected %s, received %s for snake speak", expected, actual)
	}
}