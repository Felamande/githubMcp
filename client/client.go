package client

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/Felamande/githubMcp/model"
	"github.com/google/go-github/v74/github"
)

type GithubClient struct {
	c *github.Client
}

func NewClient() *GithubClient {
	return &GithubClient{
		c: github.NewClient(nil),
	}
}

func (c *GithubClient) GetRepository(opt model.SearchOption) (r *model.SearchResult, err error) {
	if opt.ResultPerpage == 0 {
		opt.ResultPerpage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}

	if opt.DescriptionTruncateSize == 0 {
		opt.DescriptionTruncateSize = 1024
	}

	ctx := context.Background()
	opts := &github.SearchOptions{
		Sort:  opt.Sort,  // 按星标数排序
		Order: opt.Order, // 降序排列
		ListOptions: github.ListOptions{
			PerPage: opt.ResultPerpage, // 每页显示 10 个结果
			Page:    opt.Page,          // 第 1 页
		},
	}
	result, resp, err := c.c.Search.Repositories(ctx, opt.Query, opts)
	if err != nil {
		return nil, err
	}

	searches := &model.SearchResult{}
	searches.Repositories = make([]model.RepositoryInfo, 0)
	searches.TotalRepoNum = result.GetTotal()
	searches.LastPage = resp.LastPage
	searches.NextPage = resp.NextPage

	for _, repo := range result.Repositories {
		descLen := len(repo.GetDescription())
		if opt.DescriptionTruncateSize >= descLen {
			opt.DescriptionTruncateSize = descLen
		}
		repoInfo := model.RepositoryInfo{
			Name:            repo.GetName(),
			FullName:        repo.GetFullName(),
			MasterBranch:    repo.GetMasterBranch(),
			Description:     repo.GetDescription()[0:opt.DescriptionTruncateSize],
			StargazersCount: repo.GetStargazersCount(),
			ForksCount:      repo.GetForksCount(),
			Language:        repo.GetLanguage(),
			Archived:        repo.GetArchived(),
		}

		if repo.Owner != nil {
			repoInfo.Owner = repo.Owner.GetName()
		}
		if repo.Organization != nil {
			repoInfo.Organization = repo.Organization.GetCompany()
		}
		if repo.CreatedAt != nil {
			repoInfo.CreatedAt = repo.CreatedAt.Format("2006-01-02 15:04:05")
		}
		if repo.UpdatedAt != nil {
			repoInfo.UpdatedAt = repo.UpdatedAt.Format("2006-01-02 15:04:05")
		}
		searches.Repositories = append(searches.Repositories, repoInfo)
	}
	return searches, nil
}

func (c *GithubClient) ListReleases(opt model.ReleaseListOption) (*model.ReleaseListResult, error) {
	if opt.ResultPerpage == 0 {
		opt.ResultPerpage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}
	if opt.DescriptionTruncateSize == 0 {
		opt.DescriptionTruncateSize = 1024
	}
	opts := &github.ListOptions{
		PerPage: opt.ResultPerpage, // 每页显示 10 个结果
		Page:    opt.Page,          // 第 1 页
	}

	// 获取 releases 列表
	releases, resp, err := c.c.Repositories.ListReleases(context.Background(), opt.Owner, opt.Repository, opts)
	if err != nil {
		return nil, err
	}

	releasesResult := &model.ReleaseListResult{}
	releasesResult.Releases = make([]model.ReleaseInfo, 0)
	releasesResult.NextPage = resp.NextPage
	releasesResult.LastPage = resp.LastPage

	for _, release := range releases {
		descLen := len(release.GetBody())
		if opt.DescriptionTruncateSize >= descLen {
			opt.DescriptionTruncateSize = descLen
		}
		releaseResult := model.ReleaseInfo{
			Name: release.GetName(),
			Tag:  release.GetTagName(),
			// Author: release.GetAuthor(),
			IsDraft:      release.GetDraft(),
			IsPrerelease: release.GetPrerelease(),
			Description:  release.GetBody()[0:opt.DescriptionTruncateSize],
			CreatedAt:    release.GetCreatedAt().Format("2006-01-02 15:04:05"),
			PublishedAt:  release.GetPublishedAt().Format("2006-01-02 15:04:05"),
			AssetsNum:    len(release.Assets),
		}
		releasesResult.Releases = append(releasesResult.Releases, releaseResult)
	}

	return releasesResult, nil
}

