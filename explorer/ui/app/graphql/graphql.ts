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
  Time: { input: any; output: any; }
};

/** Chain represents a blockchain on the https://chainlist.org. */
export type Chain = {
  __typename?: 'Chain';
  /** Chain ID as per https://chainlist.org */
  ChainID: Scalars['BigInt']['output'];
  /** Chain name */
  Name: Scalars['String']['output'];
};

/** PageInfo represents pagination information */
export type PageInfo = {
  __typename?: 'PageInfo';
  /** Next Page Cursor */
  HasNextPage: Scalars['Boolean']['output'];
  /** Previous Page Cursor */
  HasPrevPage: Scalars['Boolean']['output'];
  /** Start Cursor */
  StartCursor: Scalars['BigInt']['output'];
};

export type Query = {
  __typename?: 'Query';
  search?: Maybe<SearchResult>;
  supportedchains: Array<Maybe<Chain>>;
  xblock?: Maybe<XBlock>;
  xblockcount?: Maybe<Scalars['BigInt']['output']>;
  xblockrange: Array<Maybe<XBlock>>;
  xmsg?: Maybe<XMsg>;
  xmsgcount?: Maybe<Scalars['BigInt']['output']>;
  xmsgrange: Array<Maybe<XMsg>>;
  /** Get XMsgs with pagination, sorted by latest (cursor goes to zero as the last page) */
  xmsgs?: Maybe<XMsgResult>;
  xreceipt?: Maybe<XReceipt>;
  xreceiptcount?: Maybe<Scalars['BigInt']['output']>;
};


export type QuerySearchArgs = {
  query: Scalars['Bytes32']['input'];
};


export type QueryXblockArgs = {
  height: Scalars['BigInt']['input'];
  sourceChainID: Scalars['BigInt']['input'];
};


export type QueryXblockrangeArgs = {
  from: Scalars['BigInt']['input'];
  to: Scalars['BigInt']['input'];
};


export type QueryXmsgArgs = {
  destChainID: Scalars['BigInt']['input'];
  sourceChainID: Scalars['BigInt']['input'];
  streamOffset: Scalars['BigInt']['input'];
};


export type QueryXmsgrangeArgs = {
  from: Scalars['BigInt']['input'];
  to: Scalars['BigInt']['input'];
};


export type QueryXmsgsArgs = {
  cursor?: InputMaybe<Scalars['BigInt']['input']>;
  limit?: InputMaybe<Scalars['BigInt']['input']>;
};


export type QueryXreceiptArgs = {
  destChainID: Scalars['BigInt']['input'];
  sourceChainID: Scalars['BigInt']['input'];
  streamOffset: Scalars['BigInt']['input'];
};

/** Search for cross-chain messages and receipts. */
export type SearchResult = {
  __typename?: 'SearchResult';
  /** Block Height */
  BlockHeight: Scalars['BigInt']['output'];
  /** Source chain ID */
  SourceChainID: Scalars['BigInt']['output'];
  /** Hash */
  TxHash: Scalars['Bytes32']['output'];
  /** Type */
  Type: SearchResultType;
};

export enum SearchResultType {
  Address = 'ADDRESS',
  Block = 'BLOCK',
  Message = 'MESSAGE',
  Receipt = 'RECEIPT'
}

/** XBlock represents a cross-chain block. */
export type XBlock = {
  __typename?: 'XBlock';
  /** Hash of the source chain block */
  BlockHash: Scalars['Bytes32']['output'];
  /** Height of the source chain block */
  BlockHeight: Scalars['BigInt']['output'];
  /** All cross-chain messages sent/emittted in the block */
  Messages: Array<XMsg>;
  /** Receipts of all submitted cross-chain messages applied in the block */
  Receipts: Array<XReceipt>;
  /** Source chain ID as per https://chainlist.org */
  SourceChainID: Scalars['BigInt']['output'];
  /** Timestamp of the source chain block */
  Timestamp: Scalars['Time']['output'];
  /** UUID of our block */
  UUID: Scalars['ID']['output'];
};

