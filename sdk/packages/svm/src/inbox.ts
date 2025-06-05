import { type AnchorProvider, Program } from '@coral-xyz/anchor'
import type { SolverInbox } from './idl/solver_inbox.js'
// import inboxIDL from './idl/solver_inbox.json' with { type: 'json' }

export type InboxProgram = Program<SolverInbox>

export function createInboxProgram(provider: AnchorProvider): InboxProgram {
  return new Program<SolverInbox>({} as SolverInbox, provider)
}
