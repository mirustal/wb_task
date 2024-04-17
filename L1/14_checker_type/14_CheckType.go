package checkertype


func CheckType(variable interface{}) string {
	switch variable.(type) {
	case int, uint, int8, int16, int32, uint32, int64, uint64:
		return "int" 
	case float32, float64:
		return "float"
	case bool:
		return "boolean"
	case string:
		return "string"
	default:
		return "unknown"
	}
}