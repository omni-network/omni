import { json } from '@remix-run/node'
import React, { RefObject, useCallback, useEffect, useMemo } from 'react'
import { XMsg } from '~/graphql/graphql'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'
import { useLoaderData, useRevalidator, useSearchParams } from '@remix-run/react'
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
import { getAddressUrl, getBaseUrl, getBlockUrl, getTxUrl } from '~/lib/sourceChains'
import debounce from 'lodash.debounce'
import { Tooltip } from 'react-tooltip'
import Button from '../shared/button'
import { PageButton } from '../shared/button-legacy'
import { copyToClipboard } from '~/lib/utils'

export default function XMsgDataTable() {
  const data = useLoaderData<typeof loader>()
  const revalidator = useRevalidator()
  const searchFieldRef = React.useRef<HTMLInputElement>(null)

  const [filterParams, setFilterParams] = React.useState<{
    address: string | null
    txHash: string | null
    sourceChain: string | null
    destChain: string | null
    status: 'Success' | 'Failed' | 'Pending' | 'All'
  }>({
    address: null,
    sourceChain: null,
    destChain: null,
    txHash: null,
    status: 'All',
  })

  const [searchParams, setSearchParams] = useSearchParams()

  const sourceChainList = data.supportedChains.map(chain => ({
    value: chain.ChainID,
    display: chain.DisplayName,
    icon: chain.Icon,
  }))

  const rows = data.xmsgs

  const columnConfig = {
    canFilter: false,
    enableColumnFilter: false,
  }

  const clearFilters = () => {
    if (searchFieldRef.current) {
      searchFieldRef.current.value = ''
    }

    setFilterParams({
      address: null,
      sourceChain: null,
      destChain: null,
      txHash: null,
      status: 'All',
    })
  }

  const hasFiltersApplied: boolean =
    Object.values(filterParams).filter(val => val !== 'All' && val !== null).length > 0

  // Listen for filter changes here and append search params
  useEffect(() => {
    const newParams = new URLSearchParams()
    for (var key in filterParams) {
      if (filterParams[key] !== null) {
        newParams.set(key, filterParams[key])
      } else {
        newParams.delete(key)
      }
    }

    setSearchParams(newParams)
    revalidator.revalidate()
  }, [filterParams])

  // here we set the filter params by clearing the old ones, and setting the current one and its value
  const searchBarInputCB = e => {
    const isAddress = e.target.value.match(/^0x[0-9a-fA-F]{40}$/)
    const isTxHash = e.target.value.match(/^0x[0-9a-fA-F]{64}$/)

    if (!isAddress && !isTxHash && e.target.value !== '') {
      // return user error cause it doesn't match either
      alert("It doesn't match")
    }

    setFilterParams(prev => {
      const params = {
        ...prev,
        address: isAddress ? e.target.value : null,
        txHash: isTxHash ? e.target.value : null,
      }

      return params
    })
  }

  const searchBarInput = useCallback(debounce(searchBarInputCB, 600), [])

  const columns = React.useMemo<ColumnDef<any>[]>(
    () => [
      {
        ...columnConfig,
        accessorKey: 'StreamOffset',
        header: () => <span>Offset</span>,
        cell: (value: any) => (
          <>
            <Tooltip className="bg-overlay" id="my-tooltip">
              <span className="text-default text-b-sm font-bold">
                Hex value:{value.getValue()}{' '}
              </span>
            </Tooltip>
            <span data-tooltip-id="my-tooltip" className=" font-bold text-b-sm">
              {Number(value.getValue())}
            </span>
          </>
        ),
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
        accessorKey: 'SourceMessageSender',
        header: () => <span>Address</span>,
        cell: (value: any) => (
          <>
            <Link
              to={`${getBaseUrl(value.row.original.SourceChainID, 'senderAddress')}/${value.getValue()}`}
              className="link"
            >
              {value.getValue() && (
                <>
                  <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
                  <span className="icon-external-link" />
                </>
              )}
            </Link>
            <span
              className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
              onClick={() => copyToClipboard(value.getValue())}
            />
          </>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'BlockHash',
        header: () => <span>Tx Hash</span>,
        cell: (value: any) => (
          <>
            <Link
              target="_blank"
              to={`${getBaseUrl(value.row.original.SourceChainID, 'blockHash')}/${value.getValue()}`}
              className="link"
            >
              <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
              <span className="icon-external-link" />
            </Link>
            <span
              className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
              onClick={() => copyToClipboard(value.getValue())}
            />
          </>
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
          <>
            <Link
              target="_blank"
              to={`${getBaseUrl(value.row.original.SourceChainID, 'destHash')}/${value.getValue()}`}
              className="link"
            >
              <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
              <span className="icon-external-link" />
            </Link>
            <span
              className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
              onClick={() => copyToClipboard(value.getValue())}
            />
          </>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'TxHash',
        header: () => <span>Tx Hash</span>,
        cell: (value: any) => (
          <>
            <Link
              target="_blank"
              to={`${getBaseUrl(value.row.original.SourceChainID, 'tx')}/${value.getValue()}`}
              className="link"
            >
              <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
              <span className="icon-external-link" />
            </Link>
            <span
              className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
              onClick={() => copyToClipboard(value.getValue())}
            />
          </>
        ),
      },
    ],
    [],
  )

  return (
    <div className="flex-none">
      <div className="flex flex-col">
        <h5 className="text-default mb-4">
          XMsgs{' '}
          <Tooltip className="bg-overlay" id="xmsg-info">
            <span className="text-default text-b-sm font-bold">
              Click to view more info about the xMsg{' '}
            </span>
          </Tooltip>
          <Link
            data-tooltip-id={'xmsg-info'}
            target="_blank"
            to="https://docs.omni.network/protocol/xmessages/xmsg"
          >
            <span className="icon-tooltip-info"></span>
          </Link>
        </h5>

        <div className={'flex mb-4 gap-2 flex-col md:flex-row'}>
          <div className="flex w-full">
            <SearchBar
              ref={searchFieldRef}
              onInput={searchBarInput}
              placeholder={'Search by address/tx hash'}
            />
          </div>
          <ChainDropdown
            onChange={e => {
              setFilterParams(prev => ({
                ...prev,
                sourceChain: e,
              }))
            }}
            placeholder="Select source"
            label="From"
            options={sourceChainList}
            value={filterParams.sourceChain}
          />
          <ChainDropdown
            onChange={e => {
              setFilterParams(prev => ({
                ...prev,
                destChain: e,
              }))
            }}
            placeholder="Select destination"
            label="To"
            options={sourceChainList}
            value={filterParams.destChain}
          />
        </div>
        <div className={`flex justify-between mb-4`}>
          <div className="">
            <FilterOptions
              value={filterParams.status}
              onSelection={status => {
                setFilterParams(prev => ({
                  ...prev,
                  status: status === 'all' ? null : status,
                }))
              }}
              options={['All', 'Success', 'Pending', 'Failed']}
            />
          </div>
          {hasFiltersApplied && (
            <Button
              onClick={clearFilters}
              kind="text"
              className={`flex justify-center items-center`}
            >
              {' '}
              <span className="icon-refresh  text-default text-[20px]" />
              <span className="text-default">Clear all filters</span>
            </Button>
          )}
        </div>
      </div>
      <div>
        <SimpleTable columns={columns} data={rows} />

        {/* Nav Buttons */}
        <div className="flex flex-row items-center justify-end mt-4">
          <PageButton
            className="rounded-full flex items-center justify-center"
            onClick={() => {}} // TODO: when clicked it needs to update the search params with the new cursor
            disabled={false} // TODO: When there is no previous cursor, we need to disable this
          >
            <span className="sr-only">Previous</span>
            <span className={`icon-chevron-med-left text-[20px]`}></span>
          </PageButton>

          {/* Page N of N */}
          <div className="flex-none flex m-3">
            <div className="flex gap-x-2 items-baseline">
              <span className="text-cb-sm text-default">
                Page <span className="">{0}</span> of <span className="">{0}</span>
              </span>
            </div>
          </div>

          <PageButton
            className="rounded-full  flex items-center justify-center"
            onClick={() => {}} // TODO: when clicked it needs to update the search params with the new cursor
            disabled={false} // TODO: When there is no next cursor, we need to disable this
          >
            <span className="sr-only">Next</span>
            <span className={`icon-chevron-med-right text-[20px]`}></span>
          </PageButton>
        </div>
      </div>
    </div>
  )
}
