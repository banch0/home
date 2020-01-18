package main

import (
	"log"
	"sort"
)

// House ...
type House struct {
	ID               int64
	Title            string
	Name             string
	Price            int64
	Currency         string
	Category         string
	DistanceToCenter int64
	City             City
	Rayon            Rayon
	Owner            string
}

// City ...
type City struct {
	Name string
}

// Rayon ...
type Rayon struct {
	ID   int8
	Name string
}

// AllHouses ...
var AllHouses = []House{
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

func main() {

	sortByPrice := SortByPrice(AllHouses, 350_000)
	log.Println("ByPrice:", sortByPrice)

	result := FindByDistrict(AllHouses, "Сино")
	log.Println("sort by district: ", result)

	districts := []string{"Сино", "Сомони"}
	result2 := FindByDistricts(AllHouses, districts)
	log.Println("sort by districts: ", result2)

	mass2 := SortByMaxPrice(AllHouses, 350000)
	log.Println("After2: ", mass2)

	byDisMin := SortByDistanceMin(AllHouses)
	log.Println("By DisMin", byDisMin)

	byDisMax := SortByDistanceMax(AllHouses)
	log.Println("By DisMax", byDisMax)

	byPrice := SortByMaxPrice(AllHouses, 300_000)
	log.Println(byPrice)

	betweenPrice := SearchByPriceBetween(AllHouses, 350_000)
	log.Println(betweenPrice)
}

// SortByPrice ... sort by max price +
func SortByPrice(massive []House, amount int64) []House {
	result := make([]House, 0)
	for _, house := range massive {
		if house.Price <= amount {
			result = append(result, house)
		}
	}
	return result
}

// MainSortFunc ...
func MainSortFunc(houses []House, predicate func(item House) bool) []House {
	result := make([]House, 0)
	for _, house := range houses {
		if predicate(house) {
			result = append(result, house)
		}
	}
	return result
}

//

// SortByMaxPrice ...
func SortByMaxPrice(houses []House, price int64) []House {
	return MainSortFunc(houses, func(house House) bool {
		return house.Price >= price
	})
}

// SortByMinPrice ...+
func SortByMinPrice(houses []House, price int64) []House {
	return MainSortFunc(houses, func(house House) bool {
		return house.Price <= price
	})
}

// end ...

// JustSortFunc ...
func JustSortFunc(houses []House, predicate func(house House) bool) []House {
	result := make([]House, len(houses))
	for _, house := range houses {
		if predicate(house) {
			result = append(result, house)
		}
	}
	return result
}

// FilterBy ...
func FilterBy(houses []House, less func(a, b House) bool) []House {
	result := make([]House, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return less(houses[i], houses[j])
	})
	return result
}

// SortByDistanceMin ...
func SortByDistanceMin(houses []House) []House {
	result := make([]House, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return houses[i].DistanceToCenter <= houses[j].DistanceToCenter
	})
	return result
}

// SortByDistanceMax ...
func SortByDistanceMax(houses []House) []House {
	result := make([]House, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return houses[i].DistanceToCenter >= houses[j].DistanceToCenter
	})
	return result
}

// SearchByPriceBetween ...
func SearchByPriceBetween(houses []House, price int64) []House {
	return MainSortFunc(houses, func(house House) bool {
		if house.Price < price {
			return false
		}
		if house.Price > price {
			return false
		}
		return true
	})
}

// working ...

// FindByDistrict ...
func FindByDistrict(houses []House, district string) []House {
	result := make([]House, 0)
	for i, m := range houses {
		if houses[i].Rayon.Name == district {
			result = append(result, m)
		}
	}
	return result
}

// FindByDistricts ...
func FindByDistricts(houses []House, districts []string) []House {
	result := make([]House, 0)
	for _, m := range houses {
		for _, d := range districts {
			if m.Rayon.Name == d {
				result = append(result, m)
			}
		}
	}
	return result
}
