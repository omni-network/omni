import { json } from '@remix-run/node'
import React from 'react'
import { XMsg } from '~/graphql/graphql'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'
import { useLoaderData } from '@remix-run/react'

export async function loader() {
  return json<XMsg[]>(new Array())
}

export default function XMsgDataTable() {
  const d = useLoaderData<typeof loader>()

  let rows = []

  const columns = React.useMemo<ColumnDef<XMsg>[]>(
    () => [
      {
        accessorKey: 'tx_hash',
        accessorFn: row => row.TxHash,
        header: () => <span>TxHash</span>,
        canFilter: false,
        enableColumnFilter: false,
      },
      {
        accessorKey: 'source_chain',
        accessorFn: row => row.SourceChainID,
        header: () => <span>Source Chain</span>,
        canFilter: false,
        enableColumnFilter: false,
      },
      {
        accessorKey: 'dest_chain',
        accessorFn: row => row.DestChainID,
        header: () => <span>Dest Chain</span>,
        canFilter: false,
        enableColumnFilter: false,
      },
      {
        accessorKey: 'time',
        accessorFn: row => '',
        header: () => <span>Updated At</span>,
        canFilter: false,
        enableColumnFilter: false,
      },
    ],
    [],
  )

  return (
    <div className="m-3">
      <div className="">
        <h1 className="prose text-xl font-semibold mb-3">XMsgs</h1>
      </div>
      <div>
        <SimpleTable columns={columns} data={rows} />
      </div>
    </div>
  )
}
