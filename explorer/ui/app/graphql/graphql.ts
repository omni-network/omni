/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  /** Address is a 20 byte Ethereum address, represented as 0x-prefixed hexadecimal. */
  Address: { input: any; output: any; }
  /**
   * BigInt is a large integer. Input is accepted as either a JSON number or as a string.
   * Strings may be either decimal or 0x-prefixed hexadecimal. Output values are all
   * 0x-prefixed hexadecimal.
   */
  BigInt: { input: any; output: any; }
  /**
   * Bytes is an arbitrary length binary string, represented as 0x-prefixed hexadecimal.
   * An empty byte string is represented as '0x'. Byte strings must have an even number of hexadecimal nybbles.
   */
  Bytes: { input: any; output: any; }
  /** Bytes32 is a 32 byte binary string, represented as 0x-prefixed hexadecimal. */
  Bytes32: { input: any; output: any; }
  /**
   * Long is a 64 bit unsigned integer. Input is accepted as either a JSON number or as a string.
   * Strings may be either decimal or 0x-prefixed hexadecimal.
   * Output values are all 0x-prefixed hexadecimal.
   */
  Long: { input: any; output: any; }
  /** Time is an time.RFC3339 encoded UTC date string or a string containing a Unix timestamp. */
  Time: { input: any; output: any; }
};

/** Chain represents a blockchain on the https://chainlist.org. */
export type Chain = {
  __typename?: 'Chain';
  /** Chain ID is a 0x prefixed hexadecimal number as per https://chainlist.org. */
  chainID: Scalars['BigInt']['output'];
  /** Display ID of the chain is a human readable base-10 number as per https://chainlist.org. */
  displayID: Scalars['Long']['output'];
  /** ID is a unique identifier of the chain. Not intended for display. */
  id: Scalars['ID']['output'];
  /** Chain logo URL */
  logoUrl: Scalars['String']['output'];
  /** Chain name */
  name: Scalars['String']['output'];
};

/** FilterInput represents a filter to apply to underlying entities. */
export type FilterInput = {
  key: Scalars['String']['input'];
  value: Scalars['String']['input'];
};

/** PageInfo represents pagination information */
export type PageInfo = {
  __typename?: 'PageInfo';
  /** Current page number */
  currentPage: Scalars['Long']['output'];
  /** hasNextPage is true if there is a next page */
  hasNextPage: Scalars['Boolean']['output'];
  /** hasPrevPage is true if there is a previous page */
  hasPrevPage: Scalars['Boolean']['output'];
  /** Total number of pages */
  totalPages: Scalars['Long']['output'];
};

/** The query type represents all of the read-only entry points into our object graph. */
export type Query = {
  __typename?: 'Query';
  /** Stats returns key Omni metrics. */
  stats: StatsResult;
  /** Returns the list of supported chains. */
  supportedChains: Array<Chain>;
  /** Retrieve a specific XMsg by source chain ID, destination chain ID and offset. */
  xmsg?: Maybe<XMsg>;
  /** Returns a paginated list of XMsgs based on the provided arguments. For forwards pagination, provide `first` and `after`. For backwards pagination, provide `last` and `before`. Defaults to the last 10 messages. `after` and `before` are the cursors of the last and first messages from the current page respectively. */
  xmsgs: XMsgConnection;
};


/** The query type represents all of the read-only entry points into our object graph. */
export type QueryXmsgArgs = {
  destChainID: Scalars['BigInt']['input'];
  offset: Scalars['BigInt']['input'];
  sourceChainID: Scalars['BigInt']['input'];
};


/** The query type represents all of the read-only entry points into our object graph. */
export type QueryXmsgsArgs = {
  after?: InputMaybe<Scalars['ID']['input']>;
  before?: InputMaybe<Scalars['ID']['input']>;
  filters?: InputMaybe<Array<FilterInput>>;
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
};

/** StatsResult represents the result of the stats query. */
export type StatsResult = {
  __typename?: 'StatsResult';
  /** The top streams by total messages sent. */
  topStreams: Array<StreamStats>;
  /** Total number of XMsgs */
  totalMsgs: Scalars['Long']['output'];
};

/** Status represents the status of an XMsg. */
export enum Status {
  Failed = 'FAILED',
  Pending = 'PENDING',
  Success = 'SUCCESS'
}

