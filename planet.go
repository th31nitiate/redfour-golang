package planet

import ("math")


type Planet struct {
    name string
    radius float64
    mass float64
}

var ngc float64 = 6.67e-11

func calcEV() float64 {
    earth := Planet{"earth", 6.371e6, 5.97e24}
    gmr := 2 * ngc * earth.mass / earth.radius
    vms := math.Sqrt(gmr)
    vkms := vms / 1000
    return roundTo(vkms, 1)
}

func roundTo(n float64, decimals uint32) float64 {
    return math.Round(n * math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}