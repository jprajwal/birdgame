package space

import "fmt"


type Cell struct {
	data interface{}
}

func NewCell(data interface{}) *Cell {
	return &Cell{data: data}
}

func (c *Cell) String() string {
	switch c.data.(type) {

	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
		return fmt.Sprintf("%v", c.data)

	case fmt.Stringer:
		v, _ := c.data.(fmt.Stringer)
		return v.String()
		
	default:
		return ""	
	}
}
