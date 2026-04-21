You need to set the `Authorization` header in the response, and then the client stores it and sends it back on subsequent requests.

**On login/token generation, add the token to the response header:**

```go
w.Header().Set("Authorization", "Bearer " + tokenString)
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(statusCode)
json.NewEncoder(w).Encode(v)
```

**Then on the client side**, store and resend it. For example with `fetch`:

```javascript
// Login request
const res = await fetch("/login", { method: "POST", body: ... });
const token = res.headers.get("Authorization"); // "Bearer <token>"
localStorage.setItem("token", token);

// Subsequent requests
fetch("/protected", {
  headers: {
    "Authorization": localStorage.getItem("token")
  }
});
```

**Or if your client is Go:**
```go
// Store token from login response
token := resp.Header.Get("Authorization")

// Send it on subsequent requests
req.Header.Set("Authorization", token)
```

One thing to keep in mind — browsers block access to custom response headers by default due to CORS. If your frontend is on a different origin, you'll need to expose the header on your server:

```go
w.Header().Set("Access-Control-Expose-Headers", "Authorization")
```

This tells the browser it's allowed to read that header from the response.