/** StreamStats represents the metrics of a stream. */
export type StreamStats = {
  __typename?: 'StreamStats';
  /** Destination chain ID */
  destChain: Chain;
  /** Number of messages sent in the stream */
  msgCount: Scalars['Long']['output'];
  /** Source chain ID */
  sourceChain: Chain;
};

/** XBlock represents a cross-chain block. */
export type XBlock = {
  __typename?: 'XBlock';
  /** The chain where the block was created. */
  chain: Chain;
  /** Hash of the source chain block */
  hash: Scalars['Bytes32']['output'];
  /** Height of the source chain block */
  height: Scalars['BigInt']['output'];
  /** ID of the block */
  id: Scalars['ID']['output'];
  /** All cross-chain messages sent/emittted in the block */
  messages: Array<XMsg>;
  /** Timestamp of the source chain block */
  timestamp: Scalars['Time']['output'];
  /** URL to view the block on the source chain */
  url: Scalars['String']['output'];
};

/** XMsg is a cross-chain message. */
export type XMsg = {
  __typename?: 'XMsg';
  /** XBlock message was emitted in */
  block: XBlock;
  /** Destination chain where the message was sent to */
  destChain: Chain;
  /** Display ID of the XMsg in the form of `<srcChainID>-<destChainID>-<offset>` */
  displayID: Scalars['String']['output'];
  /** Gas limit to use for 'call' on destination chain */
  gasLimit: Scalars['BigInt']['output'];
  /** ID of the XMsg */
  id: Scalars['ID']['output'];
  /** Monotonically incremented offset of XMsg on the source > dest stream */
  offset: Scalars['BigInt']['output'];
  /** Receipts of the message if available */
  receipt?: Maybe<XReceipt>;
  /** The address of the sender on the source chain. */
  sender: Scalars['Address']['output'];
  /** URL to view the address of the sender on the source chain. */
  senderUrl: Scalars['String']['output'];
  /** Source chain where the message was emitted */
  sourceChain: Chain;
  /** Status of the X message */
  status: Status;
  /** Target/To address to 'call' on the destination chain */
  to: Scalars['Address']['output'];
  /** URL to view the address of the target on the destination chain. */
  toUrl: Scalars['String']['output'];
  /** Hash of the source chain transaction that emitted the message */
  txHash: Scalars['Bytes32']['output'];
  /** URL to view the transaction on the source chain */
  txUrl: Scalars['String']['output'];
};

/** XMsgConnection represents a paginated list of XMsgs */
export type XMsgConnection = {
  __typename?: 'XMsgConnection';
  /** edges is a list of XMsgs with cursor information */
  edges: Array<XMsgEdge>;
  /** Page Info contains pagination information necessary for UI pagination */
  pageInfo: PageInfo;
  /** Total number of XMsgs */
  totalCount: Scalars['Long']['output'];
};

/** XMessageEdge represents a single XMsg in a paginated list */
export type XMsgEdge = {
  __typename?: 'XMsgEdge';
  /** Cursor */
  cursor: Scalars['ID']['output'];
  /** XMsg */
  node: XMsg;
};

/** XReceipt represents a cross-chain receipt. */
export type XReceipt = {
  __typename?: 'XReceipt';
  /** Destination chain where the message was sent to */
  destChain: Chain;
  /** Gas used for the cross-chain message */
  gasUsed: Scalars['BigInt']['output'];
  /** ID of the receipt */
  id: Scalars['ID']['output'];
  /** Monotonically incremented offset of Msg in the Stream */
  offset: Scalars['BigInt']['output'];
  /** Address of the relayer */
  relayer: Scalars['Address']['output'];
  /** Revert reason if the message failed */
  revertReason?: Maybe<Scalars['String']['output']>;
  /** Source chain where the message was emitted */
  sourceChain: Chain;
  /** Success indicates whether the message was successfully executed on the destination chain. */
  success: Scalars['Boolean']['output'];
  /** Timestamp of the receipt */
  timestamp: Scalars['Time']['output'];
  /** Hash of the dest chain transaction */
  txHash: Scalars['Bytes32']['output'];
  /** URL to view the transaction on the destination chain */
  txUrl: Scalars['String']['output'];
};

export type SupportedChainsQueryVariables = Exact<{ [key: string]: never; }>;


export type SupportedChainsQuery = { __typename?: 'Query', supportedChains: Array<{ __typename?: 'Chain', id: string, chainID: any, name: string, logoUrl: string }> };

