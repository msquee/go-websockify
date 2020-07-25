
# Contributing Guidelines

### Table of Contents
- [Development dependencies](#development-dependency-requirements)
  - [Required](#required)
  - [Optional but recommended](#optional-but-recommended)
- [Your first contribution](#your-first-contribution)
- [Branch name conventions](#branch-name-conventions)
  - [Acceptable branch names](#acceptable-branch-names)
  - [Unacceptable branch names](#unacceptable-branch-names)
- [Pull requests](#pull-requests)

## Development dependency requirements
`go-websockify` has a few dependencies that are required for the entire development experience.

#### Required
- `go1.14` Go: version `>=` [1.14](https://golang.org/doc/devel/release.html#go1.14)
- `yarn` [Yarn](https://yarnpkg.com/getting-started/install): used for testing a small client side demo available at [`/client`](https://github.com/msquee/go-websockify/tree/master/client)

#### Optional but recommended
- `modd` [Modd](https://github.com/cortesi/modd) development tool, this automatically recompiles and restarts the binary during development.

## Your first contribution
Checkout the repository:
```shell
$ git checkout https://github.com/msquee/go-websockify
$ cd go-websockify
```

Create a branch for your feature or pull request:
```shell
$ cd go-websockify
$ git checkout -b feature/{FEATURE_NAME}
```

## Branch name conventions
Your branch name must have a prefix that conforms with the following matrix:

| Branch prefix           | Description |
| - | - |
| `feature/{BRANCH_NAME}` | New features |
| `update/{BRANCH_NAME}`  | Updates to `README`, `LICENSE`, etc. |
| `fix/{BRANCH_NAME}`     | Bug fixes |

The name of your branch after the branch prefix must by hypenated by each word and contain no capital letters.

#### Acceptable branch names:
- `feature/my-cool-feature`
- `fix/hotfix`
- `update/documentation-for-feature`

#### Unacceptable branch names:
- `feature/MY_COOL-Feature`
- `feature/mycoolfeature`
- `fixToABug`
- `fix-to-a-bug`

## Pull requests
When your branch is ready for a review, open a pull request (PR) to merge your branch into `master`.

#### Your PR description should contain the following:
- **Concise description**. A short but concise description about what your PR accomplishes.
- **How to test your PR**. An example of how to test your feature or fix.
- **Visual elements**. Any relevant screenshots.
