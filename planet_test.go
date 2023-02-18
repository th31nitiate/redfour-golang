package planet

import "testing"

func TestEV(t *testing.T) {
    t.Log("Running escape velocity tests for earth")
    earth := Planet{"earth", 6.371e6, 5.97e24}
    if v := earth.calcEV(); v != 11.2 {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }

    t.Log("Running escape velocity tests for planet X")
    earth := Planet{"nibiru", 4.0e22, 6.21e6}
    if v := earth.calcEV(); v != 0.9 {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }
}