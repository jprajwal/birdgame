package space

type Space struct {
	field   Field2D
	objects []Object2D
}

func NewSpace(field Field2D, x, y Dimension2D) *Space {
	return &Space{
		field:   field,
		objects: make([]Object2D, 0),
	}
}

func (s *Space) GetMatterAt(x Dimension2D, y Dimension2D) Cell2D {
	if x >= s.GetRowCount() || y >= s.GetColumnCount() {
		return NewCell("")
	}
	return s.field.GetMatterAt(x, y)
}

func (s *Space) SetMatterAt(x Dimension2D, y Dimension2D, c Cell2D) {
	s.field.SetMatterAt(x, y, c)
}

func (s *Space) GetRowCount() Dimension2D {
	return s.field.GetRowCount()
}

func (s *Space) GetColumnCount() Dimension2D {
	return s.field.GetColumnCount()
}

func (s *Space) Clear() {
	s.field.Clear()
}

func (s *Space) IncludeObject2D(object Object2D) {
	s.objects = append(s.objects, object)
}

func (s *Space) IterateObjects2D() Object2DIterator {
	return nil
}

func (s *Space) RenderObjects2D() {
	iter := s.IterateObjects2D()

	for object := iter.Next(); object != nil; object = iter.Next() {
		CopyFieldToField(
			object,  // Location2D
			object,  // Field2D
			s.field, // Field2D
		)
	}
}

type ObjectIterator struct {
	space *Space
	i     int
}

func (o *ObjectIterator) Next() Object2D {
	o.i = o.i + 1
	if o.i >= len(o.space.objects) {
		return nil
	}
	return o.space.objects[o.i]
}
