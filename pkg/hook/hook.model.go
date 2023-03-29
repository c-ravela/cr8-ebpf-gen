package hook

type Hook struct {
	Name     string //Name of the hook ex: tracepoint, tp, raw_tracepoint
	HookType string //Type of hook ex: sycalls, fentry, cgroup
	Class    string //Holds the type of hook ex: tracepoint, tp, raw_tracepoint
}
