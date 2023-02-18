package planet

import "testing"

func TestEV(t *testing.T) {
    t.Log("Running Escape velocity tests")
    earth := Planet{"earth", 6.371e6, 5.97e24}
    if v := earth.calcEV(); v != 11.2 {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }

}