---
description: Iterative design through conversation - supports system design, API design, database design, and design review
---

# Design Through Conversation

## Task
Engage in iterative dialogue to complete a design solution. Automatically adapt to different design types (system architecture, API, database, etc.) and provide specialized guidance for each.

## What the AI will do:

1. ✅ **Identify Design Type**: Determine what kind of design is needed (or ask if unclear)
2. ✅ Listen to initial design requirements or ideas
3. ✅ Apply appropriate design template and checklist
4. ✅ Ask clarifying questions specific to the design type
5. ✅ Propose design solutions and alternatives
6. ✅ Discuss trade-offs and design decisions
7. ✅ Refine the design based on feedback
8. ✅ Optionally perform design review/critique
9. ✅ Output structured design documentation

## Supported Design Types

### 🏗️ System Architecture Design
**When to use**: Designing overall system structure, microservices, component interactions

**Focus areas**:
- Architecture patterns (monolith, microservices, serverless, etc.)
- Component structure and responsibilities
- Communication patterns (sync/async, event-driven, etc.)
- Data flow and system interactions
- Technology stack selection
- Scalability and performance considerations
- Deployment architecture

**Checklist**:
- [ ] Clear component boundaries defined
- [ ] Communication patterns specified
- [ ] Data flow documented
- [ ] Scalability strategy defined
- [ ] Failure modes considered
- [ ] Monitoring and observability planned

### 🔌 API Design
**When to use**: Designing RESTful APIs, GraphQL schemas, RPC interfaces

**Focus areas**:
- API style selection (REST, GraphQL, gRPC, etc.)
- Endpoint structure and naming conventions
- Request/response formats and data models
- Authentication and authorization
- Versioning strategy
- Rate limiting and quotas
- Error handling and status codes
- Documentation approach

**Checklist**:
- [ ] Consistent naming conventions
- [ ] Proper HTTP methods/status codes
- [ ] Authentication/authorization defined
- [ ] Error handling standardized
- [ ] Versioning strategy clear
- [ ] Rate limiting considered
- [ ] API documentation planned

### 🗄️ Database Design
**When to use**: Designing database schemas, data models, migration strategies

**Focus areas**:
- Database type selection (SQL, NoSQL, graph, etc.)
- Schema design and entity relationships
- Normalization vs denormalization
- Indexing strategies
- Data migration and versioning
- Query optimization
- Backup and recovery
- Sharding and partitioning

**Checklist**:
- [ ] Entities and relationships defined
- [ ] Primary and foreign keys identified
- [ ] Indexes planned for common queries
- [ ] Data types and constraints specified
- [ ] Migration strategy defined
- [ ] Backup strategy planned
- [ ] Performance considerations addressed

### 🎨 UI/UX Design
**When to use**: Designing user interfaces, user flows, interaction patterns

**Focus areas**:
- User flows and journey mapping
- UI component structure and hierarchy
- State management approach
- Responsive design strategy
- Accessibility considerations (WCAG)
- Performance (loading, animations)
- Design system integration

**Checklist**:
- [ ] User flows mapped
- [ ] Component hierarchy clear
- [ ] State management defined
- [ ] Responsive breakpoints planned
- [ ] Accessibility requirements met
- [ ] Loading states designed
- [ ] Error states handled

### 🔍 Design Review Mode
**When to use**: You already have a design and want critical feedback

**What AI will do**:
- Act as a senior architect/engineer reviewing your design
- Challenge design decisions constructively
- Identify potential issues and edge cases
- Suggest improvements and alternatives
- Validate against best practices
- Check completeness using relevant checklists

**Review aspects**:
- Completeness: Are all aspects covered?
- Consistency: Are patterns used consistently?
- Scalability: Will it scale?
- Maintainability: Is it easy to maintain?
- Security: Are security concerns addressed?
- Performance: Are there performance bottlenecks?
- Best practices: Does it follow industry standards?

## Conversation Flow

### Phase 1: Scoping (Round 1-2)
1. **Identify design type** (or ask user to specify)
2. Gather context and requirements
3. Identify design scope and constraints
4. Establish success criteria
5. Ask initial clarifying questions

### Phase 2: Exploration (Round 3-5)
1. Propose multiple design alternatives
2. Discuss pros and cons of each approach
3. Deep dive into specific design aspects
4. Address concerns and edge cases
5. Use web search for best practices if needed

### Phase 3: Refinement (Round 6-8)
1. Consolidate feedback into refined design
2. Validate design against checklist
3. Document design decisions and rationale
4. Address open questions

