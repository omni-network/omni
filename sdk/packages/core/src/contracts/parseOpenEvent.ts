import {
  type DecodeEventLogReturnType,
  type Log,
  decodeEventLog,
  parseEventLogs,
} from 'viem'
import { inboxABI } from '../constants/abis.js'
import { ParseOpenEventError } from '../errors/base.js'

export type ResolvedOrder = DecodeEventLogReturnType<
  typeof inboxABI,
  'Open'
>['args']['resolvedOrder']

export function parseOpenEvent(logs: Log[]): ResolvedOrder {
  try {
    const parsed = parseEventLogs({
      abi: inboxABI,
      logs,
      eventName: 'Open',
    })

    if (parsed.length !== 1) {
      throw new ParseOpenEventError(
        `Expected exactly one 'Open' event but found ${parsed.length}.`,
      )
    }

    const openLog = parsed[0]

    const openEvent = decodeEventLog({
      abi: inboxABI,
      eventName: 'Open',
      data: openLog.data,
      topics: openLog.topics,
    })

    return openEvent.args.resolvedOrder
  } catch (error) {
    if (error instanceof Error) {
      throw error
    }
    throw new ParseOpenEventError(`Failed to parse open event: ${error}`)
  }
}
