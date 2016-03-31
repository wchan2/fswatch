package fswatch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFswatch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fswatch Suite")
}
