package planet

import ("math")

type Planet struct {
    name string
    radius float64
    mass float64
}

type Units struct {
    unit string
    value float64
    precision uint32
}

var ngc float64 = 6.67e-11

func (planet *Planet) calcEV() float64 {
    gmr := 2 * ngc * planet.mass / planet.radius
    vms := math.Sqrt(gmr)
    return vms
}

func (units *Units) converter() float64 {
    switch units.unit {
        case "to_meters":
            return units.value * 1000
        case "to_km":
            return units.value / 1000
        case "miles_to_ls":
            return roundTo(units.value * 5.36819e-6, units.precision)
        case "inches_to_ls":
            return roundTo(units.value * 8.472522095734715723e-11, units.precision)
        case "feet_to_ls":
            return roundTo(units.value * 1.016702651488166404e-9, units.precision)
        case "meters_to_ls":
            return roundTo(units.value * 3.335638620368e-9, units.precision)
        default:
            return 0.0
    }
    return 0.0
}

func roundTo(n float64, decimals uint32) float64 {
    return math.Round(n * math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}