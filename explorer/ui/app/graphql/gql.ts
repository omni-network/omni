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
    "\n  query supportedChains {\n    supportedChains {\n    id\n    chainID\n    name\n    logoUrl\n    }\n  }\n": types.SupportedChainsDocument,
    "\nquery chainStats {\n  stats {\n    totalMsgs\n    topStreams {\n      sourceChain{\n        id\n        chainID\n        displayID\n        name\n        logoUrl\n      }\n      destChain {\n        id\n        chainID\n        displayID\n        name\n        logoUrl\n      }\n      msgCount\n    }\n  }\n}\n": types.ChainStatsDocument,
    "\n  query xmsg($sourceChainID: BigInt!, $destChainID: BigInt!, $offset: BigInt!) {\n    xmsg(sourceChainID: $sourceChainID, destChainID: $destChainID, offset: $offset) {\n      id\n      displayID\n      offset\n      sender\n      senderUrl\n      to\n      toUrl\n      gasLimit\n      sourceChain{\n        chainID\n        logoUrl\n        name\n      }\n      destChain{\n        chainID\n        logoUrl\n        name\n      }\n      txHash\n      txUrl\n      status\n      block {\n        height\n        hash\n        timestamp\n      }\n      receipt {\n        revertReason\n        txHash\n        txUrl\n        relayer\n        timestamp\n        gasUsed\n      }\n    }\n  }\n": types.XmsgDocument,
    "\nquery xmsgs($first: Int, $last: Int, $after: ID, $before: ID, $filters: [FilterInput!]) {\n  xmsgs(first: $first, last: $last, after: $after, before: $before, filters: $filters) {\n    totalCount\n    edges {\n      cursor\n      node {\n        id\n        txHash\n        offset\n        displayID\n        sender\n        senderUrl\n        to\n        toUrl\n        sourceChain{\n          chainID\n          logoUrl\n          name\n        }\n        destChain{\n          chainID\n          logoUrl\n          name\n        }\n        gasLimit\n        status\n        txHash\n        txUrl\n        block {\n          id\n          chain {\n            chainID\n          }\n          hash\n          height\n          timestamp\n        }\n        receipt {\n          txHash\n          txUrl\n          timestamp\n          success\n          offset\n          sourceChain{\n            chainID\n          }\n          destChain{\n            chainID\n          }\n          relayer\n          revertReason\n        }\n      }\n    }\n    pageInfo {\n      currentPage\n      totalPages\n      hasNextPage\n      hasPrevPage\n    }\n  }\n}\n": types.XmsgsDocument,
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
export function graphql(source: "\n  query supportedChains {\n    supportedChains {\n    id\n    chainID\n    name\n    logoUrl\n    }\n  }\n"): (typeof documents)["\n  query supportedChains {\n    supportedChains {\n    id\n    chainID\n    name\n    logoUrl\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\nquery chainStats {\n  stats {\n    totalMsgs\n    topStreams {\n      sourceChain{\n        id\n        chainID\n        displayID\n        name\n        logoUrl\n      }\n      destChain {\n        id\n        chainID\n        displayID\n        name\n        logoUrl\n      }\n      msgCount\n    }\n  }\n}\n"): (typeof documents)["\nquery chainStats {\n  stats {\n    totalMsgs\n    topStreams {\n      sourceChain{\n        id\n        chainID\n        displayID\n        name\n        logoUrl\n      }\n      destChain {\n        id\n        chainID\n        displayID\n        name\n        logoUrl\n      }\n      msgCount\n    }\n  }\n}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query xmsg($sourceChainID: BigInt!, $destChainID: BigInt!, $offset: BigInt!) {\n    xmsg(sourceChainID: $sourceChainID, destChainID: $destChainID, offset: $offset) {\n      id\n      displayID\n      offset\n      sender\n      senderUrl\n      to\n      toUrl\n      gasLimit\n      sourceChain{\n        chainID\n        logoUrl\n        name\n      }\n      destChain{\n        chainID\n        logoUrl\n        name\n      }\n      txHash\n      txUrl\n      status\n      block {\n        height\n        hash\n        timestamp\n      }\n      receipt {\n        revertReason\n        txHash\n        txUrl\n        relayer\n        timestamp\n        gasUsed\n      }\n    }\n  }\n"): (typeof documents)["\n  query xmsg($sourceChainID: BigInt!, $destChainID: BigInt!, $offset: BigInt!) {\n    xmsg(sourceChainID: $sourceChainID, destChainID: $destChainID, offset: $offset) {\n      id\n      displayID\n      offset\n      sender\n      senderUrl\n      to\n      toUrl\n      gasLimit\n      sourceChain{\n        chainID\n        logoUrl\n        name\n      }\n      destChain{\n        chainID\n        logoUrl\n        name\n      }\n      txHash\n      txUrl\n      status\n      block {\n        height\n        hash\n        timestamp\n      }\n      receipt {\n        revertReason\n        txHash\n        txUrl\n        relayer\n        timestamp\n        gasUsed\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\nquery xmsgs($first: Int, $last: Int, $after: ID, $before: ID, $filters: [FilterInput!]) {\n  xmsgs(first: $first, last: $last, after: $after, before: $before, filters: $filters) {\n    totalCount\n    edges {\n      cursor\n      node {\n        id\n        txHash\n        offset\n        displayID\n        sender\n        senderUrl\n        to\n        toUrl\n        sourceChain{\n          chainID\n          logoUrl\n          name\n        }\n        destChain{\n          chainID\n          logoUrl\n          name\n        }\n        gasLimit\n        status\n        txHash\n        txUrl\n        block {\n          id\n          chain {\n            chainID\n          }\n          hash\n          height\n          timestamp\n        }\n        receipt {\n          txHash\n          txUrl\n          timestamp\n          success\n          offset\n          sourceChain{\n            chainID\n          }\n          destChain{\n            chainID\n          }\n          relayer\n          revertReason\n        }\n      }\n    }\n    pageInfo {\n      currentPage\n      totalPages\n      hasNextPage\n      hasPrevPage\n    }\n  }\n}\n"): (typeof documents)["\nquery xmsgs($first: Int, $last: Int, $after: ID, $before: ID, $filters: [FilterInput!]) {\n  xmsgs(first: $first, last: $last, after: $after, before: $before, filters: $filters) {\n    totalCount\n    edges {\n      cursor\n      node {\n        id\n        txHash\n        offset\n        displayID\n        sender\n        senderUrl\n        to\n        toUrl\n        sourceChain{\n          chainID\n          logoUrl\n          name\n        }\n        destChain{\n          chainID\n          logoUrl\n          name\n        }\n        gasLimit\n        status\n        txHash\n        txUrl\n        block {\n          id\n          chain {\n            chainID\n          }\n          hash\n          height\n          timestamp\n        }\n        receipt {\n          txHash\n          txUrl\n          timestamp\n          success\n          offset\n          sourceChain{\n            chainID\n          }\n          destChain{\n            chainID\n          }\n          relayer\n          revertReason\n        }\n      }\n    }\n    pageInfo {\n      currentPage\n      totalPages\n      hasNextPage\n      hasPrevPage\n    }\n  }\n}\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;