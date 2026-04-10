package main

type result struct {
	intResult  int
	boolResult bool
	listResult []string
	ok         bool
}

type command struct {
	key string
	value int
	operation string
	res chan result
}

func StorageManager(ch chan command){
	switch ch {
		case 
	}
}
