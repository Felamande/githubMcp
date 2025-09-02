package model

type SearchOption struct {
	Query                   string `json:"query" jsonschema:"reqiured,description=the github repository search query string"`
	Sort                    string `json:"sort" jsonschema:"description=sort default by best match, can be [stars|fork|updated] for repository search"`
	Order                   string `json:"order" jsonschema:"description=sort order default by desc, can be [desc|asc]"`
	ResultPerpage           int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page                    int    `json:"page" jsonschema:"description=current page number of the search result,start from 1 and default to 1"`
	DescriptionTruncateSize int    `json:"description_truncate_size" jsonschema:"description=size of truncating very long release description, default for 1024"`
}

type SearchResult struct {
	TotalRepoNum int
	NextPage     int
	LastPage     int
	Repositories []RepositoryInfo
}

type RepositoryInfo struct {
	Owner           string
	Name            string
	Organization    string
	FullName        string
	MasterBranch    string
	Description     string
	StargazersCount int
	ForksCount      int
	Language        string
	CreatedAt       string
	UpdatedAt       string
	Archived        bool
}

type ReleaseListOption struct {
	Owner                   string `json:"owner" jsonschema:"reqiured,description=owner of repo you want to list releases"`
	Repository              string `json:"repository" jsonschema:"reqiured,description=name of repo you want to list releases"`
	ResultPerpage           int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page                    int    `json:"page" jsonschema:"description=current page number of the search result,start from 1 and default to 1"`
	DescriptionTruncateSize int    `json:"description_truncate_size" jsonschema:"description=size of truncating very long release description, default for 1024"`
}

type ReleaseListResult struct {
	NextPage int
	LastPage int
	Releases []ReleaseInfo
}

type ReleaseInfo struct {
	Name         string
	Tag          string
	Author       string
	IsDraft      bool
	IsPrerelease bool
	Description  string
	CreatedAt    string
	PublishedAt  string
	AssetsNum    int
}

type ReadmeOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	StartLine  int    `json:"start_line" jsonschema:"description=starting line number (1-based), default to 1"`
	EndLine    int    `json:"end_line" jsonschema:"description=ending line number, default to all lines"`
}

type ReadmeResult struct {
	Content   string
	StartLine int
	EndLine   int
	TotalLines int
}
