package main

import "fmt"

var houses = []House{
	{
		ID:               1,
		Name:             "Penthause somoni",
		Owner:            "Bob",
		Price:            350_000,
		DistanceToCenter: 400,
		City:             City{Name: "Khujand"},
		District:         District{Name: "Airport"},
	}, {
		ID:               2,
		Name:             "Appartment somoni",
		Owner:            "Tyler",
		Price:            300_000,
		DistanceToCenter: 300,
		City:             City{Name: "Dushanbe"},
		District:         District{Name: "Шохмансур"},
	}, {
		ID:               3,
		Name:             "Penthouse somoni",
		Owner:            "Marla",
		Price:            500_000,
		DistanceToCenter: 100,
		City:             City{Name: "Dushanbe"},
		District:         District{Name: "Сомони"},
	}, {
		ID:               5,
		Name:             "New House",
		Owner:            "Cornelius",
		Price:            450_000,
		DistanceToCenter: 600,
		City:             City{Name: "Dushanbe"},
		District:         District{Name: "Сино"},
	}, {
		ID:               6,
		Name:             "Old House",
		Owner:            "Bob",
		Price:            550_000,
		DistanceToCenter: 550,
		City:             City{Name: "Dushanbe"},
		District:         District{Name: "Сино"},
	}, {
		ID:               4,
		Name:             "House in center of Norak",
		Owner:            "Marla",
		Price:            2_200_000,
		DistanceToCenter: 1100,
		City:             City{Name: "Norak"},
		District:         District{Name: "Center"},
	},
}

func ExampleSortByPrice() {
	fmt.Println(SortByPrice(houses))
	// Output: [{2  Appartment somoni 300000   300 {Dushanbe} {0 Шохмансур} Tyler} {1  Penthause somoni 350000   400 {Khujand} {0 Airport} Bob} {5  New House 450000   600 {Dushanbe} {0 Сино} Cornelius} {3  Penthouse somoni 500000   100 {Dushanbe} {0 Сомони} Marla} {6  Old House 550000   550 {Dushanbe} {0 Сино} Bob} {4  House in center of Norak 2200000   1100 {Norak} {0 Center} Marla}]
}

func ExampleSortByMaxPrice() {
	fmt.Println(SortByMaxPrice(houses, 300_000))
	// Output:  [{1  Penthause somoni 350000   400 {Khujand} {0 Airport} Bob} {2  Appartment somoni 300000   300 {Dushanbe} {0 Шохмансур} Tyler} {3  Penthouse somoni 500000   100 {Dushanbe} {0 Сомони} Marla} {5  New House 450000   600 {Dushanbe} {0 Сино} Cornelius} {6  Old House 550000   550 {Dushanbe} {0 Сино} Bob} {4  House in center of Norak 2200000   1100 {Norak} {0 Center} Marla}]
}

func ExampleSortByMinPrice() {
	fmt.Println(SortByMinPrice(houses, 300_000))
	// Output: [{2  Appartment somoni 300000   300 {Dushanbe} {0 Шохмансур} Tyler}]
}

func ExampleSearchByPriceBetween() {
	fmt.Println(SearchByPriceBetween(houses, 350_000))
	// Output: [{1  Penthause somoni 350000   400 {Khujand} {0 Airport} Bob}]
}

func ExampleSortByDistanceMin() {
	fmt.Println(SortByDistanceMin(houses))
	// Output: [{3  Penthouse somoni 500000   100 {Dushanbe} {0 Сомони} Marla} {2  Appartment somoni 300000   300 {Dushanbe} {0 Шохмансур} Tyler} {1  Penthause somoni 350000   400 {Khujand} {0 Airport} Bob} {6  Old House 550000   550 {Dushanbe} {0 Сино} Bob} {5  New House 450000   600 {Dushanbe} {0 Сино} Cornelius} {4  House in center of Norak 2200000   1100 {Norak} {0 Center} Marla}]
}

func ExampleSortByDistanceMax() {
	fmt.Println(SortByDistanceMax(houses))
	// Output:  [{1  Penthause somoni 350000   400 {Khujand} {0 Airport} Bob} {2  Appartment somoni 300000   300 {Dushanbe} {0 Шохмансур} Tyler} {5  New House 450000   600 {Dushanbe} {0 Сино} Cornelius} {3  Penthouse somoni 500000   100 {Dushanbe} {0 Сомони} Marla} {4  House in center of Norak 2200000   1100 {Norak} {0 Center} Marla} {6  Old House 550000   550 {Dushanbe} {0 Сино} Bob}]
}

func ExampleFindByDistrict() {
	fmt.Println(FindByDistrict(houses, "Сино"))
	// Output: [{5  New House 450000   600 {Dushanbe} {0 Сино} Cornelius} {6  Old House 550000   550 {Dushanbe} {0 Сино} Bob}]
}

func ExampleFindByDistricts() {
	districts := []string{"Сино", "Сомони"}
	fmt.Println(FindByDistricts(houses, districts))
	// Output: [{3  Penthouse somoni 500000   100 {Dushanbe} {0 Сомони} Marla} {5  New House 450000   600 {Dushanbe} {0 Сино} Cornelius} {6  Old House 550000   550 {Dushanbe} {0 Сино} Bob}]
}