/** XMsg is a cross-chain message. */
export type XMsg = {
  __typename?: 'XMsg';
  /** XBlock message was emitted in */
  Block: XBlock;
  /** Hash of the source chain block */
  BlockHash: Scalars['Bytes32']['output'];
  /** Height of the source chain block */
  BlockHeight: Scalars['BigInt']['output'];
  /** Target/To address to 'call' on destination chain */
  DestAddress: Scalars['Address']['output'];
  /** Destination chain ID as per https://chainlist.org/ */
  DestChainID: Scalars['BigInt']['output'];
  /** Gas limit to use for 'call' on destination chain */
  DestGasLimit: Scalars['BigInt']['output'];
  /** ID of the XMsg */
  ID: Scalars['ID']['output'];
  /** Receipts of the message */
  Receipts: Array<XReceipt>;
  /** Source chain ID as per https://chainlist.org/ */
  SourceChainID: Scalars['BigInt']['output'];
  /** Sender on source chain, set to msg.Sender */
  SourceMessageSender: Scalars['Address']['output'];
  /** Monotonically incremented offset of Msg in the Steam */
  StreamOffset: Scalars['BigInt']['output'];
  /** Hash of the source chain transaction that emitted the message */
  TxHash: Scalars['Bytes32']['output'];
};

/** XMessageEdge represents a single XMsg in a paginated list */
export type XMsgEdge = {
  __typename?: 'XMsgEdge';
  /** Cursor */
  Cursor: Scalars['BigInt']['output'];
  /** XMsg */
  Node: XMsg;
};

/** XMsgResult represents a paginated list of XMsgs */
export type XMsgResult = {
  __typename?: 'XMsgResult';
  /** XMsgs */
  Edges: Array<XMsgEdge>;
  /** Page Info */
  PageInfo: PageInfo;
  /** Total number of XMsgs */
  TotalCount: Scalars['BigInt']['output'];
};

/** XReceipt represents a cross-chain receipt. */
export type XReceipt = {
  __typename?: 'XReceipt';
  /** XBlock message was emitted in */
  Block: XBlock;
  /** Destination chain ID as per https://chainlist.org */
  DestChainID: Scalars['BigInt']['output'];
  /** Gas used for the cross-chain message */
  GasUsed: Scalars['BigInt']['output'];
  /** Messages associated wit this receipt */
  Messages: Array<XMsg>;
  /** Address of the relayer */
  RelayerAddress: Scalars['Address']['output'];
  /** Source chain ID as per https://chainlist.org */
  SourceChainID: Scalars['BigInt']['output'];
  /** Monotonically incremented offset of Msg in the Steam */
  StreamOffset: Scalars['BigInt']['output'];
  /** Success of the cross-chain message */
  Success: Scalars['Boolean']['output'];
  /** Timestamp of the receipt */
  Timestamp: Scalars['Time']['output'];
  /** Hash of the source chain transaction that emitted the message */
  TxHash: Scalars['Bytes32']['output'];
  /** UUID of our block */
  UUID: Scalars['ID']['output'];
};

export type XblockQueryVariables = Exact<{
  sourceChainID: Scalars['BigInt']['input'];
  height: Scalars['BigInt']['input'];
}>;


export type XblockQuery = { __typename?: 'Query', xblock?: { __typename?: 'XBlock', SourceChainID: any, BlockHeight: any, BlockHash: any, Timestamp: any, Messages: Array<{ __typename?: 'XMsg', StreamOffset: any, SourceMessageSender: any, DestAddress: any, DestGasLimit: any, SourceChainID: any, DestChainID: any, TxHash: any }>, Receipts: Array<{ __typename?: 'XReceipt', GasUsed: any, Success: boolean, RelayerAddress: any, SourceChainID: any, DestChainID: any, StreamOffset: any, TxHash: any, Timestamp: any }> } | null };

export type XBlockRangeQueryVariables = Exact<{
  from: Scalars['BigInt']['input'];
  to: Scalars['BigInt']['input'];
}>;


export type XBlockRangeQuery = { __typename?: 'Query', xblockrange: Array<{ __typename?: 'XBlock', SourceChainID: any, BlockHash: any, BlockHeight: any, Timestamp: any } | null> };

export type XblockCountQueryVariables = Exact<{ [key: string]: never; }>;


export type XblockCountQuery = { __typename?: 'Query', xblockcount?: any | null };

export type SupportedChainsQueryVariables = Exact<{ [key: string]: never; }>;


export type SupportedChainsQuery = { __typename?: 'Query', supportedchains: Array<{ __typename?: 'Chain', ChainID: any, Name: string } | null> };

export type XMsgQueryVariables = Exact<{
  sourceChainID: Scalars['BigInt']['input'];
  destChainID: Scalars['BigInt']['input'];
  streamOffset: Scalars['BigInt']['input'];
}>;


