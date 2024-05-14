package main

import "testing"

func TestArea(t *testing.T) {
    var cube = Cube{Side: 2}
    var expected float64 = 24.0

    if cube.Area() != expected {
        t.Errorf("Expected %f, but got %f", expected, cube.Area())
    }
}

func TestCircumference(t *testing.T) {
    var cube = Cube{Side: 2}
    var expected float64 = 24.0

    if cube.Circumference() != expected {
        t.Errorf("Expected %f, but got %f", expected, cube.Circumference())
    }
}

func TestVolume(t *testing.T) {
    var cube = Cube{Side: 2}
    var expected float64 = 8.0

    if cube.Volume() != expected {
        t.Errorf("Expected %f, but got %f", expected, cube.Volume())
    }
}