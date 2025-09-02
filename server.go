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
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("get_readme", "get readme of the repository from start line to end line",
		func(opt model.ReadmeOption) (*mcpgo.ToolResponse, error) {
			readme, err := client.GetReadme(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(readme)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("get_repository_tags", "list tags of the repository",
		func(opt model.TagListOption) (*mcpgo.ToolResponse, error) {
			tags, err := client.ListTags(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(tags)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.Serve()
	if err != nil {
		panic(err)
	}
	<-done
}
