package main

import (
	"log"
	"regexp"
)

func filterDuplicatePullRequests(
	pullRequests *[]pullRequest,
) ([]pullRequest, error) {
	re := regexp.MustCompile(`[bB]ump ([a-zA-Z@/-]+)`)
	packageMap := make(map[string][]pullRequest)
	for _, pr := range *pullRequests {
		packageInfo := re.FindStringSubmatch(pr.title)
		if len(packageInfo) > 1 {
			packageName := packageInfo[1] + pr.repository
			packageMap[packageName] = append(packageMap[packageName], pr)
		} else {
			log.Printf("Failed to find package info for %s", pr.title)
		}
	}
	filteredPullRequests := []pullRequest{}
	for _, prs := range packageMap {
		if len(prs) > 1 {
			filteredPullRequests = append(filteredPullRequests, prs...)
		}
	}
	return filteredPullRequests, nil
}
