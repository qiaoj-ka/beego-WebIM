// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package models

import (
	"container/list"
)

type EventType int

//三种事件类型：加入、离开、消息
const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

//定义事件结构（事件类型、用户名、事件、内容）
type Event struct {
	Type      EventType // JOIN, LEAVE, MESSAGE
	User      string
	Timestamp int // Unix timestamp (secs)
	Content   string
}

//用来保存服务器上能够保存的消息记录，保存最新的20条
const archiveSize = 20

// Event archives.
//事件归档保存
var archive = list.New()

// NewArchive saves new event to archive list.
//将一个新的事件保存在archive中，若事件的个数已经大于等于20则删除第一个，只保留最新的20个
func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}

// GetEvents returns all events after lastReceived
//根据传过来的时间戳返回时间戳之后的所有事件消息
func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
