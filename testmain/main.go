package main

import (
	"fmt"

	"github.com/Felamande/githubMcp/client"
	"github.com/Felamande/githubMcp/model"
	// 使用合适的版本号
)

func main() {
	client := client.NewClient()
	r, err := client.ListReleases(model.ReleaseListOption{
		Owner:      "koreader",
		Repository: "koreader",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
