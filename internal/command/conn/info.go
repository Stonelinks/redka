package conn

import (
	"github.com/nalgeon/redka/internal/redis"
)

// Info returns the given string.
// Info message
// https://redis.io/commands/info
type Info struct {
	redis.BaseCmd
}

func ParseInfo(b redis.BaseCmd) (Info, error) {
	return Info{BaseCmd: b}, nil
}

func (c Info) Run(w redis.Writer, _ redis.Redka) (any, error) {

	str := `# Server
redis_version:7.2.5
`
	// num_lines := len(strings.Split(str, "\n"))
	// w.WriteArray(num_lines)
	// for _, line := range strings.Split(str, "\n") {
	// 	// w.WriteBulkString(line)
	// 	w.WriteString(fmt.Sprintf("%s\n", line))
	// }
	bytes := []byte(str)
	w.WriteAny(bytes)

	return true, nil
}
