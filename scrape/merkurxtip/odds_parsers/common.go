package odds_parsers

func checkIfMapHasFields(m map[string]interface{}, fields []string) bool {
	for _, field := range fields {
		if _, ok := m[field]; !ok {
			return false
		}
	}
	return true
}

// check if map interface is nil for each field
func checkIfMapHasInterfaces(m map[string]interface{}, fields []string) bool {
	for _, field := range fields {
		if m[field] == nil {
			return false
		}
	}
	return true
}
