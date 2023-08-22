package models

type DisplayedColumn struct {
	Name string

	// The `sticky` property in CSS positions an element based on user's scroll,
	// allowing it to be fixed within its parent after a scroll point is reached.
	IsSticky bool
}
