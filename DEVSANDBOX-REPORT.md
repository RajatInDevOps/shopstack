# DevSandbox report — final-testt

⚪ **Not enough evidence yet** — Still to run: policy audit and build & deploy.

| | |
|---|---|
| Workspace | final-testt |
| Generated | 2026-10-27 09:28 UTC by Rajat Sharma |
| Blocking findings | **0** |
| Warnings | 24 |
| Policies passed | not run |
| Pipeline stages done | 0 / 5 |
| CI engine | not scaffolded yet |

## What ran

| Capability | Result | Found by |
|---|---|---|
| Security scan | 🟡 Warnings | Trivy + Hadolint + Checkov |
| Unit tests | — not run | — |
| Integration tests | — not run | — |

## Delivery pipeline

| Stage | Status |
|---|---|
| Detection | — not run |
| Scaffolding | — not run |
| Infrastructure | — not run |
| Build & Deploy | — not run |
| Live Gateway | — not run |

## Findings

### Warnings (24)

| Finding | Capability | Found by |
|---|---|---|
| DL3066: Non-numeric user-id may not be resolvable by host system (Dockerfile line 29) | Security scan · gateway | Hadolint |
| DL3025: Use arguments JSON notation for CMD and ENTRYPOINT arguments (Dockerfile line 35) | Security scan · gateway | Hadolint |
| CVE-2025-62727 in starlette 0.41.3 — starlette: Starlette DoS via Range header merging (fixed in 0.49.1) | Security scan · gateway | Trivy |
| CVE-2026-48818 in starlette 0.41.3 — starlette: Starlette: SSRF and NTLM credential theft via UNC paths in StaticFiles on Windows (fixed in 1.1.0) | Security scan · gateway | Trivy |
| CVE-2026-54283 in starlette 0.41.3 — starlette: Starlette: request.form() limits silently ignored for application/x-www-form-urlencoded enable DoS (fixed in 1.3.1) | Security scan · gateway | Trivy |
| DL3059: Multiple consecutive `RUN` instructions. Consider consolidation. (Dockerfile line 25) | Security scan · inventory | Hadolint |
| DL3066: Non-numeric user-id may not be resolvable by host system (Dockerfile line 31) | Security scan · inventory | Hadolint |
| DL3025: Use arguments JSON notation for CMD and ENTRYPOINT arguments (Dockerfile line 34) | Security scan · inventory | Hadolint |
| CVE-2026-111943 in stdlib v1.25.14 — crypto/x509: golang: crypto/x509: Man-in-the-middle attacks via untrusted intermediate certificates (fixed in 1.26.10, 1.27.3) | Security scan · inventory | Trivy |
| CVE-2026-78667 in stdlib v1.25.14 — net/http: golang: golang: Denial of Service via crafted HTTP Range headers (fixed in 1.26.9, 1.27.2) | Security scan · inventory | Trivy |
| CVE-2026-78669 in stdlib v1.25.14 — net/http: net/http/internal/http2: golang.org/x/net/http2: net/http: Denial of Service via excessive HTTP/2 SETTINGS frames (fixed in 1.26.9, 1.27.2) | Security scan · inventory | Trivy |
| CVE-2026-97031 in stdlib v1.25.14 — crypto/tls: golang: crypto/tls: Denial of Service via multiple ECH outer extension references (fixed in 1.26.9, 1.27.2) | Security scan · inventory | Trivy |
| DL3066: Non-numeric user-id may not be resolvable by host system (Dockerfile line 19) | Security scan · orders | Hadolint |
| DL3025: Use arguments JSON notation for CMD and ENTRYPOINT arguments (Dockerfile line 25) | Security scan · orders | Hadolint |
| CVE-2026-102276 in brace-expansion 2.0.2 — brace-expansion: brace-expansion: Denial of Service via stack exhaustion from crafted brace patterns (fixed in 5.0.10, 3.0.7, 2.1.5, 1.1.19) | Security scan · orders | Trivy |
| CVE-2026-102278 in brace-expansion 2.0.2 — brace-expansion: brace-expansion: Denial of Service via uncontrolled recursion in nested brace patterns (fixed in 5.0.11, 3.0.8, 2.1.6, 1.1.20) | Security scan · orders | Trivy |
| CVE-2026-13149 in brace-expansion 2.0.2 — brace-expansion: Brace-expansion: Denial of Service due to exponential-time complexity (fixed in 5.0.7, 1.1.16, 2.1.2) | Security scan · orders | Trivy |
| CVE-2026-14257 in brace-expansion 2.0.2 — brace-expansion: Brace-expansion: Denial of Service via memory exhaustion in expand() function (fixed in 5.0.8, 3.0.3, 2.1.3, 1.1.17) | Security scan · orders | Trivy |
| CVE-2026-69152 in brace-expansion 2.0.2 — brace-expansion: DoS via unbounded intermediate arrays, bypassing the CVE-2026-14257 mitigation (fixed in 1.1.18, 2.1.4, 3.0.6, 5.0.9) | Security scan · orders | Trivy |
| CVE-2026-69192 in ip-address 10.1.0 — ip-address: ip-address: Inconsistent IP address parsing leads to Server-Side Request Forgery (SSRF) and trust-boundary bypass (fixed in 10.3.1) | Security scan · orders | Trivy |
| CVE-2026-9496 in pacote 19.0.2 — pacote: Pacote: Denial of Service via crafted spec.rawSpec value in addGitSha function (fixed in 21.5.1) | Security scan · orders | Trivy |
| CVE-2026-9496 in pacote 20.0.1 — pacote: Pacote: Denial of Service via crafted spec.rawSpec value in addGitSha function (fixed in 21.5.1) | Security scan · orders | Trivy |
| CVE-2026-33671 in picomatch 4.0.3 — picomatch: Picomatch: Regular Expression Denial of Service via crafted extglob patterns (fixed in 4.0.4, 3.0.2, 2.3.2) | Security scan · orders | Trivy |
| CVE-2026-48815 in sigstore 3.1.0 — sigstore: Sigstore: Unauthorized certificates accepted due to ignored `certificateOIDs` verification option (fixed in 4.1.1) | Security scan · orders | Trivy |

## Not checked

These capabilities have not run for this workspace: Unit tests, Integration tests.

---

_Generated by DevSandbox. Re-run the capabilities and push again to update this report._
