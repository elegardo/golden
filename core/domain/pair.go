package domain

import "reflect"

type Pair struct {
	Value1 any
	Value2 any
}

func (p *Pair) GetInt1() int {
	return any(p.Value1).(int)
}

func (p *Pair) GetInt2() int {
	return any(p.Value2).(int)
}

func (p *Pair) GetString1() string {
	return any(p.Value1).(string)
}

func (p *Pair) GetString2() string {
	return any(p.Value2).(string)
}

func (p *Pair) GetBool1() bool {
	return any(p.Value1).(bool)
}

func (p *Pair) GetBool2() bool {
	return any(p.Value2).(bool)
}

func (p *Pair) AreString() bool {
	return reflect.TypeOf(p.Value1).String() == "string" || reflect.TypeOf(p.Value2).String() == "string"
}

func (p *Pair) AreSameType() bool {
	return reflect.TypeOf(p.Value1) == reflect.TypeOf(p.Value2)
}

func NewPair(value1, value2 any) *Pair {
	return &Pair{
		Value1: value1,
		Value2: value2,
	}
}