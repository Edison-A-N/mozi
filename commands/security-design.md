---
description: Security design and threat analysis with risk-level categorization - identifies vulnerabilities and provides prioritized security recommendations
---

# Security Design & Threat Analysis

## Task
Perform comprehensive security analysis and design review for applications and systems. Identify security vulnerabilities with risk levels, provide prioritized recommendations, and help design secure architectures following security best practices.

## What the AI will do:

1. ✅ **Understand the System**: Analyze system architecture and components
2. ✅ **Identify Assets**: Determine what needs protection (data, services, users)
3. ✅ **Threat Modeling**: Identify potential threats using structured frameworks (STRIDE, OWASP)
4. ✅ **Vulnerability Assessment**: Find security weaknesses with risk levels
5. ✅ **Risk Analysis**: Evaluate likelihood and impact of threats
6. ✅ **Security Recommendations**: Provide prioritized, actionable security measures
7. ✅ **Compliance Check**: Verify against security standards (OWASP, CWE, etc.)
8. ✅ **Security Architecture**: Design security controls and defense mechanisms

## Security Risk Levels

### 🔴 **CRITICAL** (P0 - Immediate Action Required)
- **Severity**: Exploitable vulnerabilities that can lead to complete system compromise
- **Impact**: Data breach, financial loss, legal liability, reputation damage
- **Timeline**: Must be fixed immediately (within 24 hours)
- **Examples**: 
  - SQL injection allowing database access
  - Remote code execution vulnerabilities
  - Exposed credentials or API keys in code
  - Authentication bypass
  - Unencrypted sensitive data in transit

### 🟠 **HIGH** (P1 - Urgent)
- **Severity**: Serious vulnerabilities with high exploitation risk
- **Impact**: Significant data exposure, privilege escalation, service disruption
- **Timeline**: Fix within 1 week
- **Examples**:
  - Missing authentication on sensitive endpoints
  - Weak encryption algorithms
  - Cross-Site Scripting (XSS) vulnerabilities
  - Insecure direct object references (IDOR)
  - Insufficient input validation

### 🟡 **MEDIUM** (P2 - Important)
- **Severity**: Moderate vulnerabilities requiring specific conditions to exploit
- **Impact**: Limited data exposure, reduced functionality, compliance issues
- **Timeline**: Fix within 1 month
- **Examples**:
  - Missing rate limiting
  - Insufficient logging and monitoring
  - Weak password policies
  - Missing security headers
  - Information disclosure through error messages

### 🔵 **LOW** (P3 - Enhancement)
- **Severity**: Minor vulnerabilities with limited impact
- **Impact**: Minimal risk, mostly theoretical or requires complex attack chain
- **Timeline**: Fix within 3 months or next major release
- **Examples**:
  - Missing security headers (non-critical)
  - Outdated dependencies with no known exploits
  - Verbose error messages (non-sensitive info)
  - Cookie security flags missing

### 🟢 **INFO** (Best Practice / Recommendation)
- **Severity**: Not a vulnerability, but improvement opportunity
- **Impact**: Improves security posture and resilience
- **Timeline**: Consider for future implementation
- **Examples**:
  - Security documentation improvements
  - Security training recommendations
  - Tool and process improvements
  - Defense-in-depth enhancements

## Security Analysis Workflow

### Phase 1: System Understanding (Round 1-2)

**Information Gathering**:
- [ ] System architecture and technology stack
- [ ] User roles and access patterns
- [ ] Data types and sensitivity levels
- [ ] External integrations and APIs
- [ ] Deployment environment (cloud, on-prem, hybrid)
- [ ] Current security controls in place
- [ ] Compliance requirements (GDPR, HIPAA, SOC2, etc.)

**Questions to Ask**:
1. What type of system is this? (web app, API, mobile, distributed system)
2. What sensitive data does it handle? (PII, financial, health, etc.)
3. Who are the users and what are their roles?
4. What are your main security concerns?
5. What compliance requirements must be met?
6. What is the threat model? (who might attack and why?)

