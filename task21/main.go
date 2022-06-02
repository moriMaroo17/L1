package main

import "fmt"

// ThreePinsPlug interface
type ThreePinsPlug interface {
	InsertThreePins()
}

// ThreePinsSocket interface, which working only with TreePinsPlug
type ThreePinsSocket struct {
}

// Insert function, which gets ThreePinsPlug
func (u *ThreePinsSocket) Insert(plug ThreePinsPlug) {
	plug.InsertThreePins()
}

// ChargingCable which doesn't implements ThreePinsPlug interface
type ChargingCable struct {
}

// InsertInto - function for insetring cable
func (c *ChargingCable) InsertInto() {
	fmt.Println("Cable has been inserted")
}

// ThreePinsAdapter is interface for using standart cable with ThreePinsSocket
type ThreePinsAdapter struct {
	*ChargingCable
}

// InsertThreePins implements interface ThreePinsPlug and wrap InsertInto
func (a *ThreePinsAdapter) InsertThreePins() {
	a.InsertInto()
}

func main() {
	socket := ThreePinsSocket{}
	cable := ChargingCable{}
	// Doesn'w work!
	// socket.Insert(cable)
	adapter := &ThreePinsAdapter{&cable}
	socket.Insert(adapter)
}
