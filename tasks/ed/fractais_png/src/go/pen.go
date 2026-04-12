package main

import (
	"math"
	"github.com/fogleman/gg"
)

type Pen struct {
	x, y    float64
	angle   float64
	penDown bool
	dc      *gg.Context
	stack []State
}

func NewPen(width, height int) *Pen {
	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(2)
	return &Pen{
		x: float64(width) / 2, y: float64(height) / 2,
		angle: 0, penDown: true, dc: dc,
	}
}

func (t *Pen) Walk(dist float64) {
	newX := t.x + dist*math.Cos(t.angle*math.Pi/180)
	newY := t.y - dist*math.Sin(t.angle*math.Pi/180)
	if t.penDown {
		t.dc.DrawLine(t.x, t.y, newX, newY)
		t.dc.Stroke()
	}
	t.x, t.y = newX, newY
}

func (t *Pen) Left(deg float64)  { t.angle += deg }
func (t *Pen) Right(deg float64) { t.angle -= deg }

func (t *Pen) Up()   { t.penDown = false }
func (t *Pen) Down() { t.penDown = true }

func (t *Pen) Goto(x, y float64) {
	if t.penDown {
		t.dc.DrawLine(t.x, t.y, x, y)
		t.dc.Stroke()
	}
	t.x, t.y = x, y
}

func (t *Pen) SetPosition(x, y float64) {
	t.x = x
	t.y = y
}

func (t *Pen) SetHeading(angle float64) {
	t.angle = angle
}

func (t *Pen) DrawCircle(radius float64) {
	if t.penDown {
		t.dc.DrawCircle(t.x, t.y, radius)
		t.dc.Stroke()
	}
}

func (t *Pen) DrawRect(w, h float64) {
	if t.penDown {
		t.dc.DrawRectangle(t.x, t.y, w, h)
		t.dc.Stroke()
	}
}

func (t *Pen) FillCircle(radius float64) {
	if t.penDown {
		t.dc.DrawCircle(t.x, t.y, radius)
		t.dc.Fill()
	}
}

func (t *Pen) FillSquare(w, h float64) {
	t.dc.DrawRectangle(t.x, t.y, w, h)
	t.dc.Fill()
}

func (t *Pen) SetRGB(r, g, b float64) {
	t.dc.SetRGB(r/255, g/255, b/255)
}

func (t *Pen) SetLineWidth(w float64) {
	t.dc.SetLineWidth(w)
}

func (t *Pen) SavePNG(path string) {
	t.dc.SavePNG(path)
}
  
func drawTree(p *Pen, length float64){
	if length < 5{
		return
	}
	p.Walk(length)
	p.Left(30)
	drawTree(p, length*0.7)
	p.Right(60)
	drawTree(p, length*0.7)
	p.Left(30)
	p.Up()
	p.Walk(-length)
	p.Down()
}

type State struct{
	x, y, angle float64
}

func (t *Pen) Push(){
	t.stack = append(t.stack, State{t.x, t.y, t.angle})
}

func (t *Pen) Pop(){
	last := t.stack[len(t.stack)-1]
	t.x, t.y, t.angle = last.x, last.y, last.angle
	t.stack = t.stack[:len(t.stack)-1]
}

func DrawCircleFractal(p *Pen, radius float64, depth int){
	if depth <= 0 || radius < 2{
		return
	}

	p.DrawCircle(radius)

	angles := []float64{0, 90, 180, 270}

	for _, a := range angles{
		p.Push()
		p.SetHeading(a)
		p.Up()
		p.Walk(radius)
		p.Down()

		DrawCircleFractal(p, radius *0.5, depth-1)

		p.Pop()
	}
}

func drawIceFractal(p *Pen, length float64, depth int){
	if depth <= 0{
		return
	}
	p.Walk(length)

	p.Push()
	p.Left(60)
	drawIceFractal(p, length*0.5, depth-1)
	p.Pop()
}

func drawSnowflake(p *Pen, size float64, depth int){
	p.Up()
	p.Down()

	for i := 0; i < 6; i++{
		p.Push()
		p.SetHeading(float64(i * 60))
		drawIceFractal(p, size, depth)
		p.Pop()
	}
}
func drawScareTreeFractal(p *Pen, size float64, depth int){
	if depth <= 0 || size < 2{
		return
	}
	p.FillSquare(size, size)
	nextSize := size * 0.7
	offset := (nextSize - size)/2
    
	p.Push()
	p.SetPosition(p.x + offset, p.y + size)
	drawIceFractal(p, nextSize, depth-1)
	p.Pop()
}

