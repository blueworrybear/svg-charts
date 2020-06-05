package common

import (
	"fmt"
	"strings"
)

type Shape struct {
	attributes []string
}

func (s *Shape) AddAttr(style string) {
	s.attributes = append(s.attributes, style)
}

func (s *Shape) ClearAttr() {
	s.attributes = nil
}

func (s *Shape) Attributes() []string {
	if s.attributes == nil {
		return []string{}
	}
	return s.attributes
}

func (s *Shape) AttrString() string {
	return strings.Join(s.attributes, ";")
}

func (s *Shape) SetID(id string) {
	s.AddAttr(fmt.Sprintf(`id="%s"`, id))
}
