package template

import (
	"cr8-gen/pkg/hook"
	"fmt"
)

func Basic(hookPrg hook.Hook) string {
	basicTemplate := `
	package main

	import (
		"fmt"
		"os"

		"github.com/cilium/ebpf/link"
	)

	//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags $BPF_CFLAGS basic index.bpf.c -- -I../../headers

	func main(){
		basicObj := basicObjects{}
		err := loadBasicObjects(&basicObj, nil)
		if err != nil{
			fmt.Println("Error Loading the ebpf Objects: ", err)
			os.Exit(1)
		}

		//Hook
		%s

		for {

		}
	}
	`
	basicTemplate = fmt.Sprintf(basicTemplate, genHook(hookPrg))
	return basicTemplate
}

func genHook(hookTemp hook.Hook) string {

	hk := ""
	if hookTemp.Class == "tracepoint" {
		hk = `
		hook, err := link.Tracepoint("%s", "%s", basicObj.Cr8, nil)
		if err != nil{
			fmt.Println("Error attaching a the ebpf program: ", err)
			os.Exit(1)
		}
		defer hook.Close()
		`
		hk = fmt.Sprintf(hk, hookTemp.HookType, hookTemp.Name)
	} else if hookTemp.Class == "raw_tracepoint" {
		hk = `
		rtpObj := link.RawTracepointOptions{
			Name: "%s",
			Program: basicObj.Cr8,
		}
		hook, err := link.AttachRawTracepoint(rtpObj)
		if err != nil{
			fmt.Println("Error attaching a the ebpf program: ", err)
			os.Exit(1)
		}
		defer hook.Close()
		`
		hk = fmt.Sprintf(hk, hookTemp.Name)
	} else if hookTemp.Class == "kprobe" {
		hk = `
		hook, err := link.Kprobe("%s", basicObj.Cr8, nil)
		if err != nil{
			fmt.Println("Error attaching a the ebpf program: ", err)
			os.Exit(1)
		}
		defer hook.Close()
		`
		hk = fmt.Sprintf(hk, hookTemp.Name)
	}

	return hk
}
