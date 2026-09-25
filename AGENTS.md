# AGENTS.md — Development Guidelines for gossm

This document defines the architectural rules, coding standards, testing requirements, and operational workflows for AI agents and human contributors working on `gossm`.

---

## 1. Project Overview & Scope

`gossm` is an interactive CLI tool for AWS Systems Manager (SSM) Session Manager. It enables:
- Interactive shell sessions (`gossm start`)
- SSH and SCP over SSM tunnels without open inbound port 22 (`gossm ssh`, `gossm scp`)
- Multi-server Run Command execution (`gossm cmd`)
- Local and remote port forwarding (`gossm fwd`, `gossm fwdrem`)
- AWS authentication and MFA credential management (`gossm mfa`)

Target platforms: **Darwin (amd64/arm64)**, **Linux (amd64/arm64)**, and **Windows (amd64)**.

---

## 2. Tech Stack & Toolchain Constraints

- **Language & Runtime**: Go 1.23+ (configured in `.tool-versions` and `go.mod`).
- **AWS SDK**: AWS SDK for Go v2 (`github.com/aws/aws-sdk-go-v2`).
- **CLI Framework**: Cobra (`github.com/spf13/cobra`) and Viper (`github.com/spf13/viper`).
- **Testing**: Standard Go `testing` package with `github.com/stretchr/testify` (assert, require).
- **Embedded Binaries**: AWS `session-manager-plugin` binaries embedded via Go `embed` in `internal/assets/`.

---

## 3. Architecture & Code Design Rules

### 3.1 Interface-First AWS Integration (Mandatory)
- **Rule**: Never couple business logic directly to concrete AWS clients (`*ec2.Client`, `*ssm.Client`).
- **Pattern**: Always define minimal, focused interfaces for the AWS operations needed:
  ```go
  type InstanceDescriptor interface {
      DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
  }

  type SessionStarter interface {
      StartSession(ctx context.Context, params *ssm.StartSessionInput, optFns ...func(*ssm.Options)) (*ssm.StartSessionOutput, error)
      TerminateSession(ctx context.Context, params *ssm.TerminateSessionInput, optFns ...func(*ssm.Options)) (*ssm.TerminateSessionOutput, error)
  }
  ```
- **Rationale**: Enables 100% offline mock testing so AI agents and CI can verify changes instantly without live AWS credentials.

### 3.2 Elimination of Global State & Dependency Injection
- **Rule**: Do not create or rely on package-level mutable global variables (e.g., `var _credential *Credential`).
- **Pattern**: Pass dependencies explicitly via constructors or application context structs. Command handlers should receive their dependencies from Cobra's context or an app container.

### 3.3 Error Handling & Logging
- **Rule**: Never call `os.Exit()` or `panic()` from `internal/` or `pkg/` packages.
- **Rule**: `os.Exit()` is only permitted at the CLI entry point (`cmd/root.go` / `main.go`).
- **Pattern**:
  - Always wrap errors with meaningful contextual descriptions: `fmt.Errorf("starting ssm session on %s: %w", targetID, err)`.
  - Use `log/slog` for leveled structured logging (debug, info, warn, error). Do not use `fmt.Println` for diagnostic logs.
  - User-facing status messages (e.g. spinners, success prompts) belong in the presentation/CLI layer, not in core AWS service logic.

### 3.4 Context Propagation & Lifecycle
- **Rule**: Every function performing I/O, process execution, or AWS API requests must accept `ctx context.Context` as its first parameter.
- **Rule**: Respect context cancellation. Always clean up active SSM sessions (`TerminateSession`) using `defer` or cleanup signals to avoid leaking AWS resources.

---

## 4. Testing Guidelines

### 4.1 Zero Live-AWS Dependency in Unit Tests
- **Strict Requirement**: Unit tests must **NEVER** require live AWS credentials (`~/.aws/credentials`), network access, or an active AWS account.
- **Pattern**: Use in-memory mock implementations or test doubles that implement the AWS client interfaces.
- Running `go test ./...` must execute offline, deterministically, and complete in under 5 seconds.