func (c *GithubClient) GetReadme(opt model.ReadmeOption) (*model.ReadmeResult, error) {
	if opt.StartLine == 0 {
		opt.StartLine = 1
	}

	ctx := context.Background()
	contentGetOptions := (*github.RepositoryContentGetOptions)(nil)
	if opt.Ref != "" {
		contentGetOptions = &github.RepositoryContentGetOptions{
			Ref: opt.Ref,
		}
	}
	readme, _, err := c.c.Repositories.GetReadme(ctx, opt.Owner, opt.Repository, contentGetOptions)
	if err != nil {
		return nil, err
	}

	content, err := readme.GetContent()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(content, "\n")
	totalLines := len(lines)

	if opt.EndLine == 0 || opt.EndLine > totalLines {
		opt.EndLine = totalLines
	}

	if opt.StartLine < 1 {
		opt.StartLine = 1
	}
	if opt.EndLine < opt.StartLine {
		opt.EndLine = opt.StartLine
	}

	selectedLines := lines[opt.StartLine-1 : opt.EndLine]
	selectedContent := strings.Join(selectedLines, "\n")

	result := &model.ReadmeResult{
		Content:    selectedContent,
		StartLine:  opt.StartLine,
		EndLine:    opt.EndLine,
		TotalLines: totalLines,
	}

	return result, nil
}

func (c *GithubClient) ListTags(opt model.TagListOption) (*model.TagListResult, error) {
	if opt.ResultPerpage == 0 {
		opt.ResultPerpage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}

	ctx := context.Background()
	opts := &github.ListOptions{
		PerPage: opt.ResultPerpage,
		Page:    opt.Page,
	}
	tags, resp, err := c.c.Repositories.ListTags(ctx, opt.Owner, opt.Repository, opts)
	if err != nil {
		return nil, err
	}

	var tagInfos []model.TagInfo
	for _, tag := range tags {
		tagResult := model.TagInfo{
			Name:       tag.GetName(),
			ZipballURL: tag.GetZipballURL(),
			TarballURL: tag.GetTarballURL(),
		}
		if commit := tag.GetCommit(); commit != nil {
			tagResult.CommitSHA = commit.GetSHA()
		}

		tagInfos = append(tagInfos, tagResult)
	}

	result := &model.TagListResult{
		NextPage: resp.NextPage,
		LastPage: resp.LastPage,
		Tags:     tagInfos,
	}

	return result, nil
}

func (c *GithubClient) ListCommits(opt model.CommitListOption) (*model.CommitListResult, error) {
	if opt.ResultPerpage == 0 {
		opt.ResultPerpage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}

	ctx := context.Background()
	opts := &github.CommitsListOptions{
		ListOptions: github.ListOptions{
			PerPage: opt.ResultPerpage,
			Page:    opt.Page,
		},
	}

	commits, resp, err := c.c.Repositories.ListCommits(ctx, opt.Owner, opt.Repository, opts)
	if err != nil {
		return nil, err
	}

	var commitInfos []model.CommitInfo
	for _, commitResult := range commits {
		commitInfo := model.CommitInfo{
			SHA:              commitResult.GetSHA(),
			URL:              commitResult.GetHTMLURL(),
			ParentCommitHash: make([]string, 0),
		}

		if commit := commitResult.GetCommit(); commit != nil {
			commitInfo.Message = commit.GetMessage()

			if author := commit.GetAuthor(); author != nil {
				commitInfo.Author = author.GetName()
				commitInfo.AuthorEmail = author.GetEmail()
				commitInfo.Date = author.GetDate().Format(time.RFC3339)
			}
			if committer := commit.GetCommitter(); committer != nil {
				commitInfo.Committer = committer.GetName()
				commitInfo.CommitterEmail = committer.GetEmail()
			}

			for _, parent := range commit.Parents {
				commitInfo.ParentCommitHash = append(commitInfo.ParentCommitHash, parent.GetSHA())
			}

		}

		commitInfos = append(commitInfos, commitInfo)
	}

	result := &model.CommitListResult{
		NextPage: resp.NextPage,
		LastPage: resp.LastPage,
		Commits:  commitInfos,
	}

	return result, nil
}

