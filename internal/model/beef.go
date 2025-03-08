package model

import (
	"sync"
)

type BeefCounter struct {
	mu   sync.RWMutex
	Beef map[string]int `json:"beef"`
}

func NewBeefCounter() *BeefCounter {
	return &BeefCounter{
		Beef: make(map[string]int),
	}
}

func (bc *BeefCounter) AddBeef(beefType string) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.Beef[beefType]++
}

func (bc *BeefCounter) GetBeefCount() map[string]int {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	result := make(map[string]int)
	for k, v := range bc.Beef {
		result[k] = v
	}
	return result
}

type BeefTypes struct {
	ValidTypes []string
	TypesMap   map[string]bool
	Sorted     []string
}

func NewBeefTypes() *BeefTypes {
	validTypes := []string{
		"t-bone", "ribeye", "sirloin", "brisket",
		"chuck", "tenderloin", "filet mignon", "short ribs",
		"ground round", "strip steak", "flank", "tri-tip",
		"rump", "corned beef", "prime rib", "ball tip",
		"top round", "bottom round", "rib", "short loin",
		"porterhouse", "flat iron", "beef ribs", "shank", "tongue",
	}

	typesMap := make(map[string]bool)
	for _, beefType := range validTypes {
		typesMap[beefType] = true
	}

	return &BeefTypes{
		ValidTypes: validTypes,
		TypesMap:   typesMap,
		Sorted:     sortBeefTypes(validTypes),
	}
}

func sortBeefTypes(beefTypes []string) []string {
	sorted := make([]string, len(beefTypes))
	copy(sorted, beefTypes)

	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if len(sorted[i]) < len(sorted[j]) {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	return sorted
}
