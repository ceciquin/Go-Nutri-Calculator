package main

import(
	"fmt"
)


func main(){

	ns := GetNutrionalScore(NutrionalData{
		Energy: EnergyFromKcal(),
		Sugars: SugarGram(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium: SodiumMilligram(),
		Fruits: FruitsPercent(),
		Fibre: FibreGram(),
		Protein: ProteinGram(),
	}, Food)

	fmt.Printf("Nutritional score: %d\n", ns.Value)
}