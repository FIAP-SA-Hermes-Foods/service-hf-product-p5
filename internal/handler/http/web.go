package http

import (
	"errors"
	"net/http"
	"strings"
)

func router(reqMethod, path string, routesMap map[string]http.HandlerFunc) (http.HandlerFunc, error) {
	route := ""

	for k := range routesMap {
		isValidRoute, rr, m := validRoute(k, path, reqMethod)
		if isValidRoute && m == strings.ToLower(reqMethod) {
			route = rr
		}
	}

	if handler, ok := routesMap[route]; ok {
		return handler, nil
	}
	return nil, errors.New("route not found")
}

func validRoute(route, requestRoute, method string) (bool, string, string) {

	requestRoute = strings.ToLower(method) + " " + requestRoute

	isValid := false

	if requestRoute[len(requestRoute)-1] == '/' {
		requestRoute = requestRoute[:len(requestRoute)-1]
	}

	routeItems := strings.Split(route, "/")
	desiredRouteItems := strings.Split(requestRoute, "/")
	if len(routeItems) != len(desiredRouteItems) {
		return false, "", ""
	}

	idParamVal := ""

	for i := 0; i < len(routeItems); i++ {
		if strings.Contains(routeItems[i], "{") {
			idParamVal = desiredRouteItems[i]
		}
		if routeItems[i] == desiredRouteItems[i] {
			isValid = true
			continue
		}
		isValid = false
	}

	if idParamVal != "" {
		isValid = true
	}

	var methodReturn string
	if len(route) > 0 {
		methodReturn = strings.Split(route, " ")[0]
	}

	return isValid, route, methodReturn
}

func getID(handlerName, url string) string {
	index := strings.Index(url, handlerName+"/")

	if index == -1 {
		return ""
	}

	id := strings.ReplaceAll(url[index:], handlerName+"/", "")

	return id
}
