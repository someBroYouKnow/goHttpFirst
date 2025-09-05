package migrations

import "embed"

// below will find all sql files in this directory

//go:embed *.sql
var FS embed.FS
