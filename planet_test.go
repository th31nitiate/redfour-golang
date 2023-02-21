package planet

import "testing"
import "math/big"
import "fmt" //Just playing with this func to learn the big lib


//Just using part of this in order to run the code for now
func TestEV(t *testing.T) {
    t.Log("Running escape velocity tests for Nibiru")
    bNibiru := BigPlanet{"nibiru", big.NewFloat(4.0e22), big.NewFloat(6.21e6)}
    fmt.Printf("Nibiru escape velocity values: '%f' \n", bNibiru.calcBigEV())


    t.Log("Running escape velocity tests for Mars")
    bMars := BigPlanet{"mars", big.NewFloat(6.39e23), big.NewFloat(3.4e6)}
    fmt.Printf("Nibiru escape velocity values: '%f' \n", bMars.calcBigEV())


    t.Log("Running escape velocity tests for Earth")
    earth := Planet{"earth", 6.371e6, 5.97e24}
    if v := earth.calcEV(); v != 11.2 {
        t.Errorf("Incorrect value for earth defined here: '%f'", v)
    }

}

