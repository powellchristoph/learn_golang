package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2-x1
    b := y2-y1
    return math.Sqrt(a*a + b*b)
}

func circleArea(c *Circle) float64 {
    return math.Pi * c.r*c.r
}

// Circle
type Circle struct {
    x, y, r float64
}
func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}
func (c *Circle) perimeter() float64 {
    return 2 * math.Pi * c.r
}

// Rectangle
type Rectangle struct {
    x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}
func (r *Rectangle) perimeter() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return 2 * (l+w)
}

type Shape interface {
    area() float64
    perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

// Person
type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

// Android
type Android struct {
    Person
    Model string
}

type MultiShape struct {
    shapes []Shape
}
func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}

func main() {
    fmt.Println("Circle")
    c := Circle{0, 0, 5}
    fmt.Println(circleArea(&c))
    fmt.Println(c.area())

    fmt.Println("Rectangle")
    r := Rectangle{0, 0, 10, 10}
    fmt.Println(r.area())

    fmt.Println("Total Area")
    fmt.Println(totalArea(&c, &r))

    fmt.Println("Android")
    a := new(Android)
    a.Person.Talk()
    a.Talk()

    fmt.Println("Perimeter")
    fmt.Println(c.perimeter())
    fmt.Println(r.perimeter())
}