func (c *GithubClient) GetCommitBySHA(opt model.CommitBySHAOption) (*model.CommitInfo, error) {
	ctx := context.Background()

	commitResult, _, err := c.c.Repositories.GetCommit(ctx, opt.Owner, opt.Repository, opt.SHA, nil)
	if err != nil {
		return nil, err
	}

	commitInfo := &model.CommitInfo{
		SHA: commitResult.GetSHA(),
		URL: commitResult.GetHTMLURL(),
	}

	if commit := commitResult.GetCommit(); commit != nil {
		commitInfo.Message = commit.GetMessage()

		if author := commit.GetAuthor(); author != nil {
			commitInfo.Author = author.GetName()
			commitInfo.AuthorEmail = author.GetEmail()
			commitInfo.Date = author.GetDate().Format(time.RFC3339)
		}
		if committer := commit.GetCommitter(); committer != nil {
			commitInfo.Committer = committer.GetName()
			commitInfo.CommitterEmail = committer.GetEmail()
		}

	}

	return commitInfo, nil
}

func (c *GithubClient) ListBranches(opt model.BranchListOption) (*model.BranchListResult, error) {
	if opt.ResultPerpage == 0 {
		opt.ResultPerpage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}

	ctx := context.Background()
	opts := &github.BranchListOptions{
		ListOptions: github.ListOptions{
			PerPage: opt.ResultPerpage,
			Page:    opt.Page,
		},
	}

	branches, resp, err := c.c.Repositories.ListBranches(ctx, opt.Owner, opt.Repository, opts)
	if err != nil {
		return nil, err
	}

	var branchInfos []model.BranchInfo
	for _, branch := range branches {
		branchInfo := model.BranchInfo{
			Name:      branch.GetName(),
			Protected: branch.GetProtected(),
		}

		if commit := branch.GetCommit(); commit != nil {
			branchInfo.CommitSHA = commit.GetSHA()
		}

		branchInfos = append(branchInfos, branchInfo)
	}

	result := &model.BranchListResult{
		NextPage: resp.NextPage,
		LastPage: resp.LastPage,
		Branches: branchInfos,
	}

	return result, nil
}

func (c *GithubClient) ListDirectory(opt model.DirectoryListOption) (*model.DirectoryListResult, error) {
	ctx := context.Background()

	// Set up options with ref if provided
	contentGetOptions := (*github.RepositoryContentGetOptions)(nil)
	if opt.Ref != "" {
		contentGetOptions = &github.RepositoryContentGetOptions{
			Ref: opt.Ref,
		}
	}

	// Get directory contents
	_, directoryContents, _, err := c.c.Repositories.GetContents(ctx, opt.Owner, opt.Repository, opt.Path, contentGetOptions)
	if err != nil {
		return nil, err
	}

	result := &model.DirectoryListResult{
		Infos: []model.DirectoryOrFileInfo{},
	}

	// Process directory contents
	for _, content := range directoryContents {
		result.Infos = append(result.Infos, model.DirectoryOrFileInfo{
			Name:     content.GetName(),
			Path:     content.GetPath(),
			Size:     int64(content.GetSize()),
			Type:     content.GetType(),
			Encoding: content.GetEncoding(),
		})
	}

	return result, nil
}

func (c *GithubClient) ReadFile(opt model.ReadFileOption) (*model.ReadFileResult, error) {
	if opt.StartLine == 0 {
		opt.StartLine = 1
	}

	ctx := context.Background()

	// Set up options with ref if provided
	contentGetOptions := (*github.RepositoryContentGetOptions)(nil)
	if opt.Ref != "" {
		contentGetOptions = &github.RepositoryContentGetOptions{
			Ref: opt.Ref,
		}
	}

	// Get file contents
	fileContent, _, _, err := c.c.Repositories.GetContents(ctx, opt.Owner, opt.Repository, opt.Path, contentGetOptions)
	if err != nil {
		return nil, err
	}

	// Check if it's a directory
	if fileContent.GetType() != "file" {
		return nil, fmt.Errorf("path is not a file: %s", opt.Path)
	}

	content, err := fileContent.GetContent()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(content, "\n")
	totalLines := len(lines)

	if opt.EndLine == 0 || opt.EndLine > totalLines {
		opt.EndLine = totalLines
	}

	if opt.StartLine < 1 {
		opt.StartLine = 1
	}
	if opt.EndLine < opt.StartLine {
		opt.EndLine = opt.StartLine
	}

	selectedLines := lines[opt.StartLine-1 : opt.EndLine]
	selectedContent := strings.Join(selectedLines, "\n")

	result := &model.ReadFileResult{
		Content:    selectedContent,
		StartLine:  opt.StartLine,
		EndLine:    opt.EndLine,
		TotalLines: totalLines,
		Encoding:   fileContent.GetEncoding(),
	}

	return result, nil
}

