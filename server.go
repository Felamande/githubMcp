package main

import (
	"encoding/json"

	"github.com/Felamande/githubMcp/client"
	"github.com/Felamande/githubMcp/model"
	mcpgo "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

func main() {
	done := make(chan struct{})
	client := client.NewClient()
	server := mcpgo.NewServer(stdio.NewStdioServerTransport())
	err := server.RegisterTool("search_github_repository", "search github repositories using github search syntax",
		func(opt model.SearchOption) (*mcpgo.ToolResponse, error) {
			searches, err := client.GetRepository(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(searches)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}
	err = server.RegisterTool("get_repository_releases", "get releases of the repository",
		func(opt model.ReleaseListOption) (*mcpgo.ToolResponse, error) {
			releases, err := client.ListReleases(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(releases)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)

	err = server.Serve()
	if err != nil {
		panic(err)
	}
	<-done
}
