package leetcode

const (
	SubmitStatusAccepted            = "通过"
	SubmitStatusWrongAnswer         = "解答错误"
	SubmitStatusMemoryLimitExceeded = "超出内存限制"
	SubmitStatusOutputLimitExceeded = "超出输出限制"
	SubmitStatusTimeLimitExceeded   = "超出时间限制"
	SubmitStatusRuntimeError        = "执行出错"
	SubmitStatusInternalError       = "内部出错"
	SubmitStatusCompileError        = "编译出错"
	SubmitStatusTimeout             = "超时"
)

var (
	SubmitStatusMap = map[string]string{
		"A_10": SubmitStatusAccepted,
		"A_11": SubmitStatusWrongAnswer,
		"A_12": SubmitStatusMemoryLimitExceeded,
		"A_13": SubmitStatusOutputLimitExceeded,
		"A_14": SubmitStatusTimeLimitExceeded,
		"A_15": SubmitStatusRuntimeError,
		"A_16": SubmitStatusInternalError,
		"A_20": SubmitStatusCompileError,
		"A_30": SubmitStatusTimeout,
	}

	SubmitLanguageMap = map[string]string{
		"A_0":  "C++",
		"A_1":  "Java",
		"A_2":  "Python",
		"A_3":  "MySQL",
		"A_4":  "C",
		"A_5":  "C#",
		"A_6":  "JavaScript",
		"A_7":  "Ruby",
		"A_8":  "Bash",
		"A_9":  "Swift",
		"A_10": "Go",
		"A_11": "Python3",
		"A_12": "Scala",
		"A_13": "Kotlin",
		"A_14": "MS SQL Server",
		"A_15": "Oracle",
		"A_16": "HTML",
		"A_17": "Python ML",
		"A_18": "Rust",
		"A_19": "PHP",
		"A_20": "TypeScript",
		"A_21": "Racket",
		"A_22": "Erlang",
		"A_23": "Elixir",
	}

	DifficultyMap = map[string]string{
		"Easy":   "简单",
		"Medium": "中等",
		"Hard":   "困难",
	}

	CompanySlugMap = map[string]string{
		"google":    "谷歌",
		"bytedance": "字节跳动",
		"amazon":    "亚马逊",
	}
)
