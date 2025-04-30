import { expect, it } from 'vitest'
import { throwingProxy } from './throwingProxy.js'

interface Test {
  prop: string
  anotherProp: number
}

it('default: should throw an error when a property is accessed', () => {
  const proxy = throwingProxy<Test>()
  expect(() => proxy.prop).toThrowError(
    'Attempted to access property prop on context default value',
  )
  expect(() => proxy.anotherProp).toThrowError(
    'Attempted to access property anotherProp on context default value',
  )
  expect(() => (proxy as unknown as { random: string }).random).toThrowError(
    'Attempted to access property random on context default value',
  )
})
