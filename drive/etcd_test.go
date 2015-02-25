package drive

import (
	"testing"
)

func TestShouldBeNonEmptyMakeDomainValid(t *testing.T) {
	domain := makeDomain("/services/cad/domain")

	if domain != "cad" {
		t.Error("Expected cad, got", domain)
	}
}

func TestShouldBeEmptyWhenMakeDomainInvalid(t *testing.T) {
	domain := makeDomain("")

	if domain != "" {
		t.Error("Expected empty, got", domain)
	}
}
