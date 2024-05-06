package gongtable

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gongtable/dist/ng-github.com-fullstack-lang-gongtable
var NgDistNg embed.FS
