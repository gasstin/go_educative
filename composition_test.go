package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Player struct {
	u      User
	GameId int
}

func GreetingsForPlayer(p Player) string {
	return fmt.Sprintf("%s", p.u.Greetings())
}