export type ChainStatsQueryVariables = Exact<{ [key: string]: never; }>;


export type ChainStatsQuery = { __typename?: 'Query', stats: { __typename?: 'StatsResult', totalMsgs: any, topStreams: Array<{ __typename?: 'StreamStats', msgCount: any, sourceChain: { __typename?: 'Chain', id: string, chainID: any, displayID: any, name: string, logoUrl: string }, destChain: { __typename?: 'Chain', id: string, chainID: any, displayID: any, name: string, logoUrl: string } }> } };

export type XmsgQueryVariables = Exact<{
  sourceChainID: Scalars['BigInt']['input'];
  destChainID: Scalars['BigInt']['input'];
  offset: Scalars['BigInt']['input'];
}>;


export type XmsgQuery = { __typename?: 'Query', xmsg?: { __typename?: 'XMsg', id: string, displayID: string, offset: any, sender: any, senderUrl: string, to: any, toUrl: string, gasLimit: any, txHash: any, txUrl: string, status: Status, sourceChain: { __typename?: 'Chain', chainID: any, logoUrl: string, name: string }, destChain: { __typename?: 'Chain', chainID: any, logoUrl: string, name: string }, block: { __typename?: 'XBlock', height: any, hash: any, timestamp: any }, receipt?: { __typename?: 'XReceipt', revertReason?: string | null, txHash: any, txUrl: string, relayer: any, timestamp: any, gasUsed: any } | null } | null };

export type XmsgsQueryVariables = Exact<{
  first?: InputMaybe<Scalars['Int']['input']>;
  last?: InputMaybe<Scalars['Int']['input']>;
  after?: InputMaybe<Scalars['ID']['input']>;
  before?: InputMaybe<Scalars['ID']['input']>;
  filters?: InputMaybe<Array<FilterInput> | FilterInput>;
}>;


export type XmsgsQuery = { __typename?: 'Query', xmsgs: { __typename?: 'XMsgConnection', totalCount: any, edges: Array<{ __typename?: 'XMsgEdge', cursor: string, node: { __typename?: 'XMsg', id: string, txHash: any, offset: any, displayID: string, sender: any, senderUrl: string, to: any, toUrl: string, gasLimit: any, status: Status, txUrl: string, sourceChain: { __typename?: 'Chain', chainID: any, logoUrl: string, name: string }, destChain: { __typename?: 'Chain', chainID: any, logoUrl: string, name: string }, block: { __typename?: 'XBlock', id: string, hash: any, height: any, timestamp: any, chain: { __typename?: 'Chain', chainID: any } }, receipt?: { __typename?: 'XReceipt', txHash: any, txUrl: string, timestamp: any, success: boolean, offset: any, relayer: any, revertReason?: string | null, sourceChain: { __typename?: 'Chain', chainID: any }, destChain: { __typename?: 'Chain', chainID: any } } | null } }>, pageInfo: { __typename?: 'PageInfo', currentPage: any, totalPages: any, hasNextPage: boolean, hasPrevPage: boolean } } };


