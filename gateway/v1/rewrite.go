package v1

// Rewrite matching request URI to replacement.
type Rewrite struct {
	Regex       string
	Replacement string
	Flag        string
}

// Equals equals vs
func (v *Rewrite) Equals(c *Rewrite) bool {
	if v.Regex != c.Regex {
		return false
	}
	if v.Replacement != c.Replacement {
		return false
	}
	if v.Flag != c.Flag {
		return false
	}
	return true
}

func newFakeRewrite() *Rewrite {
	return &Rewrite{
		Regex:       "regex",
		Replacement: "replacement",
		Flag:        "flag",
	}
}
