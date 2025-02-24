export const throwingProxy = <T extends object>() =>
  new Proxy<T>({} as T, {
    get: (_, property) => {
      throw new Error(
        `Attempted to access property ${String(property)} on context default value`,
      )
    },
  })
