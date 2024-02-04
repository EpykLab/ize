package entries

// Define structs for storing different types of objects

// IP Address struct
type IPAddress struct {
	IP          string `json:"ip"`
	Description string `json:"description"`
}

// Domain Name struct
type DomainName struct {
	Name string `json:"name"`
}

// Username struct
type Username struct {
	Username string `json:"username"`
}

// Password struct
type Password struct {
	Password string `json:"password"`
}

// STS Token struct
type STSToken struct {
	Token string `json:"token"`
}

// API Key struct
type APIKey struct {
	Key string `json:"key"`
}

// Certificate struct
type Certificate struct {
	Certificate string `json:"certificate"`
}

// Vulnerability struct
type Vulnerability struct {
	Name string `json:"name"`
}

// Exploit struct
type Exploit struct {
	Name string `json:"name"`
}

// Payload struct
type Payload struct {
	Payload string `json:"payload"`
}

// Evidence struct
type Evidence struct {
	Evidence string `json:"evidence"`
}

// Screenshot struct
type Screenshot struct {
	URL string `json:"url"`
}

// Video struct
type Video struct {
	URL string `json:"url"`
}

// Log struct
type Log struct {
	Log string `json:"log"`
}

// Report struct
type Report struct {
	Report string `json:"report"`
}

// Configuration struct
type Configuration struct {
	Config string `json:"config"`
}

// Artifact struct
type Artifact struct {
	Artifact string `json:"artifact"`
}

// Backup struct
type Backup struct {
	Backup string `json:"backup"`
}

// MalwareSample struct
type MalwareSample struct {
	Sample string `json:"sample"`
}

// Incident struct
type Incident struct {
	Name string `json:"name"`
}
