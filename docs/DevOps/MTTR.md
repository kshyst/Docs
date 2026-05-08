# MTTR: Mean Time To... (The Four Definitions)

In incident management and maintenance, **MTTR** is a critical metric for measuring operational efficiency. However, the "R" can represent four different stages of the incident lifecycle.

---

## 🔍 The Four Faces of MTTR

### 1. Mean Time to Repair
The average time spent **actively fixing** a system once the failure is identified.
* **Focus:** Technical efficiency and technician skill.
* **Calculation:** `Total Repair Time / Number of Incidents`
* **Excludes:** Detection time or time spent waiting for parts/approvals.

### 2. Mean Time to Recovery (or Restore)
The average time from the **start of a failure** until the service is back in production.
* **Focus:** Total downtime impact on the end user.
* **Includes:** Detection, diagnosis, repair, and testing.
* **Key Goal:** This is the primary metric for High Availability (HA) teams.

### 3. Mean Time to Respond
The average time it takes for a team to **acknowledge an alert** and begin working.
* **Focus:** On-call effectiveness and alert routing.
* **Formula:** Time from `Alert Triggered` to `Work Started`.

### 4. Mean Time to Resolve
The average time to not only fix the issue but **ensure it won't happen again**.
* **Focus:** Long-term system stability.
* **Includes:** The actual fix plus root cause analysis (RCA) and permanent patches.

---

## 🚀 Comparison Summary

| Metric | Start Point | End Point | Primary Goal |
| :--- | :--- | :--- | :--- |
| **Repair** | Fix starts | System functional | Faster "wrench time" |
| **Recovery** | Failure occurs | User can use system | Minimize downtime |
| **Respond** | Alert sounds | Engineer takes ticket | Faster communication |
| **Resolve** | Failure occurs | Root cause eliminated | System hardening |

---

## 💡 Why It Matters
In modern DevOps and SRE (Site Reliability Engineering), lowering MTTR is often more vital than increasing the time between failures.
* **High MTTR:** Suggests poor monitoring, complex manual processes, or siloed teams.
* **Low MTTR:** Suggests strong automation, clear playbooks, and rapid deployment pipelines.
