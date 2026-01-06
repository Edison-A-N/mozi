---
description: Code review assistant - analyze commits and provide optimization suggestions
---

# Code Review Assistant

## Task
Help users analyze Git commit(s) changes, provide code quality assessment and optimization suggestions.

## What the AI will do:

1. ✅ **Fetch Changes**: Read the specified commit(s) changes
2. ✅ **Code Analysis**: Analyze code quality and potential issues
3. ✅ **Optimization Directions**: Provide specific improvement suggestions
4. ✅ **Best Practices Check**: Verify compliance with best practices
5. ✅ **Generate Review Report**: Output structured review results

## Review Workflow

### Phase 1: Fetch Changes

**Supported Input Formats**:
- Single commit: `<commit-hash>` or `HEAD`
- Commit range: `<commit-hash>..<commit-hash>`
- Branch comparison: `main..feature-branch`

**Fetch Change Information**:
```bash
# View single commit changes
git show <commit-hash>

# View commit range changes
git log --patch <commit1>..<commit2>

# View changed file list
git diff --name-status <commit1>..<commit2>

# View specific code changes
git diff <commit1>..<commit2>
```

### Phase 2: Change Overview Analysis

**Analysis Dimensions**:
- [ ] Number and types of changed files
- [ ] Lines of code changed (added/deleted/modified)
- [ ] Affected modules and components
- [ ] Change complexity assessment
- [ ] Commit message quality

**Output Format**:
```markdown
## 📊 Change Overview

### Statistics
- Number of Commits: X
- Files Changed: Y
- Lines Added: +XXX
- Lines Deleted: -XXX
- Main Modules Affected: [list modules]

### Commit Information
- Commit 1: [hash] - [message]
- Commit 2: [hash] - [message]
...

### Change Type Classification
- New Features: X%
- Bug Fixes: Y%
- Refactoring: Z%
- Documentation: W%
- Testing: V%
```

### Phase 3: Code Quality Analysis

**Core Checks**:

#### 3.1 Code Structure & Design
- [ ] Single Responsibility Principle for functions/methods
- [ ] SOLID principles compliance
- [ ] Code reusability
- [ ] Reasonable module division
- [ ] Clear dependency relationships
- [ ] No circular dependencies

#### 3.2 Code Readability
- [ ] Clear and consistent naming
- [ ] Adequate and accurate comments
- [ ] Complex logic explained
- [ ] Magic numbers converted to constants
- [ ] Consistent code formatting

#### 3.3 Error Handling
- [ ] Comprehensive exception handling
- [ ] Clear error messages
- [ ] Boundary condition checks
- [ ] Proper resource cleanup
- [ ] Defensive programming in place

#### 3.4 Performance Considerations
- [ ] Reasonable algorithm complexity
- [ ] No performance bottlenecks
- [ ] Optimized database queries (N+1 problem)
- [ ] No unnecessary computations
- [ ] Reasonable caching strategy

#### 3.5 Security
- [ ] Adequate input validation
- [ ] SQL injection risks
- [ ] XSS attack risks
- [ ] Sensitive information exposure
- [ ] Correct permission control

#### 3.6 Test Coverage
- [ ] Corresponding unit tests exist
- [ ] Test cases cover main scenarios
- [ ] Boundary condition tests
- [ ] Test code quality

### Phase 4: Specific Optimization Suggestions

**Optimization Categories**:

#### 🔴 Critical (Must Fix)
Issues affecting functionality, security, or system stability

**Example Issues**:
- Obvious bugs or logic errors
- Security vulnerabilities
- Resource leaks
- Concurrency safety issues

**Output Format**:
```markdown
### 🔴 Critical Issues

#### Issue 1: [Issue Title]
**Location**: `file.ext:line`
**Description**: [Detailed explanation of the issue]
**Risk**: [Explain potential impact]
**Suggested Fix**:
[Specific fix suggestions and code examples]
```

#### 🟠 Major (Strongly Recommended)
Issues affecting code quality, maintainability, or performance

**Example Issues**:
- Code duplication
- Performance issues
- Unreasonable design
- Missing error handling

**Output Format**:
```markdown
### 🟠 Major Issues

#### Issue 1: [Issue Title]
**Location**: `file.ext:line`
**Current Problem**: [Explain existing issue]
**Optimization Suggestion**: [Provide improvement approach]
**Expected Benefits**: [Explain benefits after optimization]
**Optimization Example**:
[Provide optimized code example or pseudocode]
```

#### 🟡 Minor (Recommended)
Code style, naming, comments, and other improvement suggestions

**Example Issues**:
- Unclear naming
- Missing comments
- Inconsistent code formatting
- Code that can be further simplified

