# Isolyn Core

Privacy-native perpetual trading infrastructure core components.

Isolyn Core contains the open-source reference implementation of the
Isolyn architecture, focused on infrastructure primitives, interfaces,
and separation guarantees required for private participation in
perpetual markets.

This repository is part of the Isolyn Labs ecosystem.

---

## Scope

This repository focuses on infrastructure-level components, including:

- architectural primitives enforcing separation between intelligence,
  execution, and capital attribution
- interface definitions between system layers
- reference implementations of non-custodial execution patterns
- privacy-preserving capital abstraction mechanisms (reference-level)

This repository does **not** include:
- user-facing applications
- strategy logic or alpha generation
- managed funds or custody systems
- frontend or UI components

---

## Architecture Alignment

Isolyn Core follows the system design principles defined at the
organization level:

- strict separation of concerns
- minimal trust assumptions
- no AI-controlled wallets
- no persistent execution identities

All components in this repository must preserve these guarantees.

---

## Repository Structure

The repository is organized around infrastructure separation boundaries.

```text
isolyn-core/
├── core/
│   ├── intelligence/     # intelligence interfaces
││   ├── execution/       # execution abstractions
│   └── privacy/          # privacy primitives
├── interfaces/           # cross-layer interface definitions
├── docs/                 # architecture and technical documentation
├── README.md
├── LICENSE
├── SECURITY.md
└── CONTRIBUTING.md
```
---

## Development Status

Isolyn Core is under active development.

Current focus:
- defining stable cross-layer interfaces
- documenting execution and privacy guarantees
- publishing reference implementations without custody risk

Public APIs and interfaces are expected to stabilize before any
production deployment.

---

## Contribution Guidelines

Contributions are expected to prioritize:
- correctness
- security
- clarity
- preservation of separation guarantees

All changes should align with the documented architecture and must not
introduce custody, attribution leakage, or hidden trust assumptions.

---

## License

Apache-2.0

---

© Isolyn Labs
