package main

import (
	"context"
	"fmt"

	"github.com/Felamande/githubMcp/client"
	"github.com/Felamande/githubMcp/model"
	"github.com/google/go-github/v74/github" // 使用合适的版本号
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

func listAllReleases(ctx context.Context, client *github.Client, owner, repo string) error {
	// 设置分页选项
	opts := &github.ListOptions{
		PerPage: 10, // 每页显示 10 个结果
		Page:    1,  // 第 1 页
	}

	// 获取 releases 列表
	releases, resp, err := client.Repositories.ListReleases(ctx, owner, repo, opts)
	if err != nil {
		return fmt.Errorf("获取 releases 列表失败: %w", err)
	}

	fmt.Printf("找到 %d 个 releases\n", len(releases))
	fmt.Printf("下页: %d, 最后一页: %d\n", resp.NextPage, resp.LastPage)
	fmt.Println("=====================================")

	for i, release := range releases {

		fmt.Printf("%d. %s\n", i+1, release.GetName())
		fmt.Printf("   标签: %s\n", release.GetTagName())
		fmt.Printf("   作者: %s\n", release.GetAuthor().GetLogin())
		fmt.Printf("   是否草稿: %v\n", release.GetDraft())
		fmt.Printf("   是否预发布: %v\n", release.GetPrerelease())
		fmt.Printf("   创建时间: %s\n", release.GetCreatedAt().Format("2006-01-02 15:04:05"))
		fmt.Printf("   发布时间: %s\n", release.GetPublishedAt().Format("2006-01-02 15:04:05"))
		fmt.Printf("   下载地址: %s\n", release.GetHTMLURL())
		fmt.Printf("   资源数量: %d\n", len(release.Assets))
		fmt.Println("-------------------------------------")
	}

	return nil
}

// 获取最新 release
func getLatestRelease(ctx context.Context, client *github.Client, owner, repo string) error {
	release, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return fmt.Errorf("获取最新 release 失败: %w", err)
	}

	fmt.Println("最新 Release 信息:")
	fmt.Printf("  名称: %s\n", release.GetName())
	fmt.Printf("  标签: %s\n", release.GetTagName())
	fmt.Printf("  作者: %s\n", release.GetAuthor().GetLogin())
	fmt.Printf("  发布时间: %s\n", release.GetPublishedAt().Format("2006-01-02 15:04:05"))
	fmt.Printf("  描述: %s\n", truncateString(release.GetBody(), 100))
	fmt.Printf("  下载地址: %s\n", release.GetHTMLURL())

	return nil
}

// 根据 ID 获取特定 release
func getReleaseByID(ctx context.Context, client *github.Client, owner, repo string, releaseID int64) error {
	release, _, err := client.Repositories.GetRelease(ctx, owner, repo, releaseID)
	if err != nil {
		return fmt.Errorf("获取 ID 为 %d 的 release 失败: %w", releaseID, err)
	}

	fmt.Printf("Release ID: %d\n", release.GetID())
	fmt.Printf("  名称: %s\n", release.GetName())
	fmt.Printf("  标签: %s\n", release.GetTagName())
	fmt.Printf("  发布时间: %s\n", release.GetPublishedAt().Format("2006-01-02 15:04:05"))

	return nil
}

// 根据标签名获取 release
func getReleaseByTag(ctx context.Context, client *github.Client, owner, repo, tag string) error {
	release, _, err := client.Repositories.GetReleaseByTag(ctx, owner, repo, tag)
	if err != nil {
		return fmt.Errorf("获取标签为 %s 的 release 失败: %w", tag, err)
	}

	fmt.Printf("Release 标签: %s\n", release.GetTagName())
	fmt.Printf("  名称: %s\n", release.GetName())
	fmt.Printf("  作者: %s\n", release.GetAuthor().GetLogin())
	fmt.Printf("  发布时间: %s\n", release.GetPublishedAt().Format("2006-01-02 15:04:05"))
	fmt.Printf("  描述: %s\n", truncateString(release.GetBody(), 100))

	return nil
}

// 截断字符串以便显示
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
