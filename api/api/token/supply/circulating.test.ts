import { expect, test } from 'vitest'
import { circulatingSupply } from './circulating'

test('circulating supply', () => {
  // pre first cliff returns 0
  let circ = circulatingSupply(Date.parse('4/16/24'))
  expect(circ).toBe(0)

  // test a couple cliffs
  circ = circulatingSupply(Date.parse('4/17/24'))
  expect(circ).toBe(9_895_000)
  circ = circulatingSupply(Date.parse('4/18/24'))
  expect(circ).toBe(9_895_000)

  circ = circulatingSupply(Date.parse('5/17/24'))
  expect(circ).toBe(10_509_583.33)
  circ = circulatingSupply(Date.parse('5/18/24'))
  expect(circ).toBe(10_509_583.33)

  circ = circulatingSupply(Date.parse('1/17/27'))
  expect(circ).toBe(78_661_092.88)
  circ = circulatingSupply(Date.parse('1/18/27'))
  expect(circ).toBe(78_661_092.88)

  // post last cliff returns total supply
  circ = circulatingSupply(Date.parse('4/18/28'))
  expect(circ).toBe(100_000_000)
})
