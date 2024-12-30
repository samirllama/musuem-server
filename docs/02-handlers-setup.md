1. **HTTP Handler Interface Requirement**:
   - In Go, HTTP handlers must follow a specific function signature defined by the `http.HandlerFunc` interface
   - The signature must be `func(http.ResponseWriter, *http.Request)`
   - Even if we don't use the request parameter, we must include it to satisfy this interface

2. **Future Usage**:
   - While we're not using it now, the `r *http.Request` parameter contains valuable information that we'll likely need later, such as:
```go
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    // Get query parameters
    name := r.URL.Query().Get("name")

    // Get HTTP method (GET, POST, etc.)
    method := r.Method

    // Get request headers
    userAgent := r.Header.Get("User-Agent")

    // Get form data (for POST requests)
    email := r.FormValue("email")

    // Get JSON body
    var data struct{ Name string }
    json.NewDecoder(r.Body).Decode(&data)

    // Get client's IP address
    clientIP := r.RemoteAddr
}
```

Here's a more practical example that uses the request parameter:
```go
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    // Get the 'name' from query parameter (?name=John)
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Guest"
    }

    // Write personalized response
    w.Write([]byte("Hello, " + name + " 👋"))
}
```

## Practical Examples:
1.**HTTP Method & URL Information**

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // HTTP Method (GET, POST, PUT, DELETE, etc.)
    method := r.Method

    // Full URL path
    path := r.URL.Path

    // Query parameters (?key=value)
    queryParams := r.URL.Query()
    sortBy := queryParams.Get("sortBy")
    pageSize := queryParams.Get("limit")
}
```

2. **Headers & Authentication**
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Authorization token
    authToken := r.Header.Get("Authorization")

    // Content type
    contentType := r.Header.Get("Content-Type")

    // Custom headers
    apiKey := r.Header.Get("X-API-Key")

    // User agent
    userAgent := r.Header.Get("User-Agent")
}
```

3. **Request Body Processing**
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // For JSON requests
    if r.Header.Get("Content-Type") == "application/json" {
        var requestBody struct {
            Name        string `json:"name"`
            Description string `json:"description"`
        }

        // Read and parse JSON body
        err := json.NewDecoder(r.Body).Decode(&requestBody)
        if err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
    }

    // For form submissions
    if r.Method == "POST" {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Invalid form data", http.StatusBadRequest)
            return
        }

        email := r.FormValue("email")
        password := r.FormValue("password")
    }
}
```

4. **Context & Timeouts** (Very Important in Professional Settings)
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Get context with timeout
    ctx := r.Context()

    // Check if request was cancelled by client
    select {
    case <-ctx.Done():
        return
    default:
        // Continue processing
    }
}
```

5. **Client Information**
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Client's IP address
    clientIP := r.RemoteAddr

    // Real IP behind proxy
    realIP := r.Header.Get("X-Real-IP")
    forwardedFor := r.Header.Get("X-Forwarded-For")

    // Cookies
    cookie, err := r.Cookie("session_id")
    if err == nil {
        sessionID := cookie.Value
    }
}
```

6. **File Uploads**
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Max upload size
    r.ParseMultipartForm(10 << 20) // 10MB

    // Get file from form
    file, handler, err := r.FormFile("document")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Get file details
    filename := handler.Filename
    fileSize := handler.Size
    fileType := handler.Header.Get("Content-Type")
}
```

7. **Logging & Debugging** (Common in Professional Settings)
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // Common request logging pattern
    log.Printf(
        "Request: method=%s path=%s remote_addr=%s user_agent=%s",
        r.Method,
        r.URL.Path,
        r.RemoteAddr,
        r.Header.Get("User-Agent"),
    )

    // Request ID for tracing
    requestID := r.Header.Get("X-Request-ID")
    if requestID == "" {
        requestID = generateRequestID() // custom function
    }
}
```

8. **Security Headers**
```go
func Handler(w http.ResponseWriter, r *http.Request) {
    // CORS origin
    origin := r.Header.Get("Origin")

    // CSRF token
    csrfToken := r.Header.Get("X-CSRF-Token")

    // Check if request is secure
    isSecure := r.TLS != nil
}
```

Real-world Example Combining Multiple Features:
```go
func CreateMuseumHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Validate method
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 2. Validate authentication
    authToken := r.Header.Get("Authorization")
    if !strings.HasPrefix(authToken, "Bearer ") {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // 3. Parse JSON body
    var museum struct {
        Name     string `json:"name"`
        Location string `json:"location"`
    }

    if err := json.NewDecoder(r.Body).Decode(&museum); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // 4. Get context for database operations
    ctx := r.Context()

    // 5. Log the request
    log.Printf(
        "Creating museum: name=%s location=%s requestor_ip=%s",
        museum.Name,
        museum.Location,
        r.RemoteAddr,
    )

    // ... rest of the handler logic
}
```
