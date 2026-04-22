package main

import "context"

type result struct {
	intResult  int
	boolResult bool
	listResult []string
	ok         bool
}

type command struct {
	ctx       context.Context
	key       string
	value     int
	operation string
	res       chan result
}

type Server struct {
	ch chan command
}

type createKeyValue struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}
