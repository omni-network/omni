import type { VercelRequest, VercelResponse } from '@vercel/node'

export const totalSupply = 100_000_000

export default function handler(_: VercelRequest, res: VercelResponse) {
  res.status(200).json(totalSupply)
}
