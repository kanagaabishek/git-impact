# git-impact

## Demonstration
**Git Impact** is a Git subcommand that analyzes code changes and tells you what those changes mean, not just what changed.

## It helps answer questions like:

Is this change risky?

What parts of the system are affected?

Should this be merged?
 
<img width="1625" height="384" alt="image" src="https://github.com/user-attachments/assets/8f304a95-0978-4705-8e8f-45862ca7fcd0" />


## Analyze a Specific Commit
```bash
git impact --commit a8f3d2
```

**What it does:**

Runs impact analysis on a single commit

Useful for reviewing historical changes

<img width="1384" height="372" alt="image" src="https://github.com/user-attachments/assets/22dc9ee1-f843-4f5f-ab7b-80fb80752e9d" />

## Branch / PR Analysis
```bash
git impact main..feature/auth
```

What it does:

Analyzes what would change if feature/auth is merged into main

Equivalent to a pull-request risk analysis

Example:

Overall Risk: 🔴 HIGH
• Authentication-related code modified
• Database schema changed without migration

## Installation

### Option 1: Download Binary (Recommended)

Download the binary for your platform from the Releases page for the `v0.1.0` tag.

**Linux:**
```bash
curl -L https://github.com/kanagaabishek/git-impact/releases/download/v0.1.0/git-impact-linux-amd64 -o git-impact
chmod +x git-impact
sudo mv git-impact /usr/local/bin/
```

**macOS (Intel):**
```bash
curl -L https://github.com/kanagaabishek/git-impact/releases/download/v0.1.0/git-impact-darwin-amd64 -o git-impact
chmod +x git-impact
sudo mv git-impact /usr/local/bin/
```

**macOS (Apple Silicon):**
```bash
curl -L https://github.com/kanagaabishek/git-impact/releases/download/v0.1.0/git-impact-darwin-arm64 -o git-impact
chmod +x git-impact
sudo mv git-impact /usr/local/bin/
```

**Windows:**
Download `git-impact-windows-amd64.exe` from releases and add it to your PATH.

### Option 2: Build from Source

```bash
git clone https://github.com/kanagaabishek/git-impact.git
cd git-impact
go build -o git-impact .
sudo mv git-impact /usr/local/bin/
```

## Description

`git-impact` is a small CLI tool that analyzes git diffs and gives a concise assessment of the potential impact of changes. It is intended to help developers, reviewers, and CI pipelines quickly identify risky changes before they reach production.

Key features:

- **Risk Assessment** — LOW, MEDIUM, or HIGH based on heuristics that detect schema changes, secrets, auth/security changes, API surface changes, and diff size.
- **Detailed Reasons** — A short list of findings explaining why a change raised the risk score.
- **File Classification** — Labels changed files by category (DATABASE, API, AUTH, etc.) to aid reviewers.
- **Migration Awareness** — Distinguishes schema changes that include migration files from those that don't.

## Usage

Analyze current working changes:
```bash
git impact
```

Analyze a specific commit:
```bash
git impact --commit <sha>
```

JSON output (for CI):
```bash
git impact --json
```

Show help:
```bash
git impact --help
```

## Examples

```bash
# Check your working changes before committing
git impact

# Analyze a specific commit
git impact --commit 0f671ce

# Machine-readable output for CI
git impact --json
```

## Development

Run from source:
```bash
go run main.go
```

Build locally:
```bash
go build -o git-impact .
```

Cross-compile binaries (no Docker required):
```bash
GOOS=linux   GOARCH=amd64   go build -o releases/git-impact-linux-amd64 .
GOOS=darwin  GOARCH=amd64   go build -o releases/git-impact-darwin-amd64 .
GOOS=darwin  GOARCH=arm64   go build -o releases/git-impact-darwin-arm64 .
GOOS=windows GOARCH=amd64   go build -o releases/git-impact-windows-amd64.exe .
```