### 4.2 Table-Driven Unit Tests
- Write table-driven unit tests for all domain logic:
  ```go
  func TestFindInstances(t *testing.T) {
      tests := map[string]struct {
          mockSetup func(*mocks.MockEC2Client, *mocks.MockSSMClient)
          expected  []*Target
          expectErr bool
      }{
          "successful fetch with running instances": { ... },
          "handles pagination token correctly":     { ... },
          "handles aws api rate limit error":       { ... },
          "ignores terminated instances":           { ... },
      }
      for name, tc := range tests {
          t.Run(name, func(t *testing.T) {
              // execute test
          })
      }
  }
  ```

### 4.3 Race Detection & Coverage
- Always run tests with `-race`: `go test -race ./...`.
- Ensure new code has test coverage for both happy paths and edge cases (e.g., empty lists, pagination, API errors, invalid user input).

---

## 5. Code Generation & Tooling

### 5.1 Mock Generation
- When updating or adding interfaces, regenerate mocks using:
  ```bash
  go generate ./...
  ```
- Do not manually edit files marked `// Code generated by ...; DO NOT EDIT.`

### 5.2 Session-Manager-Plugin Asset Updates
- The bundled `session-manager-plugin` binaries live in `internal/assets/plugin/<os>_<arch>/`.
- Any updates to these binaries must be automated via update scripts (`scripts/update-plugins.sh` or Go scripts) referencing official AWS checksums.

---

## 6. Standard Agent Workflow Commands

Before submitting or proposing any code changes, verify your work with these commands:

| Task | Command |
| :--- | :--- |
| **Check Go Version** | `go version` (Must be Go 1.23+) |
| **Download Dependencies** | `go mod download && go mod tidy` |
| **Run Code Generation** | `go generate ./...` |
| **Run All Unit Tests** | `go test -race -v ./...` |
| **Run Linting** | `golangci-lint run ./...` |
| **Build Binary** | `go build -o bin/gossm .` |
| **Format Code** | `go fmt ./...` |

---

## 7. Git Workflow & Issue Tracking (Mandatory)

### 7.1 No Direct Commits to `master`
- **Strict Rule**: Never commit, rebase, or push directly to the `master` branch.
- **Branch Pattern**: All changes must be addressed in a dedicated, descriptive branch created from the latest `master`.
  - Feature branches: `feat/<issue-number>-<short-description>` (e.g., `feat/12-upgrade-go-1.23`)
  - Bug fixes: `fix/<issue-number>-<short-description>` (e.g., `fix/15-ssm-session-leak`)
  - Maintenance/Chores: `chore/<issue-number>-<short-description>` (e.g., `chore/18-update-plugins`)

### 7.2 Issue-Driven Development
- **Strict Rule**: Every task, refactor, feature, or bug fix must correspond to an active **GitHub Issue**.
- If an issue does not already exist for the work to be done, one must be created or identified before branching and writing code.
- **Traceability**:
  - Reference the issue number in branch names.
  - Reference the issue number in commit messages (e.g., `feat: upgrade to Go 1.23.2 (Refs #12)`).
  - Include issue closing keywords in pull request descriptions (e.g., `Closes #12`).

---

## 8. Change Philosophy for AI Agents

1. **Be Incremental**: Break large refactors into small, verifiable commits (e.g., upgrade dependencies -> extract interfaces -> add mocks -> refactor commands).
2. **Preserve Compatibility**: Keep CLI flag names and behavioral semantics identical unless explicitly changing them.
3. **No Dead Code**: Remove deprecated code paths, unused variables, and abandoned files cleanly.
4. **Documentation**: Update [README.md](file:///Users/dirk39/Workspace/gossm/README.md) and command help strings whenever flags or behaviors are modified.

