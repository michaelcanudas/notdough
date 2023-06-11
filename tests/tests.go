package tests

import "fmt"

var registeredTests = make(map[string]func())

func RegisterTest(name string, testFunc func()) {
	registeredTests[name] = testFunc
}

func InvokeTest(name string) bool {
	test, ok := registeredTests[name]

	if !ok {
		panic("test not found!")
	}

	var passed bool
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			passed = false
		}
	}()

	test()
	passed = true
	if passed {
		println("test passed!")
	} else {
		println("test failed!")
	}
	return passed
}

func InvokeAllTests() {
	results := make(map[string]bool)
	for name, _ := range registeredTests {
		results[name] = InvokeTest(name)
	}

	fmt.Println("======================================")

	passedCount := 0
	for name, passed := range results {
		var result string

		if passed {
			result = "passed"
			passedCount++
		} else {
			result = "failed"
		}

		fmt.Printf("Test '%v' %v!\n", name, result)
	}

	fmt.Printf("%v tests passed!\n", passedCount)
	fmt.Printf("%v tests failed!\n", len(results)-passedCount)

}
