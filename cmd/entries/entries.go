package entries

// Define structs for storing different types of objects

// IP Address struct
type IPAddress struct {
	IP          string `json:"ip"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Domain Name struct
type DomainName struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Username struct
type Username struct {
	Username    string `json:"username"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Password struct
type Password struct {
	Password    string `json:"password"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// STS Token struct
type STSToken struct {
	Token       string `json:"token"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// API Key struct
type APIKey struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Certificate struct
type Certificate struct {
	Certificate string `json:"certificate"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Vulnerability struct
type Vulnerability struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Exploit struct
type Exploit struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Payload struct
type Payload struct {
	Payload     string `json:"payload"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Evidence struct
type Evidence struct {
	Evidence    string `json:"evidence"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Screenshot struct
type Screenshot struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Video struct
type Video struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Log struct
type Log struct {
	Log         string `json:"log"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Report struct
type Report struct {
	Report      string `json:"report"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Configuration struct
type Configuration struct {
	Config      string `json:"config"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Artifact struct
type Artifact struct {
	Artifact    string `json:"artifact"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Backup struct
type Backup struct {
	Backup      string `json:"backup"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// MalwareSample struct
type MalwareSample struct {
	Sample      string `json:"sample"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Incident struct
type Incident struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}
