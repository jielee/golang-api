package model_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	log "github.com/Sirupsen/logrus"
)



func TestAPIModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Model Unit Suite")
}

var _ = BeforeSuite(func() {

})

var _ = BeforeEach(func() {
	log.Println("Before Each function ")
})

var _ = AfterSuite(func() {
	log.Println("After suite function ")
})

