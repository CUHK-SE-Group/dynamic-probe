package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"kprobe/executor"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
)

type Contoller struct {
	Conf *BpfRuntimeConfig
	Opt  *Options

	templMgr *executor.TemplateMgr
}

func InitContorller(opts ...Option) (*Contoller, error) {
	ctl := &Contoller{}
	ctl.Opt = &Options{
		FuncMap:          make(template.FuncMap),
		ExecuteTemplates: make([]string, 0),
	}
	for _, f := range opts {
		f(ctl.Opt)
	}
	fmt.Printf("%+v\n", ctl.Opt)
	if ctl.Opt.ConfPath != "" {
		ctl.Conf = &BpfRuntimeConfig{}
		file, err := os.ReadFile(ctl.Opt.ConfPath)
		if err != nil {
			return nil, errors.Wrap(err, "open config failed")
		}
		err = toml.Unmarshal(file, ctl.Conf)
		if err != nil {
			return nil, errors.Wrap(err, "decode config failed")
		}
	}
	ctl.templMgr = &executor.TemplateMgr{}

	ctl.templMgr.LoadFuncMap(&ctl.Opt.FuncMap) //load function must be call before load templates
	err := ctl.templMgr.LoadTemplates(ctl.Opt.Templates...)
	if err != nil {
		return nil, err
	}
	return ctl, nil
}

func (c *Contoller) Run() {
	fmt.Println(c.Conf)
	for _, name := range c.templMgr.GetNames() {
		for _, v := range c.Opt.ExecuteTemplates {
			if v == name {
				filename := path.Base(name)
				if c.Opt.OutputStrategy == file {
					file, err := os.OpenFile(path.Join(c.Opt.OutputDir, strings.TrimSuffix(filename, ".gtpl")), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
					if err != nil {
						log.Fatalf("failed to open file: %v", err)
					}
					defer file.Close()
					err = c.templMgr.Generate(file, filename, c.Conf.EBPFProgram)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}

	// if c.Opt.GenerateCCode {
	// 	for _, name := range c.templMgr.GetNames() {
	// 		if strings.HasSuffix(name, "c.gtpl") {
	// 			filename := path.Base(name)
	// 			if c.Opt.OutputStrategy == file {
	// 				file, err := os.OpenFile(path.Join(c.Opt.OutputDir, strings.TrimSuffix(filename, ".gtpl")), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	// 				if err != nil {
	// 					log.Fatalf("failed to open file: %v", err)
	// 				}
	// 				defer file.Close()
	// 				err = c.templMgr.Generate(file, filename, c.Conf.EBPFProgram)
	// 				if err != nil {
	// 					log.Fatal(err)
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	if c.Opt.CompileCCode {
		for _, name := range c.templMgr.GetNames() {
			if strings.HasSuffix(name, "c.gtpl") {
				filename := path.Base(name)
				cmd := exec.Command("go", "run", "github.com/cilium/ebpf/cmd/bpf2go", "-target", "amd64", "-output-dir", c.Opt.OutputDir, "bpf", path.Join(c.Opt.OutputDir, strings.TrimSuffix(filename, ".gtpl")), "--", fmt.Sprintf("-I%s", c.Opt.CHeaders))
				log.Println(cmd.String())
				cmd.Env = append(os.Environ(), "GOPACKAGE=main")
				var stderr bytes.Buffer
				cmd.Stderr = &stderr

				err := cmd.Run()
				if err != nil {
					log.Fatalf("Command execution failed: %v, stderr: %s", err, stderr.String())
				}
			}
		}

	}

	// if c.Opt.GenerateCtlCode {
	// 	for _, name := range c.templMgr.GetNames() {
	// 		if strings.HasSuffix(name, "go.gtpl") {
	// 			filename := path.Base(name)
	// 			if c.Opt.OutputStrategy == file {
	// 				file, err := os.OpenFile(path.Join(c.Opt.OutputDir, strings.TrimSuffix(filename, ".gtpl")), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	// 				if err != nil {
	// 					log.Fatalf("failed to open file: %v", err)
	// 				}
	// 				defer file.Close()
	// 				err = c.templMgr.Generate(file, filename, c.Conf.EBPFProgram)
	// 				if err != nil {
	// 					log.Fatal(err)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}
