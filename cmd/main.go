package main

import (
	"context"
	"fmt"
	"part1/internal/action"
	"part1/internal/model"
	"part1/internal/service"
	"sync"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	//ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	userCreator := action.UserCreateWithStatReport{StatusChannel: make(chan model.Status)}
	cartCreator := action.CartCreateWithStatReport{StatusChannel: make(chan model.Status)}
	userUpdate := action.UserUpdateWithStatReport{StatusChannel: make(chan model.Status)}

	res := service.MergeData(ctx, userCreator.StatusChannel, cartCreator.StatusChannel, userUpdate.StatusChannel)

	wg.Add(1)
	go func() {
		defer wg.Done()
		user1 := userCreator.Do("Alex")
		user2 := userCreator.Do("Mike")
		cartCreator.Do(user1)
		cartCreator.Do(user2)
		userUpdate.User = user1
		userUpdate.Do("Boris")
	}()

	for val := range res {
		fmt.Print(val.Status + " ")
		fmt.Println(val.Id)
	}
}
