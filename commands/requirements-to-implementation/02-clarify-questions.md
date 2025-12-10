---
description: Identify what information needs clarification based on requirements and solutions
---

# Clarify Questions

## Task
Based on the broken-down requirements and design solutions, identify information points that need further clarification. **All questions must be goal-oriented and directly related to achieving the defined objectives.**

## What the AI will do:

1. ✅ **Review the Goals & Objectives** defined in the requirements analysis
2. ✅ Review requirements breakdown and solution design documents
3. ✅ Identify ambiguous, unclear, or potentially confusing points **that impact goal achievement**
4. ✅ List questions regarding technology selection, business logic, boundary conditions, etc.
5. ✅ Organize into a structured list of clarification questions
6. ✅ Mark priority and impact scope for each question
7. ✅ **Determine if questions are essential for goal achievement or can be deferred**

## Clarification Dimensions

### Business Level
- Details of business rules and logic
- User interaction flows and experience requirements
- Data formats and business constraints

### Technical Level
- Technology stack and framework choices
- Performance metrics and capacity requirements
- Integration interfaces and third-party dependencies

### Non-Functional Requirements
- Security requirements
- Availability and fault tolerance requirements
- Deployment and operations requirements

## Output Format

- **Start with Goal Alignment Check**: Verify if all questions align with the defined goals
- Use ordered list to enumerate all questions that need clarification
- Each question should include:
  - **Question Description**: Clear question statement
  - **Goal Relevance**: How this question relates to achieving the defined objectives
  - **Impact Scope**: Which features or modules are affected by this question
  - **Priority**: High/Medium/Low
  - **Blocking Status**: Whether this question blocks progress (Blocking/Non-blocking)
  - **Suggestions**: Possible solutions or alternatives with recommended defaults

## Question Filtering Criteria

Before listing a question, ask:
1. **Does this question directly impact goal achievement?** If no, defer or skip it.
2. **Is this question blocking implementation?** If no, mark as non-blocking and suggest defaults.
3. **Can we proceed with reasonable assumptions?** If yes, suggest defaults instead of asking.

## Important Notes

- **Goal-Oriented**: All questions must be directly related to achieving the defined goals
- **Stop Condition**: If no blocking questions remain, state "No blocking questions. Ready to proceed with task breakdown."
- **Use Defaults**: For non-blocking questions, provide reasonable defaults and proceed
- **Don't over-question**: Focus only on uncertainties that prevent goal achievement
- **Prioritize**: Only ask High-priority questions that affect core functionality and goal completion
- **Actionable**: Keep questions specific and actionable, avoid theoretical or abstract questions
- **Time-boxed**: Limit the number of questions (recommended: max 5-7 questions per clarification round)

