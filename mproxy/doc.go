// Package mproxy implements mysql connection pool proxy
//
// Motivations:
//
//	* While working on python based web applications, we end up tuning the
//	processes and treads model based on the infrastructure and performance
//	requirements. If we have any connection caching strategy in place, the
//	tuning become complex since we are dealing with process/threads + connection
//	cache. Having a connection pool proxy helps with managing the connection
//	cache strategy independent of the process/threads model.
//	* Connection caching at the application level would be local to a Python VM
//	process and a connection established in a process can not be shared with
//	other processes.
package mproxy
