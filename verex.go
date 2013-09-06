package verex

import "regexp"

type Verex struct {
	prefixes string
	source   string
	suffixes string
	pattern  string
}

func New() *Verex {
	return &Verex{
		prefixes: "",
		source:   "",
		suffixes: "",
		pattern:  "",
	}
}

func (v *Verex) Add(value string) *Verex {
	v.source += value
	v.pattern = v.prefixes + v.source + v.suffixes
	return v
}

func (v *Verex) StartOfLine() *Verex {
	v.prefixes = "^"
	return v.Add("")
}

func (v *Verex) EndOfLine() *Verex {
	v.suffixes = "$"
	return v.Add("")
}

func (v *Verex) Then(value string) *Verex {
	return v.Add("(?:" + value + ")")
}

func (v *Verex) Find(value string) *Verex {
	return v.Then(value)
}

func (v *Verex) Maybe(value string) *Verex {
	return v.Add("(?:" + value + ")?")
}

func (v *Verex) Anything() *Verex {
	return v.Add("(?:.*)")
}

func (v *Verex) AnythingBut(value string) *Verex {
	return v.Add("(?:[^" + value + "]*)")
}

func (v *Verex) Something() *Verex {
	return v.Add("(?:.+)")
}

func (v *Verex) SomethingBut(value string) *Verex {
	return v.Add("(?:[^" + value + "]+)")
}

func (v *Verex) LineBreak() *Verex {
	return v.Add("(?:(?:\\n)|(?:\\r\\n))")
}

func (v *Verex) Br() *Verex {
	return v.LineBreak()
}

func (v *Verex) Tab() *Verex {
	return v.Add("\\t")
}

func (v *Verex) Word() *Verex {
	return v.Add("\\w+")
}

func (v *Verex) AnyOf(value string) *Verex {
	return v.Add("[" + value + "]")
}

func (v *Verex) Any(value string) *Verex {
	return v.AnyOf(value)
}

func (v *Verex) Range(args ...string) *Verex {
	if argc := len(args); argc == 0 || argc%2 == 1 {
		panic("Range needs one or more pairs of arguments to work.")
	} else {
		value := "["
		for from, to := 0, 1; from != argc && to != argc; from, to = from + 2, to + 2 {
			value += args[from] + "-" + args[to]
		}
		value += "]"
		return v.Add(value)
	}
}

func (v *Verex) Multiple() *Verex {
    return v.Add("+")
}

func (v *Verex) WithAnyCase() *Verex {
    v.prefixes = "(?i)" + v.prefixes
    return v.Add("")
}

func (v *Verex) Compile() *regexp.Regexp {
    return regexp.MustCompile(v.pattern)
}


