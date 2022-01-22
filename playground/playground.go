package golangplayground

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

var (
	mainFile string = "main.go"
)

type Playground struct {
	cmd *exec.Cmd
	dir string
}

func (p Playground) InitGoWorkspace() {
	p.cmd = exec.Command("go", "mod", "init", "main")
	p.cmd.Dir = p.dir
	p.cmd.Run()
	p.CopySourceCode([]byte(`package main`))
}
func (p Playground) RunGoGet() {
	p.cmd = exec.Command("go", "get")
	p.cmd.Dir = p.dir
	p.cmd.Run()
}
func (p Playground) RunPlayground() (string, error) {
	p.RunGoGet()
	p.cmd = exec.Command("go", "run")
	p.cmd.Dir = p.dir
	p.cmd.Args = append(p.cmd.Args, mainFile)
	data, _ := p.cmd.CombinedOutput()
	return string(data), nil
}
func (p Playground) FormatSourceCode() (string, error) {
	p.cmd = exec.Command("gofmt")
	p.cmd.Dir = p.dir
	p.cmd.Args = append(p.cmd.Args, mainFile)
	data, _ := p.cmd.Output()
	return string(data), nil
}
func (p Playground) CopySourceCode(code []byte) {
	path := fmt.Sprintf("%s/%s", p.dir, mainFile)
	log.Println("copying source code into ", path)
	ioutil.WriteFile(path, code, 0644)
}
