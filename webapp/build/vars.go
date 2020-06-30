package build

import "encoding/json"

var (
	Time   string
	Commit string
)

type Environment struct {
	Time   string
	Commit string
}

func (e Environment) String() string {
	b, _ := json.Marshal(Env())
	return string(b)
}

func Env() Environment {
	return Environment{
		Time:   Time,
		Commit: Commit,
	}
}
