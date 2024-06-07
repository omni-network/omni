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
  [Date.parse('5/17/24'), 10_509_583],
  [Date.parse('6/17/24'), 11_124_167],
  [Date.parse('7/17/24'), 12_014_112],
  [Date.parse('8/17/24'), 12_904_058],
  [Date.parse('9/17/24'), 13_794_004],
  [Date.parse('10/17/24'), 14_683_949],
  [Date.parse('11/17/24'), 15_573_895],
  [Date.parse('12/17/24'), 16_463_841],
  [Date.parse('1/17/25'), 17_353_786],
  [Date.parse('2/17/25'), 18_243_732],
  [Date.parse('3/17/25'), 19_133_678],
  [Date.parse('4/17/25'), 36_003_068],
  [Date.parse('5/17/25'), 36_893_014],
  [Date.parse('6/17/25'), 37_782_959],
  [Date.parse('7/17/25'), 38_672_905],
  [Date.parse('8/17/25'), 39_562_851],
  [Date.parse('9/17/25'), 40_452_796],
  [Date.parse('10/17/25'), 49_332_464],
  [Date.parse('11/17/25'), 50_222_410],
  [Date.parse('12/17/25'), 51_112_355],
  [Date.parse('1/17/26'), 52_002_301],
  [Date.parse('2/17/26'), 52_892_247],
  [Date.parse('3/17/26'), 53_782_192],
  [Date.parse('4/17/26'), 62_661_860],
  [Date.parse('5/17/26'), 63_551_806],
  [Date.parse('6/17/26'), 64_441_751],
  [Date.parse('7/17/26'), 65_331_697],
  [Date.parse('8/17/26'), 66_221_643],
  [Date.parse('9/17/26'), 67_111_588],
  [Date.parse('10/17/26'), 75_991_256],
  [Date.parse('11/17/26'), 76_881_202],
  [Date.parse('12/17/26'), 77_771_147],
  [Date.parse('1/17/27'), 78_661_093],
  [Date.parse('2/17/27'), 79_551_039],
  [Date.parse('3/17/27'), 80_440_984],
  [Date.parse('4/17/27'), 89_320_652],
  [Date.parse('5/17/27'), 90_210_597],
  [Date.parse('6/17/27'), 91_100_543],
  [Date.parse('7/17/27'), 91_990_489],
  [Date.parse('8/17/27'), 92_880_434],
  [Date.parse('9/17/27'), 93_770_380],
  [Date.parse('10/17/27'), 94_660_326],
  [Date.parse('11/17/27'), 95_550_271],
  [Date.parse('12/17/27'), 96_440_217],
  [Date.parse('1/17/28'), 97_330_163],
  [Date.parse('2/17/28'), 98_220_108],
  [Date.parse('3/17/28'), 99_110_054],
  [Date.parse('4/17/28'), 100_000_000],
] as const

export default function handler(_: VercelRequest, res: VercelResponse) {
  res.status(200).json(circulatingSupply(Date.now()))
}