### Phase 2: Threat Modeling (Round 3-4)

**STRIDE Analysis**:
Apply the STRIDE framework to identify threats:

- **S**poofing: Can users impersonate others?
- **T**ampering: Can data be modified maliciously?
- **R**epudiation: Can actions be denied?
- **I**nformation Disclosure: Can data leak?
- **D**enial of Service: Can the system be overwhelmed?
- **E**levation of Privilege: Can users gain unauthorized access?

**Attack Surface Mapping**:
- [ ] Entry points (APIs, forms, file uploads)
- [ ] Authentication mechanisms
- [ ] Authorization boundaries
- [ ] Data flows and storage
- [ ] Third-party dependencies
- [ ] Network exposure

### Phase 3: Security Assessment (Round 5-7)

**OWASP Top 10 Check** (Web Applications):
- [ ] A01: Broken Access Control
- [ ] A02: Cryptographic Failures
- [ ] A03: Injection
- [ ] A04: Insecure Design
- [ ] A05: Security Misconfiguration
- [ ] A06: Vulnerable and Outdated Components
- [ ] A07: Identification and Authentication Failures
- [ ] A08: Software and Data Integrity Failures
- [ ] A09: Security Logging and Monitoring Failures
- [ ] A10: Server-Side Request Forgery (SSRF)

**Security Domains to Review**:

#### 🔐 Authentication & Authorization
- Identity verification mechanisms
- Password policies and storage
- Multi-factor authentication (MFA)
- Session management
- Token handling (JWT, OAuth)
- Role-based access control (RBAC)
- Principle of least privilege

#### 🛡️ Data Protection
- Encryption at rest and in transit
- Key management
- Sensitive data handling (PII, credentials)
- Data retention and deletion
- Secure backup procedures
- Database security

#### 🌐 API & Network Security
- API authentication and authorization
- Input validation and sanitization
- Rate limiting and throttling
- CORS configuration
- TLS/SSL configuration
- Certificate management
- Firewall rules

#### 💉 Injection Prevention
- SQL injection protection
- Command injection prevention
- XSS prevention
- LDAP injection
- XML/XXE injection
- Template injection

#### 🔍 Security Monitoring
- Logging sensitive operations
- Audit trails
- Intrusion detection
- Anomaly detection
- Security metrics and alerts

#### ⚙️ Configuration & Deployment
- Secrets management
- Environment variable security
- Default credentials removed
- Debug mode disabled in production
- Security headers configured
- Dependency scanning

### Phase 4: Risk Prioritization (Round 8)

**Risk Calculation**:
```
Risk Score = Likelihood × Impact

Likelihood:
- High (3): Easy to exploit, known attack vectors
- Medium (2): Requires some skill or specific conditions
- Low (1): Complex, requires rare circumstances

Impact:
- High (3): Critical data breach, system compromise
- Medium (2): Limited data exposure, service degradation
- Low (1): Minimal impact, isolated issue

Final Risk Level:
- 7-9 points: 🔴 CRITICAL
- 5-6 points: 🟠 HIGH
- 3-4 points: 🟡 MEDIUM
- 1-2 points: 🔵 LOW
```

### Phase 5: Security Recommendations (Round 9-10)

**Recommendation Template**:
For each identified vulnerability, provide:
```markdown
## [Risk Level Icon] [Vulnerability Title]

**Risk Level**: [CRITICAL/HIGH/MEDIUM/LOW/INFO]
**Category**: [e.g., Authentication, Injection, Data Protection]
**CWE**: [CWE-XXX if applicable]

### Description
[Clear explanation of the vulnerability]

### Location
- File: `path/to/file.ext`
- Function/Endpoint: `functionName()` or `/api/endpoint`
- Line: [if applicable]

### Threat Scenario
[How an attacker could exploit this]

### Impact
- **Confidentiality**: [High/Medium/Low/None]
- **Integrity**: [High/Medium/Low/None]
- **Availability**: [High/Medium/Low/None]

### Likelihood
[High/Medium/Low] - [Explanation]

### Proof of Concept (if applicable)
[Example attack or test case]

### Recommendations
1. **Primary Solution**: [Most effective fix]
   - Implementation steps
   - Code example if helpful
   
2. **Alternative Solutions**: 
   - [Option 2]
   - [Option 3]

3. **Quick Win** (if applicable): [Temporary mitigation]

### Prevention
[How to prevent similar issues in the future]

### References
- [Relevant OWASP guide]
- [CWE link]
- [Best practice documentation]
```

