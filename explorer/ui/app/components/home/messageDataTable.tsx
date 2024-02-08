import { json } from '@remix-run/node'
import { useLoaderData } from '@remix-run/react'
import React from 'react'
import { XMsg } from '~/graphql/graphql'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'

export async function loader() {
  return json<XMsg[]>(new Array())
}

export default function XMsgDataTable() {
  let data = useLoaderData<typeof loader>()
  console.log(data)

  let rows = [
    {
      id: 1,
      DestAddress: '0x1234',
      DestChainID: '1',
      DestGasLimit: '100000',
      SourceChainID: '2',
      SourceMessageSender: '0x1234',
      StreamOffset: '0',
      TxHash: '0x1234',
    },
    {
      id: 2,
      DestAddress: '0x5678',
      DestChainID: '2',
      DestGasLimit: '100000',
      SourceChainID: '1',
      SourceMessageSender: '0x1234',
      StreamOffset: '0',
      TxHash: '0x5678',
    },
    {
      id: 3,
      DestAddress: '0x5678',
      DestChainID: '3',
      DestGasLimit: '100000',
      SourceChainID: '1',
      SourceMessageSender: '0x1234',
      StreamOffset: '0',
      TxHash: '0x5678',
    }
  ]

  rows = [...rows, ...rows, ...rows, ...rows, ...rows, ...rows]

  const columns = React.useMemo<ColumnDef<XMsg>[]>(
    () => [
      {
        accessorKey: 'tx_hash',
        accessorFn: (row) => row.TxHash,
        header: () => <span>TxHash</span>,
        canFilter: false,
        enableColumnFilter: false,
      },
      {
        accessorKey: 'source_chain',
        accessorFn: (row) => row.SourceChainID,
        header: () => <span>Source Chain</span>,
        canFilter: false,
        enableColumnFilter: false,
      },
      {
        accessorKey: 'dest_chain',
        accessorFn: (row) => row.DestChainID,
        header: () => <span>Dest Chain</span>,
        canFilter: false,
        enableColumnFilter: false,
      },
      {
        accessorKey: 'time',
        accessorFn: (row) => "",
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
        <h1 className="prose text-xl font-semibold">XMsgs</h1>
      </div>
      <div>
        <SimpleTable columns={columns} data={rows} />
      </div>
    </div>
  )
}
