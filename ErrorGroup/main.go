package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {

	urls := []string{
		"google.com",
		"abc.com",
	}

	g, ctx := errgroup.WithContext(context.Background())

	for _, url := range urls {
		//url := url
		g.Go(func() error {
			if url == "error.com" {
				return fmt.Errorf("failed fetching %s", url)
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				fmt.Printf("successfully fetched %s \n", url)
				return nil
			}
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("error group finished with error : %v\n", err)
	} else {
		fmt.Println("done")
	}
}
