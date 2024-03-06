package policy

import _ "embed"

//go:embed schema.zed
var Schema string

//go:generate go run ../../../../scripts/policygen/main.go -destination generated.go
