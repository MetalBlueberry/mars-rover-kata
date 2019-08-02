package rover_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRover(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rover Suite")
}
