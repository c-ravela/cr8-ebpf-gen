package gen

import (
	"cr8-gen/pkg/hook"
	ctmp "cr8-gen/pkg/template/c"
	gtmp "cr8-gen/pkg/template/go"
	"fmt"
	"log"
	"os"
)

func Generate(prog_hook hook.Hook) {

	kernel_prog := ctmp.Basic(prog_hook, "Ebpf Generated!"+prog_hook.Name)
	user_prog := gtmp.Basic(prog_hook)

	if err := os.WriteFile("C:/Users/CharanRavela/Projects/cr8-gen/output/index.bpf.c", []byte(kernel_prog), 0666); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Generated kernel space file to output directory")
	}

	if err := os.WriteFile("C:/Users/CharanRavela/Projects/cr8-gen/output/main.go", []byte(user_prog), 0666); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Generated user space file to output directory")
	}
}

func Generate2(prog_hook hook.Hook, file_path string) {
	kernel_prog := ctmp.Basic(prog_hook, "Ebpf Generated! "+prog_hook.Name)
	user_prog := gtmp.Basic(prog_hook)

	if err := os.WriteFile(file_path+"/index.bpf.c", []byte(kernel_prog), 0666); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Generated kernel space file at path", file_path)
	}

	if err := os.WriteFile(file_path+"/main.go", []byte(user_prog), 0666); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Generated user space file at path", file_path)
	}
}
