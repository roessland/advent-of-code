package day02

const AnimationType = "array"

type AnimationHooks struct {
	SetLength     func(length int)
	IncreaseColor func(index int)
	SetColor      func(index int, color int)
	SetPalette    func(hexPalette []string)
}

func orNoopHooks(h *AnimationHooks) *AnimationHooks {
	if h == nil {
		return &AnimationHooks{
			SetLength:     func(int) {},
			IncreaseColor: func(int) {},
			SetColor:      func(int, int) {},
			SetPalette:    func([]string) {},
		}
	}
	return h
}
