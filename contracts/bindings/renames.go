package bindings

// abigen generates types for library structs with the format `{LibraryName}{StructName}`.
// leading to some awkward type name. we rename some types here

type Validator = XTypesValidator
type ValidatorSigTuple = XTypesSigTuple

type XMsg = XTypesMsg
type XSubmission = XTypesSubmission
type XBlockHeader = XTypesBlockHeader
