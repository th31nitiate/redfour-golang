package planet

import ("math/big";
	"math";)


type Planet struct {
    name string
    mass_kg *big.Float
    radius_m *big.Float
    orbitor_height *big.Float
}

type Units struct {
    unit string
    value *big.Float
    precision uint32
}

var Ngc = big.NewFloat(6.67e-11)

//Using new to define big floats in this funcation may
//not be the best but since it give our desired results
//We will keep for now then refactor later as required
func (bPlanet *Planet) calcBigEV() *big.Float {
    radius_m, _ := bPlanet.radius_m.Float64()
    mass_kg, _ := bPlanet.mass_kg.Float64()
    f64Radius_m := big.NewFloat(radius_m)          
    f64Mass_kg := big.NewFloat(mass_kg)
    ms := new(big.Float).Quo(f64Mass_kg, f64Radius_m)
    nmr := new(big.Float).Mul(Ngc, ms)
    gmr := new(big.Float).Mul(big.NewFloat(2), nmr)
    vkms := new(big.Float).Sqrt(gmr)
    value := Units{"to_km", vkms, 5} //-- As you can see we are not calling converter yet
    return value.bConverter()
}

func (bPlanet *Planet) calcBigOS() *big.Float {
    mass_kg, _ := bPlanet.mass_kg.Float64()
    OrbRadius := bPlanet.calcBigOR()
    f64Mass_kg := big.NewFloat(mass_kg)
    ms := new(big.Float).Quo(OrbRadius, f64Mass_kg)
    nmr := new(big.Float).Mul(Ngc, ms)
    vkms := new(big.Float).Sqrt(nmr)
    return vkms
}

func (bPlanet *Planet) calcBigOA() *big.Float {
    OrbRadius := bPlanet.calcBigOR()
    OrbSpeed := bPlanet.calcBigOS()
    //f64Mass_kg := new(big.Float).Mul(Ngc, ms)
    //ms := new(big.Float).Quo(OrbRadius, f64Mass_kg)
    nmr := new(big.Float).Mul(OrbSpeed, OrbSpeed)
    vkms := new(big.Float).Quo(nmr, OrbRadius)
    return vkms
}


func (bPlanet *Planet) calcBigOR() *big.Float {
    radius_m, _ := bPlanet.radius_m.Float64()
    orbitor_height, _ := bPlanet.orbitor_height.Float64()
    newOrbit := Units{"to_meters", big.NewFloat(orbitor_height), 5}
    OrbRadius_m := new(big.Float).Add(big.NewFloat(radius_m), newOrbit.bConverter())
    return OrbRadius_m
}

func (units *Units) bConverter() *big.Float {
    switch units.unit {
        case "to_meters":
            return new(big.Float).Mul(units.value, big.NewFloat(1000))
        case "miles_to_ls":
            return new(big.Float).Mul(units.value, big.NewFloat(5.36819e-6))
        case "inches_to_ls":
            return new(big.Float).Mul(units.value, big.NewFloat(8.472522095734715723e-11))
        case "feet_to_ls":
            return new(big.Float).Mul(units.value, big.NewFloat(1.016702651488166404e-9))
        case "meters_to_ls":
            return new(big.Float).Mul(units.value, big.NewFloat(3.335638620368e-9))
        case "to_km":
            return new(big.Float).Quo(units.value, big.NewFloat(1000))
    }
    panic("Error unknown unit")
}

func roundTo(n float64, decimals uint32) float64 {
    return math.Round(n * math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
