package api

import "mall/service"

func List() string {
	return service.GetList()
}
