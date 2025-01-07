package gogithub

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/go-github/v68/github"
	"golang.org/x/oauth2"
)

func CheckAndUpdateRepository(githubToken string, repositoryOwner string, repositoryName string, branchName string, projectPath string) error {
	// GitHub token for authentication
	client := newGitHubClient(githubToken)

	// Fetch the latest commit hash from the `main` branch on GitHub
	latestCommitHash, err := getLatestCommit(client, repositoryOwner, repositoryName, branchName)
	if err != nil {
		return err
	}
	// fmt.Printf("Latest Commit Hash on GitHub (main): %s\n", latestCommitHash)

	// Get your local commit hash (the latest commit on your current local branch)
	localCommitHash, err := getLocalCommitHash(projectPath)
	if err != nil {
		return err
	}
	// fmt.Printf("Local Commit Hash: %s\n", localCommitHash)

	// Compare the commit hashes
	if localCommitHash != latestCommitHash {
		// If hashes are different, pull the latest changes from the remote main branch
		err := gitPull(projectPath, branchName)
		if err != nil {
			return err
		}
		// fmt.Println("Pulled latest changes from the 'main' branch!")
	}
	//  else {
	// 	fmt.Println("Your local commit is up to date with the remote 'main' branch.")
	// }

	return nil
}

// Create and return a new GitHub client
func newGitHubClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc)
}

// Get the latest commit hash from a specific branch on GitHub
func getLatestCommit(client *github.Client, owner, repo, branch string) (string, error) {
	// Get the reference for the latest commit on the specified branch
	ref, _, err := client.Git.GetRef(context.Background(), owner, repo, "heads/"+branch)
	if err != nil {
		return "", err
	}

	// Return the commit hash
	return ref.GetObject().GetSHA(), nil
}

// Get the latest commit hash from your local repository
func getLocalCommitHash(projectPath string) (string, error) {
	// Run `git rev-parse HEAD` to get the local commit hash in the specified project path
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = projectPath // Set the working directory to your project path
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// Pull the latest changes from the remote main branch
func gitPull(projectPath string, branch string) error {
	// Run `git pull origin main` to pull the latest changes

	// Change the working directory to the project path
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		return err
	}

	err = os.Chdir(absPath)
	if err != nil {
		return err
	}

	// Run `git pull origin main` to pull the latest changes
	cmd := exec.Command("git", "pull", "origin", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
