# Contributing to Isolyn Core

Thank you for your interest in contributing to Isolyn Core.

Isolyn Core is an infrastructure project focused on correctness,
separation guarantees, and security. Contributions are expected to meet
a high standard of technical rigor.

---

## Project Scope

This repository contains infrastructure-level components and reference
implementations for the Isolyn architecture.

It does not include:
- user-facing applications
- trading strategies or alpha logic
- managed funds or custody systems
- frontend or UI components

---

## Design Principles

All contributions must respect the following principles:

- **Separation by design**  
  Intelligence, execution, and capital attribution must remain isolated.

- **Minimal trust assumptions**  
  No component should introduce hidden trust or custody requirements.

- **Privacy by default**  
  No change may introduce address correlation, identity linkage, or
  persistent identifiers.

- **Execution integrity**  
  Execution components must prioritize correctness and reliability over
  performance optimizations.

---

## Contribution Guidelines

When contributing code or documentation:

- keep changes focused and minimal
- document architectural assumptions explicitly
- avoid introducing dependencies without justification
- include tests or rationale where applicable

All contributions should align with the documented architecture and
preserve system guarantees.

---

## Security Considerations

If your contribution touches execution, privacy, or capital flow logic,
assume adversarial conditions by default.

Do not submit sensitive security issues via public pull requests.
Refer to `SECURITY.md` for responsible disclosure.

---

## Review Process

All pull requests are subject to architectural review.

Changes that compromise separation guarantees, introduce custody risk,
or weaken privacy properties will not be accepted.

---

## Licensing

By contributing to this repository, you agree that your contributions
will be licensed under the Apache-2.0 license.
