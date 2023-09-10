package space

type BasicObject struct {
	field *Field
	loc Location
}

func NewBasicObject(x, y Dimension2D) *BasicObject {
	return &BasicObject{
		field: NewField(1, 1),
	}
}

func (o *BasicObject) GetMatterAt(x, y Dimension2D) string {
	return o.field.GetMatterAt(x, y)
}

func (o *BasicObject) GetRowCount() Dimension2D {
	return o.field.GetRowCount()
}

func (o *BasicObject) GetColumnCount() Dimension2D {
	return o.field.GetColumnCount()
}

func (o *BasicObject) getX() Dimension2D {
	return o.loc.getX()
}

func (o *BasicObject) getY() Dimension2D {
	return o.loc.getY()
}

func (o *BasicObject) setX(x Dimension2D) {
	o.loc.setX(x)
}

func (o *BasicObject) setY(y Dimension2D) {
	o.loc.setY(y)
}

