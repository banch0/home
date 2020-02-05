package main

import (
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
	District         District
	Owner            string
}

// City ...
type City struct {
	Name string
}

// District type
type District struct {
	ID   int8
	Name string
}

// var houses = []House{
// 	{
// 		ID:               1,
// 		Name:             "Penthause somoni",
// 		Owner:            "Bob",
// 		Price:            350_000,
// 		DistanceToCenter: 400,
// 		City:             City{Name: "Khujand"},
// 		District:         District{Name: "Airport"},
// 	}, {
// 		ID:               2,
// 		Name:             "Appartment somoni",
// 		Owner:            "Tyler",
// 		Price:            300_000,
// 		DistanceToCenter: 300,
// 		City:             City{Name: "Dushanbe"},
// 		District:         District{Name: "Шохмансур"},
// 	}, {
// 		ID:               3,
// 		Name:             "Penthouse somoni",
// 		Owner:            "Marla",
// 		Price:            500_000,
// 		DistanceToCenter: 100,
// 		City:             City{Name: "Dushanbe"},
// 		District:         District{Name: "Сомони"},
// 	}, {
// 		ID:               5,
// 		Name:             "New House",
// 		Owner:            "Cornelius",
// 		Price:            450_000,
// 		DistanceToCenter: 600,
// 		City:             City{Name: "Dushanbe"},
// 		District:         District{Name: "Сино"},
// 	}, {
// 		ID:               6,
// 		Name:             "Old House",
// 		Owner:            "Bob",
// 		Price:            550_000,
// 		DistanceToCenter: 550,
// 		City:             City{Name: "Dushanbe"},
// 		District:         District{Name: "Сино"},
// 	}, {
// 		ID:               4,
// 		Name:             "House in center of Norak",
// 		Owner:            "Marla",
// 		Price:            2_200_000,
// 		DistanceToCenter: 1100,
// 		City:             City{Name: "Norak"},
// 		District:         District{Name: "Center"},
// 	},
// }

func main() {
	// d := SearchByPriceBetween(houses, 350_000, 550_000)
	// log.Println(d)
}

// SortByPrice ...
func SortByPrice(houses []House) []House {
	result := make([]House, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return houses[i].Price <= houses[j].Price
	})
	return result
}

// PrimarySortFunc ...
func PrimarySortFunc(houses []House, predicate func(item House) bool) []House {
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
	return PrimarySortFunc(houses, func(house House) bool {
		return house.Price >= price
	})
}

// SortByMinPrice ...+
func SortByMinPrice(houses []House, price int64) []House {
	return PrimarySortFunc(houses, func(house House) bool {
		return house.Price <= price
	})
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
func SearchByPriceBetween(houses []House, lower, upper int64) []House {
	return PrimarySortFunc(houses, func(house House) bool {
		if house.Price < lower {
			return false
		}
		if house.Price > upper {
			return false
		}
		return true
	})
}

// FindByDistrict with less(a, b)...
func FindByDistrict(houses []House, district string) []House {
	return FindBy(houses, func(house House) bool {
		if house.District.Name == district {
			return true
		}
		return false
	})
}

// FindByDistricts with less(a, b)...
func FindByDistricts(houses []House, districts []string) []House {
	return FindBy(houses, func(house House) bool {
		for _, district := range districts {
			if house.District.Name == district {
				return true
			}
		}
		return false
	})
}

// FindBy ...
func FindBy(houses []House, predicate func(house House) bool) []House {
	result := make([]House, 0)
	for _, house := range houses {
		if predicate(house) {
			result = append(result, house)
		}
	}
	return result
}
