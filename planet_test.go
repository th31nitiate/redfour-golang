package planet

import "testing"
import "math/big"
//import "fmt" //Just playing with this func to learn the big lib


//Just using part of this in order to run the code for now
func TestEV(t *testing.T) {
    t.Log("Running escape velocity tests for Earth")
    earth := Planet{"earth", 6.371e6, 5.97e24}
    if v := earth.calcEV(); v != 11.180501 {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }

    t.Log("Running escape velocity tests for Planet X")
    bNibiru := BigPlanet{"nibiru", big.NewFloat(4.0e22), big.NewFloat(6.21e6), big.NewFloat(0.0)}
    if v := bNibiru.calcBigEV(); v != big.NewFloat(0.926962) {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }

    t.Log("Running escape velocity tests for Mars")
    bMars := BigPlanet{"mars", big.NewFloat(6.39e23), big.NewFloat(3.4e6), big.NewFloat(0.0)}
    if v := bMars.calcBigEV(); v != big.NewFloat(5.007130) {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }


    t.Log("Running escape velocity tests for Moon")
    bMoon := BigPlanet{"moon", big.NewFloat(7.35e22), big.NewFloat(1.738e6), big.NewFloat(0.0)}
    if v := bMoon.calcBigEV(); v != big.NewFloat(2.375181) {
        t.Errorf("Issue with escape velocity please check equation: '%f'", v)
    }

    //This is a patch at best
}

func TestConverter(t *testing.T) {
    t.Log("Running miles to light seconds converter test")
    value := BigUnits{"miles_to_ls", big.NewFloat(1000), 5}
    if v := value.bConverter(); v != big.NewFloat(0.005368) {
        t.Errorf("Converter test value returned: '%f'", v)
    }

}

func TestOrbitor(t *testing.T) {
    t.Log("Running orbiter tests for a height of 100km")
    value := BigPlanet{"mars", big.NewFloat(6.39e23), big.NewFloat(3.4e6), big.NewFloat(100.0)}
    if v := value.calcBigOA(); v != big.NewFloat(0.005368) {
        t.Errorf("Converter test value returned: '%f'", v)
    }

}