package img

type Position struct {
	X int32
	Y int32
}

type Asset struct {
	Width  int32
	Height int32
	Data   []byte
}

type ImageAsset struct {
	Asset  *Asset
	Offset []Position
	Width  int32
	Height int32
}
