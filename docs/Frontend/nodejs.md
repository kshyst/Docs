# Introduction to Node.js

Node.js is a powerful tool in the modern web development stack. While JavaScript was originally designed to run only inside web browsers, Node.js allows developers to run JavaScript directly on their computers or servers.

---

## 1. What is Node.js?

At its core, **Node.js is an open-source, cross-platform JavaScript runtime environment**. 

* **Built on Chrome's V8 Engine**: It uses Google Chrome's extremely fast V8 JavaScript engine to compile and execute JavaScript code directly into native machine code.
* **Not a Framework or Language**: Node.js is neither a programming language (it runs JavaScript) nor a framework. It is the environment in which your JavaScript code is executed.
* **Server-Side Development**: It is most commonly used to build back-end services, command-line tools (CLIs), build tools, and API servers.

---

## 2. Core Concepts (How it Works)

Node.js is uniquely suited for web scale applications because of its architecture:

### Asynchronous and Event-Driven
All APIs in Node.js are asynchronous (non-blocking). When a Node.js server makes an I/O request (such as reading a file or querying a database), it doesn't wait for the operation to complete. Instead, it moves to the next task and registers a callback. Once the I/O operation finishes, Node.js triggers the callback to handle the result.

### Single-Threaded with the Event Loop
Unlike traditional web servers (like Apache) that create a new thread for every incoming request, Node.js operates on a **single main thread**. It uses the **Event Loop** (orchestrated by the C-based library `libuv`) to coordinate asynchronous tasks. 

> [!NOTE]
> This single-threaded approach makes Node.js incredibly lightweight and efficient at handling thousands of concurrent connections (like real-time chat apps or streaming platforms), but less suitable for CPU-intensive tasks (like video encoding or heavy mathematical calculations) which can block the main thread.

---

## 3. Node.js vs. Browser JavaScript

Although both execute JavaScript, the environment and capabilities are very different:

| Feature | Browser JavaScript | Node.js |
| :--- | :--- | :--- |
| **Execution Environment** | Web browsers (Chrome, Firefox, Safari) | Computer / Server |
| **Global Object** | `window` (or `self`) | `global` (and `globalThis`) |
| **DOM / HTML Access** | Yes (can manipulate `document`, `window`, etc.) | No (does not have a GUI or DOM) |
| **File System Access** | No (for security reasons) | Yes (via the built-in `fs` module) |
| **Operating System APIs** | No | Yes (via the `os` and `process` modules) |
| **Network Requests** | `fetch`, `XMLHttpRequest` | Built-in `http`, `https`, and `net` modules |
| **Package Manager** | None natively | `npm` (bundled by default) |

---

## 4. Developer Quickstart

### Running a Script
To run a JavaScript file with Node.js, create a file named `app.js` and execute it in your terminal:
```bash
node app.js
```

### Module Systems
Node.js supports two systems for splitting code into reusable modules:

1. **CommonJS (CJS)** - The traditional default:
   ```javascript
   // import
   const fs = require('fs');
   // export
   module.exports = { myFunction };
   ```
2. **ECMAScript Modules (ESM)** - The modern standard (supported by setting `"type": "module"` in `package.json`):
   ```javascript
   // import
   import fs from 'node:fs';
   // export
   export { myFunction };
   ```

### Important Global Variables
* `process`: Provides information about the running Node.js process (e.g. environment variables via `process.env`).
* `__dirname` / `__filename`: The directory name and absolute file path of the current module (CommonJS only).

---

## 5. Basic HTTP Server Example

Below is a simple web server built using the native Node.js HTTP module:

```javascript
import http from 'node:http';

const PORT = 3000;

// Create the server
const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end('Hello from Node.js!\n');
});

// Start listening for requests
server.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}/`);
});
```
