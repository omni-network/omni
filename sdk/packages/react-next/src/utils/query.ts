import type { QueryKey } from '@tanstack/react-query'

export function hashFn(queryKey: QueryKey): string {
  return JSON.stringify(queryKey, (_, value) => {
    // sort nested objs to protect against equal objs (k's + v's)
    // producing different hashes (e.g. if key order is different)
    if (isPlainObj(value))
      return Object.keys(value)
        .sort()
        .reduce((result, key) => {
          result[key] = value[key]
          return result
          // biome-ignore lint/suspicious/noExplicitAny: allowed here
        }, {} as any)
    if (typeof value === 'bigint') return value.toString()
    return value
  })
}

/**
 *
 * Checks for a plain object, important to ensure we only hash
 * objs as other types may have custom toJSON serlialsing that
 * we dont want to bypass
 *
 */
// biome-ignore lint/suspicious/noExplicitAny: allowed here
// biome-ignore lint/complexity/noBannedTypes: allowed here
function isPlainObj(o: any): o is Object {
  // biome-ignore lint/suspicious/noExplicitAny: allowed here
  function _hasObjectPrototype(obj: any) {
    return Object.prototype.toString.call(obj) === '[object Object]'
  }

  // filter all primitives besides objs, class insts, anything custom
  if (!_hasObjectPrototype(o)) return false

  // objs made with Object.create(null) have no constructor
  if (o.constructor === undefined) return true

  // _hasObjectPrototype is true for plain obj's (but also classes)
  if (!_hasObjectPrototype(o.constructor.prototype)) return false

  // filters since only base Object.prototype owns isPrototypeOf
  // biome-ignore lint/suspicious/noPrototypeBuiltins: allowed here
  if (!o.constructor.prototype.hasOwnProperty('isPrototypeOf')) return false

  return true
}
