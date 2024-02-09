import { json } from '@remix-run/node'
import { useLoaderData } from '@remix-run/react'
import React from 'react'
import { XBlock, XMsg } from '~/graphql/graphql'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'
import { graphql } from '~/graphql'
import { useQuery } from 'urql'
import { GetBlockCount, GetBlocksInRange, xblockrange } from '../queries/block'

export const loader = async () => {
  console.log('loader')
  let amt = (500).toString(16)
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

  let rows = GetBlocksInRange(1000, 0)
  let count = GetBlockCount()

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
        accessorFn: row => row.Timestamp, // TODO: format this
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
      <div className="m-auto prose text-m">
        Count: {count}
      </div>
    </div>
  )
}
