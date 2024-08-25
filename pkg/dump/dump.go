package dump

type Dumper interface {
	Dump(data []byte) (dumpedData []byte, err error)
}
