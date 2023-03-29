package main

import (
	"bufio"
	"cr8-gen/pkg/hook"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cr8-gen/pkg/gen"
)

func main() {
	hooks := []hook.Hook{}

	hook0 := hook.Hook{
		Name:     "sys_enter_execve",
		HookType: "syscalls",
		Class:    "tracepoint",
	}
	hooks = append(hooks, hook0)

	hook1 := hook.Hook{
		Name:     "sys_exit_execve",
		HookType: "syscalls",
		Class:    "tracepoint",
	}
	hooks = append(hooks, hook1)

	hook3 := hook.Hook{
		Name:  "sys_execve",
		Class: "kprobe",
	}
	hooks = append(hooks, hook3)

	hook4 := hook.Hook{
		Name:  "sys_enter",
		Class: "raw_tracepoint",
	}
	hooks = append(hooks, hook4)

	genA(hooks)
	// genall(hooks)
}

func genA(hooks []hook.Hook) {
	class := []string{
		"tracepoint",
		"raw_tracepoint",
		"kprobe",
	}

	for idx, cls := range class {
		fmt.Printf("\n%d. %s", idx+1, cls)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n\nSelect available option for hook: ")
	uscls, _ := reader.ReadString('\n')
	uscls = strings.Trim(uscls, " \n\r\t")

	subs := []string{
		"syscalls",
	}
	subcls := ""
	if uscls == "1" {
		for idx, sub := range subs {
			fmt.Printf("\n%d. %s\n", idx+1, sub)
		}

		fmt.Print("\n\nSelect available option: ")
		subcls, _ = reader.ReadString('\n')
	}

	subcls = strings.Trim(subcls, " \n\r\t")
	subn, _ := strconv.Atoi(subcls)
	switch uscls {
	case "1":
		for id, hks := range hooks {
			if subs[subn-1] == hks.HookType {
				fmt.Printf("\n%d. %s", id+1, hks.Name)
			}
		}
	case "2":
		for id, hks := range hooks {
			if class[1] == hks.Class {
				fmt.Printf("\n%d. %s", id+1, hks.Name)
			}
		}
	case "3":
		for id, hks := range hooks {
			if class[2] == hks.Class {
				fmt.Printf("\n%d. %s", id+1, hks.Name)
			}
		}
	}

	fmt.Print("\n\nSelect available option for hook: ")
	hookI, _ := reader.ReadString('\n')
	hookI = strings.Trim(hookI, " \n\r\t")

	hookId, _ := strconv.Atoi(hookI)

	gen.Generate(hooks[hookId-1])

}

func genall(hooks []hook.Hook) {
	for _, hook := range hooks {
		gen.Generate2(hook, "")
	}
}
