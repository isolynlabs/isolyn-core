# Execution Layer

This module defines executor abstractions and execution lifecycle
handling.

Execution components are responsible for placing and managing perpetual
positions using isolated, short-lived execution accounts.

Execution logic must not have visibility into user identity or capital
origin.
