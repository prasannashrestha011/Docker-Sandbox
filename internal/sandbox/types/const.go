package sandbox_type

// SandboxState represents the lifecycle state of a sandbox.
type SandboxState string

const (
	StateActive   SandboxState = "active"
	StateInActive SandboxState = "inactive"
)
