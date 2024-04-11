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

export type Query = {
  __typename?: 'Query';
  xblock?: Maybe<XBlock>;
  xblockcount?: Maybe<Scalars['BigInt']['output']>;
  xblockrange: Array<Maybe<XBlock>>;
  xmsg?: Maybe<XMsg>;
  xmsgcount?: Maybe<Scalars['BigInt']['output']>;
  xmsgrange: Array<Maybe<XMsg>>;
  xreceiptcount?: Maybe<Scalars['BigInt']['output']>;
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

/** XReceipt represents a cross-chain receipt. */
export type XReceipt = {
  __typename?: 'XReceipt';
  /** Destination chain ID as per https://chainlist.org */
  DestChainID: Scalars['BigInt']['output'];
  /** Gas used for the cross-chain message */
  GasUsed: Scalars['BigInt']['output'];
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


export type XblockQuery = { __typename?: 'Query', xblock?: { __typename?: 'XBlock', BlockHash: any } | null };

export type XBlockRangeQueryVariables = Exact<{
  from: Scalars['BigInt']['input'];
  to: Scalars['BigInt']['input'];
}>;


export type XBlockRangeQuery = { __typename?: 'Query', xblockrange: Array<{ __typename?: 'XBlock', SourceChainID: any, BlockHash: any, BlockHeight: any, Timestamp: any, Messages: Array<{ __typename?: 'XMsg', TxHash: any, DestAddress: any, DestChainID: any, SourceChainID: any }> } | null> };

export type XblockCountQueryVariables = Exact<{ [key: string]: never; }>;


export type XblockCountQuery = { __typename?: 'Query', xblockcount?: any | null };


export const XblockDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Xblock"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"sourceChainID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"height"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xblock"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"sourceChainID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"sourceChainID"}}},{"kind":"Argument","name":{"kind":"Name","value":"height"},"value":{"kind":"Variable","name":{"kind":"Name","value":"height"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"BlockHash"}}]}}]}}]} as unknown as DocumentNode<XblockQuery, XblockQueryVariables>;
export const XBlockRangeDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"XBlockRange"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"from"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"to"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BigInt"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xblockrange"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"from"},"value":{"kind":"Variable","name":{"kind":"Name","value":"from"}}},{"kind":"Argument","name":{"kind":"Name","value":"to"},"value":{"kind":"Variable","name":{"kind":"Name","value":"to"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHash"}},{"kind":"Field","name":{"kind":"Name","value":"BlockHeight"}},{"kind":"Field","name":{"kind":"Name","value":"Messages"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"TxHash"}},{"kind":"Field","name":{"kind":"Name","value":"DestAddress"}},{"kind":"Field","name":{"kind":"Name","value":"DestChainID"}},{"kind":"Field","name":{"kind":"Name","value":"SourceChainID"}}]}},{"kind":"Field","name":{"kind":"Name","value":"Timestamp"}}]}}]}}]} as unknown as DocumentNode<XBlockRangeQuery, XBlockRangeQueryVariables>;
export const XblockCountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"XblockCount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"xblockcount"}}]}}]} as unknown as DocumentNode<XblockCountQuery, XblockCountQueryVariables>;
