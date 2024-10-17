package service

import "log"

type Service interface {
	Start()
}

var (
	serviceMap = make(map[string]Service)
)

func Register(name string, service Service) {
	if _, exist := serviceMap[name]; exist {
		log.Fatalf("service (%s) exists, please change a unique service name", name)
	}
	serviceMap[name] = service
}

func New(name string) Service {
	service, ok := serviceMap[name]
	if !ok {
		names := []string{}
		for key := range serviceMap {
			names = append(names, key)
		}
		log.Fatalf("service (%s) does not exists, please choose one of the servce list: %v", name, names)
	}
	return service
}
