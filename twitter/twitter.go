https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
package main

import (
	"io"
	"os"
	"sync"
	"encoding/json"
	"bufio"
)

import "strconv"
import "fmt"
import "proj1/feed"


type TaskQueue struct {
	Mutex *sync.Mutex
	Cond *sync.Cond
	Data []map[string]interface{}
	Status feed.Feed
	Done bool
}
