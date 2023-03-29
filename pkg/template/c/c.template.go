package template

import (
	"cr8-gen/pkg/hook"
	"fmt"
)

func Basic(hookPrg hook.Hook, outstring string) string {
	basicTemplate := `
	//go:build ignore
	#include "vmlinux.h"
	#include "bpf_helpers.h"

	SEC("%s")
	int cr8(){
		bpf_printk("%s");
		return 0;
	};

	char _license[] SEC("license") = "Dual MIT/GPL";
	`
	basicTemplate = fmt.Sprintf(basicTemplate, hookGen(hookPrg), outstring)
	return basicTemplate
}

func hookGen(hooktmp hook.Hook) string {
	hk := ""
	if hooktmp.Class == "tracepoint" {
		hk = hooktmp.Class + "/" + hooktmp.HookType + "/" + hooktmp.Name
	} else {
		hk = hooktmp.Class + "/" + hooktmp.Name
	}
	return hk
}
