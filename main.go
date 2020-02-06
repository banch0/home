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

// SortByMaxPrice ...
func SortByMaxPrice(houses []House) []House {
	return SortBy(houses, func(a, b House) bool {
		return a.Price >= b.Price
	})
}

// SortByMinPrice ...
func SortByMinPrice(houses []House) []House {
	return SortBy(houses, func(a, b House) bool {
		return a.Price <= b.Price
	})
}

// SortBy ...
func SortBy(houses []House, less func(a, b House) bool) []House {
	result := make([]House, len(houses))
	copy(result, houses)

	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

// SortByDistanceMin ...
func SortByDistanceMin(houses []House) []House {
	return SortBy(houses, func(a, b House) bool {
		return a.DistanceToCenter <= b.DistanceToCenter
	})
}

// SortByDistanceMax ...
func SortByDistanceMax(houses []House) []House {
	return SortBy(houses, func(a, b House) bool {
		return a.DistanceToCenter >= b.DistanceToCenter
	})
}

// SearchByPriceBetween ...
func SearchByPriceBetween(houses []House, lower, upper int64) []House {
	return FindBy(houses, func(house House) bool {
		if house.Price < lower {
			return false
		}
		if house.Price > upper {
			return false
		}
		return true
	})
}

// FindByDistrict ...
func FindByDistrict(houses []House, district string) []House {
	return FindBy(houses, func(house House) bool {
		if house.District.Name == district {
			return true
		}
		return false
	})
}

// FindByDistricts ...
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
