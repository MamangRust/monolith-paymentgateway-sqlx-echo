# Domain Layer

This directory contains domain-level definitions used across the application. It is divided into three main parts:

- `record/`: Internal representations of data (similar to entities or data models).
- `requests/`: Request payload models for APIs (used for validation and binding).
- `response/`: Response payload models returned to clients (including structured success and error responses).

Each subdirectory is modularized by domain (auth, user, role, etc.) for scalability and ease of maintenance.
