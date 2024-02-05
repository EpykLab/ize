package entries

// Define structs for storing different types of objects

// IP Address struct
type IPAddress struct {
	Author      string `json:"author"`
	IP          string `json:"ip"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Domain Name struct
type DomainName struct {
	Author      string `json:"author"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Username struct
type Username struct {
	Author      string `json:"author"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Password struct
type Password struct {
	Author      string `json:"author"`
	Password    string `json:"password"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// STS Token struct
type STSToken struct {
	Author      string `json:"author"`
	Token       string `json:"token"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// API Key struct
type APIKey struct {
	Author      string `json:"author"`
	Key         string `json:"key"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Certificate struct
type Certificate struct {
	Author      string `json:"author"`
	Certificate string `json:"certificate"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Vulnerability struct
type Vulnerability struct {
	Author      string `json:"author"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Exploit struct
type Exploit struct {
	Author      string `json:"author"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Payload struct
type Payload struct {
	Author      string `json:"author"`
	Payload     string `json:"payload"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Evidence struct
type Evidence struct {
	Author      string `json:"author"`
	Evidence    string `json:"evidence"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Screenshot struct
type Screenshot struct {
	Author      string `json:"author"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Video struct
type Video struct {
	Author      string `json:"author"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Log struct
type Log struct {
	Author      string `json:"author"`
	Log         string `json:"log"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Report struct
type Report struct {
	Author      string `json:"author"`
	Report      string `json:"report"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Configuration struct
type Configuration struct {
	Author      string `json:"author"`
	Config      string `json:"config"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Artifact struct
type Artifact struct {
	Author      string `json:"author"`
	Artifact    string `json:"artifact"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Backup struct
type Backup struct {
	Author      string `json:"author"`
	Backup      string `json:"backup"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// MalwareSample struct
type MalwareSample struct {
	Author      string `json:"author"`
	Sample      string `json:"sample"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Incident struct
type Incident struct {
	Author      string `json:"author"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}
