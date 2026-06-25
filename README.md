# GitHub Profile Analyzer — API

Go REST API that fetches and caches GitHub profile data.

---

## What it does

Acts as a thin, cached proxy in front of the GitHub REST API:

- Fetches a user's profile, repositories, and recent push activity
- Computes top repos by stars and a language breakdown across all repos
- Groups commit activity by day for charting
- Caches every response in memory for 5 minutes to avoid hitting GitHub's 60 req/hour unauthenticated rate limit

## Tech stack

- **Go 1.22+** — standard library `net/http`, no framework
- **In-memory cache** — `sync.RWMutex`-protected map with TTL
- **No database** — fully stateless, cache resets on restart by design

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/profile/:username` | User profile — avatar, bio, followers, following, public repos |
| `GET` | `/api/repos/:username` | Top 10 repos by stars + language breakdown across all repos |
| `GET` | `/api/activity/:username` | Push events grouped by day (last ~90 days) |

### Example

```bash
curl https://your-app.railway.app/api/profile/torvalds
```

```json
{
  "login": "torvalds",
  "name": "Linus Torvalds",
  "avatar_url": "https://avatars.githubusercontent.com/u/1024025",
  "followers": 230000,
  "following": 0,
  "public_repos": 8
}
```

## Running locally

```bash
git clone https://github.com/yourusername/github-analyzer-api.git
cd github-analyzer-api
go run main.go
```

Server starts on `:8080`. Test with the included `test.http` file (VS Code REST Client extension) or:

```bash
curl http://localhost:8080/api/profile/torvalds
```

### Environment variables

| Variable | Required | Description |
|----------|----------|--------------|
| `PORT` | No | Server port. Defaults to `8080`. Set automatically by Railway in production. |
| `GITHUB_TOKEN` | No | Personal access token (read-only, public repos). Raises GitHub rate limit from 60/hr to 5000/hr. Generate at GitHub → Settings → Developer settings → Fine-grained tokens. |

## Project structure

```
.
├── main.go              # entry point, routing, server start
├── handlers/             # HTTP handlers — one per endpoint
├── github/               # GitHub API client + response types
├── cache/                # in-memory TTL cache
└── middleware/            # CORS
```


MIT
