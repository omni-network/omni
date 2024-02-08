import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { XBlock } from "~/graphql/graphql";

export async function loader() {
  return json<XBlock[]>(new Array())
}

export default function XBlockDataTable() {
  let data = useLoaderData<typeof loader>();
  console.log(data);

  let rows = [
    {
      id: 1,
      blockNumber: "1",
      number_of_transactions: "1",
      time: "2021-10-01T00:00:00Z",
    },
    {
      id: 2,
      blockNumber: "2",
      number_of_transactions: "2",
      time: "2021-10-02T00:00:00Z",
    },
  ];

  const columns: GridColDef[] = [
    {
      field: "blockNumber",
      headerName: "Block Number",
      headerClassName: "text-inherit",
      type: "string",
      minWidth: 150,
      flex: 1,
    },
    {
      field: "number_of_transactions",
      headerName: "# of TXs",
      headerClassName: "text-inherit",
      type: "string",
      minWidth: 150,
      flex: 1,
    },
    {
      field: "time",
      headerName: "Time",
      headerClassName: "text-inherit",
      type: "string",
      width: 150,
      flex: 1,
    },
  ];

  return (
    <div className="m-3">
      <h2 className="prose text-3xl antialiased leading-tight tracking-normal text-inherit">
        XBlocks
      </h2>
      <section
        id="DataGrid"
        style={{
          height: "100%",
          width: "100%",
        }}
        className="prose prose-lg"
      >
        <DataGrid
          rows={rows}
          columns={columns}
          {...data}
          initialState={{
            pagination: { paginationModel: { pageSize: 5 } },
          }}
          pageSizeOptions={[5, 10, 25]} />
      </section>
    </div>
  );
}
