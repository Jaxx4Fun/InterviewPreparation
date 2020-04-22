package effective

func aboutConstant() {
	type ByteSize float64
	const (
		// 通过赋予空白标识符来忽略第一个值
		_           = iota // ignore first value by assigning to blank identifier
		_           = iota // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		Y
	)
}
