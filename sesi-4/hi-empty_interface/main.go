package main

func main() {
	//Empty Interface
	var randomValues interface{}
	_ = randomValues
	randomValues = "Jalan Cempaka"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Asep", "Saefuddin"}
	// Empty interface (Type assertion)
	var v interface{}
	v = 20
	if value, ok := v.(int); ok == true {
		v = value * 9
	}
	// Empty interface (Map & Slice with Empty Interface)
	rs := []interface{}{1, "Asep", true, 2, "Saefuddin", true}
	rm := map[string]interface{}{
		"name":   "Asep",
		"status": true,
		"Age":    22,
	}
	_, _ = rs, rm
}
