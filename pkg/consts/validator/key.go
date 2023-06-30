package validator

type TagType int

//go:generate stringer -type=TagType -linecomment
const (
	Min TagType = iota

	Name     // name
	IP       // ip
	Password // password

	Max
)

var RuleLoaderMap = map[string]any{
	Name.String():     checkNamePattern(),
	IP.String():       checkIpPattern(),
	Password.String(): checkPasswordPattern(),
	Min.String():      checkMinPattern(),
}
