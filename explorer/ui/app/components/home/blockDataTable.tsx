import { json } from '@remix-run/node'
import { useLoaderData } from '@remix-run/react'
import React from 'react'
import { XBlock } from '~/graphql/graphql'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'
import { GetBlockCount, GetBlocksInRange } from '../queries/block'
import { dateFormatter } from '~/lib/formatting'

export const loader = async () => {
  return json<XBlock[]>([])
}

export default function XBlockDataTable() {
  useLoaderData<typeof loader>()

  const rows = GetBlocksInRange(0, 1000)
  const count = GetBlockCount()

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
        accessorFn: row => dateFormatter(row.Timestamp),
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
        <h1 className="prose text-primary font-semibold mb-3">XBlocks</h1>
      </div>
      <div>
        <SimpleTable columns={columns} data={rows} />
      </div>
      <div className="m-auto prose text-m">Count: {count}</div>
    </div>
  )
}
