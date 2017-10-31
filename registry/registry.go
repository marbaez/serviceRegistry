package registry

import (
	"time"
	"crypto/md5"
	"io"
	"fmt"
)

type Instance struct {
	instanceUrl		string
	lastHeartbeat 	time.Time
	healthUrl		string
}

type Service map[string]map[string]Instance
var registry Service;

func init() {
	registry = make(Service)
}

// AddInstance añade una nueva instancia a un servicio existente o si no existe el servicio lo crea también
func AddInstance(serviceId string, instance Instance) error {
	if registry[serviceId] == nil {
		//se trata de un nuevo servicio
		registry[serviceId] = make(map[string]Instance)
	}
	h := md5.New()
	io.WriteString(h, instance.instanceUrl)
	fmt.Printf("%x", h.Sum(nil))
	return nil;
}
