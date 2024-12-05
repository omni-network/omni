/**
 * Represents an XReceipt event - emitted when an XMsg is executed on its destination chain
 *
 * @param sourceChainId - The source chain ID the XMsg was sent
 * @param shardId - The shard ID - TODO
 * @param offset - The xMsg offset - unique identifier for this XMsg in the source -> destination XStream
 */
export type XReceipt = {
  sourceChainId: number
  shardId: number
  offset: number
}
