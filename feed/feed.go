https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
package feed

import (
	"math"
	"proj1/lock"
)

//Feed represents a user's twitter feed
// You will add to this interface the implementations as you complete them.
type Feed interface {
	Add(body string, timestamp float64)
	Remove(timestamp float64) bool
	Contains(timestamp float64) bool
}

//feed is the internal representation of a user's twitter feed (hidden from outside packages)
// You CAN add to this structure but you cannot remove any of the original fields. You must use
// the original fields in your implementation. You can assume the feed will not have duplicate posts
type feed struct {
	start *post // a pointer to the beginning post
	lock  *lock.RWLock
}

//post is the internal representation of a post on a user's twitter feed (hidden from outside packages)
// You CAN add to this structure but you cannot remove any of the original fields. You must use
// the original fields in your implementation.
type post struct {
	body      string  // the text of the post
	timestamp float64 // Unix timestamp of the post
	next      *post   // the next post in the feed
}

//NewPost creates and returns a new post value given its body and timestamp
func newPost(body string, timestamp float64, next *post) *post {
	return &post{body, timestamp, next}
}

//NewFeed creates a empy user feed
func NewFeed() Feed {
	return &feed{start: nil, lock: lock.NewRWLock()}
}

// Add inserts a new post to the feed. The feed is always ordered by the timestamp where
// the most recent timestamp is at the beginning of the feed followed by the second most
// recent timestamp, etc. You may need to insert a new post somewhere in the feed because
// the given timestamp may not be the most recent.
func (f *feed) Add(body string, timestamp float64) {
	f.lock.Lock()
	defer f.lock.UnLock()

	if f.start == nil {

		f.start = newPost(body, timestamp, nil)

	} else {

		post := f.start

		for post.next != nil {
			post = post.next
		}

		post.next = newPost(body, timestamp, nil)

	}
}

// Remove deletes the post with the given timestamp. If the timestamp
// is not included in a post of the feed then the feed remains
// unchanged. Return true if the deletion was a success, otherwise return false
func (f *feed) Remove(timestamp float64) bool {
	f.lock.Lock()
	defer f.lock.UnLock()

	flag := false

	if f.start == nil {
		return false
	}

	if IsEqual(f.start.timestamp, timestamp) {
		f.start = f.start.next
		return true
	}

	post := f.start

	for post.next != nil {

		if IsEqual(post.next.timestamp, timestamp) {

			flag = true

			post.next = post.next.next

			break

		} else {

			post = post.next
		}

	}

	return flag
}

// Contains determines whether a post with the given timestamp is
// inside a feed. The function returns true if there is a post
// with the timestamp, otherwise, false.
func (f *feed) Contains(timestamp float64) bool {
	f.lock.RLock()
	defer f.lock.RUnLock()

	post := f.start

	flag := false

	for post != nil {
		if IsEqual(post.timestamp, timestamp) {
			flag = true
			break
		} else {
			post = post.next
		}

	}

	return flag

}

func IsEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.000001
}
