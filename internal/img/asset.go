package img

type Position struct {
	X int16
	Y int16
}

type Asset struct {
	Width  int32
	Height int32
	Data   []byte
}

type ImageAsset struct {
	Offset []Position
	Width  int32
	Height int32
}
