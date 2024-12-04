/**
 * Represents an xMsg event - emitted when an xcall is made to the OmniPortal
 *
 * @param destinationChainId - The destination chain ID the xcall is aimed at
 * @param shardId - The shard ID - TODO
 * @param offset - The xMsg offset - unique identifier for this XMsg in the source -> destination XStream
 */
export type XMsg = {
  destinationChainId: number
  shardId: number
  offset: number
}
