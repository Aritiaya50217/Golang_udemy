package report

var getReport = func(id string) string {
	// connect to database
	return "Data"
}

func generateReport(id string) string {
	data := getReport(id)
	return data + " : verified."
}
