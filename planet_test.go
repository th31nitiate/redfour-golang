package planet

import "testing"

func TestEV(t *testing.T) {
    t.Log("Running Escape velocity tests")
    if v := calcEV(); v != 11.2 {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }

}