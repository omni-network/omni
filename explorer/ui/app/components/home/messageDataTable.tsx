import { json } from '@remix-run/node'
import { useLoaderData } from '@remix-run/react'
import { XBlock } from '~/graphql/graphql'

export async function loader() {
  return json<XBlock[]>(new Array())
}

export default function XMsgDataTable() {
  let data = useLoaderData<typeof loader>()
  console.log(data)

  let rows = [
    {
      id: 1,
      tx_hash: '0x1234',
      source_chain: 'eth',
      dest_chain: 'bsc',
      updated_at: '2021-10-01T00:00:00Z',
    },
    {
      id: 2,
      tx_hash: '0x5678',
      source_chain: 'bsc',
      dest_chain: 'eth',
      updated_at: '2021-10-02T00:00:00Z',
    },
  ]

  // const columns: GridColDef[] = [
  //   {
  //     field: "tx_hash",
  //     headerName: "Tx Hash",
  //     headerClassName: "text-inherit",
  //     type: "string",
  //     minWidth: 150,
  //     flex: 1,
  //   },
  //   {
  //     field: "source_chain",
  //     headerName: "Source Chain",
  //     headerClassName: "text-inherit",
  //     type: "string",
  //     minWidth: 150,
  //     flex: 1,
  //   },
  //   {
  //     field: "dest_chain",
  //     headerName: "Dest Chain",
  //     headerClassName: "text-inherit",
  //     type: "string",
  //     minWidth: 150,
  //     flex: 1,
  //   },
  //   {
  //     field: "updated_at",
  //     headerName: "Updated At",
  //     headerClassName: "text-inherit",
  //     type: "string",
  //     width: 150,
  //     flex: 1,
  //   },
  // ];

  return (
    <div className="m-3">
      <h2 className="prose text-3xl antialiased leading-tight tracking-normal text-inherit">
        XMsgs
      </h2>
      {/* <section
        id="DataGrid"
        style={{
          height: "100%",
          width: "100%",
        }}
        className=" text-inherit"
      >
        <DataGrid
          rows={rows}
          columns={columns}
          {...data}
          initialState={{
            pagination: { paginationModel: { pageSize: 5 } },
          }}
          pageSizeOptions={[5, 10, 25]} />
      </section> */}
    </div>
  )
}
