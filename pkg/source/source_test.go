package source

import (
	"testing"

	"github.com/gleich/gh_fsync/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	instance1 := replace(
		"# Contributing\n\n👋 Welcome to project_name! Thank you for showing interest in contributing to project_name, we would love to have your contribution. Below are some requirements for contributing. Please read carefully!\n\n## 🐛 Requesting Features/Reporting Bugs\n\n1. Click on the \"Issues\" tab in the repo.\n2. Make sure that the issue does exist already by searching for it.\n3. Pick the issue template.\n4. Fill in the issue template.\n\n## ➕ Adding/Changing code\n\n### ⚠️ Notice\n\nThis project uses [golangci-lint](https://github.com/golangci/golangci-lint) for code linting, please install it and format your code with `make lint-golangci`\n\n### 🧾 Process\n\n1. Make an issue (see above) and check to make sure what you are trying to add/change doesn't already exist.\n2. Create a branch with the name being the issue number. If you don't have contributor access just fork the repo.\n3. Add code.\n4. Validate code. See checking code section below 👇.\n5. Make the pull request!\n6. Now someone on the team will review your PR. Congrats!\n7. **Warning** once your PR gets merged the branch for it will automatically get deleted (only for contributors with contributor access).\n\n### ✅ Checking Code\n\n#### 🐳 Docker Container\n\nYou can check all the code inside of a docker container with all the dependencies installed by running `make docker-lint` and `make test-in-docker`. Both of these commands will build the image for you and run it. No need to install anything! Check the output to make sure you don't have any issues to resolve.\n\n#### ✍️ Manually\n\n| **Program**                                                | **Description:**                   |\n| ---------------------------------------------------------- | ---------------------------------- |\n| [golangci-lint](https://github.com/golangci/golangci-lint) | Linter for all golang files        |\n| [goreleaser](https://github.com/goreleaser/goreleaser)     | Release automation for the program |\n| [hadolint](https://github.com/hadolint/hadolint)           | Linter for all Dockerfiles         |\n\nOnce you have those installed please run `make local-test` and `make local-lint`. If you don't get any errors your all set!\n\n## ℹ️ General\n\n- When you take on an issue please set yourself as the assignee.\n",
		[]config.ReplaceOutline{{Before: "project_name", After: "Hello"}},
		[]config.ReplaceOutline{{Before: "project_name", After: "Hello World!"}},
		false,
	)
	assert.Equal(
		t,
		"# Contributing\n\n👋 Welcome to Hello World!! Thank you for showing interest in contributing to Hello World!, we would love to have your contribution. Below are some requirements for contributing. Please read carefully!\n\n## 🐛 Requesting Features/Reporting Bugs\n\n1. Click on the \"Issues\" tab in the repo.\n2. Make sure that the issue does exist already by searching for it.\n3. Pick the issue template.\n4. Fill in the issue template.\n\n## ➕ Adding/Changing code\n\n### ⚠️ Notice\n\nThis project uses [golangci-lint](https://github.com/golangci/golangci-lint) for code linting, please install it and format your code with `make lint-golangci`\n\n### 🧾 Process\n\n1. Make an issue (see above) and check to make sure what you are trying to add/change doesn't already exist.\n2. Create a branch with the name being the issue number. If you don't have contributor access just fork the repo.\n3. Add code.\n4. Validate code. See checking code section below 👇.\n5. Make the pull request!\n6. Now someone on the team will review your PR. Congrats!\n7. **Warning** once your PR gets merged the branch for it will automatically get deleted (only for contributors with contributor access).\n\n### ✅ Checking Code\n\n#### 🐳 Docker Container\n\nYou can check all the code inside of a docker container with all the dependencies installed by running `make docker-lint` and `make test-in-docker`. Both of these commands will build the image for you and run it. No need to install anything! Check the output to make sure you don't have any issues to resolve.\n\n#### ✍️ Manually\n\n| **Program**                                                | **Description:**                   |\n| ---------------------------------------------------------- | ---------------------------------- |\n| [golangci-lint](https://github.com/golangci/golangci-lint) | Linter for all golang files        |\n| [goreleaser](https://github.com/goreleaser/goreleaser)     | Release automation for the program |\n| [hadolint](https://github.com/hadolint/hadolint)           | Linter for all Dockerfiles         |\n\nOnce you have those installed please run `make local-test` and `make local-lint`. If you don't get any errors your all set!\n\n## ℹ️ General\n\n- When you take on an issue please set yourself as the assignee.\n",
		instance1,
	)
	instance2 := replace(
		"# Contributing\n\n👋 Welcome to project_name! Thank you for showing interest in contributing to project_name, we would love to have your contribution. Below are some requirements for contributing. Please read carefully!\n\n## 🐛 Requesting Features/Reporting Bugs\n\n1. Click on the \"Issues\" tab in the repo.\n2. Make sure that the issue does exist already by searching for it.\n3. Pick the issue template.\n4. Fill in the issue template.\n\n## ➕ Adding/Changing code\n\n### ⚠️ Notice\n\nThis project uses [golangci-lint](https://github.com/golangci/golangci-lint) for code linting, please install it and format your code with `make lint-golangci`\n\n### 🧾 Process\n\n1. Make an issue (see above) and check to make sure what you are trying to add/change doesn't already exist.\n2. Create a branch with the name being the issue number. If you don't have contributor access just fork the repo.\n3. Add code.\n4. Validate code. See checking code section below 👇.\n5. Make the pull request!\n6. Now someone on the team will review your PR. Congrats!\n7. **Warning** once your PR gets merged the branch for it will automatically get deleted (only for contributors with contributor access).\n\n### ✅ Checking Code\n\n#### 🐳 Docker Container\n\nYou can check all the code inside of a docker container with all the dependencies installed by running `make docker-lint` and `make test-in-docker`. Both of these commands will build the image for you and run it. No need to install anything! Check the output to make sure you don't have any issues to resolve.\n\n#### ✍️ Manually\n\n| **Program**                                                | **Description:**                   |\n| ---------------------------------------------------------- | ---------------------------------- |\n| [golangci-lint](https://github.com/golangci/golangci-lint) | Linter for all golang files        |\n| [goreleaser](https://github.com/goreleaser/goreleaser)     | Release automation for the program |\n| [hadolint](https://github.com/hadolint/hadolint)           | Linter for all Dockerfiles         |\n\nOnce you have those installed please run `make local-test` and `make local-lint`. If you don't get any errors your all set!\n\n## ℹ️ General\n\n- When you take on an issue please set yourself as the assignee.\n",
		[]config.ReplaceOutline{{Before: "project_name", After: "Hello"}},
		[]config.ReplaceOutline{{Before: "project_name", After: "Hello World!"}},
		true,
	)
	assert.Equal(
		t,
		"# Contributing\n\n👋 Welcome to Hello World!! Thank you for showing interest in contributing to Hello World!, we would love to have your contribution. Below are some requirements for contributing. Please read carefully!\n\n## 🐛 Requesting Features/Reporting Bugs\n\n1. Click on the \"Issues\" tab in the repo.\n2. Make sure that the issue does exist already by searching for it.\n3. Pick the issue template.\n4. Fill in the issue template.\n\n## ➕ Adding/Changing code\n\n### ⚠️ Notice\n\nThis project uses [golangci-lint](https://github.com/golangci/golangci-lint) for code linting, please install it and format your code with `make lint-golangci`\n\n### 🧾 Process\n\n1. Make an issue (see above) and check to make sure what you are trying to add/change doesn't already exist.\n2. Create a branch with the name being the issue number. If you don't have contributor access just fork the repo.\n3. Add code.\n4. Validate code. See checking code section below 👇.\n5. Make the pull request!\n6. Now someone on the team will review your PR. Congrats!\n7. **Warning** once your PR gets merged the branch for it will automatically get deleted (only for contributors with contributor access).\n\n### ✅ Checking Code\n\n#### 🐳 Docker Container\n\nYou can check all the code inside of a docker container with all the dependencies installed by running `make docker-lint` and `make test-in-docker`. Both of these commands will build the image for you and run it. No need to install anything! Check the output to make sure you don't have any issues to resolve.\n\n#### ✍️ Manually\n\n| **Program**                                                | **Description:**                   |\n| ---------------------------------------------------------- | ---------------------------------- |\n| [golangci-lint](https://github.com/golangci/golangci-lint) | Linter for all golang files        |\n| [goreleaser](https://github.com/goreleaser/goreleaser)     | Release automation for the program |\n| [hadolint](https://github.com/hadolint/hadolint)           | Linter for all Dockerfiles         |\n\nOnce you have those installed please run `make local-test` and `make local-lint`. If you don't get any errors your all set!\n\n## ℹ️ General\n\n- When you take on an issue please set yourself as the assignee.\n",
		instance2,
	)
}

