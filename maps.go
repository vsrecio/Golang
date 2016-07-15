package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["Demo"] = 20
	fmt.Println(x["Demo"])

	//

	myMaps := make(map[string]string)

	myMaps["AFG"] = "Afghanistan"
	myMaps["AHO"] = "Netherlands Antilles"
	myMaps["ALB"] = "Albania"

	fmt.Println(myMaps["AHO"])
	
	//
	myOtherMaps := map[string]string {
		"H" :    "Hydrogen",
		"He":    "Helium",
		"Li":    "Lithium",
	}
	fmt.Println(myOtherMaps["H"])
	//

	dominicanRepublic := map[string]map[string]string{

		"Altamira": map[string]string {
			"name"      :   "Puerto Plata",
			"population":   "19,052",
			"area-km2"  :   "179.32",
			"density"   :   "106/km2",
			"area code" :   "57091",
		},
		"Arenoso":  map[string]string {
				"name"      :   "Duarte",
				"population":   "15,122",
				"area-km2"  :   "142.20",
				"density"   :   "106/km2",
				"area code" :   "31006",
		},
		"Azua de Compostela":   map[string]string {
				"name"      :   "Azua",
				"population":   "125,487",
				"area-km2"  :   "416.30",
				"density"   :   "301/km2",
				"area code" :   "71000",
		},

	}


}
