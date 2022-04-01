package quant

const (
	_   = iota
	KiB = 1 << (10 * iota)
	MiB
	GiB
	TiB
	PiB
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
)
