package traffic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

type Users struct {
	User []User `json:"users"`
}

type User struct {
	UserInfo UserInfo `json:"userInfo"`
	Move     []Move   `json:"moves"`
}

type UserInfo struct {
	User   string `json:"userID"`
	Weight int    `json:"weight"`
}
type Move struct {
	At   string `json:"at"`
	From int    `json:"from"`
	To   int    `json:"to"`
}
type MoveTimeFormat struct {
	At   time.Time
	From int32
	To   int32
}

type MoveEvent struct {
	UserInfo UserInfo
	Move     MoveTimeFormat
}

func ReadTrafficFile() Users {
	var users Users
	byteValue, err := ioutil.ReadFile("traffic/user_traffic.json")
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(byteValue, &users); err != nil {
		fmt.Println(err)
	}

	return users
}

// TODO: Parsing and Covert Move.at into Time format
func ElevatorTraffic() []*MoveEvent {

	users := ReadTrafficFile()

	me := []*MoveEvent{}
	for _, u := range users.User {
		userInfo := u.UserInfo
		for _, m := range u.Move {

			layout := "2006-01-02T15:04:05.00"
			AtTimeFormat, _ := time.Parse(layout, "2018-01-01T"+m.At)

			me = append(me, &MoveEvent{
				UserInfo: userInfo,
				Move: MoveTimeFormat{
					At:   AtTimeFormat,
					From: int32(m.From),
					To:   int32(m.To),
				},
			})
		}
	}

	// Sort with time
	sort.SliceStable(me, func(i, j int) bool {
		t1 := me[i].Move.At
		t2 := me[j].Move.At
		return t1.Before(t2)
	})

	return me
}
