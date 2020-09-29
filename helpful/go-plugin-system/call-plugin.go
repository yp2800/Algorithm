package main

import (
	"fmt"
	"path/filepath"
	"plugin"
)

func main() {
	// 这里可以写 *.so 然后 for _, filename := range plugins 遍历
	//plugins, err := filepath.Glob("*.so")
	plugins, err := filepath.Glob("main.so")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loading plugin %s\n", plugins[0])
	// Open – Loads the plugin
	p, err := plugin.Open(plugins[0])
	if err != nil {
		panic(err)
	}
	// Lookup – Searches for a symbol name in the plugin
	add, err := p.Lookup("Add")
	if err != nil {
		panic(err)
	}
	// Checks the function signature
	addFunc, ok := add.(func(i, j int) int)
	if !ok {
		panic("Plugin has no 'Add(int, int)int' function")
	}
	// Uses the function to return results
	fmt.Printf("\n10 + 12 = %d\n", addFunc(10,12))

	sub, err := p.Lookup("Sub")
	if err != nil {
		panic(err)
	}
	subFunc, ok := sub.(func(i, j int) int)
	if !ok {
		panic("Plugin has no 'Sub(int, int)int' function")
	}
	fmt.Printf("\n22 - 8 = %d\n", subFunc(22,8))

}
