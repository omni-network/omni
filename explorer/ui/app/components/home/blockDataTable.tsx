import { json } from '@remix-run/node'
import { useLoaderData } from '@remix-run/react'
import React from 'react'
import { XBlock } from '~/graphql/graphql'
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
  const [{ data }] = useQuery({
    query: xblockrange,
    variables: {
      amount: 100,
      offset: 0,
    }
  })

  // const { data, fetching, error } = result;
  // console.log(data)
  // console.log(fetching)
  // return result

  //return json(data)
  return json([
    { id: "1", name: "Pants" },
    { id: "2", name: "Jacket" },
  ]);
}

export default function XBlockDataTable() {
  const d = useLoaderData<typeof loader>();
  console.log('xblock data table')
  console.log(d)

  const [result] = useQuery({
    query: xblockrange,
    variables: {
      amount: 100,
      offset: 0,
    }
  })
  const { data, fetching, error } = result;
  console.log(data)
  console.log(fetching)
  console.log(error)

  let rows = [
    {
      id: 1,
      UUID: '0x1234',
      BlockHash: '0x1234',
      BlockHeight: '1',
      Messages: [],
      Timestamp: '2021-10-01T00:00:00Z',
    },
    {
      id: 2,
      UUID: '0x5678',
      BlockHash: '0x5678',
      BlockHeight: '2',
      Messages: [],
      Timestamp: '2021-10-02T00:00:00Z',
    },
    {
      id: 3,
      UUID: '0x5678',
      BlockHash: '0x5678',
      BlockHeight: '3',
      Messages: [],
      Timestamp: '2021-10-02T00:00:00Z',
    },
    {
      id: 4,
      UUID: '0x5678',
      BlockHash: '0x5678',
      BlockHeight: '4',
      Messages: [],
      Timestamp: '2021-10-02T00:00:00Z',
    },
    {
      id: 5,
      UUID: '0x5678',
      BlockHash: '0x5678',
      BlockHeight: '5',
      Messages: [],
      Timestamp: '2021-10-02T00:00:00Z',
    },
    {
      id: 6,
      UUID: '0x5678',
      BlockHash: '0x5678',
      BlockHeight: '6',
      Messages: [],
      Timestamp: '2021-10-02T00:00:00Z',
    },
  ]
  rows = [...rows, ...rows, ...rows, ...rows, ...rows, ...rows]

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
