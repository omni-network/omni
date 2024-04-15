/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  query Xblock($sourceChainID: BigInt!, $height: BigInt!) {\n    xblock(sourceChainID: $sourceChainID, height: $height) {\n      SourceChainID\n      BlockHeight\n      BlockHash\n      Timestamp\n      Messages {\n        StreamOffset\n        SourceMessageSender\n        DestAddress\n        DestGasLimit\n        SourceChainID\n        DestChainID\n        TxHash\n      }\n      Receipts {\n        GasUsed\n        Success\n        RelayerAddress\n        SourceChainID\n        DestChainID\n        StreamOffset\n        TxHash\n        Timestamp\n      }\n    }\n  }\n": types.XblockDocument,
    "\n  query XBlockRange($from: BigInt!, $to: BigInt!) {\n    xblockrange(from: $from, to: $to) {\n      SourceChainID\n      BlockHash\n      BlockHeight\n      Timestamp\n    }\n  }\n": types.XBlockRangeDocument,
    "\n  query XblockCount {\n    xblockcount\n  }\n": types.XblockCountDocument,
    "\n  query SupportedChains {\n    supportedchains {\n      ChainID\n      Name\n    }\n  }\n": types.SupportedChainsDocument,
    "\n  query XMsg($sourceChainID: BigInt!, $destChainID: BigInt!, $streamOffset: BigInt!) {\n    xmsg(sourceChainID: $sourceChainID, destChainID: $destChainID, streamOffset: $streamOffset) {\n      StreamOffset\n      SourceMessageSender\n      DestAddress\n      DestGasLimit\n      SourceChainID\n      DestChainID\n      TxHash\n      BlockHeight\n      BlockHash\n      Block {\n        SourceChainID\n        BlockHeight\n        BlockHash\n        Timestamp\n      }\n      Receipts {\n        GasUsed\n        Success\n        RelayerAddress\n        SourceChainID\n        DestChainID\n        StreamOffset\n        TxHash\n        Timestamp\n      }\n    }\n  }\n": types.XMsgDocument,
    "\n  query XMsgRange($from: BigInt!, $to: BigInt!) {\n    xmsgrange(from: $from, to: $to) {\n      StreamOffset\n      SourceMessageSender\n      DestAddress\n      DestGasLimit\n      SourceChainID\n      DestChainID\n      TxHash\n      BlockHeight\n      BlockHash\n    }\n  }\n": types.XMsgRangeDocument,
    "\n  query XMsgCount {\n    xmsgcount\n  }\n": types.XMsgCountDocument,
    "\n  query XReceipt($sourceChainID: BigInt!, $destChainID: BigInt!, $streamOffset: BigInt!) {\n    xreceipt(\n      sourceChainID: $sourceChainID\n      destChainID: $destChainID\n      streamOffset: $streamOffset\n    ) {\n      GasUsed\n      Success\n      RelayerAddress\n      SourceChainID\n      DestChainID\n      StreamOffset\n      TxHash\n      Timestamp\n      Block {\n        SourceChainID\n        BlockHeight\n        BlockHash\n        Timestamp\n      }\n      Messages {\n        StreamOffset\n        SourceMessageSender\n        DestAddress\n        DestGasLimit\n        SourceChainID\n        DestChainID\n        TxHash\n      }\n    }\n  }\n": types.XReceiptDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query Xblock($sourceChainID: BigInt!, $height: BigInt!) {\n    xblock(sourceChainID: $sourceChainID, height: $height) {\n      SourceChainID\n      BlockHeight\n      BlockHash\n      Timestamp\n      Messages {\n        StreamOffset\n        SourceMessageSender\n        DestAddress\n        DestGasLimit\n        SourceChainID\n        DestChainID\n        TxHash\n      }\n      Receipts {\n        GasUsed\n        Success\n        RelayerAddress\n        SourceChainID\n        DestChainID\n        StreamOffset\n        TxHash\n        Timestamp\n      }\n    }\n  }\n"): (typeof documents)["\n  query Xblock($sourceChainID: BigInt!, $height: BigInt!) {\n    xblock(sourceChainID: $sourceChainID, height: $height) {\n      SourceChainID\n      BlockHeight\n      BlockHash\n      Timestamp\n      Messages {\n        StreamOffset\n        SourceMessageSender\n        DestAddress\n        DestGasLimit\n        SourceChainID\n        DestChainID\n        TxHash\n      }\n      Receipts {\n        GasUsed\n        Success\n        RelayerAddress\n        SourceChainID\n        DestChainID\n        StreamOffset\n        TxHash\n        Timestamp\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query XBlockRange($from: BigInt!, $to: BigInt!) {\n    xblockrange(from: $from, to: $to) {\n      SourceChainID\n      BlockHash\n      BlockHeight\n      Timestamp\n    }\n  }\n"): (typeof documents)["\n  query XBlockRange($from: BigInt!, $to: BigInt!) {\n    xblockrange(from: $from, to: $to) {\n      SourceChainID\n      BlockHash\n      BlockHeight\n      Timestamp\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query XblockCount {\n    xblockcount\n  }\n"): (typeof documents)["\n  query XblockCount {\n    xblockcount\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query SupportedChains {\n    supportedchains {\n      ChainID\n      Name\n    }\n  }\n"): (typeof documents)["\n  query SupportedChains {\n    supportedchains {\n      ChainID\n      Name\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query XMsg($sourceChainID: BigInt!, $destChainID: BigInt!, $streamOffset: BigInt!) {\n    xmsg(sourceChainID: $sourceChainID, destChainID: $destChainID, streamOffset: $streamOffset) {\n      StreamOffset\n      SourceMessageSender\n      DestAddress\n      DestGasLimit\n      SourceChainID\n      DestChainID\n      TxHash\n      BlockHeight\n      BlockHash\n      Block {\n        SourceChainID\n        BlockHeight\n        BlockHash\n        Timestamp\n      }\n      Receipts {\n        GasUsed\n        Success\n        RelayerAddress\n        SourceChainID\n        DestChainID\n        StreamOffset\n        TxHash\n        Timestamp\n      }\n    }\n  }\n"): (typeof documents)["\n  query XMsg($sourceChainID: BigInt!, $destChainID: BigInt!, $streamOffset: BigInt!) {\n    xmsg(sourceChainID: $sourceChainID, destChainID: $destChainID, streamOffset: $streamOffset) {\n      StreamOffset\n      SourceMessageSender\n      DestAddress\n      DestGasLimit\n      SourceChainID\n      DestChainID\n      TxHash\n      BlockHeight\n      BlockHash\n      Block {\n        SourceChainID\n        BlockHeight\n        BlockHash\n        Timestamp\n      }\n      Receipts {\n        GasUsed\n        Success\n        RelayerAddress\n        SourceChainID\n        DestChainID\n        StreamOffset\n        TxHash\n        Timestamp\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query XMsgRange($from: BigInt!, $to: BigInt!) {\n    xmsgrange(from: $from, to: $to) {\n      StreamOffset\n      SourceMessageSender\n      DestAddress\n      DestGasLimit\n      SourceChainID\n      DestChainID\n      TxHash\n      BlockHeight\n      BlockHash\n    }\n  }\n"): (typeof documents)["\n  query XMsgRange($from: BigInt!, $to: BigInt!) {\n    xmsgrange(from: $from, to: $to) {\n      StreamOffset\n      SourceMessageSender\n      DestAddress\n      DestGasLimit\n      SourceChainID\n      DestChainID\n      TxHash\n      BlockHeight\n      BlockHash\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query XMsgCount {\n    xmsgcount\n  }\n"): (typeof documents)["\n  query XMsgCount {\n    xmsgcount\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query XReceipt($sourceChainID: BigInt!, $destChainID: BigInt!, $streamOffset: BigInt!) {\n    xreceipt(\n      sourceChainID: $sourceChainID\n      destChainID: $destChainID\n      streamOffset: $streamOffset\n    ) {\n      GasUsed\n      Success\n      RelayerAddress\n      SourceChainID\n      DestChainID\n      StreamOffset\n      TxHash\n      Timestamp\n      Block {\n        SourceChainID\n        BlockHeight\n        BlockHash\n        Timestamp\n      }\n      Messages {\n        StreamOffset\n        SourceMessageSender\n        DestAddress\n        DestGasLimit\n        SourceChainID\n        DestChainID\n        TxHash\n      }\n    }\n  }\n"): (typeof documents)["\n  query XReceipt($sourceChainID: BigInt!, $destChainID: BigInt!, $streamOffset: BigInt!) {\n    xreceipt(\n      sourceChainID: $sourceChainID\n      destChainID: $destChainID\n      streamOffset: $streamOffset\n    ) {\n      GasUsed\n      Success\n      RelayerAddress\n      SourceChainID\n      DestChainID\n      StreamOffset\n      TxHash\n      Timestamp\n      Block {\n        SourceChainID\n        BlockHeight\n        BlockHash\n        Timestamp\n      }\n      Messages {\n        StreamOffset\n        SourceMessageSender\n        DestAddress\n        DestGasLimit\n        SourceChainID\n        DestChainID\n        TxHash\n      }\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;
