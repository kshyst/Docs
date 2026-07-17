# Supply Chain Attacks

A **Supply Chain Attack** is a cyberattack that targets less secure elements in a supply network to compromise a final organization or product. Instead of attacking a high-security target directly, the attacker infiltrates a trusted third-party vendor, software library, or hardware supplier that the target relies on.

---

## 1. How It Works

A supply chain attack typically unfolds in three stages:

1. **Infiltration**: The attacker identifies and compromises a weaker link in the chain (e.g., an open-source library, a third-party IT management tool, or a vendor's update server).
2. **Integration**: The compromised component is distributed or merged into the target system. Because the source is trusted, security controls (like firewalls or code reviews) often bypass or fail to detect it.
3. **Execution**: The malicious code runs on the target's systems (or their clients' systems), allowing the attacker to steal data, establish backdoors, or deploy ransomware.

---

## 2. Common Attack Vectors in Software

Modern software relies heavily on open-source ecosystems and automated build pipelines, making it a prime target for supply chain poisoning.

### Typosquatting
Attackers publish malicious packages to public registries (like npm, PyPI, or Cargo) using names that are slight misspellings of popular packages.
* *Example*: Publishing `reqeusts` or `pythn-dateutil` to trick developers who make typos during installation.

### Dependency Confusion
Many companies use private internal packages (e.g., `@mycompany/auth`). If the private registry is misconfigured, or if an attacker registers the exact same package name on a public registry (like npmjs.com) with a higher version number (e.g., `99.0.0`), build systems may pull the public malicious version instead of the internal one.

### Upstream Poisoning (Account Hijacking)
Attackers gain control of popular open-source repositories by:
* Stealing maintainer credentials (e.g., via phishing or session hijacking).
* Social engineering (gaining the trust of a project maintainer over time to be granted commit/publish rights).

### CI/CD Pipeline Compromise
If an attacker compromises the build servers or automated delivery pipelines, they can inject malicious code directly into the compiled binaries or installation packages before they are signed and distributed.

---

## 3. Notable Real-World Examples

### SolarWinds Orion (2020)
* **What Happened**: State-sponsored attackers gained access to SolarWinds' software build system. They injected a backdoor (known as *SUNBURST*) into a signed DLL update of the Orion network management software.
* **Impact**: Over 18,000 customers, including multiple US government agencies and Fortune 500 companies, downloaded the poisoned update, giving the attackers deep, stealthy access.

### xz Utils Backdoor (2024)
* **What Happened**: An attacker spent years building a persona ("Jia Tan") to gain co-maintainer status on the widely used open-source compression library `xz`. They eventually merged an extremely sophisticated, multi-stage backdoor into the library, targeting the OpenSSH server (`sshd`) on Linux systems.
* **Impact**: Discovered by a developer noticing micro-benchmarking anomalies before the backdoored version reached most stable enterprise Linux distributions.

### event-stream npm Package (2018)
* **What Happened**: A malicious developer offered to help the owner of the popular `event-stream` package. After being granted publish access, the new maintainer added a dependency to a malicious package designed to steal Bitcoin from Copay wallet users.

---

## 4. Defensive Best Practices & Mitigation

Securing the supply chain requires a defense-in-depth approach:

### Use Lockfiles and Checksum Verification
Always commit lockfiles (e.g., `package-lock.json`, `go.sum`, `yarn.lock`) to your repository. This ensures that every build, environment, and CI/CD agent installs the exact same dependency versions and validates their cryptographic hashes.

### Software Composition Analysis (SCA)
Integrate SCA tools into your CI/CD pipeline to automatically scan dependencies for known vulnerabilities and licensing issues.
* *Tools*: `npm audit`, `snyk`, `trivy`, `OWASP Dependency-Check`.

### Namespace / Scope Isolation
When using private registries, enforce strict routing rules or namespaces (such as scoped npm packages) to ensure internal package requests are never forwarded to public registries.

### Software Bill of Materials (SBOM)
Generate and maintain an SBOM (a formal, structured record containing the details and supply chain relationships of various components used in building software) to quickly audit if your system is affected when a new vulnerability is disclosed.
