# Progressive Delivery

**Progressive Delivery** is the evolution of Continuous Delivery. While CD focuses on making code _ready_ to ship, Progressive Delivery focuses on how that code is actually _exposed_ to users to minimize risk and maximize feedback.

In 2026, the industry has shifted away from "big bang" releases toward a "blast radius" philosophy: if a microservice update is buggy, it should only affect a tiny fraction of users and ideally fix itself automatically.

---

## 1. The Core Philosophy: "Separate Deployment from Release"

The most important rule in Progressive Delivery is that **Deployment $\neq$ Release**.

- **Deployment:** Moving the bits to the server. The code is "there" but invisible (dark).
    
- **Release:** Turning the feature "on" for users.
    

By separating these, you can deploy a risky change on a Tuesday afternoon and release it on a Wednesday morning—or release it gradually over three days.

---

## 2. Key Techniques & Strategies

### A. Canary Releases (Traffic-Based)

Named after the "canary in a coal mine," you route a small percentage of traffic (e.g., 5%) to the new version while the rest stays on the stable version.

- **Modern Standard:** You don't just watch a dashboard. You use **Analysis Templates** that query your metrics (Prometheus/Datadog). If the error rate exceeds a threshold, the system automatically triggers an **Auto-Rollback**.
    

### B. Feature Flags (User-Based)

Instead of using networking to shift traffic, you use code.

- **The Logic:** `if (user.isBetaTester) { showNewCheckout() } else { showOldCheckout() }`
    
- **Use Case:** Ideal for microservices with complex logic that can't be easily split at the load balancer. You can target specific groups (e.g., users in a certain region or internal staff).
    

### C. Blue-Green Deployment (Environment-Based)

You maintain two identical production environments ("Blue" and "Green").

- **Workflow:** You deploy to "Green" while users are on "Blue." Once Green is fully validated, you flip the router switch.
    
- **Benefit:** Provides an instantaneous fallback. If Green fails, you flip back to Blue in milliseconds.
    

---

## 3. The "Automated Gates" Workflow

In 2026, maturity is measured by how "predictive" your pipeline is. A typical automated rollout looks like this:

1. **Phase 1:** Deploy to one pod (0% traffic). Run smoke tests.
    
2. **Phase 2:** Route **5%** traffic. Wait 5 minutes.
    
    - **Gate:** Is the HTTP 500 error rate $< 0.1\%$?
        
3. **Phase 3:** Route **25%** traffic. Wait 10 minutes.
    
    - **Gate:** Is $p99$ latency $< 200ms$?
        
4. **Phase 4:** Scale to **100%** and retire the old version.
    

---

## 4. Comparison Summary

|**Strategy**|**Control Level**|**Risk Mitigation**|**Infrastructure Cost**|
|---|---|---|---|
|**Canary**|Network / Traffic|**High** (Limited blast radius)|Low (Uses same cluster)|
|**Feature Flags**|Application Code|**Very High** (User-specific)|Low (Software-defined)|
|**Blue-Green**|Infrastructure|**Moderate** (Full switch)|**High** (Double resources)|

---

## 5. The Tech Stack

- **Rollout Controllers:** **Argo Rollouts** or **Flux (Flagger)** (Native to Kubernetes).
    
- **Traffic Shifting:** **Istio**, **Linkerd**, or eBPF-based solutions like **Cilium**.
    
- **Observability:** **Prometheus** for metrics and **OpenTelemetry** for tracing.
    
- **Feature Management:** **LaunchDarkly**, **Unleash**, or **Flagsmith**.
    

### The "Final Boss" of DevOps

Progressive Delivery allows your team to move at high speed because the **system** acts as the safety net. Developers can merge code confidently, knowing that if their microservice behaves poorly, the pipeline will "kill" the release before the majority of customers even see it.

Between **Canary releases** and **Feature Flags**, which approach sounds more aligned with how your current microservices are architected?