package interfaces

import "fmt"

type Vehicle interface {
	SpeedUp(forcePoint float32)
	SlowDown(forcePoint float32)
	Print()
}

type Car struct {
	currentSpeed float32
}

type Bicycle struct {
	currentSpeed float32
}

func TestInterfaces() {
	var car Car = Car{currentSpeed: 0}
	car.Print()
	car.SpeedUp(5)
	car.Print()
	car.SpeedUp(12)
	car.Print()
	car.SlowDown(10)
	car.Print()
	car.SlowDown(3)
	car.Print()

	var bicycle Bicycle = Bicycle{currentSpeed: 0}
	bicycle.Print()
	bicycle.SpeedUp(5)
	bicycle.Print()
	bicycle.SpeedUp(12)
	bicycle.Print()
	bicycle.SlowDown(10)
	bicycle.Print()
	bicycle.SlowDown(3)
	bicycle.Print()

}

func (car *Car) SpeedUp(forcePoint float32) {
	fmt.Printf("Accelerating..  = [%v] forcePoint\n", forcePoint)
	car.currentSpeed = car.currentSpeed + forcePoint*5
}

func (car *Car) SlowDown(forcePoint float32) {
	fmt.Printf("Decelerating..  = [%v] forcePoint\n", forcePoint)
	if d := car.currentSpeed - forcePoint*2.3; d <= 0 {
		car.currentSpeed = 0
	} else {
		car.currentSpeed = car.currentSpeed - forcePoint*2.3

	}
}

func (car *Car) Print() {
	fmt.Printf("Car currentSpeed = [%v]\n", car.currentSpeed)
}

func (b *Bicycle) SpeedUp(forcePoint float32) {
	fmt.Printf("Accelerating..  = [%v] forcePoint\n", forcePoint)
	b.currentSpeed = b.currentSpeed + forcePoint
}

func (b *Bicycle) SlowDown(forcePoint float32) {
	fmt.Printf("Decelerating..  = [%v] forcePoint\n", forcePoint)

	if d := b.currentSpeed - forcePoint*3.8; d <= 0 {
		b.currentSpeed = 0
	} else {
		b.currentSpeed = b.currentSpeed - forcePoint*2.3

	}
}

func (b *Bicycle) Print() {
	fmt.Printf("Bicycle currentSpeed = [%v]\n", b.currentSpeed)
}
