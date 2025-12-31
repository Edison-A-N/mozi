---
description: Systematic bug fixing workflow - helps LLM accurately diagnose and fix bugs through structured analysis
---

# Bug Fix Assistant

## Task
Guide the LLM through a systematic bug fixing process to accurately diagnose and resolve issues. The user will provide error information, and the AI will follow a structured approach to understand, locate, and fix the bug.

## What the AI will do:

1. ✅ **Understand the Bug**: Analyze the provided error information thoroughly
2. ✅ **Gather Context**: Collect relevant code, logs, and system information
3. ✅ **Reproduce the Issue**: Understand the steps to reproduce (if provided)
4. ✅ **Root Cause Analysis**: Identify the underlying cause, not just symptoms
5. ✅ **Propose Solution**: Suggest fix with explanation and alternatives
6. ✅ **Implement Fix**: Apply the fix with proper testing considerations
7. ✅ **Verify Fix**: Ensure the fix resolves the issue without side effects
8. ✅ **Document**: Explain what was fixed and why

## Bug Fixing Workflow

### Phase 1: Bug Understanding (Round 1)

**Collect Information**:
- [ ] Error message/stack trace
- [ ] Expected vs actual behavior
- [ ] Steps to reproduce
- [ ] Environment details (OS, versions, dependencies)
- [ ] Recent changes that might have triggered the bug
- [ ] Frequency (always, sometimes, edge case)

**Initial Questions to Ask**:
1. What is the exact error message or unexpected behavior?
2. When did this bug start occurring?
3. Can you consistently reproduce it? If so, what are the steps?
4. What were you expecting to happen?
5. Have there been any recent code or environment changes?
6. Does this affect all users or specific scenarios?

### Phase 2: Context Gathering (Round 2-3)

**Code Investigation**:
- [ ] Read the file(s) mentioned in stack trace
- [ ] Understand the code flow leading to the error
- [ ] Check related functions and dependencies
- [ ] Review recent git commits if relevant
- [ ] Look for similar patterns in the codebase

**Search Strategy**:
```
1. Start with error location (stack trace)
2. Trace backwards through call stack
3. Check input validation and data flow
4. Look for edge cases and boundary conditions
5. Review error handling and logging
```

**Common Bug Categories to Consider**:
- **Logic Errors**: Wrong algorithm, incorrect conditions
- **Data Issues**: Type mismatches, null/undefined, validation
- **Concurrency**: Race conditions, deadlocks
- **Resource Issues**: Memory leaks, file handles, connections
- **Integration**: API changes, dependency versions
- **Configuration**: Environment variables, settings
- **Edge Cases**: Boundary values, empty inputs, special characters

### Phase 3: Root Cause Analysis (Round 4-5)

**Analysis Checklist**:
- [ ] Identified the exact line/function causing the error
- [ ] Understood why the error occurs (not just where)
- [ ] Traced the data flow to find incorrect values
- [ ] Checked for related issues in the same area
- [ ] Considered impact on other parts of the system
- [ ] Identified if this is a symptom of a larger problem

**Ask Yourself**:
1. Is this the root cause or just a symptom?
2. Why did this code work before (if it did)?
3. What assumptions were made that are now invalid?
4. Are there similar bugs lurking elsewhere?
5. What would prevent this type of bug in the future?

### Phase 4: Solution Design (Round 6)

**Solution Evaluation**:
- [ ] Propose primary solution with explanation
- [ ] Consider alternative approaches
- [ ] Evaluate trade-offs (complexity, performance, maintainability)
- [ ] Assess impact on existing functionality
- [ ] Plan for testing the fix
- [ ] Consider if refactoring is needed

**Solution Template**:
```markdown
## Proposed Fix

### Root Cause
[Clear explanation of why the bug occurs]

### Solution
[Describe the fix approach]

### Why This Works
[Explain the reasoning]

### Alternatives Considered
1. [Alternative 1]: [Why not chosen]
2. [Alternative 2]: [Why not chosen]

### Testing Strategy
- Unit tests to add/modify
- Integration tests needed
- Manual testing steps

### Potential Side Effects
[Any risks or impacts to consider]
```

### Phase 5: Implementation (Round 7-8)

**Implementation Checklist**:
- [ ] Apply the fix to the identified code
- [ ] Add/update error handling if needed
- [ ] Add input validation if missing
- [ ] Update related code if necessary
- [ ] Add logging for debugging if helpful
- [ ] Update comments/documentation
- [ ] Consider backward compatibility

**Code Quality**:
- Follow existing code style and patterns
- Don't introduce new technical debt
- Keep changes minimal and focused
- Maintain or improve readability

### Phase 6: Verification (Round 9)

**Verification Steps**:
- [ ] Fix addresses the reported error
- [ ] No new errors introduced
- [ ] Edge cases handled
- [ ] Error messages are clear and helpful
- [ ] Logging is appropriate
- [ ] Performance impact is acceptable

**Testing Recommendations**:
```markdown
## Testing Checklist

### Unit Tests
- [ ] Test the fixed function directly
- [ ] Test edge cases that caused the bug
- [ ] Test related functionality

### Integration Tests
- [ ] Test the full flow that was broken
- [ ] Test with realistic data
- [ ] Test error scenarios

### Manual Testing
- [ ] Reproduce the original bug (should be fixed)
- [ ] Test normal use cases
- [ ] Test boundary conditions
```

## Bug Type Specific Guidance

### 🔴 Runtime Errors (NullPointerException, undefined, etc.)
**Focus on**:
- Null/undefined checks
- Input validation
- Initialization order
- Optional chaining/safe navigation

**Common Causes**:
- Missing null checks
- Async timing issues
- Incorrect initialization
- API response changes

