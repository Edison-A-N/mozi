---
description: Fast-track code changes for features/bugfixes - assess complexity, list impacts, implement core code only
---

# Smart Code Change Assistant (`/impl`)

> **🎯 Scope: CORE CODE ONLY**
> - ✅ Production code with proper error handling
> - ❌ No tests, docs, or examples (user adds later)
> - This applies throughout execution, even after confirmations

## Task
Assess change complexity, then implement **core code only**. Focus on code quality and reviewability.

## What the AI will do:

1. ✅ **Quick Assessment**: Evaluate complexity based on review difficulty
2. ✅ **Impact Analysis**: List affected files and key functions (for complex changes)
3. ✅ **Clarify When Needed**: Ask about refactoring, breaking changes, or unclear requirements
4. ✅ **Implement Core Code**: Write quality code that solves the problem
5. ❌ **Skip Non-Essentials**: No documentation, tests, or examples

---

## Phase 1: Complexity Assessment

### Complexity Levels (Based on Review Difficulty)

**🟢 Simple** (Easy to review - immediate understanding):
- Single file modification
- <100 lines of changes
- Obvious logic changes
- No API/interface changes
- No dependency changes
- Self-contained change

**🟡 Medium** (Requires careful review - need to think through):
- 2-5 files modified
- 100-300 lines of changes
- Some API/interface changes (backward compatible)
- Logic flow changes across functions
- May add dependencies
- Moderate cognitive load

**🔴 Complex** (Hard to review - need to trace through):
- 5+ files modified
- >300 lines of changes
- Breaking changes to public APIs
- Architecture/design changes
- Multiple modules interaction
- Data migration needed
- High cognitive load - easy to miss issues

### Assessment Output Format

```markdown
## 📊 Change Assessment

**Complexity**: [🟢 Simple / 🟡 Medium / 🔴 Complex]

**Impact Scope**:
- `path/to/file1.go` - `FunctionA()`, `FunctionB()`
- `path/to/file2.go` - `StructX`, `MethodY()`
- `config/settings.json` - configuration changes

**Dependencies**: 
- Add: `github.com/pkg/errors v0.9.1` - for better error handling
- Update: `go.mod` - require Go 1.21+

**Risk Assessment**:
- ⚠️ Breaking change: `UserService.Login()` signature changed
- ⚠️ Data migration: need to update existing records
- ✅ Backward compatible: old clients still work with feature flag

**Estimated Scope**: ~X files, ~Y lines
```

---

## Phase 2: Clarification (Ask Before Proceeding)

### Questions to Ask When Encountering:

**🔄 Refactoring Opportunities**:
```
Found that [module/function] has [issue]. Options:
A. Minimal change - work around the issue (add X lines)
B. Refactor while fixing - cleaner but touches Y files
Which approach do you prefer?
```

**💥 Breaking Changes**:
```
⚠️ This change will break the existing API:
- Old: `DoSomething(a, b)`
- New: `DoSomething(config Config)`

Options:
A. Breaking change - clients need to update
B. Add new method - keep old one deprecated
C. Feature flag - gradual migration

Please confirm approach.
```

**🤔 Multiple Solutions**:
```
Found multiple ways to solve this:

Option A: [approach] - Pros: [...] Cons: [...]
Option B: [approach] - Pros: [...] Cons: [...]

I recommend Option A because [reason]. Proceed?
```

**📈 Complexity Exceeds Expectation**:
```
Initial assessment: Medium
After investigation: Complex

Reason: [explanation]

New scope: [details]

Do you want to:
A. Proceed with full implementation
B. Break into smaller changes
C. Reconsider approach
```

**❓ Unclear Requirements**:
```
Need clarification on:
1. [Question about requirement]
2. [Question about edge case]
3. [Question about desired behavior]

Please provide guidance.
```

---

## Phase 3: Implementation Principles

> 💡 **Remember**: Implement core code only. User adds tests/docs later.

### ✅ DO (Quality is Priority)

