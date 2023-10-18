package dispacement

import "testing"

func TestZeros(t *testing.T)	{
	expected := 0.00
	fn:= GenDisplaceFn(0, 0, 0)
	time := 3.0
	actual := fn(time)

	if expected != actual 	{
		t.Errorf("Expected %f, received %f for time %f", expected, actual, time)
	}
}

func TestValues(t *testing.T)	{
	expected := 52.00
	fn:= GenDisplaceFn(10, 2, 1)
	time := 3.0
	actual := fn(time)

	if expected != actual 	{
		t.Errorf("Expected %f, received %f for time %f", expected, actual, time)
	}
}

func TestValues2(t *testing.T)	{
	expected := 136.00
	fn:= GenDisplaceFn(10, 2, 1)
	time := 5.0
	actual := fn(time)

	if expected != actual 	{
		t.Errorf("Expected %f, received %f for time %f", expected, actual, time)
	}
}