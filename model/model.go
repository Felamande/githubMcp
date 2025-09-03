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

type AssetInfo struct {
	ID    int64
	URL   string
	Name  string
	Label string
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
	Assets       []AssetInfo
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

type GetCommitBySHAOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	SHA        string `json:"sha" jsonschema:"required,description=the SHA hash of the commit"`
}

type GetTagByNameOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	TagName    string `json:"tag_name" jsonschema:"required,description=name of the tag"`
}

type GetIssueByNumberOption struct {
	Owner       string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository  string `json:"repository" jsonschema:"required,description=name of the repository"`
	IssueNumber int    `json:"issue_number" jsonschema:"required,description=the issue number"`
}

type GetBranchByNameOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	BranchName string `json:"branch_name" jsonschema:"required,description=name of the branch"`
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

type ListIssuesOption struct {
	Owner         string   `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository    string   `json:"repository" jsonschema:"required,description=name of the repository"`
	State         string   `json:"state" jsonschema:"description=filter by issue state: open, closed, or all"`
	Labels        []string `json:"labels" jsonschema:"description=filter by labels (comma separated)"`
	Assignee      string   `json:"assignee" jsonschema:"description=filter by assignee username"`
	Creator       string   `json:"creator" jsonschema:"description=filter by issue creator username"`
	Mentioned     string   `json:"mentioned" jsonschema:"description=filter by mentioned username"`
	Milestone     string   `json:"milestone" jsonschema:"description=filter by milestone number or title"`
	Sort          string   `json:"sort" jsonschema:"description=issue list sort can be created, updated and comments. Default value is created"`
	Since         string   `json:"since" jsonschema:"description=filter issues updated after this timestamp"`
	ResultPerpage int      `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page          int      `json:"page" jsonschema:"description=current page number, start from 1 and default to 1"`
}

type SearchIssuesOption struct {
	Query         string `json:"query" jsonschema:"required,description=github issues search query"`
	Sort          string `json:"sort" jsonschema:"description=sort by: created, updated, comments"`
	Order         string `json:"order" jsonschema:"description=sort order: asc or desc"`
	ResultPerpage int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page          int    `json:"page" jsonschema:"description=current page number, start from 1 and default to 1"`
}

type ListIssueCommentsOption struct {
	Owner         string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository    string `json:"repository" jsonschema:"required,description=name of the repository"`
	IssueNumber   int    `json:"issue_number" jsonschema:"required,description=the issue number"`
	Since         string `json:"since" jsonschema:"description=filter comments updated after this timestamp"`
	ResultPerpage int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page          int    `json:"page" jsonschema:"description=current page number, start from 1 and default to 1"`
}

type ListIssueLabelsOption struct {
	Owner         string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository    string `json:"repository" jsonschema:"required,description=name of the repository"`
	ResultPerpage int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page          int    `json:"page" jsonschema:"description=current page number, start from 1 and default to 1"`
}

type IssuesListResult struct {
	TotalCount int
	NextPage   int
	LastPage   int
	Issues     []IssueInfo
}

type IssueInfo struct {
	Number    int
	Title     string
	State     string
	Body      string
	Labels    []string
	Assignee  string
	Assignees []string
	Milestone *MilestoneInfo
	Creator   string
	CreatedAt string
	UpdatedAt string
	ClosedAt  string
	URL       string
	HTMLURL   string
	Comments  int
}

type IssueCommentsResult struct {
	NextPage int
	LastPage int
	Comments []IssueCommentInfo
}

type IssueCommentInfo struct {
	ID        int64
	Body      string
	User      string
	CreatedAt string
	UpdatedAt string
	URL       string
	HTMLURL   string
}

type LableListResult struct {
	NextPage int
	LastPage int
	Labels   []LabelInfo
}

type LabelInfo struct {
	Name        string
	Color       string
	Description string
}

type UserInfo struct {
	Login     string
	ID        int64
	AvatarURL string
	HTMLURL   string
	Type      string
}

