---
description: Break down tasks into step-by-step design, focusing on splitting into actionable Task/TODO items
---

# Task Breakdown

## Task
Based on requirements and solution design, break down work into executable Task/TODO items. Tasks should be implementation-oriented, with each task resulting in a testable module or feature.

## What the AI will do:

1. ✅ Analyze requirements and solution design
2. ✅ Organize tasks by dependencies and priorities
3. ✅ Break down large tasks into small, independently completable subtasks
4. ✅ Set clear acceptance criteria for each task
5. ✅ Identify dependencies between tasks
6. ✅ Create structured TODO list

## Breakdown Principles

### Task Granularity
- Each task should be an independently completable work unit
- Tasks should be implementation-oriented, targeting a testable module or feature
- Tasks should have clear inputs and outputs
- Each task should result in a working, testable implementation

### Task Organization
- Group by module or functional domain
- Consider dependencies, complete foundational tasks first
- Distinguish between core functionality and enhancements

### Task Description
- Use clear, specific verbs (create, implement, test, refactor, etc.)
- Include explicit acceptance criteria
- Note key information like tech stack, file paths, etc.

## Output Format

Use TODO list format, for example:

```markdown
## Module/Functional Group Name

### TODO
- [ ] Task 1: Description (Priority: High/Medium/Low)
  - Acceptance Criteria: ...
  - Dependencies: ...
- [ ] Task 2: Description
  - Acceptance Criteria: ...
  - Dependencies: Task 1
```

## Important Notes

- **Implementation-oriented**: Each task should focus on delivering a working, testable module or feature
- **Testable outcomes**: Tasks should result in code that can be tested and verified
- Ensure no circular dependencies between tasks
- Prioritize breaking down tasks on the critical path
- Set clear completion criteria for each task
- Consider testing and validation tasks
- Maintain task traceability

