package main

const PILLARWIDTH = 3
const PILLARHEIGHT = 1

type Pillar struct {
	obj         BasicObject
	gapSize     int
	gapPosition int
	ph          int
	drawn       bool
}

func NewPillar(gapPosition, gapSize, x, y int) *Pillar {
	data := make2DSlice[rune](PILLARHEIGHT, PILLARWIDTH)
	return &Pillar{
		obj:         *NewBasicObject(data, x, y),
		gapSize:     gapSize,
		gapPosition: gapPosition,
		ph:          PILLARHEIGHT,
		drawn:       false,
	}
}

func (b *Pillar) GetPositionX() int {
	return b.obj.GetPositionX()
}

func (b *Pillar) GetPositionY() int {
	return b.obj.GetPositionY()
}

func (b *Pillar) SetPositionX(x int) {
	b.obj.SetPositionX(x)
}

func (b *Pillar) SetPositionY(y int) {
	b.obj.SetPositionY(y)
}

func (p *Pillar) isGapBeginning(i int) bool {
	return i == (p.gapPosition - int(p.gapSize/2))
}

func (p *Pillar) isGapEnding(i int) bool {
	return i == (p.gapPosition + int(p.gapSize/2))
}

func (p *Pillar) fillPillar(data [][]rune, r rune, i int) {
	for j := 1; j < PILLARWIDTH-1; j++ {
		data[i][j] = r
	}
}

func (p *Pillar) drawPillar(data [][]rune, x int) {
	for i := 0; i < x; i++ {
		if p.isGapBeginning(i) {
			for !p.isGapEnding(i) {
				i += 1
			}
			i -= 1
			continue
		}
		LOG.Printf("drawPillar: i = %v", i)
		data[i][0] = '|'
		data[i][PILLARWIDTH-1] = '|'
		p.fillPillar(data, '#', i)
	}
}

func (b *Pillar) Draw(f Field) {
	fh := f.Height()
	LOG.Printf("Pillar.Draw(): field height: %v, gapPosition: %v, gapSize: %v", fh, b.gapPosition, b.gapSize)
	if b.gapSize >= fh {
		return
	}
	if fh > b.ph || !b.drawn {
		data := make2DSlice[rune](fh, PILLARWIDTH)
		b.drawPillar(data, fh)
		LOG.Printf("pillar position: (%v, %v), field size: (%v, %v)", b.obj.GetPositionX(), b.obj.GetPositionY(), f.Width(), f.Height())
		b.obj = *NewBasicObject(
			data,
			b.obj.GetPositionX(),
			b.obj.GetPositionY(),
		)
		b.ph = fh
		b.drawn = true
	}
	LOG.Printf("pillar.basicObject data size: width: %v, height: %v", b.obj.data.width, b.obj.data.height)
	b.obj.Draw(f)
}
