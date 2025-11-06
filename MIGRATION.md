# DQD Migration Guide

This document outlines the migration strategy from the Java/Javalin backend to a modern Go backend with a React frontend.

## Overview

DQD is being migrated from a monolithic Java application with a basic HTML frontend to a modern architecture:

- **Old Stack**: Java 17 + Javalin + Single HTML file + Vanilla JS
- **New Stack**: Go + Chi Router + React + TypeScript + Vite

## Architecture Changes

### Before (Java)

```
┌─────────────────────────────────────┐
│   Java Application (Javalin)       │
│   ├── REST API Endpoints            │
│   ├── Business Logic                │
│   ├── HTML Generation               │
│   └── Static HTML (index.html)     │
└─────────────────────────────────────┘
        ↓
    Port 8080
```

### After (Go + React)

```
┌──────────────────────┐    ┌──────────────────────┐
│   React Frontend     │    │   Go Backend         │
│   ├── React Router   │    │   ├── Chi Router     │
│   ├── Components     │ ←→ │   ├── Handlers       │
│   ├── API Client     │    │   ├── Services       │
│   └── Tailwind CSS  │    │   └── Models         │
└──────────────────────┘    └──────────────────────┘
      Port 3000                   Port 8080
```

## Migration Phases

### Phase 1: Foundation (✅ Complete)

**Goal**: Set up new infrastructure alongside existing Java app

**What's Done**:
- ✅ React frontend scaffolded with Vite
- ✅ Go backend initialized with Chi router
- ✅ Basic project structure
- ✅ `/api/about.json` endpoint implemented
- ✅ CORS middleware
- ✅ Health check endpoint
- ✅ Development tooling (Makefile, scripts)

**Result**: Both stacks can run in parallel

### Phase 2: File Upload & Simple Endpoints (🔄 In Progress)

**Goal**: Implement file upload handling and simple analysis

**Tasks**:
- [ ] Go: Multipart form data handling
- [ ] Go: File type detection and validation
- [ ] Go: Archive extraction (ZIP, TAR.GZ)
- [ ] React: Complete Simple Profile Analysis page
- [ ] React: File upload progress indicators
- [ ] React: Error handling and user feedback

**APIs to Implement**:
- `POST /api/simple-profile` (uses existing Java code via FFI or port logic)

### Phase 3: Profile JSON Analysis (📋 Planned)

**Goal**: Port profile.json parsing and analysis

**Tasks**:
- [ ] Go: JSON parsing (use `encoding/json`)
- [ ] Go: Port ProfileJSON DTOs (20+ classes)
- [ ] Go: Profile analysis logic
- [ ] Go: HTML report generation (use `html/template`)
- [ ] React: Detailed profile analysis page
- [ ] React: Visualizations (charts, timelines)

**APIs to Implement**:
- `POST /api/profile` - Detailed analysis
- `POST /api/profiles` - Profile comparison

**Considerations**:
- Java has ~1,814 lines of DTO code
- Complex nested structures
- Consider code generation from JSON schema

### Phase 4: Queries JSON Analysis (📋 Planned)

**Goal**: Port queries.json bulk analysis

**Tasks**:
- [ ] Go: Archive reading (multiple files)
- [ ] Go: Time-window filtering
- [ ] Go: Query pattern detection
- [ ] Go: Aggregation and statistics
- [ ] React: Queries analysis page with filters

**APIs to Implement**:
- `POST /api/queriesjson`

### Phase 5: Schema Generation (📋 Planned)

**Goal**: Port reproduction/schema generation

**Tasks**:
- [ ] Go: Profile parsing for schema extraction
- [ ] Go: SQL DDL generation
- [ ] Go: ZIP file creation
- [ ] Go: YAML parsing for column definitions
- [ ] React: Schema generation form

**APIs to Implement**:
- `POST /api/reproduction`

### Phase 6: System Analysis (📋 Planned)

**Goal**: Port iostat and top analysis

**Tasks**:
- [ ] Go: Text parsing for iostat format
- [ ] Go: Text parsing for top format
- [ ] Go: Statistics calculation
- [ ] React: IOStat visualization
- [ ] React: Top analysis visualization

**APIs to Implement**:
- `POST /api/iostat`
- `POST /api/ttop`

### Phase 7: Production Ready (📋 Planned)

**Goal**: Polish and deployment

**Tasks**:
- [ ] Comprehensive testing
- [ ] Performance optimization
- [ ] Security audit
- [ ] Docker containers
- [ ] CI/CD pipelines
- [ ] Monitoring and logging
- [ ] Documentation
- [ ] Migration guide for users

## Running Both Stacks in Parallel

### Option 1: Side-by-Side (Recommended for Development)

**Java Backend**:
```bash
cd /home/user/dqd
./mvnw clean package
java -jar target/dqd-0.12.3.jar server -p 9090
```

**Go Backend**:
```bash
cd backend-go
PORT=8080 make run
```

**React Frontend**:
```bash
cd frontend
npm run dev  # Runs on port 3000
```

