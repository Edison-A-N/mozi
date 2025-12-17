---
description: Comprehensive architecture assessment - evaluates system architecture, identifies technical debt, and provides prioritized improvement recommendations
---

# Architecture Assessment

## Task
Perform comprehensive architecture evaluation for systems and applications. Analyze architecture quality, identify technical debt, assess scalability and maintainability, and provide prioritized improvement recommendations with actionable roadmaps.

## What the AI will do:

1. ✅ **Understand Architecture**: Analyze system architecture, components, and design patterns
2. ✅ **Assess Architecture Quality**: Evaluate against architecture principles and best practices
3. ✅ **Identify Technical Debt**: Find architectural issues and accumulated debt
4. ✅ **Evaluate Scalability**: Assess ability to handle growth and increased load
5. ✅ **Evaluate Maintainability**: Assess code organization, documentation, and ease of changes
6. ✅ **Performance Analysis**: Identify performance bottlenecks and optimization opportunities
7. ✅ **Risk Assessment**: Identify architectural risks and potential failure points
8. ✅ **Provide Recommendations**: Offer prioritized, actionable improvement suggestions
9. ✅ **Create Roadmap**: Develop implementation roadmap with timelines

## Architecture Assessment Dimensions

### 🏗️ Architecture Quality

**Assessment Areas**:
- Architecture patterns and styles (monolith, microservices, serverless, etc.)
- Component design and boundaries
- Separation of concerns
- Dependency management
- Interface design
- Data flow and system interactions
- Technology stack appropriateness

**Quality Indicators**:
- [ ] Clear architectural boundaries
- [ ] Appropriate abstraction levels
- [ ] Loose coupling between components
- [ ] High cohesion within components
- [ ] Consistent patterns throughout system
- [ ] Well-defined interfaces
- [ ] Proper layering (presentation, business, data)

### 📈 Scalability

**Assessment Areas**:
- Horizontal vs vertical scaling capability
- Database scalability strategy
- Caching strategy
- Load balancing approach
- State management
- Resource utilization
- Bottleneck identification

**Scalability Indicators**:
- [ ] Can scale horizontally (add more instances)
- [ ] Stateless design where possible
- [ ] Database can handle growth (sharding, read replicas)
- [ ] Caching strategy in place
- [ ] No single points of failure
- [ ] Efficient resource usage
- [ ] Async processing for long-running tasks

### 🔧 Maintainability

**Assessment Areas**:
- Code organization and structure
- Documentation quality
- Testing coverage and strategy
- Code complexity
- Technical debt level
- Onboarding difficulty
- Change impact analysis

**Maintainability Indicators**:
- [ ] Clear module boundaries
- [ ] Comprehensive documentation
- [ ] Good test coverage
- [ ] Low cyclomatic complexity
- [ ] Consistent coding standards
- [ ] Easy to understand and modify
- [ ] Clear dependency graph

### ⚡ Performance

**Assessment Areas**:
- Response time characteristics
- Throughput capacity
- Resource efficiency
- Database query optimization
- Caching effectiveness
- Network efficiency
- Algorithm efficiency

**Performance Indicators**:
- [ ] Acceptable response times
- [ ] Efficient database queries
- [ ] Proper caching usage
- [ ] Optimized algorithms
- [ ] Efficient resource usage
- [ ] Minimal unnecessary data transfer
- [ ] Proper connection pooling

### 🛡️ Reliability & Resilience

**Assessment Areas**:
- Error handling strategies
- Failure recovery mechanisms
- Circuit breakers and retries
- Monitoring and alerting
- Disaster recovery
- Data consistency
- Transaction management

**Reliability Indicators**:
- [ ] Comprehensive error handling
- [ ] Graceful degradation
- [ ] Retry mechanisms with backoff
- [ ] Circuit breakers for external services
- [ ] Proper monitoring and alerting
- [ ] Backup and recovery procedures
- [ ] Data consistency guarantees

### 🔐 Security Architecture

**Assessment Areas**:
- Authentication and authorization
- Data protection
- API security
- Network security
- Secrets management
- Security monitoring
- Compliance considerations

**Security Indicators**:
- [ ] Proper authentication mechanisms
- [ ] Role-based access control
- [ ] Data encryption (at rest and in transit)
- [ ] Secure API design
- [ ] Secrets management
- [ ] Security monitoring
- [ ] Compliance requirements met

