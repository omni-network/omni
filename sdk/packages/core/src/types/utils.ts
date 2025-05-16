/**
 * @description Create partial of T, but w/ required keys K
 *
 * @example
 * PartialBy<{ foo: string, bar: number }, 'foo'>
 * => { foo?: string, bar: number }
 *
 */
export type PartialBy<T, K extends keyof T> = Omit<T, K> &
  ExactPartial<Pick<T, K>>

export type ExactPartial<type> = {
  [key in keyof type]?: type[key] | undefined
}
