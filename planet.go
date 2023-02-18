package planet

import ("math"; "fmt")

type Planet struct {
    name string
    radius float64
    mass float64
}

type Units struct {
    unit string
    value float64
    precision int
}

var ngc float64 = 6.67e-11

func (planet *Planet) calcEV() float64 {
    gmr := 2 * ngc * planet.mass / planet.radius
    vms := math.Sqrt(gmr)
    return vms
}

func (units *Units) converter() float64 {
    switch units.unit {
        case "meters":
            roundTo(units.value * 5.36819e-6, units.precision)
            return
        case "miles":
            return roundTo(units.value * 3.335638620368e-9, units.precision)
        case "kilometers":
            return roundTo(units.value / 1000, units.precision)
        default:
            return 0.0
    }
    return 0.0
}

func roundTo(n float64, decimals uint32) float64 {
    return math.Round(n * math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}