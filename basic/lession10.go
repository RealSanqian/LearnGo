package main

import "time"

type PPerson struct {
	age  int
	name string
}

func (p PPerson) GetName() string {
	return p.name
}

func GetTime() time.Time {
	return time.Now()
}

func GetTime2() (time.Time, error) {
	return time.Now(), nil
}
