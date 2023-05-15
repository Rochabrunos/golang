package main

import (
	"fmt"
	"net"
	"net/url"
)
/*
URLs provide a uniform way to locate resources
*/

func main() {
	//This example url include: scheme, authentication info
	//host, port, path, query params, and query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	//contains all infromation about authentication
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	//Host contains both the hostname and the port	
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	//the fragment after #
	fmt.Println(u.Fragment)
	//print query params in a string k=v format
	fmt.Println(u.RawQuery)
	//maps the query
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}