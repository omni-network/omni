/**
 * Creates a proxy that throws an error when any property is accessed,
 * useful utility to ensure context isn't accessed before initialized,
 * and avoids having to define default values.
 */
export const throwingProxy = <T extends object>() =>
  new Proxy<T>({} as T, {
    get: (_, property) => {
      throw new Error(
        `Attempted to access property ${String(property)} on context default value`,
      )
    },
  })
