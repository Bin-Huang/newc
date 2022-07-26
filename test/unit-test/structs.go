package unittest

// StructRef ...
//go:generate go run ../../../newc
type StructRef struct {
	Debug bool
}

// StructValue ...
//go:generate go run ../../../newc --value
type StructValue struct {
	Debug bool
}

// StructWithInit ...
//go:generate go run ../../../newc --init
type StructWithInit struct {
	Debug bool
}

func (s *StructWithInit) init() {
	s.Debug = true
}

// StructValueWithInit ...
//go:generate go run ../../../newc --value --init
type StructValueWithInit struct {
	Debug bool
}

func (s *StructValueWithInit) init() {
	s.Debug = true
}

// Skipeed ...
//go:generate go run ../../../newc --value --init
type Skipeed struct {
	Msg    string `bson:"msg" json:"msg"`
	Status int    `bson:"status" json:"status" newc:"-"`
}

func (e *Skipeed) init() {
	e.Status = 403
}
