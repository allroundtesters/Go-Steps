package cli

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type FlagType string

const (
	FlagString FlagType = "string"
	FlagBool   FlagType = "bool"
	FlagSlice  FlagType = "slice"
	FlagArray  FlagType = "array"
)

type Flag struct {
	Name      string
	ShortHand string
	Usage     string
	Default   string
	Type      FlagType
	IsValid   func(string) bool
}

type Values map[string][] string

func (v *Values) GetString(key string) string {
	r := (*v)[key]
	if len(r) == 0 {
		return ""
	} else {
		return r[0]
	}
}

func (v *Values) GetInt64(key string) (int64, error) {
	ns := v.GetString(key)
	if ns == "" {
		return 0, nil
	}
	n, err := strconv.ParseInt(ns, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("%s invalid: not a integer", key)
	}
	return n, nil
}

func (v *Values) GetStringArray(key string) []string {
	return (*v)[key]
}

func (v *Values) GetBool(key string) bool {
	r := strings.ToLower(v.GetString(key))
	return r == "true" || r == "yes" || r == "y" || r == "1"
}

func (v *Values) GetStringSlice(key string) []string {
	if strings.TrimSpace(v.GetString(key)) == "" {
		return nil
	}
	res := strings.Split(v.GetString(key), "||")
	if len(res) == 1 && strings.Contains(res[0], ",") {
		return strings.Split(res[0], ",")
	}
	return res
}

type Arg struct {
	Name       string
	IsValid    func(string) bool
	Weight     int
	AllowEmpty bool
}

func orderArgs(a ...Arg) args {
	for i := range a {
		if a[i].Weight == 0 {
			a[i].Weight = i
		}
	}
	res := args(a)
	sort.Sort(res)
	return res
}

type args []Arg

func (s args) Len() int {
	return len(s)
}

func (s args) Less(i, j int) bool {
	return s[i].Weight < s[j].Weight
}

func (s args) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Command struct {
	Name         string
	Ctx          []Arg
	Args         []Arg
	OptionArgs   []Arg
	VariadicArgs []Arg
	Short        string
	Long         string
	Example      string
	Flags        []Flag
	Aliases      []string
	Hidden       bool
	PreRun       func(c *Command, args *[]string) error
}

type CommandModifier func(*Command, interface{})

func CommandWithoutExtractFlags(c *Command, run interface{}) {}

func CommandWithExtractFlags(c *Command, run interface{}) {
	var extractFlags = []Flag{}
	switch run.(type) {
	case RunGetFunc:
		extractFlags = []Flag{
			{
				Name:    "format",
				Default: "plain",
				Usage:   "Output format: plain | json| yaml ",
				Type:    FlagString,
			}, {
				Name:  "verbose",
				Usage: "Display all object fields",
				Type:  FlagBool,
			}, {
				Name:      "quiet",
				ShortHand: "q",
				Default:   "",
				Usage:     "Only display specified object fields.",
				Type:      FlagString,
			},
		}
	}
	case RunListFunc:
		extractFlags = []Flag{
		{
			Name: "filter",
			Default: "",
			Usage: "Filter output based on condition",
			Type: FlagString
		},
		{
			Name:    "format",
			Default: "table",
			Usage:   "Output format: table|json|yaml",
			Type:    FlagString,
		},
		{
			Name: "quiet",
			Default: "q",
			Usage: "Only display object's key",
			Type: FlagString
		},{
			Name: "fields",
			Default: "",
			Usage: "Only display specified object"
			Type: FloatString,
		},{
			Name: "verbose",
			Default: "",
			Type: FlagBool
		},
	}
	case RunDeleteFunc:
		extraFlags =[]Flag{
		{
			Name: "force",
			Default: "false",
			Usage: "Force delete without confirmation"
			Type: FlagBool,
		},
		}
}