## Output Format

### Security Assessment Report

```markdown
# Security Assessment Report: [System Name]

**Date**: [Date]
**Assessed By**: AI Security Analysis
**System Version**: [Version]
**Scope**: [What was reviewed]

---

## Executive Summary

### Risk Overview
- 🔴 Critical Issues: [X]
- 🟠 High Issues: [Y]
- 🟡 Medium Issues: [Z]
- 🔵 Low Issues: [W]
- 🟢 Recommendations: [V]

### Top 3 Security Concerns
1. [Most critical issue]
2. [Second critical issue]
3. [Third critical issue]

### Overall Security Posture
[Brief assessment: Strong/Adequate/Needs Improvement/Weak]

---

## 1. System Overview

### Architecture
[Brief description of system architecture]

### Assets
- **Critical Assets**: [List]
- **Sensitive Data**: [Types of data]
- **User Roles**: [List]

### Threat Actors
- **External**: [e.g., hackers, competitors]
- **Internal**: [e.g., malicious employees]
- **Accidental**: [e.g., user errors]

---

## 2. Critical Findings (🔴 P0)

[List all critical vulnerabilities using the template above]

---

## 3. High Priority Findings (🟠 P1)

[List all high-priority vulnerabilities]

---

## 4. Medium Priority Findings (🟡 P2)

[List all medium-priority vulnerabilities]

---

## 5. Low Priority Findings (🔵 P3)

[List all low-priority vulnerabilities]

---

## 6. Best Practice Recommendations (🟢)

[List security enhancements and improvements]

---

## 7. Security Architecture Recommendations

### Proposed Security Controls

#### Preventive Controls
- [Controls to prevent security incidents]

#### Detective Controls
- [Controls to detect security incidents]

#### Corrective Controls
- [Controls to respond to security incidents]

### Defense in Depth Strategy
[Multi-layer security approach]

### Zero Trust Principles
[How to implement zero trust architecture]

---

## 8. Compliance & Standards

### OWASP Top 10 Compliance
[Checklist with status]

### Regulatory Compliance
- [ ] GDPR (if applicable)
- [ ] HIPAA (if applicable)
- [ ] PCI DSS (if applicable)
- [ ] SOC 2 (if applicable)

---

## 9. Implementation Roadmap

### Phase 1: Immediate (0-1 week)
- [ ] Fix all CRITICAL issues
- [ ] [Specific tasks]

### Phase 2: Short-term (1-4 weeks)
- [ ] Fix all HIGH issues
- [ ] [Specific tasks]

### Phase 3: Medium-term (1-3 months)
- [ ] Fix all MEDIUM issues
- [ ] Implement key security controls
- [ ] [Specific tasks]

### Phase 4: Long-term (3-6 months)
- [ ] Fix LOW issues
- [ ] Implement best practices
- [ ] [Specific tasks]

---

## 10. Security Testing Recommendations

### Automated Testing
- [ ] Static Application Security Testing (SAST)
- [ ] Dynamic Application Security Testing (DAST)
- [ ] Dependency scanning (SCA)
- [ ] Container security scanning

### Manual Testing
- [ ] Penetration testing
- [ ] Security code review
- [ ] Threat modeling workshop

### Continuous Security
- [ ] Security CI/CD integration
- [ ] Regular vulnerability scanning
- [ ] Security metrics dashboard

---

## 11. Security Checklist for Development

### Design Phase
- [ ] Threat model created
- [ ] Security requirements defined
- [ ] Privacy impact assessed

### Implementation Phase
- [ ] Secure coding guidelines followed
- [ ] Input validation implemented
- [ ] Output encoding applied
- [ ] Authentication/authorization enforced
- [ ] Secrets not hardcoded

### Testing Phase
- [ ] Security tests included
- [ ] Penetration testing performed
- [ ] Dependency vulnerabilities checked

### Deployment Phase
- [ ] Security headers configured
- [ ] TLS/SSL properly configured
- [ ] Secrets properly managed
- [ ] Logging and monitoring enabled
- [ ] Security incident response plan ready

---

## 12. Appendix

### Tools Recommended
- [Security testing tools]
- [Monitoring tools]
- [Secrets management tools]

### References
- [Relevant security standards]
- [Best practice guides]
- [Training resources]

### Glossary
[Security terms used in the report]
```

