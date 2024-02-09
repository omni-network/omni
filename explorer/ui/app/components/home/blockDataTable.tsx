import { json } from '@remix-run/node'
import { useLoaderData } from '@remix-run/react'
import React from 'react'
import { XBlock, XMsg } from '~/graphql/graphql'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'
import { graphql } from '~/graphql'
import { Client, gql, useQuery } from 'urql'

const xblock = graphql(`
  query Xblock($sourceChainID: BigInt!, $height: BigInt!) {
      xblock(sourceChainID: $sourceChainID, height: $height) {
        BlockHash
      }
    }
`);

const xblockrange = graphql(`
  query XBlockRange($amount: BigInt!, $offset: BigInt!) {
    xblockrange(amount: $amount, offset: $offset) {
      SourceChainID
      BlockHash
      BlockHeight
      Messages {
        DestAddress
      }
      Timestamp
    }
  }
`);

const xblockcount = graphql(`
  query XblockCount {
    xblockcount
  }
`);
export const loader = async () => {
  console.log('loader')
  let amt = (1000).toString(16)
  let offset = (0).toString(16)
  const [{ data }] = useQuery({
    query: xblockrange,
    variables: {
      amount: amt,
      offset: offset,
    },
  })

  return json(data)
}

export default function XBlockDataTable() {
  const d = useLoaderData<typeof loader>();
  console.log(d)

  let amt = "0x" + (1000).toString(16)
  let offset = "0x" + (0).toString(16)

  const [result] = useQuery({
    query: xblockrange,
    variables: {
      amount: amt,
      offset: offset,
    }
  })
  const { data, fetching, error } = result;

  var rows: XBlock[] = []
  data?.xblockrange.map((xblock: any) => {
    var msgs: XMsg[] = []
    let block = {
      id: xblock.BlockHeight,
      UUID: "",
      SourceChainID: xblock.SourceChainID,
      BlockHash: xblock.BlockHash,
      BlockHeight: xblock.BlockHeight,
      Messages: msgs,
      Timestamp: xblock.Timestamp,
      Receipts: [],
    }

    xblock.Messages.map((msg: any) => {
      let xmsg = {
        DestAddress: "",
        DestChainID: "",
        DestGasLimit: "",
        SourceChainID: "",
        SourceMessageSender: "",
        StreamOffset: "",
        TxHash: "",
      }
      msgs.push(xmsg)
    })

    block.Messages = msgs
    rows.push(block)
  })

  const columns = React.useMemo<ColumnDef<XBlock>[]>(
    () => [
      {
        accessorKey: 'blockNumber',
        accessorFn: row => row.BlockHeight,
        header: () => <span>Block Number</span>,
        canFilter: false,
        enableColumnFilter: false,
      },
      {
        accessorKey: 'number_of_transactions',
        accessorFn: row => row.Messages.length,
        header: () => <span># of Txs</span>,
        canFilter: false,
        enableColumnFilter: false,
      },
      {
        accessorKey: 'time',
        accessorFn: row => row.Timestamp,
        header: () => <span>Time</span>,
        canFilter: false,
        enableColumnFilter: false,
      },
    ],
    [],
  )

  return (
    <div className="m-3">
      <div className="">
        <h1 className="prose text-xl font-semibold mb-3">XBlocks</h1>
      </div>
      <div>
        <SimpleTable columns={columns} data={rows} />
      </div>
    </div>
  )
}
