import type { VercelRequest, VercelResponse } from '@vercel/node'
import { totalSupply } from './total'

export function circulatingSupply(timestamp: number): number {
  const index = cliffs.findIndex(([date]) => date > timestamp)

  // timestamp is before the first cliff
  if (index === 0) return 0

  // timestamp is after the last cliff
  if (index === -1) return totalSupply

  const [_, supply] = cliffs[index - 1]
  return supply
}

const cliffs = [
  [Date.parse('4/17/24'), 9_895_000],
  [Date.parse('5/17/24'), 10_509_583.33],
  [Date.parse('6/17/24'), 11_124_166.67],
  [Date.parse('7/17/24'), 12_014_112.32],
  [Date.parse('8/17/24'), 12_904_057.97],
  [Date.parse('9/17/24'), 13_794_003.62],
  [Date.parse('10/17/24'), 14_683_949.28],
  [Date.parse('11/17/24'), 15_573_894.93],
  [Date.parse('12/17/24'), 16_463_840.58],
  [Date.parse('1/17/25'), 17_353_786.23],
  [Date.parse('2/17/25'), 18_243_731.88],
  [Date.parse('3/17/25'), 19_133_677.54],
  [Date.parse('4/17/25'), 36_003_068.19],
  [Date.parse('5/17/25'), 36_893_013.84],
  [Date.parse('6/17/25'), 37_782_959.49],
  [Date.parse('7/17/25'), 38_672_905.14],
  [Date.parse('8/17/25'), 39_562_850.8],
  [Date.parse('9/17/25'), 40_452_796.45],
  [Date.parse('10/17/25'), 49_332_464.1],
  [Date.parse('11/17/25'), 50_222_409.75],
  [Date.parse('12/17/25'), 51_112_355.41],
  [Date.parse('1/17/26'), 52_002_301.06],
  [Date.parse('2/17/26'), 52_892_246.71],
  [Date.parse('3/17/26'), 53_782_192.36],
  [Date.parse('4/17/26'), 62_661_860.01],
  [Date.parse('5/17/26'), 63_551_805.67],
  [Date.parse('6/17/26'), 64_441_751.32],
  [Date.parse('7/17/26'), 65_331_696.97],
  [Date.parse('8/17/26'), 66_221_642.62],
  [Date.parse('9/17/26'), 67_111_588.28],
  [Date.parse('10/17/26'), 75_991_255.93],
  [Date.parse('11/17/26'), 76_881_201.58],
  [Date.parse('12/17/26'), 77_771_147.23],
  [Date.parse('1/17/27'), 78_661_092.88],
  [Date.parse('2/17/27'), 79_551_038.54],
  [Date.parse('3/17/27'), 80_440_984.19],
  [Date.parse('4/17/27'), 89_320_651.84],
  [Date.parse('5/17/27'), 90_210_597.49],
  [Date.parse('6/17/27'), 91_100_543.14],
  [Date.parse('7/17/27'), 91_990_488.8],
  [Date.parse('8/17/27'), 92_880_434.45],
  [Date.parse('9/17/27'), 93_770_380.1],
  [Date.parse('10/17/27'), 94_660_325.75],
  [Date.parse('11/17/27'), 95_550_271.41],
  [Date.parse('12/17/27'), 96_440_217.06],
  [Date.parse('1/17/28'), 97_330_162.71],
  [Date.parse('2/17/28'), 98_220_108.36],
  [Date.parse('3/17/28'), 99_110_054.01],
  [Date.parse('4/17/28'), 100_000_000],
] as const

export default function handler(_: VercelRequest, res: VercelResponse) {
  res.status(200).json(circulatingSupply(Date.now()))
}