export type XMsgQuery = { __typename?: 'Query', xmsg?: { __typename?: 'XMsg', StreamOffset: any, SourceMessageSender: any, DestAddress: any, DestGasLimit: any, SourceChainID: any, DestChainID: any, TxHash: any, BlockHeight: any, BlockHash: any, Block: { __typename?: 'XBlock', SourceChainID: any, BlockHeight: any, BlockHash: any, Timestamp: any }, Receipts: Array<{ __typename?: 'XReceipt', GasUsed: any, Success: boolean, RelayerAddress: any, SourceChainID: any, DestChainID: any, StreamOffset: any, TxHash: any, Timestamp: any }> } | null };

export type XMsgRangeQueryVariables = Exact<{
  from: Scalars['BigInt']['input'];
  to: Scalars['BigInt']['input'];
}>;


export type XMsgRangeQuery = { __typename?: 'Query', xmsgrange: Array<{ __typename?: 'XMsg', StreamOffset: any, SourceMessageSender: any, DestAddress: any, DestGasLimit: any, SourceChainID: any, DestChainID: any, TxHash: any, BlockHeight: any, BlockHash: any } | null> };

export type XMsgCountQueryVariables = Exact<{ [key: string]: never; }>;


export type XMsgCountQuery = { __typename?: 'Query', xmsgcount?: any | null };

export type XReceiptQueryVariables = Exact<{
  sourceChainID: Scalars['BigInt']['input'];
  destChainID: Scalars['BigInt']['input'];
  streamOffset: Scalars['BigInt']['input'];
}>;


export type XReceiptQuery = { __typename?: 'Query', xreceipt?: { __typename?: 'XReceipt', GasUsed: any, Success: boolean, RelayerAddress: any, SourceChainID: any, DestChainID: any, StreamOffset: any, TxHash: any, Timestamp: any, Block: { __typename?: 'XBlock', SourceChainID: any, BlockHeight: any, BlockHash: any, Timestamp: any }, Messages: Array<{ __typename?: 'XMsg', StreamOffset: any, SourceMessageSender: any, DestAddress: any, DestGasLimit: any, SourceChainID: any, DestChainID: any, TxHash: any }> } | null };


