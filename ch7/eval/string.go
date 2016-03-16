package eval

import (
	"strconv"
	"fmt"
)

func (v Var) String() string {
	return fmt.Sprintf("[%s]", string(v))
}

func (l literal) String() string {
	return fmt.Sprintf("[%s]", string(strconv.FormatFloat(float64(l), 'f', 4, 64)))
}

func (u unary) String() string {
	return fmt.Sprintf("[%s%s]", string(u.op), u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("[%s %s %s]", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
	o := ""
	o += fmt.Sprintf("%s(", c.fn)
	for _, arg := range c.args {
		o += arg.String()
	}
	o += ")"
	return o
}
