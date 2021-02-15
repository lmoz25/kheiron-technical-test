package prefix_calculator_test

import (
	"os"
	"testing"

	"github.com/houqp/gtest"
)

type Tests struct{}

type TestData struct {
}

func (s *Tests) Setup(t *testing.T) {
	os.Setenv("BANTER_BUS_CONFIG_PATH", "config.test.yml")
	// config, err := core.NewConfig()
	// if err != nil {
	// 	fmt.Printf("Failed to load config %s", err)
	// }
	// logger := core.SetupLogger(ioutil.Discard)
	// core.UpdateLogLevel(logger, "DEBUG")

}

func (s *Tests) Teardown(t *testing.T) {}

func (s *Tests) BeforeEach(t *testing.T) {}

func (s *Tests) AfterEach(t *testing.T) {}

func TestSampleTests(t *testing.T) {
	gtest.RunSubTests(t, &Tests{})
}

func InsertData(db core.Repository, dataFilePath string, collection string) {}
