package client

import (
	"context"

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

		repoInfo := model.RepositoryInfo{
			Name:            repo.Name,
			FullName:        repo.FullName,
			MasterBranch:    repo.MasterBranch,
			Description:     repo.Description,
			StargazersCount: repo.StargazersCount,
			ForksCount:      repo.ForksCount,
			Language:        repo.Language,
			Archived:        repo.Archived,
		}
		if repo.Description != nil && len(*repo.Description) > 1024 {
			strShort := (*repo.Description)[:1024] + "......"
			repoInfo.Description = &strShort
		}
		if repo.Owner != nil {
			repoInfo.Owner = repo.Owner.Name
		}
		if repo.Organization != nil {
			repoInfo.Organization = repo.Organization.Company
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
