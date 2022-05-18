package method

type Result struct {
	Code int
	Res  int
}

// Add 加法
func Add(a, b int) Result {
	res := a + b
	return Result{Code: 200, Res: res}
}
