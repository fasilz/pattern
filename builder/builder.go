package builder

// Tester is an interface for conducting a test
type Tester interface {
	Collect() Tester
	Extract() Tester
	Analyze() Tester
	Result() Tester
}

// Runner is an interfac for building/running a process
type Runner interface {
	Run()
}

// TestRunner implements the Runner interface
type TestRunner struct {
	t Tester
}

// value semantic
func New(t Tester) TestRunner {
	return TestRunner{}
}

// Run test in sequence
func (t TestRunner) Run() {
	t.Collect().Extract().Analyze().Result()
}
