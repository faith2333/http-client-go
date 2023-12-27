package http_client_go

import "strings"

const (
	MethodGET    Method = "GET"
	MethodPOST   Method = "POST"
	MethodDELETE Method = "DELETE"
	MethodPUT    Method = "PUT"
	MethodOPTION Method = "OPTION"
)

var AllMethod = []Method{
	MethodGET,
	MethodPOST,
	MethodDELETE,
	MethodPUT,
	MethodOPTION,
}

type Method string

func (m Method) String() string {
	return string(m)
}

func (m Method) Upper() Method {
	return Method(strings.ToUpper(m.String()))
}

func (m Method) IsGET() bool {
	return m.Upper() == MethodGET
}

func (m Method) IsPOST() bool {
	return m.Upper() == MethodPOST
}

func (m Method) IsDELETE() bool {
	return m.Upper() == MethodDELETE
}

func (m Method) IsPUT() bool {
	return m.Upper() == MethodPUT
}

func (m Method) IsOPTION() bool {
	return m.Upper() == MethodOPTION
}

func (m Method) WasSupported() bool {
	m = m.Upper()
	for _, mInner := range AllMethod {
		if m == mInner {
			return true
		}
	}

	return false
}
