import { describe, expect, it } from 'vitest'
import { toJSON } from './util.js'

describe('toJSON', () => {
  it('default: correctly serialize types', () => {
    expect(toJSON('hello')).toBe('"hello"')
    expect(toJSON(123)).toBe('123')
    expect(toJSON(true)).toBe('true')
    expect(toJSON(null)).toBe('null')
    // bigint to hex string
    const bigInt = 12345678901234567890n
    expect(toJSON(bigInt)).toBe('"0xab54a98ceb1f0ad2"')
  })

  it('default: correctly serialize BigInt despite toJSON patch', () => {
    // default BigInt has no toJSON method
    // biome-ignore lint/suspicious/noExplicitAny: expected
    expect(typeof (BigInt(1) as any).toJSON).toBe('undefined')

    // we explicitly set toJSON on BigInt, mocking extension behaviour

    // biome-ignore lint/suspicious/noExplicitAny: expected
    ;(BigInt.prototype as any).toJSON = () => 'BigIntOverwritten'
    const bigInt = BigInt(100)
    // JSON.Stringify is using the patched toJSON on BigInt
    expect(JSON.stringify(bigInt)).toBe('"BigIntOverwritten"')

    // our toJSON bypasses this
    expect(toJSON(bigInt)).toBe('"0x64"')
  })

  it('params: convert BigInts within an array to hex strings', () => {
    const arr = [1n, 20n, 300n]
    expect(toJSON(arr)).toBe('["0x1","0x14","0x12c"]')
  })

  it('params: handle mixed types including BigInts in an array', () => {
    const arr = [1, 'hello', 100n, true, 200n]
    expect(toJSON(arr)).toBe('[1,"hello","0x64",true,"0xc8"]')
  })

  it('params: handle nested arrays with BigInts', () => {
    const arr = [1, [2, 3n], [4n, [5, 6n]]]
    expect(toJSON(arr)).toBe('[1,[2,"0x3"],["0x4",[5,"0x6"]]]')
  })

  it('params: convert BigInt values in an object to hex strings', () => {
    const obj = { count: 10n, total: 5000n }
    expect(toJSON(obj)).toBe('{"count":"0xa","total":"0x1388"}')
  })

  it('params: handle objects with mixed types including BigInts', () => {
    const obj = { id: 123, name: 'item', value: 99n, active: true }
    expect(toJSON(obj)).toBe(
      '{"id":123,"name":"item","value":"0x63","active":true}',
    )
  })

  it('params: handle nested objects with BigInts', () => {
    const obj = {
      level1: {
        val1: 1n,
        level2: {
          val2: 2n,
          text: 'nested',
        },
      },
      val3: 3n,
    }
    expect(toJSON(obj)).toBe(
      '{"level1":{"val1":"0x1","level2":{"val2":"0x2","text":"nested"}},"val3":"0x3"}',
    )
  })

  it('params: handle objects containing arrays with BigInts', () => {
    const obj = {
      name: 'data',
      values: [1n, 2, 3n],
      details: { count: 5n },
    }
    expect(toJSON(obj)).toBe(
      '{"name":"data","values":["0x1",2,"0x3"],"details":{"count":"0x5"}}',
    )
  })

  it('params: handle arrays containing objects with BigInts', () => {
    const arr = [
      { id: 1n, value: 10n },
      { id: 2, value: 20 },
      { id: 3n, value: 30n },
    ]
    expect(toJSON(arr)).toBe(
      '[{"id":"0x1","value":"0xa"},{"id":2,"value":20},{"id":"0x3","value":"0x1e"}]',
    )
  })

  it('params: handle complex nested structures with BigInts', () => {
    const complex = {
      a: [1, { b: 2n, c: [3, 4n] }],
      d: { e: 5n, f: 'text' },
      g: 6n,
      h: null,
      i: [7n],
    }
    expect(toJSON(complex)).toBe(
      '{"a":[1,{"b":"0x2","c":[3,"0x4"]}],"d":{"e":"0x5","f":"text"},"g":"0x6","h":null,"i":["0x7"]}',
    )
  })

  it('params: handle empty arrays', () => {
    expect(toJSON([])).toBe('[]')
  })

  it('params: handle empty objects', () => {
    expect(toJSON({})).toBe('{}')
  })
})