**Code Quality**:
- ✅ Proper error handling (return errors, don't ignore)
- ✅ Input validation and boundary checks
- ✅ Follow existing code patterns and style
- ✅ Graceful degradation for failures
- ✅ Resource cleanup (defer close, context cancellation)
- ✅ Concurrency safety if applicable (locks, channels)

**Minimal Comments** (only when necessary):
- ✅ Non-obvious logic or algorithms
- ✅ Important assumptions or constraints
- ✅ Why a particular approach was chosen (if not obvious)
- ✅ TODOs for known limitations
- ❌ Not for self-evident code

**Code Structure**:
- ✅ Keep functions focused and reasonable size
- ✅ Use meaningful variable names
- ✅ Extract complex logic into named functions
- ✅ Maintain consistent naming conventions

**Quality Over Brevity**:
```go
// ❌ BAD: Too brief, ignoring errors
func Process(data string) string {
    result, _ := Parse(data)
    return Transform(result)
}

// ✅ GOOD: Quality matters, handle errors
func Process(data string) (string, error) {
    if data == "" {
        return "", errors.New("empty data")
    }
    
    result, err := Parse(data)
    if err != nil {
        return "", fmt.Errorf("parse failed: %w", err)
    }
    
    transformed := Transform(result)
    return transformed, nil
}
```

### ❌ DON'T (Skip Non-Essentials)

**Documentation**:
- ❌ No README updates
- ❌ No API documentation
- ❌ No usage guides
- ❌ No changelog entries
- ❌ No excessive comments

**Tests**:
- ❌ No unit tests
- ❌ No integration tests
- ❌ No test fixtures
- ❌ No mock implementations

**Examples**:
- ❌ No example code
- ❌ No demo applications
- ❌ No sample configurations

**Over-Engineering**:
- ❌ No premature optimization
- ❌ No unnecessary abstractions
- ❌ No "nice to have" features not in requirements

---

## Phase 4: Execution Strategy

> 💡 **Reminder**: Only implement core code. Skip tests, docs, and examples.

### For 🟢 Simple Changes:
```
✅ Quick assessment
✅ Implement immediately (core code only)
✅ Done
```

### For 🟡 Medium Changes:
```
✅ Assessment with impact scope
✅ Ask if breaking changes or refactoring needed
✅ Implement with quality focus
✅ Provide basic verification guidance
```

### For 🔴 Complex Changes:
```
✅ Detailed assessment
✅ List all affected files and functions
✅ Highlight risks and breaking changes
✅ Ask for confirmation before proceeding
✅ Implement in logical order
✅ Update dependencies if needed
✅ Provide verification guidance
```

---

## Phase 5: Completion Summary

### Output Format

```markdown
## ✅ Implementation Complete

**Changes Made**:
- `file1.go` - Added `FunctionX()`, modified `FunctionY()`
- `file2.go` - Updated struct `Config` with new field
- `go.mod` - Added dependency `pkg/errors`

**Dependencies Updated**:
- Added: `github.com/pkg/errors v0.9.1`
- Updated: Go version requirement to 1.21+

**Breaking Changes** (if any):
- ⚠️ `UserService.Login(user, pass)` → `Login(ctx, creds)`
- Migration: [brief note on how to adapt]

**Verification**:
```bash
# Build and run
go build ./...
./app

# Or specific test
curl http://localhost:8080/api/endpoint
```

**Notes**:
- [Any important considerations]
- [Known limitations or edge cases]
- [Suggestions for testing focus areas]

**Next Steps**: User adds tests and updates documentation as needed.
```

---

## Important Guidelines

### Complexity Assessment:
- Think about **review difficulty**, not just line count
- Consider **cognitive load**: How much do I need to keep in mind?
- Consider **risk**: What's the blast radius if something goes wrong?

### When to Ask:
- ✅ Breaking changes detected
- ✅ Multiple equally valid approaches
- ✅ Refactoring needed
- ✅ Requirements unclear
- ✅ Complexity exceeds initial estimate
- ✅ Security or data safety concerns

### Code Quality Standards:
- Quality is **not negotiable**
- More lines for quality is **acceptable**
- Proper error handling is **required**
- Edge case handling is **required**
- Following project patterns is **required**

### What User Will Do:
- User verifies the code works
- User adds tests if needed
- User writes documentation if needed
- User handles deployment

---

## Usage Examples

### Example 1: Simple Feature

```
User: Add a rate limit to the login endpoint - max 5 attempts per minute

AI: 
## 📊 Change Assessment
**Complexity**: 🟢 Simple
**Impact Scope**: 
- `api/auth_handler.go` - `Login()` method
**Dependencies**: Add `golang.org/x/time/rate`

Proceeding with implementation...

[implements code with proper error handling]

## ✅ Implementation Complete
**Changes Made**: 
- `api/auth_handler.go` - Added rate limiter to Login()
- `go.mod` - Added golang.org/x/time/rate

**Next Steps**: Add tests for rate limiting behavior.
```

### Example 2: Medium Feature with Question

```
User: Add caching for user profiles

AI:
## 📊 Change Assessment
**Complexity**: 🟡 Medium
**Impact Scope**:
- `service/user_service.go` - `GetProfile()`, `UpdateProfile()`
- `cache/redis_client.go` - new cache layer
- `config/config.go` - cache settings

Found that `GetProfile()` is called from 8 places. Options:
A. Add caching in service layer - transparent to callers
B. Add caching in repository layer - more control
Recommend A for simplicity. Proceed?

[waits for confirmation]

User: Yes, option A

AI: 
[implements caching layer with proper error handling]

## ✅ Implementation Complete
**Changes Made**:
- `service/user_service.go` - Added cache layer
- `cache/redis_client.go` - Created Redis client
- `config/config.go` - Added cache configuration

**Next Steps**: Add tests for cache behavior and update docs.
```

### Example 3: Complex with Breaking Change

```
User: Refactor authentication to support OAuth2

AI:
## 📊 Change Assessment
**Complexity**: 🔴 Complex
**Impact Scope**: 6 files, ~400 lines
- `auth/authenticator.go` - complete rewrite
- `api/middleware.go` - update auth middleware
- `model/user.go` - add OAuth fields
- `config/auth_config.go` - new config structure
- `service/user_service.go` - update user creation
- `migrations/` - new migration file

⚠️ **Breaking Changes**:
- Old: `Authenticate(username, password)`
- New: `Authenticate(ctx, authMethod, credentials)`

**Migration Path**: 
- Keep old method as `AuthenticateBasic()` (deprecated)
- New method supports both basic and OAuth
- Feature flag: `USE_NEW_AUTH`

**Dependencies**:
- Add: `golang.org/x/oauth2 v0.15.0`

This is a significant change. Please confirm:
1. OK with breaking change?
2. Migration strategy acceptable?
3. Proceed with implementation?

User: Yes, proceed

AI:
[implements OAuth2 authentication with proper error handling]

## ✅ Implementation Complete
**Changes Made**:
- `auth/authenticator.go` - Implemented OAuth2 + basic auth
- `api/middleware.go` - Updated middleware
- `model/user.go` - Added OAuth fields
- `config/auth_config.go` - Added config structure
- `service/user_service.go` - Updated user creation
- `migrations/20240115_oauth.sql` - Database migration
- `go.mod` - Added golang.org/x/oauth2 v0.15.0

**Breaking Changes Applied**:
- ⚠️ Old `Authenticate()` → `AuthenticateBasic()` (deprecated)
- ✅ New `Authenticate()` supports both methods
- ✅ Feature flag `USE_NEW_AUTH` controls behavior

**Next Steps**: Add tests, update docs, and test migration in staging.
```

---

## Integration with Other Commands

- Use `/bug-fix` for systematic debugging when issue is unclear
- Use `/requirements-to-implementation/01-requirements-analysis` for large features needing design
- Use `/design` when architectural decisions are needed
- Use this command (`/impl`) for straightforward implementation tasks

---

## When to Use This Command

Use `/impl` when:
- ✅ You have a clear feature or bug to implement
- ✅ You want fast, quality implementation without overhead
- ✅ Requirements are mostly clear (or AI can ask to clarify)
- ✅ You'll handle testing and docs yourself

## When NOT to Use

- ❌ Need full design discussion (use `/design`)
- ❌ Requirements are vague (use `/requirements-to-implementation/*`)
- ❌ Just debugging (use `/bug-fix`)
- ❌ Want comprehensive tests and docs generated
