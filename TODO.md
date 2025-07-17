## 🧱 Core Middleware Features

- Global Error Handler Middleware
  - Add `Recover()` to catch panics
  - Allow user-defined `OnError(func)` handler
- ✅ Logger Middleware
  - Log method, path, status, duration
  - Store in `middlewares/logger.go`
- ✅ Built-in Middleware Support
  - Implement `server.Use(...)`
  - Support middleware chaining
  - Allow scoped group middleware

---

## 🧰 Developer Experience Enhancements

- Static File Server
  - Add `server.Static("/public", "./assets")`
- ✅ Response Helper Methods
  - Add `res.JSON(status, data)`
  - Add `res.Text(status, text)`
  - Add `res.HTML(status, html)`
- ✅ Query & Form Helpers
  - `req.Query("key")`
  - `req.Form("key")`
  - `req.BodyJSON(&struct)`

---

## 🧱 Routing Enhancements

- Route Grouping
  - `server.Group("/api", middleware...)`
  - Return isolated `*Router` instance
- 🔜 Trie-Based Routing
  - Replace current slice-matching
  - Optimize for O(k) lookup performance

---

## 🖼️ View Support

- Template Rendering
  - Add `res.Render("file.html", data)`
  - Use Go’s `html/template` package
  - Support global layout or base templates

---

## 🔐 Session Management (Optional)

- Cookie + Memory-backed Session
  - Set & Get session values
  - Automatic cookie creation
  - Flash message support (optional)

---

## ⚡ Real-Time Capability

- WebSocket Support
  - `server.WS("/chat", handler)`
  - Simple wrapper over `net/http` `Upgrade()`
  - Broadcast to group (optional feature)

---

## 📚 Developer Tools

- Auto Route Documentation
  - Track registered routes
  - Expose `/routes` or `/docs` endpoint
  - Optional HTML rendering
- CLI Tooling (Advanced / Final Step)
  - `squirrel new myapp` – scaffold project
  - `squirrel serve` – dev server
  - `squirrel routes` – print registered routes