### 🟠 Logic Errors (Wrong output, incorrect behavior)
**Focus on**:
- Algorithm correctness
- Conditional logic
- Loop boundaries
- State management

**Common Causes**:
- Off-by-one errors
- Wrong operators (&&/||, ==/===)
- Incorrect assumptions
- Missing edge case handling

### 🟡 Performance Issues (Slow, timeout, memory)
**Focus on**:
- Algorithm complexity
- Database queries (N+1 problem)
- Memory leaks
- Caching opportunities

**Common Causes**:
- Inefficient algorithms
- Missing indexes
- Unbounded loops
- Resource not released

### 🟢 Integration Errors (API, database, external service)
**Focus on**:
- API contract changes
- Network error handling
- Timeout configuration
- Data format validation

**Common Causes**:
- API version mismatch
- Missing error handling
- Incorrect assumptions about external service
- Configuration issues

### 🔵 Concurrency Issues (Race conditions, deadlocks)
**Focus on**:
- Shared state access
- Lock ordering
- Atomic operations
- Thread safety

**Common Causes**:
- Unsynchronized access
- Incorrect lock scope
- Async/await misuse
- Event ordering assumptions

### 🟣 Configuration/Environment Issues
**Focus on**:
- Environment variables
- Configuration files
- Dependency versions
- System resources

**Common Causes**:
- Missing environment variables
- Wrong configuration values
- Version incompatibilities
- Insufficient resources

## Output Format

After completing the bug fix, provide a summary:

```markdown
# Bug Fix Summary

## 🐛 Issue
[Brief description of the bug]

## 🔍 Root Cause
[Explanation of why it occurred]

## ✅ Solution Applied
[What was changed and why]

## 📝 Files Modified
- `path/to/file1.ext` - [what was changed]
- `path/to/file2.ext` - [what was changed]

## 🧪 Testing Recommendations
[How to verify the fix]

## ⚠️ Notes
[Any important considerations, side effects, or follow-up needed]

## 🔮 Prevention
[How to prevent similar bugs in the future]
```

## Usage Examples

### Example 1: Runtime Error
```
User: /bug-fix
Getting "TypeError: Cannot read property 'name' of undefined" in user profile page

AI: I'll help you fix this bug. Let me gather information:

1. Can you provide the full stack trace?
2. Which user profile page (specific route/component)?
3. Does this happen for all users or specific ones?

[After getting info]
Let me examine the code...
[reads relevant files]

Root Cause: The user object is undefined when the profile loads before 
the API call completes. This is a race condition.

Proposed Fix: Add null check and loading state...
```

### Example 2: Logic Error
```
User: /bug-fix
The discount calculation is wrong. Items with 10% discount show 20% off

AI: Let me investigate this logic error.

Questions:
1. Can you show me an example calculation (input/output)?
2. Is this for all items or specific categories?
3. When did this start happening?

[After analysis]
Found the issue: The discount is being applied twice - once in the 
cart calculation and once in the checkout. Let me fix this...
```

### Example 3: Performance Issue
```
User: /bug-fix
The search endpoint times out when searching for common terms

Error: Request timeout after 30s
Database logs show slow query

AI: This is a performance issue. Let me analyze:

1. Checking the search query implementation...
2. Looking at database indexes...
3. Reviewing query patterns...

Root Cause: Missing index on the search column causing full table scan.
Also found N+1 query problem in result processing.

I'll fix both issues...
```

## Important Guidelines

### DO:
- ✅ Ask clarifying questions before diving into code
- ✅ Read the actual error messages and stack traces carefully
- ✅ Trace through the code flow systematically
- ✅ Consider multiple possible causes
- ✅ Explain the root cause clearly
- ✅ Test the fix thoroughly
- ✅ Keep changes minimal and focused
- ✅ Add comments explaining non-obvious fixes
- ✅ Suggest preventive measures

### DON'T:
- ❌ Jump to conclusions without evidence
- ❌ Fix symptoms without addressing root cause
- ❌ Make unnecessary changes while fixing
- ❌ Ignore error handling
- ❌ Skip testing verification
- ❌ Leave debugging code in production
- ❌ Introduce new patterns inconsistent with codebase
- ❌ Over-engineer the solution

## Advanced Techniques

### Debugging Strategies:
1. **Binary Search**: Comment out code sections to isolate the issue
2. **Rubber Duck**: Explain the code flow step-by-step
3. **Diff Analysis**: Compare working vs broken versions
4. **Logging**: Add strategic logs to trace execution
5. **Reproduce Minimal**: Create minimal reproduction case

### When to Refactor:
- The bug reveals poor code structure
- Multiple similar bugs exist in the area
- The fix would be cleaner with refactoring
- Technical debt is blocking the fix

**But**: Keep refactoring separate from bug fix if possible

### When to Escalate:
- Bug is in external dependency
- Requires architectural changes
- Security implications
- Data corruption risk
- Unclear requirements

## Integration with Other Commands

- Use `/design` if the bug reveals design issues
- Use `/requirements-to-implementation/*` for larger fixes requiring new features
- Use `/ask` for quick questions during debugging
- Follow `code-quality-standards.mdc` when implementing fixes

## When to Use This Command

Use `/bug-fix` when:
- You have a specific error or unexpected behavior
- You need systematic debugging assistance
- You want to ensure thorough root cause analysis
- You need help understanding complex bugs
- You want to avoid introducing new bugs while fixing

## When NOT to Use This Command

Don't use `/bug-fix` when:
- You want to add new features (use `/requirements-to-implementation/*`)
- You need design advice (use `/design`)
- You just have questions (use `/ask`)
- The issue is unclear (gather more info first)

