package netx

import (
	"errors"
	"fmt"
	"strings"
)

/*
 * @author: Chen Chiheng
 * @date: 2023/4/6 22:03
 * @description:
 */

// Route 路由。。
type Route struct {
	ServiceID   string // ServiceID 服务ID。
	ServiceType string // ServiceType 服务类型。
	MethodName  string // MethodName 方法名称。
}

func (r *Route) String() string {
	return fmt.Sprintf("%s.%s.%s", r.ServiceID, r.ServiceType, r.MethodName)
}

// ParseRoute 路由解析。
func ParseRoute(routeStr string) (*Route, error) {
	values := strings.Split(routeStr, ".")
	if len(values) < 3 {
		return nil, errors.New("route:invalid route string")
	}
	return &Route{
		ServiceID:   values[0],
		ServiceType: values[2],
		MethodName:  values[1],
	}, nil
}
