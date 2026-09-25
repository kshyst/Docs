# Digital Signatures and Electronic Contracts in Iran Glossary

## 1. Contracts, consent, and legal concepts

| Term | Persian equivalent | Meaning |
| --- | --- | --- |
| Electronic contract | قرارداد الکترونیکی | An agreement formed and recorded electronically. Its electronic format does not by itself establish lawful subject matter, the parties’ authority, or compliance with required formalities. |
| Electronic record / data message | سابقه الکترونیکی / داده‌پیام | Information generated, sent, received, stored, or processed electronically. A contract, acceptance event, or electronic receipt can be a data message. |
| Electronic signature | امضای الکترونیکی | An electronic mark or action connected to a record and used to identify its signer. This is a broader concept than a cryptographic digital signature. |
| Digital signature | امضای دیجیتال | A cryptographic signature created using a private key and verified using the corresponding public key. It supports verification of signed content and its connection to the signing key. |
| Secure electronic signature | امضای الکترونیکی مطمئن | The legal category defined by Article 10: unique to the signer, identifying the signer, produced by them or under their exclusive control, and linked to the data so changes are detectable. A product label does not establish compliance. |
| Secure electronic record | سابقه الکترونیکی مطمئن | Under Article 11 of the Law, an electronic record stored using a secure information system and remaining accessible and understandable when needed. |
| Consent / acceptance | رضایت / پذیرش | A person’s agreement to specified terms. A checkbox or SMS confirmation may provide evidence of acceptance, but does not automatically establish a secure electronic signature. |
| Signatory | امضاکننده | The person who signs, or acts with the relevant legal authority to sign for another party. |
| Signing intent | قصد امضا | The signer’s deliberate decision to sign the particular document. Logging in or completing KYC does not authorize all later agreements. |
| Exclusive control | کنترل یا اراده انحصاری امضاکننده | The requirement that signing occurs under the signer’s control. Protection must cover both access to the key and authorization to use it. |
| Signatory authority | اختیار امضا / حق امضا | The legal authority to bind a person or organization. Identifying a company employee does not by itself prove that they may sign for the company. |
| Evidentiary value | ارزش اثباتی | The weight a record has as evidence. Electronic form alone is not a reason to reject it; reliability matters. |
| Admissibility | قابلیت پذیرش به‌عنوان دلیل | Whether evidence may be considered in proceedings. Admissibility does not mean that the evidence conclusively proves a claim. |
| Non-repudiation | انکارناپذیری | A security objective of providing evidence against a later denial of an action. It is not a guarantee that a signature or contract can never be challenged legally. |
| Integrity | تمامیت داده | Assurance that protected content has not been altered without detection. Integrity does not prove that the original content was true or lawful. |
| Authenticity | اصالت / انتساب | Confidence that a record or action originates from the claimed source. Signature verification contributes to this, together with identity and certificate checks. |
| Confidentiality | محرمانگی | Restricting access to information. A digital signature does not automatically encrypt or hide a contract. |
| Official instrument | سند رسمی | A legal category involving legally specified authorities and formalities. A privately signed PDF does not automatically become a سند رسمی. |
| Mandatory registration | ثبت الزامی / ثبت رسمی | A separate legal requirement to register certain transactions. Electronic signing does not replace registration where it is required. |
| Investment participation agreement | قرارداد مشارکت در سرمایه‌گذاری | An agreement defining participation rights and obligations. Whether it transfers property rights depends on its actual terms and structure. |
| Counter-signature / other party’s signature | امضای طرف مقابل | In this glossary, the additional signature required from another contracting party, such as Jafund. Some cryptographic standards use “countersignature” more narrowly for a signature over an existing signature. |

