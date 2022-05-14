package report

import (
	"fmt"
	"strings"
	"testing"
)

func TestBasicReport(t *testing.T) {
	getReport = func(id string) string {
		return "Data"
	}
	report := generateReport("12DB")
	if !strings.Contains(report, " :verify") {
		t.Error("failed")
	}
	fmt.Println("got report :", report)
}
