// Command verifypr provides a tool to verify omni PRs against the conventional commit template.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	cc "github.com/leodido/go-conventionalcommits"
	"github.com/leodido/go-conventionalcommits/parser"
)

var (
	descRegex = regexp.MustCompile(`^[a-z][-\w\s]+$`) // e.g. "add foo-bar"

	scopeRegex = regexp.MustCompile(`^[*\w]+(/[*\w]+)?$`) // e.g. "*" or "foo" or "foo/bar"
)

// run runs the verification.
func run() error {
	pr, err := prFromEnv()
	if err != nil {
		return err
	}

	// Skip dependabot PRs.
	if strings.Contains(pr.Title, "deps") && strings.Contains(pr.Body, "dependabot") {
		return nil
	}

	log.Printf("Verifying omni PR against template\n")
	log.Printf("PR Title: %s\n", pr.Title)
	log.Printf("## PR Body:\n%s\n####\n", pr.Body)

	// Convert PR title and body to conventional commit message.
	commitMsg := fmt.Sprintf("%s\n\n%s", pr.Title, pr.Body)

	return verify(commitMsg)
}

type PR struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	ID    string `json:"node_id"`
}

// prFromEnv returns the PR by parsing it from "GITHUB_PR" env var or an error.
func prFromEnv() (PR, error) {
	const prEnv = "GITHUB_PR"
	prJSON, ok := os.LookupEnv(prEnv)
	if !ok || strings.TrimSpace(prJSON) == "" {
		return PR{}, errors.New("env variable not set")
	}

	var pr PR
	if err := json.Unmarshal([]byte(prJSON), &pr); err != nil {
		return PR{}, err
	}

	if pr.Title == "" || pr.Body == "" || pr.ID == "" {
		return PR{}, errors.New("pr field not set")
	}

	return pr, nil
}

// verify returns an error if the commit message doesn't correspond to the omni conventional commit template.
func verify(commitMsg string) error {
	// Fix line endings, since conventional commit parser doesn't support CRLF.
	commitMsg = strings.ReplaceAll(commitMsg, "\r\n", "\n")

	// Parse conventional commit message.
	m := parser.NewMachine()
	m.WithTypes(cc.TypesConventional)
	msg, err := m.Parse([]byte(commitMsg))
	if err != nil {
		return fmt.Errorf("parse conventional commit message: %w", err)
	}

	commit, ok := msg.(*cc.ConventionalCommit)
	if !ok {
		return errors.New("message is not a conventional commit")
	}

	// Verify conventional commit message is valid.
	if !commit.Ok() {
		return errors.New("conventional commit not ok")
	}

	// Verify title is valid.
	if err := verifyDescription(commit.Description); err != nil {
		return err
	}

	// Verify body is non-empty.
	if commit.Body == nil || *commit.Body == "" {
		return errors.New("body empty")
	}

	// Verify footer is valid.
	if err := verifyFooter(commit); err != nil {
		return err
	}

	// Verify scope is valid.
	if err := verifyScope(commit); err != nil {
		return err
	}

	return nil
}

func verifyDescription(description string) error {
	const maxLen = 50
	if len(description) > maxLen {
		return errors.New("description too long")
	}

	if !descRegex.MatchString(description) {
		return errors.New("description doesn't match regex")
	}

	return nil
}

func verifyFooter(commit *cc.ConventionalCommit) error {
	if len(commit.Footers) == 0 {
		return errors.New("missing `task` footer")
	}
	if len(commit.Footers) > 1 {
		return errors.New("invalid number of footers, only `task` required")
	}

	if len(commit.Footers["task"]) != 1 {
		return errors.New("invalid number of task footers, only one required")
	}

	task := commit.Footers["task"][0]
	if task == "" {
		return errors.New("task footer empty")
	} else if task == "none" {
		// None is also fine
	} else if !strings.HasPrefix(task, "https://app.asana.com/") {
		return errors.New("only asana tasks allowed")
	}

	return nil
}

func verifyScope(commit *cc.ConventionalCommit) error {
	if commit.Scope == nil {
		return errors.New("scope not set")
	}

	scope := *commit.Scope

	if scope == "" {
		return errors.New("scope empty")
	}

	if !scopeRegex.MatchString(scope) {
		return errors.New("scope doesn't match regex")
	}

	return nil
}
