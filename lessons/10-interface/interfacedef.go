package main


type Greeting interface {
	SayHelloTo(person string) string
}