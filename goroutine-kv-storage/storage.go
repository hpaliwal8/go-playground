package main

func storageManager(ch chan command) {
	data := make(map[string]int)
	for com := range ch {
		tempResult := result{}
		switch com.operation {
		case "get":
			val, present := data[com.key]
			if present {
				tempResult.intResult = val
				tempResult.ok = true
			}
			com.res <- tempResult
		case "post":
			_, present := data[com.key]
			if present {
				tempResult.boolResult = true
			}
			data[com.key] = com.value
			tempResult.ok = true
			com.res <- tempResult
		case "delete":
			_, present := data[com.key]
			if present {
				delete(data, com.key)
				tempResult.boolResult = true
				tempResult.ok = true
			}
			com.res <- tempResult
		case "list":
			for key := range data {
				tempResult.listResult = append(tempResult.listResult, key)
			}
			tempResult.ok = true
			com.res <- tempResult
		}

	}
}
