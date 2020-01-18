package main

import "log"

var houses = []House{
	{
		ID:               1,
		Name:             "Penthause somoni",
		Owner:            "Bob",
		Price:            350_000,
		DistanceToCenter: 400,
		City:             City{Name: "Khujand"},
		Rayon:            Rayon{Name: "Airport"},
	}, {
		ID:               2,
		Name:             "Appartment somoni",
		Owner:            "Tyler",
		Price:            300_000,
		DistanceToCenter: 300,
		City:             City{Name: "Dushanbe"},
		Rayon:            Rayon{Name: "Шохмансур"},
	}, {
		ID:               3,
		Name:             "Penthouse somoni",
		Owner:            "Marla",
		Price:            500_000,
		DistanceToCenter: 100,
		City:             City{Name: "Dushanbe"},
		Rayon:            Rayon{Name: "Сомони"},
	}, {
		ID:               5,
		Name:             "New House",
		Owner:            "Cornelios",
		Price:            450_000,
		DistanceToCenter: 600,
		City:             City{Name: "Dushanbe"},
		Rayon:            Rayon{Name: "Сино"},
	}, {
		ID:               5,
		Name:             "Old House",
		Owner:            "Bob",
		Price:            550_000,
		DistanceToCenter: 550,
		City:             City{Name: "Dushanbe"},
		Rayon:            Rayon{Name: "Сино"},
	}, {
		ID:               4,
		Name:             "House in center of Norak",
		Owner:            "Marla",
		Price:            200_000,
		DistanceToCenter: 1100,
		City:             City{Name: "Norak"},
		Rayon:            Rayon{Name: "Center"},
	},
}

func ExampleSortByMaxPrice() {
	byPrice := SortByMaxPrice(houses, 300_000)
	log.Println(byPrice)
}

func ExampleSortByMinPrice() {
	byPrice := SortByMinPrice(houses, 450_000)
	log.Println(byPrice)
}

func ExampleSortByPrice() {
	sortBy := SortByPrice(houses, 300_000)
	log.Println(sortBy)
}

func ExampleSearchByPriceBetween() {
	betweenPrice := SearchByPriceBetween(houses, 350_000)
	log.Println(betweenPrice)
}

func ExampleFindByDistrict() {
	result := FindByDistrict(houses, "Center")
	log.Println(result)
}

func ExampleFindByDistricts() {
	districts := []string{"Сино", "Сомони"}
	result := FindByDistricts(houses, districts)
	log.Println(result)
}

func ExampleSortByDistanceMin() {
	sortBy := SortByDistanceMin(houses)
	log.Println(sortBy)
}

func ExampleSortByDistanceMax() {
	sortBy := SortByDistanceMax(houses)
	log.Println(sortBy)
}
