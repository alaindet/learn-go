package utils

func AsBoolean(val any) bool {
	switch val.(type) {
	case bool:
		return val.(bool)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return val != 0
	case string:
		return val != ""
	case nil:
		return false
	default:
		return true
	}
}
