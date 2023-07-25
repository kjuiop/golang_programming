package main

import (
	"fmt"
	"strings"
)

type CommandBuilder struct {
	cmdArgs []string
}

func NewCommandBuilder() *CommandBuilder {
	return &CommandBuilder{}
}

func (cb *CommandBuilder) WithHWAccel(accelType string) *CommandBuilder {
	cb.cmdArgs = append(cb.cmdArgs, "-hwaccel", accelType)
	return cb
}

func (cb *CommandBuilder) WithHWAccelOutputFormat(format string) *CommandBuilder {
	cb.cmdArgs = append(cb.cmdArgs, "-hwaccel_output_format", format)
	return cb
}

func (cb *CommandBuilder) Build() string {
	return "ffmpeg " + joinArgs(cb.cmdArgs)
}

func joinArgs(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	command := NewCommandBuilder().
		WithHWAccel("cuda").
		WithHWAccelOutputFormat("cuda").
		Build()

	fmt.Println(command)
}
