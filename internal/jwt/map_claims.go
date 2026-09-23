package jwt

// MapClaims is a claims type that uses map[string]interface{} for JSON decoding.
// Used as the default claims type for parsing and creating tokens.
type MapClaims map[string]interface{}