### Phase 4: Review (Optional)
1. Switch to review mode if requested
2. Critically evaluate the design
3. Suggest final improvements

## Output Format

The final design document adapts to the design type but generally includes:

```markdown
# Design Document: [Title]

## 1. Overview
- Design Type: [System/API/Database/UI/etc.]
- Purpose and goals
- Scope and constraints
- Key stakeholders

## 2. Requirements Summary
- Functional requirements
- Non-functional requirements
- Constraints and assumptions

## 3. Design Decisions
- Major choices made
- Alternatives considered
- Trade-offs and rationale

## 4. Detailed Design
[Type-specific sections, e.g.:]
- For System: Component diagram, data flow, deployment
- For API: Endpoints, request/response formats, authentication
- For Database: Schema, relationships, indexes
- For UI: Component hierarchy, user flows, state management

## 5. Non-Functional Considerations
- Performance requirements
- Security considerations
- Scalability approach
- Monitoring and observability

## 6. Implementation Considerations
- Technology recommendations
- Integration points
- Migration strategy (if applicable)
- Testing strategy

## 7. Design Checklist Results
[Type-specific checklist with ✅/⚠️/❌]

## 8. Open Questions & Future Considerations
- Unresolved issues
- Future enhancements
- Technical debt considerations

## 9. Appendix
- Diagrams (ASCII art, mermaid)
- References and resources
```

## Usage Examples

### Example 1: System Architecture
```
User: /design
I need to design a real-time notification system that can handle 100k concurrent users

AI: I'll help you design this system. Based on your description, this is a **System Architecture Design**.

Let me start by understanding your requirements:
1. What types of notifications? (push, email, SMS, in-app?)
2. What's the expected message volume per second?
3. Do you need message persistence and delivery guarantees?
...
```

### Example 2: API Design
```
User: /design
I want to design a RESTful API for a blog platform

AI: Great! This is an **API Design** task. Let me help you create a well-structured API.

First, let me understand the core entities:
1. What are the main resources? (posts, comments, users, tags?)
2. What operations do you need? (CRUD, search, filtering?)
3. Authentication requirements? (JWT, OAuth, API keys?)
...
```

### Example 3: Design Review
```
User: /design
I have a database schema for an e-commerce platform. Can you review it?
[paste schema]

AI: I'll perform a **Design Review** of your database schema.

Let me analyze it systematically:

✅ Strengths:
- Good normalization for user and product tables
- Proper foreign key relationships

⚠️ Concerns:
- Missing indexes on frequently queried columns (e.g., orders.user_id)
- No soft delete strategy for products
...
```

### Example 4: Mixed Design
```
User: /design
I'm building a microservices-based social media app. I need help with both the architecture and API design.

AI: This involves **multiple design types**. Let's tackle them systematically:

1. First, let's design the **System Architecture** (microservices structure)
2. Then, we'll design the **APIs** for each service

Starting with architecture:
- What are the core domains? (user management, posts, messaging, etc.)
- What's your deployment target? (Kubernetes, serverless, etc.)
...
```

## Important Notes

- **Adaptive**: Automatically adapts to the design type
- **No Code Implementation**: Focus purely on design - do not write implementation code
- **Iterative**: Expect multiple rounds of conversation to refine the design
- **Question-Driven**: Ask targeted questions based on design type
- **Checklist-Driven**: Use type-specific checklists to ensure completeness
- **Visual Aids**: Use diagrams (ASCII art, mermaid) when helpful
- **Trade-off Analysis**: Always discuss pros/cons of different approaches
- **Web Search Enabled**: Research best practices and patterns
- **Review Mode**: Can switch to critical review mode when requested
- **Multi-Type Support**: Can handle designs that span multiple types

## Mode Switching

During the conversation, you can:
- **Switch design types**: "Let's also design the API for this"
- **Enter review mode**: "Can you review what we have so far?"
- **Deep dive**: "Let's focus on the database design now"
- **Get checklist**: "Show me the checklist for API design"

## When to Use This Command

Use `/design` when you:
- Need to think through any design problem
- Want to explore different approaches
- Need structured design documentation
- Want expert feedback on a design
- Are planning a refactor
- Need to validate a design against best practices

## When NOT to Use This Command

Don't use `/design` when you:
- Just want to ask a quick question (use `/ask`)
- Are ready to implement immediately (use `/requirements-to-implementation/*`)
- Need working code examples (use regular chat)

