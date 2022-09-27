package main

type chessTopping struct{}

func AddChessTopping() IPizza {
	return &chessTopping{}
}

func (c *chessTopping) GetPrice(add func() int) func() int {
	return func() int {
		return 5 + add()
	}
}
