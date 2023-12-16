package house

import (
	"awesomeProject/awesomeProject/appliance"
	"awesomeProject/awesomeProject/family"
	"awesomeProject/awesomeProject/furniture"
	"awesomeProject/awesomeProject/kids"
	"awesomeProject/awesomeProject/pets"
	"fmt"
)

type House struct {
	Sizes         []float64
	FamilyInfo    []family.Family
	KidsInfo      []kids.Kids
	FurnitureInfo []furniture.Furniture
	ApplianceInfo []appliance.Appliance
	PetsInfo      []pets.Pets
}

func CreateAwesomeHouse() House {
	return House{
		Sizes:         []float64{150.0, 200.0, 100.0},
		FamilyInfo:    family.AddFamilyInfo(),
		KidsInfo:      kids.AddKidsInfo(),
		FurnitureInfo: furniture.AddFurnitureInfo(),
		ApplianceInfo: appliance.AddApplianceInfo(),
		PetsInfo:      pets.AddPetsInfo(),
	}
}

func DisplayAwesomeHouse() {
	house := CreateAwesomeHouse()

	fmt.Printf("\nMy awesome house:\n")
	for _, size := range house.Sizes {
		fmt.Printf("%.2f sq.m\n", size)
	}

	displayFamilyInfo(house.FamilyInfo, "\nFamily:")
	displayKidsInfo(house.KidsInfo, "\nKids:")
	displayFurnitureInfo(house.FurnitureInfo, "\nFurniture:")
	displayApplianceInfo(house.ApplianceInfo, "\nAppliance:")
	displayPetsInfo(house.PetsInfo, "\nPets:")

}

func displayFamilyInfo(familyInfo []family.Family, header string) {
	fmt.Println(header)
	for _, obj := range familyInfo {
		fmt.Printf("- %s Name %s Age: %d years\n", obj.Member, obj.Name, obj.Age)
	}
}

func displayKidsInfo(kidsInfo []kids.Kids, header string) {
	fmt.Println(header)
	for _, obj := range kidsInfo {
		fmt.Printf("- %s: Name %s Age: %d years\n", obj.Member, obj.Name, obj.Age)
	}
}

func displayFurnitureInfo(furnitureInfo []furniture.Furniture, header string) {
	fmt.Println(header)
	for _, obj := range furnitureInfo {
		fmt.Printf("- %s Size: %0.1f m\n", obj.Name, obj.Size)
	}
}

func displayApplianceInfo(applianceInfo []appliance.Appliance, header string) {
	fmt.Println(header)
	for _, obj := range applianceInfo {
		fmt.Printf("- %s Brand: %s, Count: %d\n", obj.Name, obj.Brand, obj.Count)
	}
}

func displayPetsInfo(petsInfo []pets.Pets, header string) {
	fmt.Println(header)
	for _, obj := range petsInfo {
		fmt.Printf("- Owner: %s Pet: %s шт.\n", obj.Owner, obj.Name)
	}
}
