package main

import (
	"fmt"
	"log"
	"os"
)

/**
composition vs inheritance
*/

type User struct {
	Id       int
	Name     string
	Location string
}

//
//type Player struct {
//	Id int
//	Name string
//	Location string
//	GameId int
//}

type Player struct {
	*User
	GameId int
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Job struct {
	Command string
	*log.Logger
}

func NewPlayer(id int, name, location string, gameId int) *Player {
	return &Player{
		User: &User{
			Id:       id,
			Name:     name,
			Location: location,
		}, GameId: gameId,
	}
}

func main() {
	//error,pointer and reference
	p := &Player{
		User: &User{
			Id:       42,
			Name:     "Pa",
			Location: "LA",
		},
		GameId: 100,
	}
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v", p)
	p = &Player{
		&User{Id: 42, Name: "Matt", Location: "LA"},
		90404}
	fmt.Printf(
		"Id: %d, Name: %s, Location: %s, Game id: %d\n",
		p.Id, p.Name, p.Location, p.GameId)
	// Directly set a field defined on the Player struct
	p.Id = 11
	fmt.Printf("%+v", p)
	fmt.Println(p.Greetings())

	job := &Job{"demo", log.New(os.Stdout, "Job: ", log.Ldate)}
	job.Print("starting now...")

	p1 := NewPlayer(42, "Matt", "LA", 90404)
	fmt.Println(p1.Greetings())
}
