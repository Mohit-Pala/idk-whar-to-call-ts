package models

// ps -eo pid,ppid,user,group,etime,pcpu,pmem,vsz,rss,tty,stat,cmd
type PsCommand struct {
	Pid               int     `json:"pid"`              // pid
	Ppid              int     `json:"ppid"`             // parent pid
	User              string  `json:"user"`             // user
	Group             string  `json:"group"`            // group id?
	ElapsedTime       string  `json:"elapsed_time"`     // elapsed time
	CpuUsage          float64 `json:"cpu_usage"`        // cpu uagae perecenm
	MemUsage          float64 `json:"mem_usage"`        // memory usage percent
	VirtMem           int64   `json:"virtual_mem"`      // virtual memory size
	RssWhatDoesTsMean int64   `json:"dunno_whar_ts_is"` // resident set size
	Tty               string  `json:"tty"`              // tty
	Status            string  `json:"status"`           // status
	Cmd               string  `json:"cmd"`              // command
}

