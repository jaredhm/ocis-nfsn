package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSearch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Userlog service Suite")
}