**Output Format**:
```markdown
### 🟡 Minor Suggestions

#### Suggestion 1: [Suggestion Title]
**Location**: `file.ext:line`
**Current State**: [Brief description]
**Suggestion**: [Specific recommendation]
**Example**: [Provide example if needed]
```

#### 🟢 Good Practices (Worth Learning)
Excellent practices and highlights in the code

**Examples**:
- Elegant design pattern application
- Good error handling
- Clear code structure
- Comprehensive test coverage

**Output Format**:
```markdown
### 🟢 Good Practices Found

#### Practice 1: [Highlight Title]
**Location**: `file.ext:line`
**Description**: [Explain why this is a good practice]
**Learning Value**: [What scenarios can this be applied to]
```

### Phase 5: Best Practices Checklist

**General Best Practices**:

#### Code Organization
- [ ] Appropriate function length (typically < 50 lines)
- [ ] Single and clear class responsibilities
- [ ] Low coupling between modules
- [ ] Reasonable interface design

#### Naming Conventions
- [ ] Descriptive variable names
- [ ] Function names clearly express intent
- [ ] Constants use uppercase
- [ ] Class names use nouns

#### Comments & Documentation
- [ ] Complex logic has explanatory comments
- [ ] Public APIs have documentation
- [ ] TODO/FIXME are tracked
- [ ] Comments are accurate and valuable

#### Error Handling
- [ ] Don't ignore exceptions
- [ ] Meaningful error messages
- [ ] Proper resource release
- [ ] Boundary condition handling

#### Performance Awareness
- [ ] Avoid unnecessary nested loops
- [ ] Reasonable data structure usage
- [ ] Memory usage consideration
- [ ] Avoid redundant calculations

#### Security Awareness
- [ ] Validate user input
- [ ] Avoid hardcoded sensitive information
- [ ] Use secure APIs
- [ ] Proper permission handling

### Phase 6: Language-Specific Checks

#### Go
- [ ] Error handling follows Go conventions
- [ ] Proper use of defer
- [ ] Correct Context passing
- [ ] No goroutine leaks
- [ ] Follows Effective Go

#### JavaScript/TypeScript
- [ ] Proper use of async/await
- [ ] Type definitions present (TS)
- [ ] Avoid callback hell
- [ ] Correct Promise handling
- [ ] Follows ES6+ standards

#### Python
- [ ] Follows PEP 8
- [ ] Proper exception usage
- [ ] Uses type hints
- [ ] Follows Pythonic style
- [ ] Resource management (with statement)

#### Java
- [ ] Proper exception hierarchy usage
- [ ] Follows naming conventions
- [ ] Proper generic usage
- [ ] Avoids null pointers
- [ ] Thread safety

## Review Report Template

```markdown
# 🔍 Code Review Report

## 📋 Review Information
- **Review Scope**: [commit range or single commit]
- **Review Time**: [timestamp]
- **Reviewer**: AI Assistant

---

## 📊 Change Overview
[Results from Phase 2 analysis]

---

## 🎯 Review Results Summary

### Issue Statistics
- 🔴 Critical: X items
- 🟠 Major: Y items
- 🟡 Minor: Z items
- 🟢 Good Practices: W items

### Overall Rating
- **Code Quality**: ⭐⭐⭐⭐☆ (4/5)
- **Maintainability**: ⭐⭐⭐⭐☆ (4/5)
- **Test Coverage**: ⭐⭐⭐☆☆ (3/5)
- **Documentation**: ⭐⭐⭐☆☆ (3/5)
- **Performance**: ⭐⭐⭐⭐☆ (4/5)

### Overall Impression
[2-3 paragraphs summarizing overall impression and main concerns]

---

## 🔴 Critical Issues
[Critical issues identified in Phase 4]

---

## 🟠 Major Issues
[Major issues identified in Phase 4]

---

## 🟡 Minor Suggestions
[Minor suggestions from Phase 4]

---

## 🟢 Good Practices
[Good practices found in Phase 4]

---

## 📝 Best Practices Checklist
[Checklist results from Phase 5]

---

## 💡 Optimization Roadmap

### Immediate Fixes (1-2 days)
1. [Critical Issue 1]
2. [Critical Issue 2]
...

### Short-term Optimization (Within 1 week)
1. [Major Issue 1]
2. [Major Issue 2]
...

### Long-term Improvements (Can be planned for next iteration)
1. [Refactoring Suggestion 1]
2. [Architecture Optimization Suggestion]
...

---

## 🎓 Learning & Improvement Suggestions

### For This Change
[Specific learning points for this commit]

### General Improvement Directions
[Suggestions applicable to future development]

---

## ✅ Checklist Before Merge

- [ ] All Critical issues fixed
- [ ] Major issues assessed with handling plan
- [ ] Tests passed
- [ ] Documentation updated
- [ ] Code review comments addressed
- [ ] CI/CD checks passed

---

## 📌 Follow-up Actions

### Need Immediate Action
- [ ] Action 1
- [ ] Action 2

### Can Follow Up Later
- [ ] Action 3
- [ ] Action 4

---

## 💬 Additional Notes
[Other content that needs explanation]
```

