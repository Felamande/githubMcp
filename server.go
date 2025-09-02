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

	err = server.RegisterTool("list_commits", "list commits of the repository",
		func(opt model.CommitListOption) (*mcpgo.ToolResponse, error) {
			commits, err := client.ListCommits(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(commits)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("get_commit_by_sha", "get commit details by SHA hash",
		func(opt model.CommitBySHAOption) (*mcpgo.ToolResponse, error) {
			commit, err := client.GetCommitBySHA(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(commit)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("list_branches", "list branches of the repository",
		func(opt model.BranchListOption) (*mcpgo.ToolResponse, error) {
			branches, err := client.ListBranches(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(branches)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("list_directory", "list directories and files in a repository directory",
		func(opt model.DirectoryListOption) (*mcpgo.ToolResponse, error) {
			directory, err := client.ListDirectory(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(directory)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("read_file", "read file content with line range support",
		func(opt model.ReadFileOption) (*mcpgo.ToolResponse, error) {
			fileContent, err := client.ReadFile(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(fileContent)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("find_tags", "find tags matching a regex pattern",
		func(opt model.FindTagsOption) (*mcpgo.ToolResponse, error) {
			tags, err := client.FindTags(opt)
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

	err = server.RegisterTool("find_branches", "find branches matching a regex pattern",
		func(opt model.FindBranchesOption) (*mcpgo.ToolResponse, error) {
			branches, err := client.FindBranches(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(branches)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("search_code", "search code across GitHub repositories",
		func(opt model.SearchCodeOption) (*mcpgo.ToolResponse, error) {
			codeResults, err := client.SearchCode(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(codeResults)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("list_issues", "list repository issues with filtering",
		func(opt model.ListIssuesOption) (*mcpgo.ToolResponse, error) {
			issues, err := client.ListIssues(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(issues)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("search_issues", "search issues across GitHub",
		func(opt model.SearchIssuesOption) (*mcpgo.ToolResponse, error) {
			issues, err := client.SearchIssues(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(issues)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("list_issue_comments", "list comments for a specific issue",
		func(opt model.ListIssueCommentsOption) (*mcpgo.ToolResponse, error) {
			comments, err := client.ListIssueComments(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(comments)
			return mcpgo.NewToolResponse(mcpgo.NewTextContent(string(out))), nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTool("list_issue_labels", "list all labels available in a repository",
		func(opt model.ListIssueLabelsOption) (*mcpgo.ToolResponse, error) {
			labels, err := client.ListIssueLabels(opt)
			if err != nil {
				return nil, err
			}
			out, err := json.Marshal(labels)
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