## Assessment Workflow

### Phase 1: Architecture Discovery (Round 1-3)

**Information Gathering**:
- [ ] System architecture overview
- [ ] Technology stack and versions
- [ ] Component structure and responsibilities
- [ ] Data flow and system interactions
- [ ] Deployment architecture
- [ ] Infrastructure setup
- [ ] Current performance metrics
- [ ] Known issues and pain points

**Questions to Ask**:
1. What is the overall system architecture? (monolith, microservices, serverless, etc.)
2. What technologies and frameworks are used?
3. How are components organized and what are their responsibilities?
4. How does data flow through the system?
5. What is the deployment architecture?
6. What are the current performance characteristics?
7. What are known pain points or issues?
8. What are the business requirements and constraints?

**Discovery Activities**:
- Read codebase structure and organization
- Analyze dependency relationships
- Review configuration files
- Examine deployment configurations
- Check documentation
- Review monitoring dashboards (if available)

### Phase 2: Architecture Analysis (Round 4-6)

**Pattern Analysis**:
- [ ] Identify architecture patterns used
- [ ] Evaluate pattern appropriateness
- [ ] Check for anti-patterns
- [ ] Assess consistency of patterns
- [ ] Identify missing patterns

**Component Analysis**:
- [ ] Evaluate component boundaries
- [ ] Assess coupling and cohesion
- [ ] Review interface design
- [ ] Check dependency management
- [ ] Evaluate component responsibilities

**Data Flow Analysis**:
- [ ] Map data flow through system
- [ ] Identify data transformation points
- [ ] Assess data consistency strategies
- [ ] Evaluate caching strategies
- [ ] Check for data bottlenecks

**Technology Stack Evaluation**:
- [ ] Assess technology choices
- [ ] Check for outdated dependencies
- [ ] Evaluate technology fit
- [ ] Identify technology risks
- [ ] Assess vendor lock-in

### Phase 3: Quality Assessment (Round 7-9)

