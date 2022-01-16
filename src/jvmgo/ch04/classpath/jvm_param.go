package classpath

type JvmParam struct {
	Xss int64
}

func (e *JvmParam) ParseParam(Xss int64) {
	e.Xss = Xss
}
