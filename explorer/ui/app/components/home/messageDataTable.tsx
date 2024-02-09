import { json } from '@remix-run/node'
import { useLoaderData } from '@remix-run/react'
import React from 'react'
import { XMsg } from '~/graphql/graphql'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'
import { useQuery } from 'urql'
import { xblockrange } from '../queries/block'

export async function loader() {
  return json<XMsg[]>(new Array())
}

export default function XMsgDataTable() {
  // let data = useLoaderData<typeof loader>()
  // console.log(data)

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

  var rows: XMsg[] = []
  data?.xblockrange.map((xblock: any) => {

    if (xblock.Messages.length == 0) {
      return
    }

    xblock.Messages.map((msg: any) => {
      let xmsg = {
        DestAddress: msg.DestAddress,
        DestChainID: msg.DestChainID,
        DestGasLimit: "",
        SourceChainID: msg.SourceChainID,
        SourceMessageSender: "",
        StreamOffset: "",
        TxHash: msg.TxHash,
      }
      rows.push(xmsg)
    })

  })

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
