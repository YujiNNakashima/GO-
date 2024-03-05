package structembedding

type Position struct {
	posX float64
	posY float64
}

func (p *Position) Move(posX, posY float64) {
	p.posX += posX
	p.posY += posY
}

func (p *Position) Teleport(posX, posY float64) {
	p.posX = posX
	p.posY = posY
}

type Ball struct {
	name string
	*Position
}

func NewBall(name string, initialX, initialY float64) *Ball {
	return &Ball{
		name: name,
		Position: &Position{
			posX: initialX,
			posY: initialY,
		},
	}
}

type Square struct {
	name string
	*Position
}

func NewSquare(name string, initialX, initialY float64) *Square {
	return &Square{
		name: name,
		Position: &Position{
			posX: initialX,
			posY: initialY,
		},
	}
}
