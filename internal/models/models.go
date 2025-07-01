package models

type ABTestSpec struct {
	ServiceName   string
	RouteName     string
	StructName    string
	Fields        []Field
	Modifications []Modification
}

type Modification struct {
	Name  string // имя поля
	Type  string // тип модификации: multiply / set / custom
	Value string // значение для умножения или присвоения
	Code  string // кастомный код
}

type Field struct {
	Name string
	Type string
	Tag  string
}
