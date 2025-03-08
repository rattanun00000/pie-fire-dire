package service

import (
	"pie-fire-dire/internal/model"
	"pie-fire-dire/pkg/util"
	"regexp"
	"sync"
	"time"
)

type BeefService struct {
	beefTypes     *model.BeefTypes
	cachedResult  map[string]int
	cacheTime     time.Time
	cacheMutex    sync.RWMutex
	cacheDuration time.Duration
}

func NewBeefService() *BeefService {
	return &BeefService{
		beefTypes:     model.NewBeefTypes(),
		cacheDuration: 5 * time.Minute,
	}
}

func (s *BeefService) CountBeefTypes(text string) map[string]int {
	counter := model.NewBeefCounter()

	text = util.CleanText(text)

	text = " " + text + " "

	for _, beefType := range s.beefTypes.Sorted {
		pattern := regexp.MustCompile(`\b` + beefType + `\b`)

		matches := pattern.FindAllStringIndex(text, -1)

		for range matches {
			counter.AddBeef(beefType)
		}
	}

	return counter.GetBeefCount()
}

func (s *BeefService) GetCachedResult() (map[string]int, bool) {
	s.cacheMutex.RLock()
	defer s.cacheMutex.RUnlock()

	if s.cachedResult != nil && time.Since(s.cacheTime) < s.cacheDuration {
		return s.cachedResult, true
	}

	return nil, false
}

func (s *BeefService) SetCachedResult(result map[string]int) {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()

	s.cachedResult = result
	s.cacheTime = time.Now()
}
