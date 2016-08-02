package xacml

// A version match is '.'-separated, like a version string.  A
// number represents a direct numeric match.  A '*' means that
// any single number is valid.  A '+' means that any number, and
// any subsequent numbers, are valid.  In this manner, the
// following four patterns would all match the version string
// '1.2.3': '1.2.3', '1.*.3', '1.2.*' and ‘1.+'.
type VersionMatchType string

// TODO: validation
