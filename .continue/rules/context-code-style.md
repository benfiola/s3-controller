---
name: Context (Style)
alwaysApply: true
description: Describes the intended code style of the project
---

# (Rule) Code Style

Code should be idiomatic to Go while optimizing for clarity and long-term maintainability.
When you return to this codebase in a year, the logic should be immediately understandable
without requiring cross-file context or mental gymnastics.

**Principles:**

- Prefer explicit over implicit. Use clear variable names and straightforward control flow.
- Optimize for readability first. Only optimize for performance after profiling identifies bottlenecks.
- Favor composition and simple interfaces over complex abstractions.
- Follow Go idioms: explicit error handling, small interfaces, minimal hidden state.
- Avoid magic behavior or side effects that aren't obvious from reading a single function.

**Anti-patterns to avoid:**

- Clever one-liners that require context to understand
- Global state or configuration mutation
- Functions with multiple responsibilities (testability suffers)
- Unexplained goroutine spawning or concurrent behavior
