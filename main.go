package main

import (
	"fmt"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.
	s := "gopher"
	var age int8 = 20
	fmt.Println("Hello and welcome, %s!", s)
	fmt.Println("Age %s!", age)
	for i := 1; i <= 5; i++ {
		//TIP You can try debugging your code. We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>. To start your debugging session,
		// right-click your code in the editor and select the <b>Debug</b> option.
		fmt.Println("i =", 100/i)
	}
	if age > 20 {
		fmt.Println("Age is greater than 20")
	} else {
		fmt.Println("Age is greater than 20")
	}

	result := add(5, 5)
	fmt.Println(result)
}

func add(a int, b int) int {
	return a + b
}
