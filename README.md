# GitHub MCP Server

A Model Context Protocol (MCP) server for GitHub that provides comprehensive access to GitHub APIs through MCP tools.

## Features

### Repository Tools
- **`search_github_repository`** - Search GitHub repositories using GitHub search syntax
- **`get_repository_releases`** - Get releases of a repository
- **`get_readme`** - Get README content with line range support
- **`get_repository_tags`** - List repository tags
- **`list_branches`** - List repository branches
- **`list_directory`** - List directories and files in a repository
- **`read_file`** - Read file content with line range support

### Code Search
- **`search_code`** - Search code across GitHub repositories with text matching

### Issue Management
- **`list_issues`** - List repository issues with filtering (state, labels, assignee, etc.)
- **`search_issues`** - Search issues across GitHub
- **`list_issue_comments`** - List comments for a specific issue
- **`list_issue_labels`** - List all labels available in a repository

### Pull Request Tools
- **`list_pull_requests`** - List repository PRs with filtering and sorting
- **`get_pull_request`** - Get detailed PR information including diff stats
- **`search_pull_requests`** - Search PRs across GitHub repositories

### Advanced Filtering
- **`find_tags`** - Find tags matching a regex pattern
- **`find_branches`** - Find branches matching a regex pattern

## Installation

```bash
go build -o githubMcp
```

## Usage

Run the MCP server:
```bash
./githubMcp
```

The server communicates via stdio using the MCP protocol and can be integrated with MCP clients like Claude Code.

## API Examples

### Search Repositories
```json
{
  "query": "language:go stars:>1000",
  "sort": "stars",
  "order": "desc",
  "result_per_page": 10
}
```

### Get Repository Details
```json
{
  "owner": "github",
  "repository": "hub"
}
```

### Search Code
```json
{
  "query": "func main language:go",
  "result_per_page": 5
}
```

### List Issues with Filtering
```json
{
  "owner": "owner",
  "repository": "repo",
  "state": "open",
  "labels": ["bug", "enhancement"]
}
```

### Get Pull Request Details
```json
{
  "owner": "owner",
  "repository": "repo", 
  "number": 123
}
```

## Configuration

The server uses the official GitHub Go client and supports:
- Public repository access (no authentication required)
- Pagination support for all list operations
- Rich filtering and sorting options
- Comprehensive error handling

## Development

### Project Structure
- `server.go` - MCP server implementation and tool registration
- `client/client.go` - GitHub API client implementation
- `model/model.go` - Data structures and request/response models

### Dependencies
- `github.com/google/go-github/v74` - GitHub Go API client
- `github.com/metoro-io/mcp-golang` - MCP framework for Go

## License

MIT License