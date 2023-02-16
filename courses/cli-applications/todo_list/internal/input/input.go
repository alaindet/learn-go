package input

import "flag"

type CLIFlags struct {
	flags map[string]any
}

func NewCLIFlags() *CLIFlags {
	flags := make(map[string]any, 0)
	return &CLIFlags{flags}
}

func (f *CLIFlags) String(name, val, desc string) {
	f.flags[name] = flag.String(name, val, desc)
}

func (f *CLIFlags) Int(name string, val int, desc string) {
	f.flags[name] = flag.Int(name, val, desc)
}

func (f *CLIFlags) Bool(name string, val bool, desc string) {
	f.flags[name] = flag.Bool(name, val, desc)
}

func (f *CLIFlags) Parse() {
	flag.Parse()
}

func (f *CLIFlags) GetString(key string) string {
	return *f.flags[key].(*string)
}

func (f *CLIFlags) GetInt(key string) int {
	return *f.flags[key].(*int)
}

func (f *CLIFlags) GetBool(key string) bool {
	return *f.flags[key].(*bool)
}
