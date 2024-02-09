package bindings

// abigen generates types for library structs with the format `{LibraryName}{StructName}`.
// leading to some awkward go code like `bindings.ValidatorsValidator`
// we rename those types here to make them more idiomatic

type Validator = ValidatorsValidator
type ValidatorSigTuple = ValidatorsSigTuple
