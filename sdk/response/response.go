package response

func Success(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}

func FailOrError(httpCode int, msg string, data map[string]interface{}) map[string]interface{} {
	switch httpCode / 100 {
	case 4: //FAIL 4xx
		return map[string]interface{}{
			"status":  "fail",
			"message": msg,
			"data":    data,
		}
	case 5: //ERROR 5xx
		return map[string]interface{}{
			"status":  "error",
			"message": msg,
		}
	default:
		return map[string]interface{}{
			"status":  "error",
			"message": "INTERNAL SERVER ERROR",
		}
	}
}
