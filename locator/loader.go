package locator

type Loader interface {
	Load(locator IPLocator) error
}

type CSVLoader struct {
	locationsFile, networksFile string
}

func (c CSVLoader) Load(loc IPLocator) error {
	// todo: read the CSV files and write into the repo.
	return nil
}
