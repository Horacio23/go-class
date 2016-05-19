package main

import "fmt"

type Client interface {
	Draw(string) string
}

type ClientFunc func(s string) string

func (c ClientFunc) Draw(s string) string {
	return c(s)
}

type Decorator func(Client) Client

func Logging(log string) Decorator {
	// This is the decorator returned By using the func Loggin
	return func(c Client) Client {
		//This is the client that gets returned by calling the Decorator func
		// It gets executed by calling Draw on the Client, since Draw just calls the func itself
		return ClientFunc(func(s string) string {
			fmt.Println("He the log is: ", log)
			return c.Draw(s)
		})
	}
}

func main() {
	var f ClientFunc

	f = func(s string) string {
		return "This is it boy:" + s
	}

	d := Logging("log this")
	c := d(f)

	fmt.Println(c.Draw("Ok this is fine"))
}