## Usage Examples

### Example 1: Web Application Security Review
```
User: /security-design
Review the security of our e-commerce web application. We handle credit cards and user data.

AI: I'll perform a comprehensive security assessment of your e-commerce application.

First, let me understand your system:
1. What payment processor do you use? (Stripe, PayPal, direct card processing?)
2. How is user authentication implemented?
3. What framework/technology stack?
4. Do you store credit card data, or is it tokenized?
5. What are your PCI DSS requirements?

[After gathering info]

I've identified several security concerns. Here are the most critical:

🔴 CRITICAL: SQL Injection Vulnerability in Search Function
**Risk Level**: CRITICAL (P0)
**Location**: `api/search.js:45`
...
```

### Example 2: API Security Assessment
```
User: /security-design
We're launching a REST API for our mobile app. Can you review the security?

AI: I'll assess your API security comprehensively.

Questions:
1. What authentication method? (JWT, OAuth 2.0, API keys?)
2. What sensitive data is exposed through the API?
3. Is there rate limiting?
4. What's your API versioning strategy?

[After analysis]

Security Assessment:

🔴 CRITICAL Issues: 1
🟠 HIGH Issues: 3
🟡 MEDIUM Issues: 5

Top Concerns:
1. Missing authentication on `/api/users` endpoint (CRITICAL)
2. No rate limiting - vulnerable to DoS (HIGH)
3. Sensitive data in logs (HIGH)
...
```

### Example 3: Cloud Infrastructure Security
```
User: /security-design
Review our AWS infrastructure security. We use ECS, RDS, and S3.

AI: I'll review your AWS infrastructure security posture.

Current Setup Analysis:
- ECS containers
- RDS database
- S3 storage
- [analyzing configuration]

🔴 CRITICAL: S3 Buckets Publicly Accessible
**Risk Level**: CRITICAL
**Impact**: All stored data is accessible to anyone
**Recommendation**: Immediately set bucket policies to private
...

🟠 HIGH: RDS Security Groups Too Permissive
**Risk Level**: HIGH
**Impact**: Database accessible from any IP
...
```

### Example 4: Secure Architecture Design
```
User: /security-design
Help me design a secure architecture for a healthcare application handling patient data.

AI: I'll help you design a secure, HIPAA-compliant architecture.

Key Requirements:
- HIPAA compliance mandatory
- PHI (Protected Health Information) handling
- Strong access controls
- Audit logging
- Encryption everywhere

Proposed Secure Architecture:

1. **Network Security**:
   - VPC with private subnets
   - WAF in front of application
   - DDoS protection
   ...

2. **Data Protection**:
   - End-to-end encryption
   - Encryption at rest (AES-256)
   - Key management via HSM/KMS
   ...

[Detailed design with security controls]
```

## Security Frameworks & Standards

### STRIDE Threat Model
- **Spoofing**: Identity verification
- **Tampering**: Data integrity
- **Repudiation**: Audit trails
- **Information Disclosure**: Data confidentiality
- **Denial of Service**: Availability
- **Elevation of Privilege**: Access control

### OWASP Resources
- OWASP Top 10 (Web/API/Mobile)
- OWASP ASVS (Application Security Verification Standard)
- OWASP Testing Guide
- OWASP Cheat Sheet Series