func TestGetContent(t *testing.T) {
	instance1 := getSourceContent(
		"https://raw.githubusercontent.com/gleich/jsx/master/public/robots.txt",
	)
	assert.Equal(
		t,
		"# https://www.robotstxt.org/robotstxt.html\nUser-agent: *\nDisallow:\n",
		instance1,
	)
	instance2 := getSourceContent(
		"https://raw.githubusercontent.com/gleich/CongressPresenation/master/.metadata",
	)
	assert.Equal(
		t,
		"# This file tracks properties of this Flutter project.\n# Used by Flutter tool to assess capabilities and perform upgrades etc.\n#\n# This file should be version controlled and should not be manually edited.\n\nversion:\n  revision: 8735ab1e35346ae20b6c347d259b07b1589756a5\n  channel: master\n\nproject_type: app\n",
		instance2,
	)
}

func TestRawURL(t *testing.T) {
	instance1 := rawURL("https://github.com/gleich/go_template/blob/master/CONTRIBUTING.md")
	assert.Equal(
		t,
		"https://raw.githubusercontent.com/gleich/go_template/master/CONTRIBUTING.md",
		instance1,
	)
	instance2 := rawURL("https://github.com/gleich/Dot-Files/blob/master/LICENSE.md")
	assert.Equal(
		t,
		"https://raw.githubusercontent.com/gleich/Dot-Files/master/LICENSE.md",
		instance2,
	)
}
