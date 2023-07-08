// Package internal is one of two magical Go package identifiers.
// Basically, Go has no real package hierarchy. However, using "internal", only packages sharing the same
// root package, can import from internal.
// This is very useful to disallow public identifiers to be imported from other modules, similar to how Java
// modules work with the "exports" keyword.
// Keep in mind, that any go-module can be imported by any from your other go modules, which may be good or evil.
// For example, if you want to couple on a kind of shared kernel, you may define them in this module outside of
// the internal package, but mind your own backwards-compatibility guidelines.
//
// Logically, this internal package represents the "super" domain of all contained bounded contexts.
//
// In conclusion, we recommend to put everything into internal, to avoid a direct coupling between domain-driven
// applications. Communication and protocols must be made explicitly and independent of the actual code.
package internal