Sources: [Electronic Commerce Law — Persian](https://nezamat.ir/post-34221/), [Electronic Commerce Law — English](https://wipolex-res.wipo.int/edocs/lexdocs/laws/en/ir/ir008en.html).

## 2. Iranian legal framework

| Term | Meaning |
| --- | --- |
| Electronic Commerce Law — قانون تجارت الکترونیکی | The principal law discussed here, covering electronic records, signatures, and related matters. |
| Article 10 of the Law | Defines the requirements for a secure electronic signature. |
| Articles 14–15 of the Law | Give secure records and signatures stronger evidentiary treatment; forgery and legal invalidity remain possible grounds of challenge. |
| Article 31 of the Law | Establishes the framework for certification service providers and certificate lifecycle services. |
| Article 32 of the Law | Delegates detailed establishment rules and responsibilities for certification service providers to a Cabinet-approved regulation. It is not itself a technical signing specification. |
| Article 32 implementing regulation — آیین‌نامه اجرایی ماده ۳۲ | The regulation approved on ۱۳۸۶/۰۶/۱۱ describing the certification hierarchy, licensing, responsibilities, and certificate operations. |
| Article 8 of the regulation | Sets operational duties for intermediate CAs, including correct information, exclusive signer control, record preservation, and protections for signing data. |
| Article 11 of the regulation | Addresses internal institutional certificates issued without root authorization and excludes them from that framework and Article 10 secure-signature treatment. This is different from Article 11 of the Law. |
| Article 18 of the regulation | Sets conditions for recognition of certificates issued by foreign certification authorities. Foreign-provider popularity alone does not establish recognition. |
| Article 19 of the regulation | Covers certificate revocation, including cases involving private-key compromise. |
| License / authorization — مجوز | Permission to perform a regulated role, such as operating an intermediate CA or RA. A software vendor’s business registration is not necessarily that authorization. |
| Assurance level — سطح اطمینان | A policy-defined level of confidence supported by identity checks and security controls. The ICeCD’s November 2025 presentation lists Bronze/1, Silver/2, Gold/3, and Platinum/4. The applicable policy determines their requirements and suitability. |

Sources: [Law](https://wipolex-res.wipo.int/edocs/lexdocs/laws/en/ir/ir008en.html), [Article 32 regulation](https://law-ir.ir/law/cat/1780/), [ICeCD presentation, November 2025](https://www.asiapki.org/download/presentation/202511/4.%20Iran%20-%20ICeCD%20-%20Updates.pdf). The presentation is a dated infrastructure overview, not a current license registry.

## 3. Organizations and participants

| Term | Persian equivalent | Meaning |
| --- | --- | --- |
| PKI — Public Key Infrastructure | زیرساخت کلید عمومی | The combination of technology, organizations, policies, and procedures used to manage keys, certificates, and trust. |
| Digital Certificate Policy Council | شورای سیاستگذاری گواهی الکترونیکی | The policy and oversight body in the Iranian framework. |
| CA — Certification Authority | مرکز صدور گواهی | An organization that issues and digitally signs certificates under its policies and authority. |
| Root CA | مرکز ریشه | The authority at the top of a certificate hierarchy. Its trusted certificate anchors validation of certificates below it. |
| Intermediate CA | مرکز میانی | A CA certified under a root that issues and manages certificates within its authorized scope. |
| RA — Registration Authority | دفتر ثبت‌نام | Handles applicant identity checks, documents, and certificate requests under an issuer’s rules. It is distinct from the CA that cryptographically issues the certificate. |
| Issuer | صادرکننده گواهی | The CA that signed a particular certificate. The commercial signing vendor and the issuer may be different organizations. |
| Certificate holder / subscriber | دارنده گواهی | The person or organization to whom the certificate is issued, subject to its subscriber terms. |
| Certificate subject | موضوع گواهی | The entity identified in the certificate. For an investor certificate, this should correspond to the intended investor. |
| Relying party | طرف اعتمادکننده | A party that relies on a certificate or signature. Jafund acts as a relying party when accepting an investor’s signature. |
| Signing platform / API provider | سامانه یا ارائه‌دهنده API امضا | Software or a service that presents documents and coordinates signing. It may integrate with another organization’s CA or RA services. |
| TSA — Time-Stamping Authority | مرجع مهر زمانی | A service that issues cryptographically signed timestamp tokens under a defined trust policy. Its trust and authorization must also be evaluated. |
| ICeCD | مرکز توسعه تجارت الکترونیکی | The Iranian e-Commerce Development Center, identified as the author of the November 2025 infrastructure presentation referenced here. |
| RCA | مرکز ریشه | Shorthand used here for the root certification authority; `rca.gov.ir` was the root website referenced in the discussion. |
| GICA | مرکز صدور گواهی الکترونیکی میانی عام | The Governmental General Intermediate Certification Authority referenced through `gica.ir`. Mentioning it does not establish that a particular service or certificate is suitable for Jafund. |

Sources: [Article 32 regulation](https://law-ir.ir/law/cat/1780/), [ICeCD presentation](https://www.asiapki.org/download/presentation/202511/4.%20Iran%20-%20ICeCD%20-%20Updates.pdf), [RFC 3647](https://www.rfc-editor.org/rfc/rfc3647.html).

## 4. Keys, certificates, and policies

| Term | Meaning |
| --- | --- |
| Asymmetric cryptography | Cryptography using a mathematically related public/private key pair. |
| Key pair | The related private key and public key. |
| Private key | Secret cryptographic material used to create signatures. Customer signing keys must not be available for unrestricted use by Jafund. |
| Public key | The key used to verify signatures created by the corresponding private key. It is not secret. |
| Digital certificate / electronic certificate | A signed credential binding a public key to an identified subject, with information such as issuer, validity period, and serial number. It is not itself a signed contract. |
| X.509 | A widely used certificate format and framework for public-key certificates. |
| Certificate serial number | An identifier assigned by the issuer. Interpret it together with the issuer; it is not necessarily globally unique by itself. |
| Certificate chain / certification path | The sequence linking the signer’s certificate through intermediate certificates to a trusted root. The RA and API vendor are not automatically links in this chain. |
| Trust anchor | A root key/certificate or equivalent information explicitly trusted by the verifier. Trust cannot be established merely because a document includes a root certificate. |
| Trust store | The controlled collection of trust anchors accepted by a verification system. |
| CP — Certificate Policy | The rules and requirements defining the intended use and assurance of a class of certificates. |
| CPS — Certification Practice Statement | The issuer’s description of how it operates and implements certificate requirements. |
| Certificate policy identifier / OID | An object identifier used to identify a certificate policy. The policy text explains what that identifier means. |
| Key usage / certificate purpose | Certificate fields and policies restricting what the key or certificate may be used for. A certificate usable for one purpose is not automatically suitable for contract signing. |
| Certificate validity period | The interval during which a certificate is intended to be valid, subject to revocation, policy, and other checks. |
| Expiration | Reaching the end of a certificate’s validity period. Historical signature validation requires evidence about the relevant signing time and certificate status. |
| Revocation | Cancellation of a certificate before its scheduled expiry, for example after key compromise. |
| CRL — Certificate Revocation List | An issuer-signed list of revoked certificates, published and refreshed according to policy. |
| OCSP — Online Certificate Status Protocol | A protocol for obtaining signed certificate-status information. It is one possible component of revocation checking. |
| Key compromise | Loss of assurance that a private key remains protected from unauthorized use. |
| Certificate renewal | Obtaining a new certificate for continued use under the issuer’s rules. It may be accompanied by a new key pair. |
| Re-keying | Creating a replacement key pair and obtaining the corresponding certificate. |

Sources: [RFC 5280 — X.509 certificates and CRLs](https://www.rfc-editor.org/rfc/rfc5280.html), [RFC 3647 — CP/CPS framework](https://www.rfc-editor.org/rfc/rfc3647.html), [RFC 6960 — OCSP](https://www.rfc-editor.org/rfc/rfc6960.html).

## 5. Identity, enrollment, and signing mechanisms

| Term | Meaning |
| --- | --- |
| Enrollment | The process of applying for a certificate, verifying identity, establishing keys, and issuing/accepting the credential. Enrollment is separate from signing individual contracts. |
| Identity proofing | Establishing who a person is using suitable evidence and checks. |
| KYC — Know Your Customer | Customer identity and related checks performed by a business. Existing Jafund KYC does not automatically satisfy a CA’s enrollment policy. |
| Authentication | Checking that someone controls a credential associated with an account or identity. |
| Authorization | Deciding whether an identified actor may perform a particular action. |
| Transaction authorization | Explicit authorization for a specific operation with defined details, such as a contract and investment amount. |
| OTP — One-Time Password | A short-lived code intended for a limited use. An OTP can help authenticate or authorize a signing operation; it is not automatically a cryptographic contract signature. |
| PIN | A secret code used to activate a device or credential. Knowledge of a PIN alone does not establish lawful representative authority. |
| Hardware token / smart card | A device that can protect a private key and perform signing operations after activation. |
| Mobile signing | A signing experience using a mobile device. The term alone does not specify where the key lives or how it is protected. |
| Remote signing | A signing service where cryptographic operations occur in remote protected infrastructure. Its approved policy and authorization design must establish the required signer control. |
| HSM — Hardware Security Module | Specialized hardware for protecting keys and performing cryptographic operations. An HSM does not, by itself, prove that only the customer can authorize signing. |
| Key activation | The action that permits a protected key to perform an operation, such as entering a PIN or approving a provider challenge. |
| Signing session / challenge | A limited signing attempt bound to a person and a specific document or transaction. It should expire and resist reuse for another operation. |
| What You See Is What You Sign | A design principle that the signer can identify and approve the significant content actually being signed. |

Sources: [RFC 3647](https://www.rfc-editor.org/rfc/rfc3647.html), [OWASP Transaction Authorization](https://cheatsheetseries.owasp.org/cheatsheets/Transaction_Authorization_Cheat_Sheet.html). Iranian acceptance of a particular mechanism requires its applicable approved policy and authorization evidence.

## 6. Signed documents and long-term evidence

| Term | Meaning |
| --- | --- |
| Document hash / digest | A cryptographic fingerprint of exact document bytes. It helps detect changes; by itself it proves neither who accepted the document nor when. |
| Signed artifact | The actual signed electronic file or container, including the signature data needed for verification. |
| Signature verification | Checking mathematically whether a signature matches the protected content and public key. |
| Signature validation | The broader assessment of signature verification, identity, certificate path, policy, timing, revocation evidence, and applicable acceptance rules. |
| PDF — Portable Document Format | A document format commonly used to present contracts. A PDF is not necessarily digitally signed. |
| PAdES — PDF Advanced Electronic Signatures | A family of specifications for digital signatures in PDF documents. PAdES is not automatic certification of compliance with Iranian law. |
| PAdES B-B | The basic baseline signature level. |
| PAdES B-T | A baseline signature with trusted evidence of the signature’s existence at a particular time. |
| PAdES B-LT | A level incorporating validation material, such as certificates and revocation information, to support long-term validation. |
| PAdES B-LTA | A level adding archival timestamp protection to long-term validation material. Continued preservation can require later timestamp renewal. |
| Trusted timestamp / timestamp token | Cryptographically protected evidence that particular data existed at a stated time. It does not independently prove consent or the truth of the document. |
| RFC 3161 | The specification for an Internet PKI timestamp protocol. |
| Long-term validation | Preserving enough trustworthy information to evaluate a signature later, including after certificate expiry or service changes. |
| Validation report | A record of what a verifier checked and concluded, including relevant time and policy. Preserve the underlying signed files and evidence as well. |
| Evidence package | The collected signed documents, certificates, timestamps, status evidence, identity/authority references, and relevant event records for an agreement. |
| Audit trail | A chronological record of events such as preparation, presentation, authorization, signing, verification, and delivery. |
| Immutable archive / retention protection | Storage controls designed to prevent alteration or deletion during the required retention period. Merely calling a database row “immutable” does not enforce this. |
| Append-only event record | An event log where new entries are added rather than earlier entries being rewritten. Its protection depends on enforced permissions and storage controls. |
| Delivery evidence | Records showing how and when a signed agreement was made available or delivered to a party. Delivery is distinct from signing. |

Sources: [ETSI PAdES specification](https://www.etsi.org/deliver/etsi_EN/319100_319199/31914201/01.01.01_60/en_31914201v010101p.pdf), [RFC 3161](https://www.rfc-editor.org/rfc/rfc3161.html), [RFC 5280](https://www.rfc-editor.org/rfc/rfc5280.html).

## 7. Application and Jafund implementation terms

These entries describe the proposed integration concepts discussed, not a claim that they are already implemented in Jafund.

| Term | Meaning |
| --- | --- |
| API — Application Programming Interface | The programmatic interface through which Jafund communicates with a signing provider. |
| Backend / server-side enforcement | Checks performed by the server before accepting a signature or executing an investment. A frontend checkbox alone cannot enforce them. |
| Contract template | The approved reusable agreement text and structure into which transaction-specific information is inserted. |
| Template version / contract revision | An identifier for the exact wording or document revision. Changed material terms require the appropriate new acceptance or signature. |
| Transaction snapshot | A preserved set of transaction details used to prepare the contract, such as investor, plan, amount, pricing rule, and fees. |
| Contract binding | The enforced connection between the signed contract and the precise investment operation it authorizes. |
| IRR | The currency code for the Iranian rial. Contracts and APIs should distinguish rials from tomans explicitly. |
| Reserved quote / quote expiry | A price or set of terms held for a defined period. This avoids signing one price and silently executing another. |
| Investment intent | A record of a requested investment before all steps required for completion have finished. |
| Payment gateway | The external payment service used to collect funds. Payment success does not independently establish contract signing. |
| Settlement | In the discussed flow, the later step that completes investment execution after the necessary payment and business checks. |
| Callback / webhook | A server-to-server notification from a provider. Verify its authenticity and correlate it with the intended operation. |
| Redirect | Navigation of the user’s browser to another page after an operation. A “success” redirect is not sufficient proof that a signature is valid. |
| Idempotency | Designing retries so that repeating the same authorized request does not create duplicate investments or signing operations. |
| Replay attack | Reusing previously valid authorization data or messages to cause an unintended operation. |
| TOCTOU — Time of Check to Time of Use | A flaw where details or conditions change between authorization and execution. Binding execution to the approved snapshot helps prevent it. |
| Proof of concept | A limited trial demonstrating the complete flow and exported evidence before committing to a provider or production integration. |

Implementation reference: [OWASP Transaction Authorization](https://cheatsheetseries.owasp.org/cheatsheets/Transaction_Authorization_Cheat_Sheet.html).

## 8. Quick distinctions

- **Certificate issuance identifies a key holder; contract signing authorizes a particular document.**
- **The CA issues the certificate; the RA performs enrollment duties; the platform coordinates the signing experience.**
- **Signature verification checks the cryptographic result; validation evaluates whether it should be accepted.**
- **A hash detects changes; a signature connects content to a key; a certificate connects a public key to an identity.**
- **A timestamp supports timing evidence; it does not independently prove intent or consent.**
- **An HSM protects keys; the authorization design determines who may cause them to sign.**
- **A valid signature does not cure an unlawful contract or replace mandatory registration.**
- **A signed document, a payment receipt, and an executed investment are separate records that should be linked.**
