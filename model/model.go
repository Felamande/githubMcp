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
	Ref        string `json:"ref" jsonschema:"description=the name of the commit/branch/tag, default uses repository's default branch"`
	StartLine  int    `json:"start_line" jsonschema:"description=starting line number (1-based), default to 1"`
	EndLine    int    `json:"end_line" jsonschema:"description=ending line number, default to all lines"`
}

type ReadmeResult struct {
	Content    string
	StartLine  int
	EndLine    int
	TotalLines int
}

type TagListOption struct {
	Owner         string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository    string `json:"repository" jsonschema:"required,description=name of the repository"`
	ResultPerpage int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page          int    `json:"page" jsonschema:"description=current page number of the search result,start from 1 and default to 1"`
}

type TagListResult struct {
	NextPage int
	LastPage int
	Tags     []TagInfo
}

type TagInfo struct {
	Name       string
	CommitSHA  string
	ZipballURL string
	TarballURL string
}

type CommitListOption struct {
	Owner         string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository    string `json:"repository" jsonschema:"required,description=name of the repository"`
	ResultPerpage int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page          int    `json:"page" jsonschema:"description=current page number of the search result,start from 1 and default to 1"`
}

type CommitListResult struct {
	NextPage int
	LastPage int
	Commits  []CommitInfo
}

type CommitInfo struct {
	SHA              string
	Message          string
	Author           string
	AuthorEmail      string
	Committer        string
	CommitterEmail   string
	Date             string
	URL              string
	ParentCommitHash []string
}

type CommitBySHAOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	SHA        string `json:"sha" jsonschema:"required,description=the SHA hash of the commit"`
}

type BranchListOption struct {
	Owner         string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository    string `json:"repository" jsonschema:"required,description=name of the repository"`
	ResultPerpage int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page          int    `json:"page" jsonschema:"description=current page number of the search result,start from 1 and default to 1"`
}

type BranchListResult struct {
	NextPage int
	LastPage int
	Branches []BranchInfo
}

type BranchInfo struct {
	Name      string
	CommitSHA string
	Protected bool
}

type DirectoryListOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	Path       string `json:"path" jsonschema:"description=directory path to list, default to root directory"`
	Ref        string `json:"ref" jsonschema:"description=the name of the commit/branch/tag, default uses repository's default branch"`
}

type DirectoryListResult struct {
	Infos []DirectoryOrFileInfo
}

type DirectoryOrFileInfo struct {
	Name     string
	Path     string
	Size     int64
	Type     string
	Encoding string
}

type ReadFileOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	Path       string `json:"path" jsonschema:"required,description=file path to read"`
	Ref        string `json:"ref" jsonschema:"description=the name of the commit/branch/tag, default uses repository's default branch"`
	StartLine  int    `json:"start_line" jsonschema:"description=starting line number (1-based), default to 1"`
	EndLine    int    `json:"end_line" jsonschema:"description=ending line number, default to all lines"`
}

type ReadFileResult struct {
	Content    string
	StartLine  int
	EndLine    int
	TotalLines int
	Encoding   string
}

type FindTagsOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	Pattern    string `json:"pattern" jsonschema:"required,description=regex pattern to match against tag names"`
}

type FindBranchesOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	Pattern    string `json:"pattern" jsonschema:"required,description=regex pattern to match against branch names"`
}

type FindTagsResult struct {
	Tags []TagInfo
}

type FindBranchesResult struct {
	Branches []BranchInfo
}

type SearchCodeOption struct {
	Query         string `json:"query" jsonschema:"required,description=the github code search query string"`
	Sort          string `json:"sort" jsonschema:"description=sort default by best match, can be [indexed] for code search"`
	Order         string `json:"order" jsonschema:"description=sort order default by desc, can be [desc|asc]"`
	ResultPerpage int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page          int    `json:"page" jsonschema:"description=current page number of the search result,start from 1 and default to 1"`
}

type SearchCodeResult struct {
	TotalCount int
	NextPage   int
	LastPage   int
	CodeFiles  []CodeFileInfo
}

type CodeFileInfo struct {
	Name        string
	Path        string
	Repository  string
	Owner       string
	HTMLURL     string
	TextMatches []TextMatch
}

type TextMatch struct {
	Fragment   string
	Matches    []MatchDetail
	ObjectType string
	ObjectURL  string
	Property   string
}

type MatchDetail struct {
	Indices []int
	Text    string
}
