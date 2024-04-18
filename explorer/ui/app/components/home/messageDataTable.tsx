import { json } from '@remix-run/node'
import React, { useEffect } from 'react'
import { XMsg } from '~/graphql/graphql'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'
import { useLoaderData } from '@remix-run/react'
import { dateFormatter, hashShortener } from '~/lib/formatting'
import Tag from '../shared/tag'
import RollupIcon from '../shared/rollupIcon'
import { Link } from '@remix-run/react'
import LongArrow from '~/assets/images/LongArrow.svg'
import { loader } from '~/routes/_index'
import SearchBar from '../shared/search'
import Dropdown from '../shared/dropdown'
import ChainDropdown from './chainDropdown'
import FilterOptions from '../shared/filterOptions'
import { getAddressUrl, getBlockUrl, getTxUrl } from '~/lib/sourceChains'

export default function XMsgDataTable() {
  const data = useLoaderData<typeof loader>()

  const filterOptions = [
    { display: 'Source address', value: 'sourceAddress' },
    {
      display: 'Source tx hash',
      value: 'sourceTxHash',
    },
    {
      display: 'Destination address',
      value: 'destAddress',
    },
    {
      display: 'Destination tx hash',
      value: 'destTxHash',
    },
  ]

  const sourceChainList = data.supportedChains.map(chain => ({
    value: chain.ChainID,
    display: chain.DisplayName,
    icon: chain.Icon,
  }))

  const [searchValue, setSearchValue] = React.useState<string>('')
  const [searchPlaceholder, setSearchPlaceholder] = React.useState<string>()

  const rows = data.xmsgs

  const columnConfig = {
    canFilter: false,
    enableColumnFilter: false,
  }

  const columns = React.useMemo<ColumnDef<any>[]>(
    () => [
      {
        ...columnConfig,
        accessorKey: 'StreamOffset',
        header: () => <span>Nounce</span>,
        cell: (value: any) => <span className=" font-bold text-b-sm">{value.getValue()}</span>,
      },
      {
        ...columnConfig,
        accessorKey: 'timeStamp',
        header: () => <span>Age</span>,
        cell: (value: any) => (
          <span className="text-subtlest font-bold text-b-xs">
            {' '}
            {dateFormatter(value.getValue())}
          </span>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'status',
        header: () => <span>Status</span>,
        cell: (value: any) => <Tag status={value.getValue()} />,
      },
      {
        ...columnConfig,
        accessorKey: 'SourceChainID',
        header: () => <span></span>,
        cell: (value: any) => <RollupIcon chainId={value.getValue()} />,
      },
      {
        ...columnConfig,
        accessorKey: 'fromAddress',
        header: () => <span>Address</span>,
        cell: (value: any) => (
          <Link to="/" className="link">
            {value.getValue() && (
              <>
                <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
                <span className="icon-external-link" />
              </>
            )}
          </Link>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'BlockHash',
        header: () => <span>Block Hash</span>,
        cell: (value: any) => (
          <Link
            target="_blank"
            to={getBlockUrl(
              value.row.original.SourceChainID,
              value.row.original.BlockHash,
              data.supportedChains,
            )}
            className="link"
          >
            <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
            <span className="icon-external-link" />
          </Link>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'Empty',
        header: () => <span></span>,
        cell: (value: any) => <img src={LongArrow} alt="" />,
      },
      {
        ...columnConfig,
        accessorKey: 'DestChainID',
        header: () => <span></span>,
        cell: (value: any) => <RollupIcon chainId={value.getValue()} />,
      },
      {
        ...columnConfig,
        accessorKey: 'DestAddress',
        header: () => <span>Address</span>,
        cell: (value: any) => (
          <Link
            target="_blank"
            to={getAddressUrl(
              value.row.original.SourceChainID,
              value.row.original.DestAddress,
              data.supportedChains,
            )}
            className="link"
          >
            <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
            <span className="icon-external-link" />
          </Link>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'TxHash',
        header: () => <span>Tx Hash</span>,
        cell: (value: any) => (
          <Link
            target="_blank"
            to={getTxUrl(
              value.row.original.SourceChainID,
              value.row.original.TxHash,
              data.supportedChains,
            )}
            className="link"
          >
            <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
            <span className="icon-external-link" />
          </Link>
        ),
      },
    ],
    [],
  )

  return (
    <div className="flex-none">
      <div className="flex flex-col">
        <h5 className="text-default mb-4">XMsgs</h5>
        <div className={'flex mb-4 gap-2 flex-col md:flex-row'}>
          <div className="flex w-full">
            <Dropdown
              position="left"
              options={filterOptions}
              onChange={value => {
                setSearchPlaceholder(
                  `Search by ${(filterOptions.find(option => option.value === value)?.display || filterOptions[0].display).toLowerCase()}`,
                )
              }}
              defaultValue={filterOptions[0].value}
            />
            <SearchBar placeholder={searchPlaceholder} />
          </div>
          <ChainDropdown placeholder="Select source" label="From" options={sourceChainList} />
          <ChainDropdown placeholder="Select destination" label="To" options={sourceChainList} />
        </div>
      </div>
      <div>
        <SimpleTable
          headChildren={
            <div className={`flex justify-between `}>
              <div className="table-highlight  w-[21.856%] min-w-[221px]"></div>
              <div className={`px-6 py-3`}>
                <FilterOptions
                  onSelection={e => {
                    console.log('Filter selected', e)
                  }}
                  options={['All', 'Success', 'Pending', 'Failed']}
                />
              </div>
            </div>
          }
          columns={columns}
          data={rows}
        />
      </div>
    </div>
  )
}
