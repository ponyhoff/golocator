package main

import "github.com/ponyhoff/golocator/locator"

type LocatorService struct {
	loc locator.IPLocator
}

func (ls LocatorService) GetLocation(ipAddr string) (map[string]interface{}, error) {
	location, err := ls.loc.GetLocationByAddress(ipAddr)
	if err != nil {
		return nil, err
	}

	return location.ToMapInterface(), nil
}
