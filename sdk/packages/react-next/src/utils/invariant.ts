export function invariant(
  condition: boolean,
  message = 'Invariant violation: condition not met',
): asserts condition {
  if (!condition) {
    throw new Error(message)
  }
}
