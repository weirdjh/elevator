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

type MoveEvent struct {
	UserInfo UserInfo
	Move     Move
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

func ElevatorTraffic() []*MoveEvent {

	users := ReadTrafficFile()

	me := []*MoveEvent{}
	for _, u := range users.User {
		userInfo := u.UserInfo
		for _, m := range u.Move {
			me = append(me, &MoveEvent{
				UserInfo: userInfo,
				Move:     m,
			})
		}
	}

	// Sort with time
	sort.SliceStable(me, func(i, j int) bool {
		layout := "15:04:05.00"
		t1, _ := time.Parse(layout, me[i].Move.At)
		t2, _ := time.Parse(layout, me[j].Move.At)

		return t1.Before(t2)
	})

	return me
}
