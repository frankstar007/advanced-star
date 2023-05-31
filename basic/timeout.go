/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/3/17
 * @contact frankstarye@tencent.com
 **/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ch := make(chan string, 1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	go do(ctx, ch)

	select {
	case  <- ctx.Done():
		fmt.Printf("context cancelled...%v\n", ctx.Err())
	case result := <- ch:
		fmt.Printf("received..%s\n", result)
	}


}




func do(ctx context.Context, ch chan string) {
	fmt.Println("start do something")
	time.Sleep(time.Second * 5)
	fmt.Println("finished something")
	ch <- "done"
}