export const SupportedChainsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"supportedChains"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"supportedChains"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"chainID"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"logoUrl"}}]}}]}}]} as unknown as DocumentNode<SupportedChainsQuery, SupportedChainsQueryVariables>;
export const ChainStatsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"chainStats"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"stats"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"totalMsgs"}},{"kind":"Field","name":{"kind":"Name","value":"topStreams"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"sourceChain"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"chainID"}},{"kind":"Field","name":{"kind":"Name","value":"displayID"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"logoUrl"}}]}},{"kind":"Field","name":{"kind":"Name","value":"destChain"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"chainID"}},{"kind":"Field","name":{"kind":"Name","value":"displayID"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"logoUrl"}}]}},{"kind":"Field","name":{"kind":"Name","value":"msgCount"}}]}}]}}]}}]} as unknown as DocumentNode<ChainStatsQuery, ChainStatsQueryVariables>;
export const XmsgDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"xmsg"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"sourceChainID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"destChainID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"offset"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xmsg"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"sourceChainID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"sourceChainID"}}},{"kind":"Argument","name":{"kind":"Name","value":"destChainID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"destChainID"}}},{"kind":"Argument","name":{"kind":"Name","value":"offset"},"value":{"kind":"Variable","name":{"kind":"Name","value":"offset"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"displayID"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}},{"kind":"Field","name":{"kind":"Name","value":"sender"}},{"kind":"Field","name":{"kind":"Name","value":"senderUrl"}},{"kind":"Field","name":{"kind":"Name","value":"to"}},{"kind":"Field","name":{"kind":"Name","value":"toUrl"}},{"kind":"Field","name":{"kind":"Name","value":"gasLimit"}},{"kind":"Field","name":{"kind":"Name","value":"sourceChain"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chainID"}},{"kind":"Field","name":{"kind":"Name","value":"logoUrl"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"destChain"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chainID"}},{"kind":"Field","name":{"kind":"Name","value":"logoUrl"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"txHash"}},{"kind":"Field","name":{"kind":"Name","value":"txUrl"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"block"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"height"}},{"kind":"Field","name":{"kind":"Name","value":"hash"}},{"kind":"Field","name":{"kind":"Name","value":"timestamp"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receipt"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"revertReason"}},{"kind":"Field","name":{"kind":"Name","value":"txHash"}},{"kind":"Field","name":{"kind":"Name","value":"txUrl"}},{"kind":"Field","name":{"kind":"Name","value":"relayer"}},{"kind":"Field","name":{"kind":"Name","value":"timestamp"}},{"kind":"Field","name":{"kind":"Name","value":"gasUsed"}}]}}]}}]}}]} as unknown as DocumentNode<XmsgQuery, XmsgQueryVariables>;
export const XmsgsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"xmsgs"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"first"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"last"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"after"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"before"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"filters"}},"type":{"kind":"ListType","type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"FilterInput"}}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xmsgs"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"first"},"value":{"kind":"Variable","name":{"kind":"Name","value":"first"}}},{"kind":"Argument","name":{"kind":"Name","value":"last"},"value":{"kind":"Variable","name":{"kind":"Name","value":"last"}}},{"kind":"Argument","name":{"kind":"Name","value":"after"},"value":{"kind":"Variable","name":{"kind":"Name","value":"after"}}},{"kind":"Argument","name":{"kind":"Name","value":"before"},"value":{"kind":"Variable","name":{"kind":"Name","value":"before"}}},{"kind":"Argument","name":{"kind":"Name","value":"filters"},"value":{"kind":"Variable","name":{"kind":"Name","value":"filters"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"totalCount"}},{"kind":"Field","name":{"kind":"Name","value":"edges"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"cursor"}},{"kind":"Field","name":{"kind":"Name","value":"node"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"txHash"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}},{"kind":"Field","name":{"kind":"Name","value":"displayID"}},{"kind":"Field","name":{"kind":"Name","value":"sender"}},{"kind":"Field","name":{"kind":"Name","value":"senderUrl"}},{"kind":"Field","name":{"kind":"Name","value":"to"}},{"kind":"Field","name":{"kind":"Name","value":"toUrl"}},{"kind":"Field","name":{"kind":"Name","value":"sourceChain"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chainID"}},{"kind":"Field","name":{"kind":"Name","value":"logoUrl"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"destChain"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chainID"}},{"kind":"Field","name":{"kind":"Name","value":"logoUrl"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"gasLimit"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"txHash"}},{"kind":"Field","name":{"kind":"Name","value":"txUrl"}},{"kind":"Field","name":{"kind":"Name","value":"block"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"chain"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chainID"}}]}},{"kind":"Field","name":{"kind":"Name","value":"hash"}},{"kind":"Field","name":{"kind":"Name","value":"height"}},{"kind":"Field","name":{"kind":"Name","value":"timestamp"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receipt"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"txHash"}},{"kind":"Field","name":{"kind":"Name","value":"txUrl"}},{"kind":"Field","name":{"kind":"Name","value":"timestamp"}},{"kind":"Field","name":{"kind":"Name","value":"success"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}},{"kind":"Field","name":{"kind":"Name","value":"sourceChain"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chainID"}}]}},{"kind":"Field","name":{"kind":"Name","value":"destChain"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"chainID"}}]}},{"kind":"Field","name":{"kind":"Name","value":"relayer"}},{"kind":"Field","name":{"kind":"Name","value":"revertReason"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"pageInfo"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currentPage"}},{"kind":"Field","name":{"kind":"Name","value":"totalPages"}},{"kind":"Field","name":{"kind":"Name","value":"hasNextPage"}},{"kind":"Field","name":{"kind":"Name","value":"hasPrevPage"}}]}}]}}]}}]} as unknown as DocumentNode<XmsgsQuery, XmsgsQueryVariables>;