func (c *GithubClient) FindTags(opt model.FindTagsOption) (*model.FindTagsResult, error) {
	ctx := context.Background()

	tags, _, err := c.c.Repositories.ListTags(ctx, opt.Owner, opt.Repository, &github.ListOptions{
		PerPage: 1000, // Get more tags to search through
	})
	if err != nil {
		return nil, err
	}

	pattern, err := regexp.Compile(opt.Pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %v", err)
	}

	var matchedTags []model.TagInfo
	for _, tag := range tags {
		if pattern.MatchString(tag.GetName()) {
			tagResult := model.TagInfo{
				Name:       tag.GetName(),
				ZipballURL: tag.GetZipballURL(),
				TarballURL: tag.GetTarballURL(),
			}
			if commit := tag.GetCommit(); commit != nil {
				tagResult.CommitSHA = commit.GetSHA()
			}
			matchedTags = append(matchedTags, tagResult)
		}
	}

	result := &model.FindTagsResult{
		Tags: matchedTags,
	}

	return result, nil
}

func (c *GithubClient) FindBranches(opt model.FindBranchesOption) (*model.FindBranchesResult, error) {
	ctx := context.Background()

	branches, _, err := c.c.Repositories.ListBranches(ctx, opt.Owner, opt.Repository, &github.BranchListOptions{
		ListOptions: github.ListOptions{
			PerPage: 1000, // Get more branches to search through
		},
	})
	if err != nil {
		return nil, err
	}

	pattern, err := regexp.Compile(opt.Pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %v", err)
	}

	var matchedBranches []model.BranchInfo
	for _, branch := range branches {
		if pattern.MatchString(branch.GetName()) {
			branchInfo := model.BranchInfo{
				Name:      branch.GetName(),
				Protected: branch.GetProtected(),
			}
			if commit := branch.GetCommit(); commit != nil {
				branchInfo.CommitSHA = commit.GetSHA()
			}
			matchedBranches = append(matchedBranches, branchInfo)
		}
	}

	result := &model.FindBranchesResult{
		Branches: matchedBranches,
	}

	return result, nil
}

func (c *GithubClient) SearchCode(opt model.SearchCodeOption) (*model.SearchCodeResult, error) {
	if opt.ResultPerpage == 0 {
		opt.ResultPerpage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}

	ctx := context.Background()
	opts := &github.SearchOptions{
		Sort:  opt.Sort,
		Order: opt.Order,
		ListOptions: github.ListOptions{
			PerPage: opt.ResultPerpage,
			Page:    opt.Page,
		},
		TextMatch: true, // Enable text matches for highlighting
	}

	result, resp, err := c.c.Search.Code(ctx, opt.Query, opts)
	if err != nil {
		return nil, err
	}

	searchResult := &model.SearchCodeResult{
		TotalCount: result.GetTotal(),
		NextPage:   resp.NextPage,
		LastPage:   resp.LastPage,
		CodeFiles:  make([]model.CodeFileInfo, 0),
	}

	for _, codeResult := range result.CodeResults {
		codeFile := model.CodeFileInfo{
			Name:       codeResult.GetName(),
			Path:       codeResult.GetPath(),
			Repository: codeResult.Repository.GetName(),
			Owner:      codeResult.Repository.GetOwner().GetLogin(),
			HTMLURL:    codeResult.GetHTMLURL(),
		}

		// Process text matches for highlighting
		for _, textMatch := range codeResult.TextMatches {
			match := model.TextMatch{
				Fragment:   textMatch.GetFragment(),
				ObjectType: textMatch.GetObjectType(),
				ObjectURL:  textMatch.GetObjectURL(),
				Property:   textMatch.GetProperty(),
			}

			for _, m := range textMatch.Matches {
				matchDetail := model.MatchDetail{
					Indices: m.Indices,
					Text:    m.GetText(),
				}
				match.Matches = append(match.Matches, matchDetail)
			}

			codeFile.TextMatches = append(codeFile.TextMatches, match)
		}

		searchResult.CodeFiles = append(searchResult.CodeFiles, codeFile)
	}

	return searchResult, nil
}

