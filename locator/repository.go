package locator

import (
	"errors"
	"github.com/ponyhoff/golocator/btreestorage"
	"sync"
)

type LocatorRepository interface {
	FetchLocation(locationID string) (*Location, error)
	PersistLocation(Location) bool

	FetchNetwork(mask string) (*Network, error)
	PersistNetwork(Network) bool
}

type MemoryRepository struct {
	networks  map[string]Network
	locations *btreestorage.BTStorage

	cacheMutex sync.Mutex
}

func (m MemoryRepository) FetchLocation(locationID string) (*Location, error) {
	data, found := m.locations.Query(locationID)

	if !found {
		return nil, errors.New("entry not found")
	}

	l, ok := data.(Location)
	if !ok {
		return nil, errors.New("failed to type assert to location")
	}

	return &l, nil
}

func (m MemoryRepository) PersistLocation(l Location) bool {
	return m.locations.Put(l.LocationID, l)
}

func (m MemoryRepository) FetchNetwork(mask string) (*Network, error) {
	n, found := m.networks[mask]

	if !found {
		return nil, errors.New("entry not found")
	}

	return &n, nil
}

func (m *MemoryRepository) PersistNetwork(n Network) bool {

	m.cacheMutex.Lock()
	defer m.cacheMutex.Unlock()

	if m.networks == nil {
		m.networks = make(map[string]Network)
	}

	m.networks[n.Mask] = n
	return true
}