**How it Works**:
- React dev server proxies `/api/*` requests to Go backend (port 8080)
- For unimplemented endpoints, configure frontend to fallback to Java (port 9090)

### Option 2: Proxy Layer

Use nginx or Caddy to route requests:

```nginx
# Implemented in Go
location /api/about.json {
    proxy_pass http://localhost:8080;
}

# Still in Java
location /api/profile {
    proxy_pass http://localhost:9090;
}
```

## API Compatibility

The Go backend maintains **100% API compatibility** with the Java backend:

| Endpoint | Method | Request | Response | Status |
|----------|--------|---------|----------|--------|
| `/api/about.json` | GET | - | JSON | ✅ Go |
| `/api/profile` | POST | multipart | HTML | ❌ Java |
| `/api/simple-profile` | POST | multipart | HTML | ❌ Java |
| `/api/profiles` | POST | multipart | HTML | ❌ Java |
| `/api/queriesjson` | POST | multipart | HTML | ❌ Java |
| `/api/reproduction` | POST | multipart | ZIP | ❌ Java |
| `/api/iostat` | POST | multipart | HTML | ❌ Java |
| `/api/ttop` | POST | multipart | HTML | ❌ Java |

## Testing Strategy

### Unit Tests
- Go: Standard `testing` package
- React: Vitest + React Testing Library

### Integration Tests
- API contract tests
- End-to-end tests with Playwright

### Migration Tests
- Compare outputs from Java vs Go for same inputs
- Regression testing suite

## Performance Targets

| Metric | Java | Go Target | Notes |
|--------|------|-----------|-------|
| Cold Start | ~3s | <1s | Server startup |
| Profile Analysis | ~500ms | <300ms | Single profile.json |
| Memory Usage | ~300MB | <100MB | Idle |
| Binary Size | ~50MB (JAR) | ~15MB | Self-contained |

## Code Organization

### Mapping Java Packages to Go

| Java Package | Go Package | Notes |
|--------------|------------|-------|
| `com.dremio.support.diagnostics` | `internal/` | Root |
| `.cmds` | `cmd/` | Entry points |
| `.server` | `internal/handlers` | HTTP handlers |
| `.shared` | `internal/shared` | Utilities |
| `.shared.dto` | `internal/models` | Data models |
| `.profilejson` | `internal/services/profile` | Business logic |
| `.queriesjson` | `internal/services/queries` | Business logic |
| `.repro` | `internal/services/repro` | Business logic |

## Breaking Changes

### For Users

**None**. The API remains compatible.

### For Developers

- New languages (Go, TypeScript)
- New build tools (Go toolchain, Vite)
- New project structure

## Rollback Plan

If issues arise, the Java backend remains available:

1. Deploy Java backend on port 8080
2. Update frontend proxy to point to Java
3. Investigate and fix Go issues
4. Re-deploy when ready

## Timeline (Estimated)

- **Phase 1**: ✅ Complete (Week 1)
- **Phase 2**: 🔄 Current (Weeks 2-3)
- **Phase 3**: Weeks 4-6
- **Phase 4**: Weeks 7-8
- **Phase 5**: Weeks 9-10
- **Phase 6**: Weeks 11-12
- **Phase 7**: Weeks 13-14

**Total**: ~3-4 months for complete migration

## Success Criteria

- [ ] All endpoints implemented in Go
- [ ] Feature parity with Java version
- [ ] <30% performance improvement
- [ ] <50% memory reduction
- [ ] 100% API compatibility
- [ ] Comprehensive test coverage (>80%)
- [ ] Documentation complete
- [ ] Zero critical bugs

## Questions & Decisions

### Open Questions

1. **Report Format**: Continue generating HTML strings or use templates?
   - **Recommendation**: Use Go templates (`html/template`)

2. **Java Interop**: Call Java code from Go for complex logic?
   - **Recommendation**: Pure Go port (cleaner, faster)

3. **Database**: Add persistent storage?
   - **Recommendation**: Defer to Phase 7

4. **Authentication**: Add user auth?
   - **Recommendation**: Defer to Phase 7

### Decisions Made

- ✅ Use Chi router (lightweight, idiomatic)
- ✅ Use Vite for frontend (fast, modern)
- ✅ Use Tailwind CSS (utility-first, customizable)
- ✅ Maintain API compatibility
- ✅ Run both backends in parallel during migration

## Resources

- [Go Documentation](https://go.dev/doc/)
- [Chi Router](https://go-chi.io/)
- [React Documentation](https://react.dev/)
- [Vite Documentation](https://vitejs.dev/)
- [Original DQD Repo](https://github.com/rsvihladremio/dqd)

## Contributing to Migration

See individual README files:
- `/backend-go/README.md` - Go backend
- `/frontend/README.md` - React frontend

## Support

For questions or issues during migration:
- GitHub Issues: https://github.com/rsvihladremio/dqd/issues
- Internal Slack: #dqd-migration