type MilestoneInfo struct {
	Number      int
	Title       string
	Description string
	State       string
	DueOn       string
}

type ListPROption struct {
	Owner         string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository    string `json:"repository" jsonschema:"required,description=name of the repository"`
	State         string `json:"state" jsonschema:"description=filter by PR state: open, closed, or all"`
	Head          string `json:"head" jsonschema:"description=filter by head branch or user:branch"`
	Base          string `json:"base" jsonschema:"description=filter by base branch name"`
	Sort          string `json:"sort" jsonschema:"description=sort by: created, updated, popularity, long-running"`
	Direction     string `json:"direction" jsonschema:"description=sort direction: asc or desc"`
	ResultPerpage int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page          int    `json:"page" jsonschema:"description=current page number, start from 1 and default to 1"`
}

type GetPullRequestByNumberOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	Number     int    `json:"number" jsonschema:"required,description=the pull request number"`
}

type SearchPROption struct {
	Query         string `json:"query" jsonschema:"required,description=github pull request search query"`
	Sort          string `json:"sort" jsonschema:"description=sort by: created, updated, comments"`
	Order         string `json:"order" jsonschema:"description=sort order: asc or desc"`
	ResultPerpage int    `json:"result_per_page" jsonschema:"description=results per page, default to 10"`
	Page          int    `json:"page" jsonschema:"description=current page number, start from 1 and default to 1"`
}

type PRListResult struct {
	TotalCount int
	NextPage   int
	LastPage   int
	PRs        []PRInfo
}

type PRInfo struct {
	Number             int
	Title              string
	State              string
	Body               string
	Labels             []string
	Assignee           string
	Assignees          []string
	RequestedReviewers []string
	Milestone          *MilestoneInfo
	Creator            string
	CreatedAt          string
	UpdatedAt          string
	ClosedAt           string
	MergedAt           string
	URL                string
	HTMLURL            string
	Comments           int
	Additions          int
	Deletions          int
	ChangedFiles       int
	Mergeable          bool
	MergeableState     string
	Merged             bool
	BaseRef            string
	HeadRef            string
	Draft              bool
	ReviewComments     int
	Commits            int
}

type GetCommitFilesBySHAOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	SHA        string `json:"sha" jsonschema:"required,description=the SHA hash of the commit"`
}

type CommitFileInfo struct {
	SHA              string `json:"sha,omitempty"`
	Filename         string `json:"filename,omitempty"`
	Additions        int    `json:"additions,omitempty"`
	Deletions        int    `json:"deletions,omitempty"`
	Changes          int    `json:"changes,omitempty"`
	Status           string `json:"status,omitempty"`
	Patch            string `json:"patch,omitempty"`
	BlobURL          string `json:"blob_url,omitempty"`
	RawURL           string `json:"raw_url,omitempty"`
	ContentsURL      string `json:"contents_url,omitempty"`
	PreviousFilename string `json:"previous_filename,omitempty"`
}

type CommitFilesResult struct {
	Files []CommitFileInfo `json:"files"`
}

type CompareCommitsOption struct {
	Owner      string `json:"owner" jsonschema:"required,description=owner of the repository"`
	Repository string `json:"repository" jsonschema:"required,description=name of the repository"`
	Base       string `json:"base" jsonschema:"required,description=base commit SHA or branch name"`
	Head       string `json:"head" jsonschema:"required,description=head commit SHA or branch name"`
}

type CompareCommitsResult struct {
	TotalCommits     int           `json:"total_commits"`
	AheadBy          int           `json:"ahead_by"`
	BehindBy         int           `json:"behind_by"`
	Commits          []CommitInfo  `json:"commits"`
	Files            []CommitFileInfo `json:"files"`
	HTMLURL          string        `json:"html_url"`
	PermalinkURL     string        `json:"permalink_url"`
	DiffURL          string        `json:"diff_url"`
	PatchURL         string        `json:"patch_url"`
	Status           string        `json:"status"`
}
