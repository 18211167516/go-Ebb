package ebb

import (
	"time"
	"log"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("Logger [%d] %s in %v", c.HttpCode, c.Request.RequestURI, time.Since(t))
	}
}

func Testlogger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("TEST [%d] %s in %v", c.HttpCode, c.Request.RequestURI, time.Since(t))
	}
}