## Usage Examples

### Example 1: Review Single Commit
```
User: /review
Please review the latest commit

AI: Sure, let me check the latest commit.

[Execute git show HEAD]

I will analyze this commit's changes...

[Analysis process]

# 🔍 Code Review Report

## 📋 Review Information
- **Commit**: abc1234 - "Add user authentication feature"
- **Author**: John Doe
- **Date**: 2024-01-15

## 📊 Change Overview
...
```

### Example 2: Review Commit Range
```
User: /review
Please review all changes in main..feature/user-auth branch

AI: Sure, let me analyze all changes in this branch.

[Execute git log --patch main..feature/user-auth]

There are 5 commits involving 15 files...

[Detailed analysis]
```

### Example 3: Focus on Specific Aspects
```
User: /review
Please review the last 3 commits, focusing on performance and security

AI: Sure, I will focus on performance and security aspects.

[Execute git log --patch HEAD~3..HEAD]

Regarding performance, I found the following issues:
1. ...

Regarding security, note the following:
1. ...
```

### Example 4: Quick Review
```
User: /review 
Quick review of commit abc1234, just tell me the main issues

AI: Quick review results:

🔴 Critical: 1 item
- SQL injection risk: User input not validated and directly concatenated into query

🟠 Major: 2 items
- Missing error handling: API call may fail but not handled
- Performance issue: Database query in loop (N+1)

Recommend fixing Critical issue first.
```

## Important Guidelines

### DO:
- ✅ Comprehensively analyze code changes
- ✅ Provide specific optimization suggestions and examples
- ✅ Differentiate issue priorities
- ✅ Point out excellent practices in code
- ✅ Provide actionable improvement roadmap
- ✅ Consider code context and overall architecture
- ✅ Balance criticism and encouragement
- ✅ Provide learning value

### DON'T:
- ❌ Be overly harsh or nitpick details
- ❌ Only point out problems without solutions
- ❌ Ignore strengths in code
- ❌ Provide vague suggestions
- ❌ Ignore project realities
- ❌ Force personal preferences
- ❌ Ignore existing code style
- ❌ Review scope too large leading to lack of depth

## Advanced Review Techniques

### 1. Architecture-Level Review
- Assess impact of changes on overall architecture
- Check for technical debt introduction
- Consider scalability and maintainability
- Evaluate dependencies between modules

### 2. Business Logic Review
- Understand business requirements
- Verify implementation meets requirements
- Check boundary conditions and exception scenarios
- Consider user experience

### 3. Code Evolution Review
- Compare with historical versions
- Assess refactoring effectiveness
- Check for code regression
- Track technical debt changes

### 4. Team Collaboration Review
- Check code style consistency
- Assess readability and understandability
- Consider onboarding difficulty for new team members
- Check documentation and comment completeness

## Review Output Adjustment

Adjust review detail level based on user needs:

### Quick Review
- Only output issue statistics and main problems
- Suitable for quick checks

### Standard Review
- Complete review report
- Includes all analysis dimensions
- Suitable for regular reviews

### Deep Review
- Detailed code analysis
- Includes refactoring suggestions
- Architecture-level assessment
- Suitable for important changes or quality improvements

### Focused Review
- Only focus on specific aspects (performance, security, testing, etc.)
- Deep analysis of specific issues
- Suitable for reviews with clear objectives

## Integration with Other Commands

- Use `/bug-fix` to fix bugs when found
- Use `/design` to design refactoring solutions
- Use `/ask` for discussion questions
- Use `/git-help/git-commit` to generate commit messages

## When to Use This Command

Use `/review` when:
- Need to perform code review on commits
- Want to understand optimization directions for code changes
- Quality check before merging branches
- Learning excellent practices from others' code
- Assessing code quality and technical debt
- Preparing for refactoring

## When NOT to Use This Command

Don't use `/review` when:
- Need to fix specific bugs (use `/bug-fix`)
- Need to design new features (use `/design`)
- Just want to ask questions (use `/ask`)
- Need to execute code changes (use other implementation commands)
