package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*7)
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Reserva cancelada devido timeout.")
	case <-time.After(time.Second * 5):
		fmt.Println("Reserva feita.")
	}
}
