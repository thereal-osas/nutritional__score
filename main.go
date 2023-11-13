package main

import "fmt"

func main() {
	nutriScore := GetNutritionalScore(NutritionalData{
		Energy: EnergyFromKcal(0),
		Sugars: SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(2),
		Sodium: SodiumMilligram(500),
		Fruits: FruitsPercent(60),
		Fibre: FibreGram(4),
		Protein: ProteinGram(2),
	}, Food) 

	fmt.Printf("Nutritonal Score: %d\n", nutriScore.Value)
	fmt.Printf("NutriScore: %s\n", nutriScore.GetNutriScore())
}