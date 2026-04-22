package main

import "log"

func storageManager(ch chan command) {
	data := make(map[string]int)
	for com := range ch {
		tempResult := result{}
		select {
		case <-com.ctx.Done():
		default:
			switch com.operation {
			case "get":
				val, present := data[com.key]
				if present {
					tempResult.intResult = val
					tempResult.ok = true
				}
				select {
				case com.res <- tempResult:
				case <-com.ctx.Done():
					log.Printf("request cancelled, discarding result for key %s", com.key)
				}

			case "post":
				_, present := data[com.key]
				if present {
					tempResult.boolResult = true
				}
				data[com.key] = com.value
				tempResult.ok = true
				select {
				case com.res <- tempResult:
				case <-com.ctx.Done():
					log.Printf("request cancelled, discarding result for key %s", com.key)
				}

			case "delete":
				_, present := data[com.key]
				if present {
					delete(data, com.key)
					tempResult.boolResult = true
					tempResult.ok = true
				}
				select {
				case com.res <- tempResult:
				case <-com.ctx.Done():
					log.Printf("request cancelled, discarding result for key %s", com.key)
				}
			case "list":
				for key := range data {
					tempResult.listResult = append(tempResult.listResult, key)
				}
				tempResult.ok = true
				select {
				case com.res <- tempResult:
				case <-com.ctx.Done():
					log.Printf("request cancelled, discarding list query result")
				}
			}
		}

	}
}