export const XblockDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Xblock"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"sourceChainID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"height"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xblock"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"sourceChainID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"sourceChainID"}}},{"kind":"Argument","name":{"kind":"Name","value":"height"},"value":{"kind":"Variable","name":{"kind":"Name","value":"height"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHeight"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHash"}},{"kind":"Field","name":{"kind":"Name","value":"Timestamp"}},{"kind":"Field","name":{"kind":"Name","value":"Messages"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"StreamOffset"}},{"kind":"Field","name":{"kind":"Name","value":"SourceMessageSender"}},{"kind":"Field","name":{"kind":"Name","value":"DestAddress"}},{"kind":"Field","name":{"kind":"Name","value":"DestGasLimit"}},{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"DestChainID"}},{"kind":"Field","name":{"kind":"Name","value":"TxHash"}}]}},{"kind":"Field","name":{"kind":"Name","value":"Receipts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"GasUsed"}},{"kind":"Field","name":{"kind":"Name","value":"Success"}},{"kind":"Field","name":{"kind":"Name","value":"RelayerAddress"}},{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"DestChainID"}},{"kind":"Field","name":{"kind":"Name","value":"StreamOffset"}},{"kind":"Field","name":{"kind":"Name","value":"TxHash"}},{"kind":"Field","name":{"kind":"Name","value":"Timestamp"}}]}}]}}]}}]} as unknown as DocumentNode<XblockQuery, XblockQueryVariables>;
export const XBlockRangeDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"XBlockRange"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"from"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"to"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xblockrange"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"from"},"value":{"kind":"Variable","name":{"kind":"Name","value":"from"}}},{"kind":"Argument","name":{"kind":"Name","value":"to"},"value":{"kind":"Variable","name":{"kind":"Name","value":"to"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHash"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHeight"}},{"kind":"Field","name":{"kind":"Name","value":"Timestamp"}}]}}]}}]} as unknown as DocumentNode<XBlockRangeQuery, XBlockRangeQueryVariables>;
export const XblockCountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"XblockCount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xblockcount"}}]}}]} as unknown as DocumentNode<XblockCountQuery, XblockCountQueryVariables>;
export const SupportedChainsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"SupportedChains"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"supportedchains"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"ChainID"}},{"kind":"Field","name":{"kind":"Name","value":"Name"}}]}}]}}]} as unknown as DocumentNode<SupportedChainsQuery, SupportedChainsQueryVariables>;
export const XMsgDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"XMsg"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"sourceChainID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"destChainID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"streamOffset"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xmsg"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"sourceChainID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"sourceChainID"}}},{"kind":"Argument","name":{"kind":"Name","value":"destChainID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"destChainID"}}},{"kind":"Argument","name":{"kind":"Name","value":"streamOffset"},"value":{"kind":"Variable","name":{"kind":"Name","value":"streamOffset"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"StreamOffset"}},{"kind":"Field","name":{"kind":"Name","value":"SourceMessageSender"}},{"kind":"Field","name":{"kind":"Name","value":"DestAddress"}},{"kind":"Field","name":{"kind":"Name","value":"DestGasLimit"}},{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"DestChainID"}},{"kind":"Field","name":{"kind":"Name","value":"TxHash"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHeight"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHash"}},{"kind":"Field","name":{"kind":"Name","value":"Block"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHeight"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHash"}},{"kind":"Field","name":{"kind":"Name","value":"Timestamp"}}]}},{"kind":"Field","name":{"kind":"Name","value":"Receipts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"GasUsed"}},{"kind":"Field","name":{"kind":"Name","value":"Success"}},{"kind":"Field","name":{"kind":"Name","value":"RelayerAddress"}},{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"DestChainID"}},{"kind":"Field","name":{"kind":"Name","value":"StreamOffset"}},{"kind":"Field","name":{"kind":"Name","value":"TxHash"}},{"kind":"Field","name":{"kind":"Name","value":"Timestamp"}}]}}]}}]}}]} as unknown as DocumentNode<XMsgQuery, XMsgQueryVariables>;
export const XMsgRangeDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"XMsgRange"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"from"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"to"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xmsgrange"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"from"},"value":{"kind":"Variable","name":{"kind":"Name","value":"from"}}},{"kind":"Argument","name":{"kind":"Name","value":"to"},"value":{"kind":"Variable","name":{"kind":"Name","value":"to"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"StreamOffset"}},{"kind":"Field","name":{"kind":"Name","value":"SourceMessageSender"}},{"kind":"Field","name":{"kind":"Name","value":"DestAddress"}},{"kind":"Field","name":{"kind":"Name","value":"DestGasLimit"}},{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"DestChainID"}},{"kind":"Field","name":{"kind":"Name","value":"TxHash"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHeight"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHash"}}]}}]}}]} as unknown as DocumentNode<XMsgRangeQuery, XMsgRangeQueryVariables>;
export const XMsgCountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"XMsgCount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xmsgcount"}}]}}]} as unknown as DocumentNode<XMsgCountQuery, XMsgCountQueryVariables>;
export const XReceiptDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"XReceipt"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"sourceChainID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"destChainID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"streamOffset"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xreceipt"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"sourceChainID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"sourceChainID"}}},{"kind":"Argument","name":{"kind":"Name","value":"destChainID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"destChainID"}}},{"kind":"Argument","name":{"kind":"Name","value":"streamOffset"},"value":{"kind":"Variable","name":{"kind":"Name","value":"streamOffset"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"GasUsed"}},{"kind":"Field","name":{"kind":"Name","value":"Success"}},{"kind":"Field","name":{"kind":"Name","value":"RelayerAddress"}},{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"DestChainID"}},{"kind":"Field","name":{"kind":"Name","value":"StreamOffset"}},{"kind":"Field","name":{"kind":"Name","value":"TxHash"}},{"kind":"Field","name":{"kind":"Name","value":"Timestamp"}},{"kind":"Field","name":{"kind":"Name","value":"Block"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHeight"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHash"}},{"kind":"Field","name":{"kind":"Name","value":"Timestamp"}}]}},{"kind":"Field","name":{"kind":"Name","value":"Messages"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"StreamOffset"}},{"kind":"Field","name":{"kind":"Name","value":"SourceMessageSender"}},{"kind":"Field","name":{"kind":"Name","value":"DestAddress"}},{"kind":"Field","name":{"kind":"Name","value":"DestGasLimit"}},{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"DestChainID"}},{"kind":"Field","name":{"kind":"Name","value":"TxHash"}}]}}]}}]}}]} as unknown as DocumentNode<XReceiptQuery, XReceiptQueryVariables>;