package utils

import "log"

// ServiceDownNotifier this function will be start service post down tasks
func ServiceDownNotifier(err error, serviceName string) {
	log.Default().Printf("service %s has been downed by %s", serviceName, err.Error())
}
