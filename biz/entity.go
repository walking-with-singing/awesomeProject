package biz

type Point struct{
	X,Y int
}
type Circular struct{
	point *Point
	P Point
	R int
}

func InitCircular() Circular{
	c:=Circular{}
	c.point=&Point{2,3}
	c.P=Point{4,5}
	c.R=1
	return c
}
