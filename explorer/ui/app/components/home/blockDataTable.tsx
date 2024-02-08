import { json } from '@remix-run/node'
import { useLoaderData } from '@remix-run/react'
import React from 'react'
import { XBlock } from '~/graphql/graphql'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'

export async function loader() {
  return json<XBlock[]>(new Array())
}

export default function XBlockDataTable() {
  let data = useLoaderData<typeof loader>()
  console.log(data)

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