func (c *GithubClient) ListIssues(opt model.ListIssuesOption) (*model.IssuesListResult, error) {
	if opt.ResultPerpage == 0 {
		opt.ResultPerpage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}
	if opt.State == "" {
		opt.State = "open"
	}

	ctx := context.Background()
	opts := &github.IssueListByRepoOptions{
		State:       opt.State,
		Labels:      opt.Labels,
		Assignee:    opt.Assignee,
		Creator:     opt.Creator,
		Mentioned:   opt.Mentioned,
		Milestone:   opt.Milestone,
		ListOptions: github.ListOptions{
			PerPage: opt.ResultPerpage,
			Page:    opt.Page,
		},
	}

	if opt.Since != "" {
		if sinceTime, err := time.Parse(time.RFC3339, opt.Since); err == nil {
			opts.Since = sinceTime
		}
	}

	issues, resp, err := c.c.Issues.ListByRepo(ctx, opt.Owner, opt.Repository, opts)
	if err != nil {
		return nil, err
	}

	result := &model.IssuesListResult{
		TotalCount: len(issues),
		NextPage:   resp.NextPage,
		LastPage:   resp.LastPage,
		Issues:     make([]model.IssueInfo, 0),
	}

	for _, issue := range issues {
		if issue.IsPullRequest() {
			continue // Skip pull requests
		}

		issueInfo := model.IssueInfo{
			Number:    issue.GetNumber(),
			Title:     issue.GetTitle(),
			State:     issue.GetState(),
			Body:      issue.GetBody(),
			Comments:  issue.GetComments(),
			CreatedAt: issue.GetCreatedAt().Format(time.RFC3339),
			UpdatedAt: issue.GetUpdatedAt().Format(time.RFC3339),
			URL:       issue.GetURL(),
			HTMLURL:   issue.GetHTMLURL(),
		}

		if issue.ClosedAt != nil {
			issueInfo.ClosedAt = issue.ClosedAt.Format(time.RFC3339)
		}

		// Process labels
		for _, label := range issue.Labels {
			issueInfo.Labels = append(issueInfo.Labels, model.LabelInfo{
				Name:        label.GetName(),
				Color:       label.GetColor(),
				Description: label.GetDescription(),
			})
		}

		// Process assignee
		if issue.Assignee != nil {
			issueInfo.Assignee = &model.UserInfo{
				Login:     issue.Assignee.GetLogin(),
				ID:        issue.Assignee.GetID(),
				AvatarURL: issue.Assignee.GetAvatarURL(),
				HTMLURL:   issue.Assignee.GetHTMLURL(),
				Type:      issue.Assignee.GetType(),
			}
		}

		// Process assignees
		for _, assignee := range issue.Assignees {
			issueInfo.Assignees = append(issueInfo.Assignees, model.UserInfo{
				Login:     assignee.GetLogin(),
				ID:        assignee.GetID(),
				AvatarURL: assignee.GetAvatarURL(),
				HTMLURL:   assignee.GetHTMLURL(),
				Type:      assignee.GetType(),
			})
		}

		// Process creator
		if issue.User != nil {
			issueInfo.Creator = model.UserInfo{
				Login:     issue.User.GetLogin(),
				ID:        issue.User.GetID(),
				AvatarURL: issue.User.GetAvatarURL(),
				HTMLURL:   issue.User.GetHTMLURL(),
				Type:      issue.User.GetType(),
			}
		}

		// Process milestone
		if issue.Milestone != nil {
			issueInfo.Milestone = &model.MilestoneInfo{
				Number:      issue.Milestone.GetNumber(),
				Title:       issue.Milestone.GetTitle(),
				Description: issue.Milestone.GetDescription(),
				State:       issue.Milestone.GetState(),
			}
			if issue.Milestone.DueOn != nil {
				issueInfo.Milestone.DueOn = issue.Milestone.DueOn.Format(time.RFC3339)
			}
		}

		result.Issues = append(result.Issues, issueInfo)
	}

	return result, nil
}