### Compliance Standards
- **GDPR**: EU data protection
- **HIPAA**: Healthcare data (US)
- **PCI DSS**: Payment card data
- **SOC 2**: Service organization controls
- **ISO 27001**: Information security management

### Security Principles

#### Principle of Least Privilege
Grant minimum permissions necessary

#### Defense in Depth
Multiple layers of security controls

#### Fail Securely
Default to secure state on failure

#### Zero Trust
Never trust, always verify

#### Security by Design
Build security in from the start, not bolt on later

## Important Guidelines

### DO:
- ✅ Prioritize findings by actual risk (likelihood × impact)
- ✅ Provide actionable, specific recommendations
- ✅ Include proof of concept or test cases
- ✅ Consider business context and constraints
- ✅ Suggest both quick wins and long-term solutions
- ✅ Reference security standards (OWASP, CWE)
- ✅ Think like an attacker (threat modeling)
- ✅ Consider entire attack surface
- ✅ Provide implementation examples

### DON'T:
- ❌ Report theoretical issues without practical impact
- ❌ Overwhelm with too many low-priority findings
- ❌ Provide generic advice without specifics
- ❌ Ignore business requirements and constraints
- ❌ Skip the "why" - explain the risk clearly
- ❌ Forget to consider defense in depth
- ❌ Only focus on code - consider infrastructure, process, people
- ❌ Assume perfect security is achievable - balance risk and cost

## Advanced Security Topics

### Secure Code Review Checklist
```
Authentication & Session Management:
- [ ] Password hashing (bcrypt, Argon2)
- [ ] MFA available
- [ ] Session timeout configured
- [ ] Secure session storage
- [ ] CSRF protection

Input Validation:
- [ ] All inputs validated server-side
- [ ] Whitelist validation used
- [ ] Input length limits enforced
- [ ] Special characters handled

Output Encoding:
- [ ] XSS prevention (HTML encoding)
- [ ] JSON encoding
- [ ] URL encoding
- [ ] Context-aware encoding

Cryptography:
- [ ] Strong algorithms (AES-256, RSA-2048+)
- [ ] No hardcoded keys
- [ ] Proper key storage (KMS, Vault)
- [ ] Random IV generation
- [ ] Secure random number generation

Error Handling:
- [ ] No sensitive info in errors
- [ ] Generic error messages to users
- [ ] Detailed logs (server-side only)
- [ ] No stack traces exposed

API Security:
- [ ] Authentication required
- [ ] Authorization checked
- [ ] Rate limiting implemented
- [ ] Input validation
- [ ] API versioning
```

### Security Testing Strategy
1. **SAST** (Static Analysis): Find vulnerabilities in source code
2. **DAST** (Dynamic Analysis): Test running application
3. **IAST** (Interactive Analysis): Combine SAST + DAST
4. **SCA** (Software Composition Analysis): Check dependencies
5. **Penetration Testing**: Manual security testing
6. **Red Team**: Adversarial testing

### Incident Response Preparation
- Define security incident types
- Establish incident response team
- Create communication plan
- Document response procedures
- Regular incident response drills
- Post-incident review process

## Integration with Other Commands

- Use `/design` for overall system architecture with security in mind
- Use `/bug-fix` for fixing identified security vulnerabilities
- Use `/requirements-to-implementation/*` when implementing security features
- Use `/review` for security-focused code review

## When to Use This Command

Use `/security-design` when:
- Starting a new project requiring security design
- Reviewing existing application security
- Preparing for security audit or compliance review
- Responding to security incidents
- Implementing security best practices
- Threat modeling a system
- Need prioritized security recommendations

## When NOT to Use This Command

Don't use `/security-design` when:
- Performing general code review (use `/review`)
- Fixing non-security bugs (use `/bug-fix`)
- Designing non-security aspects (use `/design`)
- Just need quick security advice (use `/ask`)

---

**Remember**: Security is not a one-time activity but a continuous process. Regular security reviews, updates, and monitoring are essential for maintaining a secure system.

