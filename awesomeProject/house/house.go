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
	RoomsCount    int
	FloorsCount   int
	Area          float64
	FamilyInfo    []family.Family
	KidsInfo      []kids.Kids
	FurnitureInfo []furniture.Furniture
	ApplianceInfo []appliance.Appliance
	PetsInfo      []pets.Pets
}

func CreateHouse() House {
	return House{
		RoomsCount:    10,
		FloorsCount:   3,
		Area:          100.55,
		FamilyInfo:    family.AddFamilyInfo(),
		KidsInfo:      kids.AddKidsInfo(),
		FurnitureInfo: furniture.AddFurnitureInfo(),
		ApplianceInfo: appliance.AddApplianceInfo(),
		PetsInfo:      pets.AddPetsInfo(),
	}
}

func DisplayHouseInfo(house House) {
	fmt.Printf("Описание дома:\n")
	fmt.Printf("Количество комнат: %d\n", house.RoomsCount)
	fmt.Printf("Количество этажей: %d\n", house.FloorsCount)
	fmt.Printf("Площадь дома: %.2f кв. м\n", house.Area)

	displayFamilyInfo(house.FamilyInfo, "Описание членов семьи:")
	displayKidsInfo(house.KidsInfo, "Описание родственников:")
	displayFurnitureInfo(house.FurnitureInfo, "Описание мебели:")
	displayApplianceInfo(house.ApplianceInfo, "Описание техники:")
	displayPetsInfo(house.PetsInfo, "Описание одежды:")
}

func displayFamilyInfo(familyInfo []family.Family, header string) {
	fmt.Println(header)
	for _, obj := range familyInfo {
		fmt.Printf("- %s: %s\n", obj.Member, obj.Name)
	}
}

func displayKidsInfo(kidsInfo []kids.Kids, header string) {
	fmt.Println(header)
	for _, obj := range kidsInfo {
		fmt.Printf("- %s: %s\n", obj.Member, obj.Name)
	}
}

func displayFurnitureInfo(furnitureInfo []furniture.Furniture, header string) {
	fmt.Println(header)
	for _, obj := range furnitureInfo {
		fmt.Printf("- %s: %.2f кв. м, %d шт.\n", obj.Name, obj.Size, obj.Count)
	}
}

func displayApplianceInfo(applianceInfo []appliance.Appliance, header string) {
	fmt.Println(header)
	for _, obj := range applianceInfo {
		fmt.Printf("- %s: %s, %d шт.\n", obj.Name, obj.Brand, obj.Count)
	}
}

func displayPetsInfo(petsInfo []pets.Pets, header string) {
	fmt.Println(header)
	for _, obj := range petsInfo {
		fmt.Printf("- %s: %d, %d шт.\n", obj.Name, obj.Size, obj.Count)
	}
}