**Architecture Principles Check**:
- [ ] SOLID principles adherence
- [ ] DRY (Don't Repeat Yourself)
- [ ] KISS (Keep It Simple, Stupid)
- [ ] YAGNI (You Aren't Gonna Need It)
- [ ] Separation of Concerns
- [ ] Single Responsibility Principle

**Technical Debt Identification**:
- [ ] Code duplication
- [ ] Outdated dependencies
- [ ] Missing tests
- [ ] Poor abstractions
- [ ] Tight coupling
- [ ] Legacy code issues
- [ ] Documentation gaps

**Risk Assessment**:
- [ ] Single points of failure
- [ ] Scalability bottlenecks
- [ ] Security vulnerabilities
- [ ] Performance risks
- [ ] Maintenance risks
- [ ] Technology risks

### Phase 4: Recommendations (Round 10-11)

**Prioritization Framework**:
```
Priority = Impact × Urgency × Effort

Impact:
- High (3): Critical for business/performance
- Medium (2): Important improvement
- Low (1): Nice to have

Urgency:
- High (3): Immediate action needed
- Medium (2): Should address soon
- Low (1): Can plan for later

Effort:
- Low (1): Quick win (< 1 week)
- Medium (2): Moderate effort (1-4 weeks)
- High (3): Significant effort (> 1 month)

Priority Levels:
- 18-27: 🔴 CRITICAL (P0)
- 12-17: 🟠 HIGH (P1)
- 6-11: 🟡 MEDIUM (P2)
- 1-5: 🔵 LOW (P3)
```

**Recommendation Template**:
```markdown
## [Priority Icon] [Recommendation Title]

**Priority**: [CRITICAL/HIGH/MEDIUM/LOW]
**Category**: [Scalability/Maintainability/Performance/Security/etc.]
**Impact**: [High/Medium/Low]
**Effort**: [Low/Medium/High]
**Timeline**: [Estimated time]

### Current State
[Description of current architecture/issue]

### Problem Statement
[What problem does this address?]

### Proposed Solution
[Detailed solution description]

### Benefits
- [Benefit 1]
- [Benefit 2]
- [Benefit 3]

### Implementation Steps
1. [Step 1]
2. [Step 2]
3. [Step 3]

### Risks & Mitigation
- **Risk**: [Potential risk]
  - **Mitigation**: [How to mitigate]

### Success Metrics
[How to measure success]

### References
[Relevant patterns, articles, best practices]
```

## Output Format

### Architecture Assessment Report

```markdown
# Architecture Assessment Report: [System Name]

**Date**: [Date]
**Assessed By**: AI Architecture Assessment
**System Version**: [Version]
**Assessment Scope**: [What was assessed]

---

## Executive Summary

### Overall Architecture Rating
- **Architecture Quality**: ⭐⭐⭐⭐☆ (4/5)
- **Scalability**: ⭐⭐⭐☆☆ (3/5)
- **Maintainability**: ⭐⭐⭐⭐☆ (4/5)
- **Performance**: ⭐⭐⭐⭐☆ (4/5)
- **Reliability**: ⭐⭐⭐☆☆ (3/5)
- **Security**: ⭐⭐⭐⭐☆ (4/5)

### Key Findings
- 🔴 Critical Issues: [X]
- 🟠 High Priority: [Y]
- 🟡 Medium Priority: [Z]
- 🔵 Low Priority: [W]

### Top 3 Recommendations
1. [Most critical recommendation]
2. [Second priority recommendation]
3. [Third priority recommendation]

### Overall Assessment
[2-3 paragraphs summarizing architecture health, main strengths, and key concerns]

---

## 1. Architecture Overview

### Current Architecture
[Description of current architecture]

### Architecture Diagram
[ASCII art or mermaid diagram]

### Technology Stack
- **Languages**: [List]
- **Frameworks**: [List]
- **Databases**: [List]
- **Infrastructure**: [List]
- **Third-party Services**: [List]

### Component Structure
[Description of major components and their responsibilities]

---

## 2. Architecture Quality Assessment

### Strengths ✅
- [Strength 1 with explanation]
- [Strength 2 with explanation]
- [Strength 3 with explanation]

### Weaknesses ⚠️
- [Weakness 1 with explanation]
- [Weakness 2 with explanation]
- [Weakness 3 with explanation]

### Architecture Patterns Used
- [Pattern 1]: [Assessment]
- [Pattern 2]: [Assessment]

### Anti-patterns Identified
- [Anti-pattern 1]: [Impact and recommendation]
- [Anti-pattern 2]: [Impact and recommendation]

---

## 3. Scalability Assessment

### Current Scalability Characteristics
[Analysis of current scalability]

### Scalability Strengths ✅
- [Strength 1]
- [Strength 2]

### Scalability Concerns ⚠️
- [Concern 1]
- [Concern 2]

### Bottlenecks Identified
- [Bottleneck 1]: [Impact and solution]
- [Bottleneck 2]: [Impact and solution]

### Scaling Strategy Recommendations
[Recommendations for improving scalability]

---

## 4. Maintainability Assessment

### Code Organization
[Assessment of code structure]

### Documentation Quality
[Assessment of documentation]

### Test Coverage
[Assessment of testing]

### Technical Debt
- **Total Debt Level**: [High/Medium/Low]
- **Key Debt Areas**:
  - [Debt area 1]
  - [Debt area 2]

### Complexity Analysis
[Assessment of code complexity]

---

## 5. Performance Assessment

### Current Performance Characteristics
- **Response Time**: [Metrics]
- **Throughput**: [Metrics]
- **Resource Usage**: [Metrics]

### Performance Strengths ✅
- [Strength 1]
- [Strength 2]

### Performance Issues ⚠️
- [Issue 1]
- [Issue 2]

### Optimization Opportunities
[Specific optimization recommendations]

---

## 6. Reliability & Resilience Assessment

### Error Handling
[Assessment of error handling]

### Failure Recovery
[Assessment of recovery mechanisms]

### Monitoring & Observability
[Assessment of monitoring]

### Single Points of Failure
[List identified SPOFs]

---

## 7. Security Architecture Assessment

### Security Strengths ✅
- [Strength 1]
- [Strength 2]

### Security Concerns ⚠️
- [Concern 1]
- [Concern 2]

### Compliance Status
[Assessment against compliance requirements]

---

## 8. Critical Recommendations (🔴 P0)

[List all critical recommendations using the template]

---

## 9. High Priority Recommendations (🟠 P1)

[List all high-priority recommendations]

---

## 10. Medium Priority Recommendations (🟡 P2)

[List all medium-priority recommendations]

---

## 11. Low Priority Recommendations (🔵 P3)

[List all low-priority recommendations]

---

## 12. Implementation Roadmap

### Phase 1: Critical Fixes (0-1 month)
- [ ] [Critical recommendation 1]
- [ ] [Critical recommendation 2]
- **Timeline**: [Dates]
- **Resources**: [Required resources]

### Phase 2: High Priority Improvements (1-3 months)
- [ ] [High priority recommendation 1]
- [ ] [High priority recommendation 2]
- **Timeline**: [Dates]
- **Resources**: [Required resources]

### Phase 3: Medium Priority Enhancements (3-6 months)
- [ ] [Medium priority recommendation 1]
- [ ] [Medium priority recommendation 2]
- **Timeline**: [Dates]
- **Resources**: [Required resources]

### Phase 4: Long-term Improvements (6-12 months)
- [ ] [Long-term recommendation 1]
- [ ] [Long-term recommendation 2]
- **Timeline**: [Dates]
- **Resources**: [Required resources]

---

## 13. Risk Assessment

### High Risk Areas
- [Risk 1]: [Impact and mitigation]
- [Risk 2]: [Impact and mitigation]

### Medium Risk Areas
- [Risk 1]: [Impact and mitigation]
- [Risk 2]: [Impact and mitigation]

### Risk Mitigation Strategy
[Overall risk mitigation approach]

---

## 14. Architecture Evolution Path

### Short-term (0-6 months)
[Architecture improvements for short term]

### Medium-term (6-12 months)
[Architecture evolution for medium term]

### Long-term (12+ months)
[Vision for long-term architecture]

---

## 15. Metrics & KPIs

### Architecture Health Metrics
- [Metric 1]: [Current value] → [Target value]
- [Metric 2]: [Current value] → [Target value]

### Success Criteria
[How to measure success of improvements]

---

## 16. Appendix

### Architecture Diagrams
[Additional diagrams if needed]

### References
- [Architecture pattern references]
- [Best practice guides]
- [Technology documentation]

### Glossary
[Architecture terms used in the report]
```

## Usage Examples

### Example 1: Full Architecture Assessment
```
User: /architecture-assessment
Please assess the architecture of our e-commerce platform

AI: I'll perform a comprehensive architecture assessment of your e-commerce platform.

Let me start by understanding your system:
1. What is the overall architecture? (monolith, microservices?)
2. What technologies are you using?
3. What are the main components?
4. What are current pain points?

[After gathering information]

# Architecture Assessment Report: E-Commerce Platform

## Executive Summary
- Architecture Quality: ⭐⭐⭐☆☆ (3/5)
- Scalability: ⭐⭐☆☆☆ (2/5)
- Maintainability: ⭐⭐⭐⭐☆ (4/5)

🔴 Critical Issues: 2
🟠 High Priority: 5

Top Recommendations:
1. Database scalability bottleneck needs immediate attention
2. Monolithic architecture limiting horizontal scaling
...
```

### Example 2: Focused Assessment
```
User: /architecture-assessment
Assess the scalability of our API service

AI: I'll focus on scalability assessment of your API service.

Questions:
1. Current request volume?
2. Expected growth?
3. Current scaling strategy?
4. Performance bottlenecks?

[After analysis]

## Scalability Assessment

### Current State
- Handles 1K req/s
- Single database instance
- No caching layer

### Concerns
🔴 CRITICAL: Database is single point of failure
🟠 HIGH: No horizontal scaling capability
🟡 MEDIUM: Missing caching layer

### Recommendations
1. Implement database read replicas (CRITICAL)
2. Add load balancer and multiple API instances (HIGH)
3. Implement Redis caching layer (MEDIUM)
...
```

### Example 3: Technical Debt Assessment
```
User: /architecture-assessment
Identify technical debt in our codebase

AI: I'll analyze your codebase for technical debt.

[Analyzing code structure, dependencies, tests, documentation]

## Technical Debt Assessment

### Debt Summary
- **Total Debt Level**: Medium-High
- **Estimated Payoff Time**: 3-4 months

### Key Debt Areas

🔴 CRITICAL:
- Outdated dependencies with security vulnerabilities
- Missing test coverage (30% coverage)

🟠 HIGH:
- Code duplication across modules
- Tight coupling between components

🟡 MEDIUM:
- Inconsistent error handling
- Missing documentation

### Debt Payoff Roadmap
[Prioritized plan to address technical debt]
```

### Example 4: Architecture Pattern Evaluation
```
User: /architecture-assessment
Evaluate if we should migrate from monolith to microservices

AI: I'll assess your current architecture and evaluate migration feasibility.

[Analyzing current architecture, team size, deployment, etc.]

## Architecture Pattern Evaluation

### Current State: Monolithic Architecture
**Strengths**:
- Simple deployment
- Easy development
- Good for current scale

**Weaknesses**:
- Limited independent scaling
- Tight coupling
- Deployment bottlenecks

### Migration Feasibility
**Recommendation**: Not recommended at this time

**Reasons**:
1. Current scale doesn't justify complexity
2. Team size too small for microservices
3. No clear service boundaries identified

### Alternative Recommendations
1. Modular monolith approach (better fit)
2. Improve current architecture
3. Revisit in 6-12 months when scale increases
```

## Assessment Frameworks

### Architecture Quality Attributes
- **Modifiability**: Ease of making changes
- **Testability**: Ease of testing
- **Deployability**: Ease of deployment
- **Scalability**: Ability to handle growth
- **Performance**: Response time and throughput
- **Security**: Protection against threats
- **Reliability**: System availability
- **Usability**: Developer and user experience

### Architecture Patterns Reference
- **Monolithic**: Single deployable unit
- **Microservices**: Independent services
- **Serverless**: Function-as-a-Service
- **Event-Driven**: Event-based communication
- **Layered**: Separation by layers
- **Hexagonal**: Ports and adapters
- **CQRS**: Command Query Responsibility Segregation
- **Event Sourcing**: Event-based state

### Anti-patterns to Watch For
- **Big Ball of Mud**: No clear structure
- **God Object**: Too many responsibilities
- **Spaghetti Architecture**: Unclear dependencies
- **Vendor Lock-in**: Over-dependence on vendor
- **Premature Optimization**: Optimizing too early
- **Copy-Paste Programming**: Code duplication
- **Magic Numbers**: Hardcoded values
- **Tight Coupling**: High interdependency

## Important Guidelines

### DO:
- ✅ Assess architecture holistically
- ✅ Consider business context and constraints
- ✅ Provide actionable, prioritized recommendations
- ✅ Balance ideal vs. practical solutions
- ✅ Consider team size and capabilities
- ✅ Evaluate cost vs. benefit
- ✅ Provide clear implementation roadmaps
- ✅ Reference architecture patterns and best practices
- ✅ Consider migration paths for improvements

### DON'T:
- ❌ Recommend changes without understanding context
- ❌ Over-engineer solutions
- ❌ Ignore business constraints
- ❌ Suggest patterns that don't fit team size
- ❌ Provide vague recommendations
- ❌ Ignore migration complexity
- ❌ Focus only on code without considering infrastructure
- ❌ Overlook operational aspects

## Advanced Assessment Techniques

### Architecture Decision Records (ADRs)
- Document key architectural decisions
- Record alternatives considered
- Track decision rationale
- Maintain decision history

### Architecture Fitness Functions
- Define measurable architecture goals
- Create automated checks
- Monitor architecture drift
- Alert on violations

### Dependency Analysis
- Map component dependencies
- Identify circular dependencies
- Assess dependency health
- Evaluate dependency risks

### Performance Profiling
- Identify performance bottlenecks
- Profile critical paths
- Measure resource usage
- Optimize hot paths


## When to Use This Command

Use `/architecture-assessment` when:
- Starting a new project and need architecture guidance
- Evaluating existing architecture health
- Planning major refactoring or migration
- Identifying technical debt
- Preparing for scale
- Assessing architecture before major changes
- Need prioritized improvement roadmap
- Evaluating architecture patterns

## When NOT to Use This Command

Don't use `/architecture-assessment` when:
- Need quick code review (use `/review`)
- Fixing specific bugs (use `/bug-fix`)
- Designing new features (use `/design`)
- Just have questions (use `/ask`)
- Need immediate implementation (use `/requirements-to-implementation/*`)

---

**Remember**: Architecture assessment is an ongoing process. Regular assessments help maintain architecture health and guide evolution. Balance ideal architecture with practical constraints and business needs.

