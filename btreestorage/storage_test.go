package btreestorage

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

const (
	testDumpFileName string = "test_dump.dat"
)

type StorageTestSuite struct {
	suite.Suite

	s BTStorage
}

func (st *StorageTestSuite) SetupTest() {
	st.s = *New()
}

func (st *StorageTestSuite) TestStorageFromGOBS() {

	st.s.Put("test", "value")

	n, err := st.s.DumpToFile(testDumpFileName)
	st.NoError(err)
	st.True(n > 0)

	file, err := os.Open(testDumpFileName)
	st.NoError(err)

	newStorage, err := FromGOBs(file)
	st.NoError(err)
	st.NotEmpty(newStorage)

	res, found := newStorage.Query("test")
	st.True(found)
	st.NotEmpty(res)

	d, ok := res.(string)
	st.True(ok)
	st.Equal("value", d)

	// clear the test file
	st.removeTestFile()
}

func (st *StorageTestSuite) TestCanPutAndQuery() {
	st.s.Put("foo", "bar")
	st.s.Put("key", "value")

	d, found := st.s.Query("foo")
	st.True(found)
	st.NotEmpty(d)

	sd, ok := d.(string)
	st.True(ok)
	st.Equal("bar", sd)
}

func (st *StorageTestSuite) removeTestFile() {
	err := os.Remove(testDumpFileName)
	st.NoError(err)
}

func TestMemoryStorageSuite(t *testing.T) {
	suite.Run(t, new(StorageTestSuite))
}
