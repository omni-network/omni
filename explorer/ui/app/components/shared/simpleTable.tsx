import {
  Column,
  Table as ReactTable,
  PaginationState,
  useReactTable,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  ColumnDef,
  OnChangeFn,
  flexRender,
} from '@tanstack/react-table'
import {
  ChevronDoubleLeftIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
  ChevronDoubleRightIcon,
} from '@heroicons/react/24/solid'
import { Button, PageButton } from './button'
import { Link } from '@remix-run/react'

export default function SimpleTable({
  data,
  columns,
  headChildren,
}: {
  data: any[]
  columns: ColumnDef<any>[]
  headChildren?: Array<React.ReactNode> | React.ReactNode
}) {
  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
  })

  return (
    <div className="overflow-x-auto ">
      <div className="w-full bg-raised rounded-lg min-w-[919px]">
        {headChildren}

        <table className="min-w-full">
          <thead className="">
            {table.getHeaderGroups().map(headerGroup => (
              <tr key={headerGroup.id}>
                {headerGroup.headers.map((header, headerIndex) => {
                  return (
                    <th
                      className={`pt-[15px] pb-[7px] px-4 py-2 text-start ${headerIndex < 3 && 'table-highlight'}`}
                      key={header.id}
                      colSpan={header.colSpan}
                    >
                      {header.isPlaceholder ? null : (
                        <div>
                          <span className="text-b-sm text-subtlest font-bold text-start">
                            {flexRender(header.column.columnDef.header, header.getContext())}
                          </span>
                          {header.column.getCanFilter() ? (
                            <div>
                              <Filter column={header.column} table={table} />
                            </div>
                          ) : null}
                        </div>
                      )}
                    </th>
                  )
                })}
              </tr>
            ))}
          </thead>
          <tbody className="">
            {table.getRowModel().rows.map(row => {
              return (
                <tr key={row.id} className="border-border-subtle border-t-[1px]">
                  {row.getVisibleCells().map((cell, cellIndex) => {
                    return (
                      <td
                        key={cell.id}
                        className={`px-4 py-2 whitespace-nowrap ${cellIndex < 3 && 'table-highlight'} ${cellIndex === 0 && 'text-center'}`}
                        role="cell"
                      >
                        {flexRender(cell.column.columnDef.cell, cell.getContext())}
                      </td>
                    )
                  })}
                </tr>
              )
            })}
          </tbody>
        </table>
      </div>
      <div className="m-auto">
        {/* Pagination */}
        <div className="flex items-center mt-4">
          {/* Page Size Dropdown */}
          <div className="flex-none flex items-center ">
            <label className="relative">
              <select
                className="appearance-none cursor-pointer bg-bg-input-default rounded-full text-cb-md px-4 py-3 pr-8 text-default "
                value={table.getState().pagination.pageSize}
                onChange={e => {
                  table.setPageSize(Number(e.target.value))
                }}
              >
                {[5, 10, 20].map(pageSize => (
                  <option key={pageSize} value={pageSize}>
                    Show {pageSize}
                  </option>
                ))}
              </select>

              <span
                className={
                  'pointer-events-none icon-chevron-lrg-down absolute right-2 top-[9px] text-default'
                }
              ></span>
            </label>
          </div>

          {/* middle element */}
          <div className="grow"></div>

          {/* Nav Buttons */}
          <div className="flex flex-row items-center jusify-between">
            <PageButton
              className="rounded-full  flex items-center justify-center"
              onClick={() => table.setPageIndex(0)}
              disabled={!table.getCanPreviousPage()}
            >
              <span className="sr-only">First</span>
              <span className={`icon-rewind text-[20px]`}></span>
            </PageButton>
            <PageButton
              className="rounded-full flex items-center justify-center"
              onClick={() => table.previousPage()}
              disabled={!table.getCanPreviousPage()}
            >
              <span className="sr-only">Previous</span>
              <span className={`icon-chevron-med-left text-[20px]`}></span>
            </PageButton>

            {/* Page N of N */}
            <div className="flex-none flex m-3">
              <div className="flex gap-x-2 items-baseline">
                <span className="text-cb-sm text-default">
                  Page <span className="">{table.getState().pagination.pageIndex + 1}</span> of{' '}
                  <span className="">{table.getPageCount() == 0 ? 1 : table.getPageCount()}</span>
                </span>
              </div>
            </div>

            <PageButton
              className="rounded-full  flex items-center justify-center"
              onClick={() => table.nextPage()}
              disabled={!table.getCanNextPage()}
            >
              <span className="sr-only">Next</span>
              <span className={`icon-chevron-med-right text-[20px]`}></span>
            </PageButton>
            <PageButton
              className="rounded-full  flex items-center justify-center"
              onClick={() => table.setPageIndex(table.getPageCount() - 1)}
              disabled={!table.getCanNextPage()}
            >
              <span className="sr-only">Last</span>
              <span className={`icon-fast-forward text-[20px]`}></span>
            </PageButton>
          </div>
        </div>
      </div>
    </div>
  )
}

function Filter({ column, table }: { column: Column<any, any>; table: ReactTable<any> }) {
  const firstValue = table.getPreFilteredRowModel().flatRows[0]?.getValue(column.id)

  const columnFilterValue = column.getFilterValue()

  return typeof firstValue === 'number' ? (
    <div className="flex space-x-2">
      <input
        type="number"
        value={(columnFilterValue as [number, number])?.[0] ?? ''}
        onChange={e => column.setFilterValue((old: [number, number]) => [e.target.value, old?.[1]])}
        placeholder={`Min`}
        className="w-24 border shadow rounded"
      />
      <input
        type="number"
        value={(columnFilterValue as [number, number])?.[1] ?? ''}
        onChange={e => column.setFilterValue((old: [number, number]) => [old?.[0], e.target.value])}
        placeholder={`Max`}
        className="w-24 border shadow rounded"
      />
    </div>
  ) : (
    <input
      type="text"
      value={(columnFilterValue ?? '') as string}
      onChange={e => column.setFilterValue(e.target.value)}
      placeholder={`Search...`}
      className="w-36 border shadow rounded"
    />
  )
}
