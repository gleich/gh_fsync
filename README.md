<!-- DO NOT REMOVE - contributor_list:data:start:["Matt-Gleich"]:end -->

# gh_fsync

🔄 GitHub action to sync files across repos in GitHub

![build](https://github.com/Matt-Gleich/gh_fsync/workflows/build/badge.svg)
![test](https://github.com/Matt-Gleich/gh_fsync/workflows/test/badge.svg)
![lint](https://github.com/Matt-Gleich/gh_fsync/workflows/lint/badge.svg)
![release](https://github.com/Matt-Gleich/gh_fsync/workflows/release/badge.svg)

## ❓ What is gh_fsync

gh_fsync is a simple, configurable, and blazing fast way to sync files in your repository with files from another repository.

## ⚙️ Configuration

First create a file in one of the following locations inside of your repository:

- `/fsync.yml`
- `/fsync.yaml`
- `/.fsync.yml`
- `/.fsync.yaml`
- `/.github/fsync.yml`
- `/.github/fsync.yml`

Now that you have your file lets go over the syntax.

### 🌍 Global replace

Replace certain text for all source files listed under the files section. Below is an example:

```yaml
replace:
  - before: project_name
    after: gh_fsync
```

### 📁 Files

List all of the files you want to sync. `path:` is the file name in your repo and `source` is the URL on GitHub for the file. Below is an example:

```yaml
files:
  - path: CONTRIBUTING.md
    source: https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md
```

You can even replace values specific to a file. Below is an example:

```yaml
files:
  - path: CONTRIBUTING.md
    source: https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md
    replace:
      - before: project_name
        after: gh_fsync2
```

A file specific replace will override any global replace with the name `before`. So in the case shown below the replace of `project_name` for the `CONTRIBUTING.md` file will overrule the global replace defined above:

```yaml
replace:
  - before: project_name
    after: gh_fsync
files:
  - path: CONTRIBUTING.md
    source: https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md
    replace:
      - before: project_name
        after: gh_fsync2
```

### 💬 Commit message

Define the commit message to use when updating the files. The default commit message is `Update via sync`. Below is an example

```yaml
commit_message: 🔄 Update file via sync
```

### ✨ Example

```yaml
commit_message: 🔄 Update file via sync
replace:
  - before: project_name
    after: gh_fsync
files:
  - path: CONTRIBUTING.md
    source: https://github.com/Matt-Gleich/go_template/blob/master/CONTRIBUTING.md
    replace:
      - before: project_name
        after: gh_fsync2
  - path: LICENSE.md
    source: https://github.com/Matt-Gleich/go_template/blob/master/LICENSE.md
    replace:
      - before: author_name
        after: Matthew Gleich
```

### 🤖 GitHub action

Use the following for the GitHub action:

```yaml
path: gh_fsync

on:
  push:
  schedule:
    - cron: '*/10 * * * *' # Runs every 10 minutes

jobs:
  file_sync:
    runs-on: ubuntu-latest
    steps:
      - path: Checkout Repository
        uses: actions/checkout@v2
      - path: gh_fsync
        uses: Matt-Gleich/gh_fsync@master
```

## 🙌 Contributing

Before contributing please read the [CONTRIBUTING.md file](https://github.com/Matt-Gleich/gh_fsync/blob/master/CONTRIBUTING.md)

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@Matt-Gleich](https://github.com/Matt-Gleich)**

<!-- DO NOT REMOVE - contributor_list:end -->
