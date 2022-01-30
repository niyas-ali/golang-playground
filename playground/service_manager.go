package golangplayground

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type ServiceManager struct {
	connections map[string]Playground
}

func (sm ServiceManager) Add(ip string) {
	if _, ok := sm.connections[ip]; !ok {
		log.Printf("%s ip connected", ip)
		path := fmt.Sprintf("bin/%s", ip)
		playground := Playground{
			dir: path,
		}
		sm.connections[ip] = playground
		os.MkdirAll(path, os.ModePerm)
		playground.InitGoWorkspace()
	}
}
func (sm ServiceManager) GetPlayground(ip string) (Playground, error) {
	if v, ok := sm.connections[ip]; ok {
		return v, nil
	}
	return Playground{}, errors.New("no playground associated with this ip")
}
func (sm ServiceManager) Cleanup() {
	for range time.Tick(time.Minute * 5) {
		for key, p := range sm.connections {
			path := fmt.Sprintf("%s/%s", p.dir, mainFile)
			stat, _ := os.Stat(path)
			if time.Since(stat.ModTime()).Minutes() > float64(20) {
				fmt.Println("removed dir", p.dir)
				os.RemoveAll(p.dir)
				delete(sm.connections, key)
			}
		}
	}
}
func NewServiceManager() ServiceManager {
	sm := ServiceManager{}
	sm.connections = make(map[string]Playground)
	go sm.Cleanup()
	return sm
}
