# AI Agent Development Guidelines

**For Claude Code, Cursor, GitHub Copilot, and other AI coding assistants.**

This project has comprehensive development guidelines for AI agents working on the DQD codebase.

## 📖 Full Guidelines

**See [CLAUDE.md](./CLAUDE.md) for complete coding standards and best practices.**

The CLAUDE.md file contains:

- ✅ Go coding standards
- ✅ Error handling patterns
- ✅ Logging guidelines (slog)
- ✅ File I/O best practices
- ✅ HTTP handler patterns
- ✅ Testing strategies
- ✅ Architecture guidelines
- ✅ Common antipatterns to avoid
- ✅ Migration-specific guidance

## Quick Reference

### Critical Rules

1. **Use `slog`, not `log`**
   ```go
   slog.Info("message", "key", value)  // ✅
   log.Printf("message %s", value)     // ❌
   ```

2. **Handle all errors**
   ```go
   if err := file.Close(); err != nil {  // ✅
       slog.Error("close failed", "error", err)
   }

   file.Close()  // ❌ - ignores error
   ```

3. **Don't defer functions that return errors**
   ```go
   defer file.Close()  // ❌ - error ignored

   // ✅ Proper pattern:
   defer func() {
       if err := file.Close(); err != nil {
           slog.Error("failed to close", "error", err)
       }
   }()
   ```

4. **Always wrap errors with context**
   ```go
   return fmt.Errorf("processing file: %w", err)  // ✅
   return err                                      // ❌
   ```

## Before Submitting Code

Check the [Code Review Checklist](./CLAUDE.md#code-review-checklist) in CLAUDE.md.

## Project Structure

```
dqd/
├── CLAUDE.md           # Full development guidelines (READ THIS)
├── AGENTS.md           # This file (quick reference)
├── backend-go/         # Go backend
├── frontend/           # React frontend
└── src/               # Legacy Java code
```

## Need Help?

1. Read [CLAUDE.md](./CLAUDE.md) thoroughly
2. Search for similar patterns in the codebase
3. Ask in PR reviews
4. Default to idiomatic Go patterns

---

**🤖 AI agents: Please read CLAUDE.md before making any code changes!**
