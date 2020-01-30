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
	Districts        Districts
	Owner            string
}

// City ...
type City struct {
	Name string
}

// Districts type
type Districts struct {
	ID   int8
	Name string
}

func main() {}

// SortByPrice ...
func SortByPrice(massive []House, amount int64) []House {
	result := make([]House, 0)
	for _, house := range massive {
		if house.Price <= amount {
			result = append(result, house)
		}
	}
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
		if house.Districts.Name == district {
			return true
		}
		return false
	})
}

// FindByDistricts with less(a, b)...
func FindByDistricts(houses []House, districts []string) []House {
	return FindBy(houses, func(house House) bool {
		for _, district := range districts {
			if house.Districts.Name == district {
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
