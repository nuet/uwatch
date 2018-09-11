package falcon

type Strategy struct {
	Metric		string
	Tags		string
	MaxStep		int
	Priority 	int
	Note    	string
	Func 		string
	Op      	string
	Value   	string
}

func GetBaseTMP() []Strategy{
	return []Strategy{
		{"df.bytes.free.percent", "", 0, 3, "磁盘剩余20%告警", "all(#3)", "<", "20"},
		{"df.bytes.free.percent", "", 0, 0, "磁盘剩余10%告警", "all(#3)", "<", "10"},
		{"df.bytes.free.percent", "fstype=ext4,mount=/boot", 3, 0, "/boot分区空间不足10%，严重", "all(#3)", "<", "10"},
		{"df.bytes.free.percent", "fstype=ext4,mount=/boot", 3, 3, "/boot分区空间不足20%，警告", "all(#3)", "<", "20"},
		{"df.bytes.free.percent", "fstype=ext4,mount=/data", 3, 0, "/data分区空间不足10%，严重", "all(#3)", "<", "10"},
		{"df.bytes.free.percent", "fstype=ext4,mount=/data", 3, 3, "/data分区空间不足20%，警告", "all(#3)", "<", "20"},
		{"df.bytes.free.percent", "fstype=ext4,mount=/", 3, 0, "/分区空间不足10%，严重", "all(#3)", "<", "10"},
		{"df.bytes.free.percent", "fstype=ext4,mount=/", 3, 3, "/分区空间不足20%，警告", "all(#3)", "<", "20"},
		{"df.inodes.free.percent", "fstype=ext4,mount=/boot", 3, 0, "/boot分区inodes不足10%，严重", "all(#3)", "<", "10"},
		{"df.inodes.free.percent", "fstype=ext4,mount=/boot", 3, 0, "/boot分区inodes不足20%，警告", "all(#3)", "<", "20"},
		{"df.inodes.free.percent", "fstype=ext4,mount=/data", 3, 0, "/data分区inodes不足10%，严重", "all(#3)", "<", "10"},
		{"df.inodes.free.percent", "fstype=ext4,mount=/data", 3, 3, "/data分区inodes不足20%，警告", "all(#3)", "<", "20"},
		{"df.inodes.free.percent", "fstype=ext4,mount=/", 3, 0, "/分区inodes不足10%，严重", "all(#3)", "<", "10"},
		{"df.inodes.free.percent", "fstype=ext4,mount=/", 3, 0, "/分区inodes不足20%，警告", "all(#3)", "<", "20"},
		{"kernel.maxfiles", "", 3, 2, "内核参数配置过低", "all(#3)", "<", "1024"},
		{"kernel.maxproc", "", 3, 2, "内核参数配置过低", "all(#3)", "<", "256"},
		{"load.15min", "", 3, 3, "load大于60，警告", "all(#3)", ">", "60"},
		{"load.15min", "", 3, 0, "load大于65，严重", "all(#3)", ">", "65"},
		{"load.1min", "", 3, 3, "load大于80，警告", "all(#3)", ">", "80"},
		{"load.1min", "", 3, 0, "load大于80，严重", "all(#3)", ">", "80"},
		{"load.5min", "", 3, 3, "load大于80，警告", "all(#3)", ">", "80"},
		{"load.5min", "", 3, 0, "load大于80，严重", "all(#3)", ">", "80"},
		{"mem.memfree.percent", "", 3, 2, "内存可用率低于10%", "all(#3)", "<", "10"},
		{"mem.swapfree.percent", "", 3, 2, "swap小于50%告警", "all(#3)", "<", "50"},
		{"mem.swapfree.percent", "", 3, 2, "swap小于20%严重", "all(#3)", "<", "20"},
		{"system.boottime", "", 0, 0, "", "all(#3)", "==", "0"},
		{"system.cpu.switches", "", 0, 0, "", "all(#3)", ">", "0"},
		{"agent.alive", "", 3, 0, "agent上报异常告警", "all(#3)", "<", "1"},
		{"ads.server", "", 3, 0, "ads上报异常告警", "all(#1)", "<", "1"},
	}
}