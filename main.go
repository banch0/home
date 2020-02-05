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

func main() {}

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
func SearchByPriceBetween(houses []House, price int64) []House {
	return PrimarySortFunc(houses, func(house House) bool {
		if house.Price < price {
			return false
		}
		if house.Price > price {
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
