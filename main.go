package main

import "fmt"

func main() {
	nutriScore := GetNutritionalScore(NutritionalData{
		Energy: EnergyFromKcal(),
		Sugars: SugarGram(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium: SodiumMilligram(),
		Fruits: FruitsPercent(),
		Fibre: FibreGram(),
		Protein: ProteinGram(),
	}, Food) 

	fmt.Printf("Nutritonal Score: %d\n", nutriScore.Value)
	fmt.Printf("NutriScore: %s\n", nutriScore.GetNutriScore())
}