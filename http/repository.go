package main

type repository interface {
	AllKeys() map[string]int
}