func (c *GithubClient) SearchIssues(opt model.SearchIssuesOption) (*model.IssuesListResult, error) {
	if opt.ResultPerpage == 0 {
		opt.ResultPerpage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}

	ctx := context.Background()
	opts := &github.SearchOptions{
		Sort:  opt.Sort,
		Order: opt.Order,
		ListOptions: github.ListOptions{
			PerPage: opt.ResultPerpage,
			Page:    opt.Page,
		},
	}

	result, resp, err := c.c.Search.Issues(ctx, opt.Query, opts)
	if err != nil {
		return nil, err
	}

	searchResult := &model.IssuesListResult{
		TotalCount: result.GetTotal(),
		NextPage:   resp.NextPage,
		LastPage:   resp.LastPage,
		Issues:     make([]model.IssueInfo, 0),
	}

	for _, issue := range result.Issues {
		if issue.IsPullRequest() {
			continue // Skip pull requests
		}

		issueInfo := model.IssueInfo{
			Number:    issue.GetNumber(),
			Title:     issue.GetTitle(),
			State:     issue.GetState(),
			Body:      issue.GetBody(),
			Comments:  issue.GetComments(),
			CreatedAt: issue.GetCreatedAt().Format(time.RFC3339),
			UpdatedAt: issue.GetUpdatedAt().Format(time.RFC3339),
			URL:       issue.GetURL(),
			HTMLURL:   issue.GetHTMLURL(),
		}

		if issue.ClosedAt != nil {
			issueInfo.ClosedAt = issue.ClosedAt.Format(time.RFC3339)
		}

		// Process labels
		for _, label := range issue.Labels {
			issueInfo.Labels = append(issueInfo.Labels, model.LabelInfo{
				Name:        label.GetName(),
				Color:       label.GetColor(),
				Description: label.GetDescription(),
			})
		}

		// Process creator
		if issue.User != nil {
			issueInfo.Creator = model.UserInfo{
				Login:     issue.User.GetLogin(),
				ID:        issue.User.GetID(),
				AvatarURL: issue.User.GetAvatarURL(),
				HTMLURL:   issue.User.GetHTMLURL(),
				Type:      issue.User.GetType(),
			}
		}

		searchResult.Issues = append(searchResult.Issues, issueInfo)
	}

	return searchResult, nil
}

func (c *GithubClient) ListIssueComments(opt model.ListIssueCommentsOption) ([]model.IssueCommentInfo, error) {
	if opt.ResultPerpage == 0 {
		opt.ResultPerpage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}

	ctx := context.Background()
	opts := &github.IssueListCommentsOptions{
		ListOptions: github.ListOptions{
			PerPage: opt.ResultPerpage,
			Page:    opt.Page,
		},
	}

	if opt.Since != "" {
		if sinceTime, err := time.Parse(time.RFC3339, opt.Since); err == nil {
			opts.Since = &sinceTime
		}
	}

	comments, _, err := c.c.Issues.ListComments(ctx, opt.Owner, opt.Repository, opt.IssueNumber, opts)
	if err != nil {
		return nil, err
	}

	var commentInfos []model.IssueCommentInfo
	for _, comment := range comments {
		commentInfo := model.IssueCommentInfo{
			ID:        comment.GetID(),
			Body:      comment.GetBody(),
			CreatedAt: comment.GetCreatedAt().Format(time.RFC3339),
			UpdatedAt: comment.GetUpdatedAt().Format(time.RFC3339),
			URL:       comment.GetURL(),
			HTMLURL:   comment.GetHTMLURL(),
		}

		if comment.User != nil {
			commentInfo.User = model.UserInfo{
				Login:     comment.User.GetLogin(),
				ID:        comment.User.GetID(),
				AvatarURL: comment.User.GetAvatarURL(),
				HTMLURL:   comment.User.GetHTMLURL(),
				Type:      comment.User.GetType(),
			}
		}

		commentInfos = append(commentInfos, commentInfo)
	}

	return commentInfos, nil
}

func (c *GithubClient) ListIssueLabels(opt model.ListIssueLabelsOption) ([]model.LabelInfo, error) {
	if opt.ResultPerpage == 0 {
		opt.ResultPerpage = 10
	}
	if opt.Page == 0 {
		opt.Page = 1
	}

	ctx := context.Background()
	opts := &github.ListOptions{
		PerPage: opt.ResultPerpage,
		Page:    opt.Page,
	}

	labels, _, err := c.c.Issues.ListLabels(ctx, opt.Owner, opt.Repository, opts)
	if err != nil {
		return nil, err
	}

	var labelInfos []model.LabelInfo
	for _, label := range labels {
		labelInfo := model.LabelInfo{
			Name:        label.GetName(),
			Color:       label.GetColor(),
			Description: label.GetDescription(),
		}
		labelInfos = append(labelInfos, labelInfo)
	}

	return labelInfos, nil
}
