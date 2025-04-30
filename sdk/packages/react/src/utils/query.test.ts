import { expect, it } from 'vitest'
import { hashFn } from './query.js'

it('default: should hash query keys consistently', () => {
  const key1 = ['test']
  const key2 = ['test']
  expect(hashFn(key1)).toBe(hashFn(key2))
  expect(hashFn(key1)).toBe('["test"]')
})

it('params: should hash query keys with numbers', () => {
  const key = ['test', 1]
  expect(hashFn(key)).toBe('["test",1]')
})

it('params: should hash query keys with booleans', () => {
  const key = ['test', true, false]
  expect(hashFn(key)).toBe('["test",true,false]')
})

it('params: should hash query keys with null and undefined', () => {
  const key = ['test', null, undefined]
  expect(hashFn(key)).toBe('["test",null,null]')
})

it('params: should sort keys of plain objects', () => {
  const key1 = ['test', { id: 1, name: 'foo' }]
  const key2 = ['test', { name: 'foo', id: 1 }]
  expect(hashFn(key1)).toBe(hashFn(key2))
  expect(hashFn(key1)).toBe('["test",{"id":1,"name":"foo"}]')
})

it('params: should sort keys of nested plain objects', () => {
  const key1 = ['test', { foo: { name: 'bar', id: 1 }, bar: { foo: 'test' } }]
  const key2 = ['test', { bar: { foo: 'test' }, foo: { id: 1, name: 'bar' } }]
  expect(hashFn(key1)).toBe(hashFn(key2))
  expect(hashFn(key1)).toBe(
    '["test",{"bar":{"foo":"test"},"foo":{"id":1,"name":"bar"}}]',
  )
})

it('params: should handle empty objects', () => {
  const key = ['data', {}]
  expect(hashFn(key)).toBe('["data",{}]')
})

it('params: should convert BigInt values to strings', () => {
  const key = ['test', { amount: BigInt(9007199254740991), id: 123n }]
  expect(hashFn(key)).toBe('["test",{"amount":"9007199254740991","id":"123"}]')
})

it('params: should not sort keys of arrays within objects', () => {
  const key = ['test', { ids: [3, 1, 2] }]
  expect(hashFn(key)).toBe('["test",{"ids":[3,1,2]}]')
})

it('params: should not treat class instances as plain objects (no key sorting)', () => {
  class TestClass {
    b = 2
    a = 1
    // custom toJson shouldn't be bypassed by hashFn
    toJSON() {
      return '123'
    }
  }
  const inst = new TestClass()
  const key = ['test', inst]
  // falls back toJSON for class instances
  expect(hashFn(key)).toBe('["test","123"]')
})

it('params: should handle objects created with Object.create(null)', () => {
  const obj = Object.create(null)
  obj.b = 2
  obj.a = 1
  const key = ['nullProto', obj]
  expect(hashFn(key)).toBe('["nullProto",{"a":1,"b":2}]')
})

it('params: should handle complex nested structures', () => {
  const key = [
    'complex',
    123,
    {
      c: 'foo',
      a: [1, { y: null, x: true }],
      b: { nested: BigInt(10), another: 'val' },
    },
    true,
  ]
  const expected =
    '["complex",123,{"a":[1,{"x":true,"y":null}],"b":{"another":"val","nested":"10"},"c":"foo"},true]'
  expect(hashFn(key)).toBe(expected)
})
