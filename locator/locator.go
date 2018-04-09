package locator

import (
	"net"
)

type IPLocator interface {
	GetLocationByAddress(string) (*Location, error)
}

type L struct {
	repo MemoryRepository
}

func NewLocator() IPLocator {
	return L{MemoryRepository{}}
}

func (l L) GetLocationByAddress(addr string) (*Location, error) {
	ipInstance := net.ParseIP(addr)
	ipInstance.DefaultMask()

	network, err := l.repo.FetchNetwork(ipInstance.DefaultMask().String())
	if err != nil {
		return nil, err
	}

	if network.Location != nil {
		return network.Location, nil
	}

	loc, err := l.repo.FetchLocation(network.LocationID)
	if err != nil {
		return nil, err
	}

	network.Location = loc
	l.repo.PersistNetwork(*network)

	return loc, nil
}

