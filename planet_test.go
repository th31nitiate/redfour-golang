package planet

import "testing"
import "math/big"
//import "fmt" //Just playing with this func to learn the big lib


//Just using part of this in order to run the code for now
func TestEV(t *testing.T) {
    t.Log("Running escape velocity tests for Earth")
    earth := Planet{"earth", 6.371e6, 5.97e24}
    if v := earth.calcEV(); v != 11.2 {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }

    t.Log("Running escape velocity tests for Planet X")
    bNibiru := BigPlanet{"nibiru", big.NewFloat(4.0e22), big.NewFloat(6.21e6)}
    if v := bNibiru.calcBigEV(); v != big.NewFloat(0.9) {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }

    t.Log("Running escape velocity tests for Mars")
    bMars := BigPlanet{"mars", big.NewFloat(6.39e23), big.NewFloat(3.4e6)}
    if v := bMars.calcBigEV(); v != big.NewFloat(5.0) {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }


    t.Log("Running escape velocity tests for Moon")
    bMoon := BigPlanet{"moon", big.NewFloat(7.35e22), big.NewFloat(1.738e6)}
    if v := bMoon.calcBigEV(); v != big.NewFloat(0.8) {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }

    //This is a patch at best

}

