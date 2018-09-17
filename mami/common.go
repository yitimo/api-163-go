package mami

// Res common responce model
type Res struct {
	State bool
	Code  int32
	Msg   string
	Data  interface{}
}
