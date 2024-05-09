import React, { RefObject, useCallback, useEffect, useMemo } from 'react'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'
import { useLoaderData, useRevalidator, useSearchParams } from '@remix-run/react'
import { dateFormatter, hashShortener } from '~/lib/formatting'
import Tag from '../shared/tag'
import RollupIcon from '../shared/rollupIcon'
import { Link } from '@remix-run/react'
import LongArrow from '~/assets/images/LongArrow.svg'
import { XmsgResponse } from '~/routes/_index'
import SearchBar from '../shared/search'
import ChainDropdown from './chainDropdown'
import FilterOptions from '../shared/filterOptions'
import { getBaseUrl } from '~/lib/sourceChains'
import debounce from 'lodash.debounce'
import { Tooltip } from 'react-tooltip'
import Button from '../shared/button'
import { PageButton } from '../shared/button-legacy'
import { copyToClipboard } from '~/lib/utils'

type Status = 'Success' | 'Failed' | 'Pending' | 'All'

export default function XMsgDataTable() {
  const data = useLoaderData<XmsgResponse>()
  const revalidator = useRevalidator()
  const searchFieldRef = React.useRef<HTMLInputElement>(null)

  const pageLoaded = React.useRef<boolean>(false)

  const [searchParams, setSearchParams] = useSearchParams()

  const [filterParams, setFilterParams] = React.useState<{
    address: string | null
    txHash: string | null
    sourceChain: string | null
    destChain: string | null
    status: Status
    before: string | null
    after: string | null
  }>({
    address: searchParams.get('address') ?? null,
    sourceChain: searchParams.get('sourceChain') ?? null,
    destChain: searchParams.get('destChain') ?? null,
    txHash: searchParams.get('txHash') ?? null,
    status: (searchParams.get('status') as Status) ?? 'All',
    before: searchParams.get('before') ?? null,
    after: searchParams.get('after') ?? null,
  })

  const sourceChainList = []
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
      after: null,
      before: null,
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

    if (pageLoaded.current) {
      revalidator.revalidate()
    } else {
      pageLoaded.current = true
    }
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
        accessorKey: 'node.displayID',
        header: () => <span>ID</span>,
        cell: (value: any) => {
          return (
            <>
              <Link
                data-tooltip-id={`${value.getValue()}-full-id-tooltip`}
                to={`xmsg/${value.getValue()}`}
                className="link"
              >
                {hashShortener(value.getValue())}
              </Link>
              <Tooltip className="tooltip" id={`${value.getValue()}-full-id-tooltip`}>
                <label className="text-default text-b-sm font-bold">
                  {value.getValue()}
                </label>
              </Tooltip>
              {/* <span  className="icon-tooltip-info"></span> */}
            </>
          )
        },
      },
      // cant see the data
      {
        ...columnConfig,
        accessorKey: 'node.receipt.timestamp',
        header: () => <span>Age</span>,
        cell: (value: any) => (
          <span className="text-subtlest font-bold text-b-xs">
            {' '}
            {dateFormatter(new Date(value.getValue()))}
          </span>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'node.status',
        header: () => <span>Status</span>,
        cell: (value: any) => <Tag status={value.getValue()} />,
      },
      {
        ...columnConfig,
        accessorKey: 'node.sourceChainID',
        header: () => <span></span>,
        cell: (value: any) => <RollupIcon chainId={value.getValue()} />,
      },
      {
        ...columnConfig,
        accessorKey: 'node.sender',
        header: () => (
          <div className="flex items-center">
            <span>Source Address</span>
            <Tooltip className="tooltip" id="address-info">
              <label className="text-default text-b-sm font-bold">
                Sender on the source chain, <br /> set to msg.Sender
              </label>
            </Tooltip>
            <span data-tooltip-id={'address-info'} className="icon-tooltip-info"></span>
          </div>
        ),
        cell: (value: any) => {
          return (
            <>
              <Link to={`${value.row.original.node.senderUrl}`} className="link">
                {value.getValue() && (
                  <>
                    <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
                    <span className="icon-external-link" />
                  </>
                )}
              </Link>
              <span
                data-tooltip-id="tooltip-clipboard"
                className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
                onClick={() => copyToClipboard(value.getValue())}
              />
            </>
          )
        },
      },
      {
        ...columnConfig,
        accessorKey: 'node.txHash',
        header: () => (
          <div className="flex items-center">
            <span>Tx Hash</span>
            <Tooltip className="tooltip" id="tx-hash-info">
              <label className="text-default text-b-sm font-bold">
                Hash of the source chain <br /> transaction that emitted the message
              </label>
            </Tooltip>
            <span data-tooltip-id={'tx-hash-info'} className="icon-tooltip-info"></span>
          </div>
        ),
        cell: (value: any) => {
          return (
            <>
              {value.getValue() && (
                <>
                  {' '}
                  <Link
                    target="_blank"
                    to={`${value.row.original.node.txHashUrl}`}
                    className="link"
                  >
                    <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
                    <span className="icon-external-link" />
                  </Link>
                  <span
                    data-tooltip-id="tooltip-clipboard"
                    className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
                    onClick={() => copyToClipboard(value.getValue())}
                  />{' '}
                </>
              )}
            </>
          )
        },
      },
      {
        ...columnConfig,
        accessorKey: 'Empty',
        header: () => <span></span>,
        cell: (value: any) => <img src={LongArrow} alt="" />,
      },
      {
        ...columnConfig,
        accessorKey: 'node.destChainID',
        header: () => <span></span>,
        cell: (value: any) => <RollupIcon chainId={value.getValue()} />,
      },
      {
        ...columnConfig,
        accessorKey: 'node.to',
        header: () => (
          <div className="flex items-center">
            <span>Destination Address</span>
            <Tooltip className="tooltip" id="receiver-address-info">
              <label className="text-default text-b-sm font-bold">
                Contract address on the destination <br /> chain that receives the call
              </label>
            </Tooltip>
            <span data-tooltip-id={'receiver-address-info'} className="icon-tooltip-info"></span>
          </div>
        ),
        cell: (value: any) => (
          <>
            <Link target="_blank" to={`${value.row.original.node.toUrl}`} className="link">
              <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
              <span className="icon-external-link" />
            </Link>
            <span
              data-tooltip-id="tooltip-clipboard"
              className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
              onClick={() => copyToClipboard(value.getValue())}
            />
          </>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'node.receipt.txHash',
        header: () => (
          <div className="flex items-center">
            <span>Tx Hash</span>
            <Tooltip className="tooltip" id="receiver-tx-hash-info">
              <label className="text-default text-b-sm font-bold">
                Hash of the transaction executed <br /> on the destination chain by the relayer
              </label>
            </Tooltip>
            <span data-tooltip-id={'receiver-tx-hash-info'} className="icon-tooltip-info"></span>
          </div>
        ),
        cell: (value: any) => {
          return (
            <>
              {value.getValue() && (
                <>
                  {' '}
                  <Link
                    target="_blank"
                    to={`${value.row.original.node.receipt.txHashUrl}`}
                    className="link"
                  >
                    <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
                    <span className="icon-external-link" />
                  </Link>
                  <span
                    data-tooltip-id="tooltip-clipboard"
                    className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
                    onClick={() => copyToClipboard(value.getValue())}
                  />{' '}
                </>
              )}
              {!value.getValue() && '----'}
            </>
          )
        },
      },
    ],
    [],
  )
  return (
    <div className="flex-none">
      <div className="flex flex-col">
        <h5 className="text-default mb-4">
          XMsgs{' '}
          <Tooltip className="tooltip" id="xmsg-info">
            <label className="text-default text-b-sm font-bold">
              XMsgs are cross-rollup messages. <br /> Click to learn more
            </label>
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
        <div className={`flex justify-between mb-4 flex-col md:flex-row`}>
          <div className={`flex justify-between mb-4 flex-col md:flex-row`}>
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
          <Button
            disabled={!hasFiltersApplied}
            onClick={clearFilters}
            kind="text"
            className={`flex items-center ${!hasFiltersApplied && 'opacity-40'}`}
          >
            {' '}
            <span className="icon-refresh text-default text-[20px]" />
            <span className={`text-default`}>Clear all filters</span>
          </Button>
        </div>
      </div>
      <div>
        {/* <div className='rounded-xl bg-raised p-10 min-h-[650px]'>
          <h4 className='text-default mb-4'>No result found.</h4>
          <p className='text-default text-b'>Please edit your filter selection and try again.</p>
        </div> */}

        {!data && (
          <div role="status" className="animate-pulse overflow-x-auto">
            <div className="w-full bg-raised rounded-lg min-w-[919px] h-96"></div>
          </div>
        )}

        {data && (
          <>
            {rows.length === 0 && (
              <div className="bg-raised p-5">
                <h3>No results found.</h3>
                <p>Please edit your filter selection and try again.</p>
              </div>
            )}
            {rows.length > 0 && <SimpleTable columns={columns} data={rows} />}
          </>
        )}

        {/* Nav Buttons */}
        <div className="flex flex-row items-center justify-end mt-4">
          <PageButton
            className="rounded-full flex items-center justify-center"
            onClick={() => {
              setFilterParams(prev => ({ ...prev, after: null, before: data.xmsgs[0].cursor}))
            }} // TODO: when clicked it needs to update the search params with the new cursor
            disabled={!data.pageInfo.hasPrevPage} // TODO: When there is no previous cursor, we need to disable this
          >
            <span className="sr-only">Previous</span>
            <span className={`icon-chevron-med-left text-[20px]`}></span>
          </PageButton>

          {/* Page N of N */}
          <div className="flex-none flex m-3">
            <div className="flex gap-x-2 items-baseline">
              <span className="text-cb-sm text-default">
                Page <span className="">{data?.pageInfo?.currentPage}</span> of{' '}
                <span className="">{data?.pageInfo?.totalPages}</span>
              </span>
            </div>
          </div>

          <PageButton
            className="rounded-full  flex items-center justify-center"
            onClick={() => {
              setFilterParams(prev => ({ ...prev, after: data.xmsgs[9].cursor, before: null}))
            }} // TODO: when clicked it needs to update the search params with the new cursor
            disabled={!data.pageInfo.hasNextPage} // TODO: When there is no next cursor, we need to disable this
          >
            <span className="sr-only">Next</span>
            <span className={`icon-chevron-med-right text-[20px]`}></span>
          </PageButton>
        </div>
      </div>

      {/* tooltip for offset */}
      <Tooltip className="tooltip" id={'tooltip-offset'} />
      <Tooltip className="tooltip" id="tooltip-clipboard">
        <span className="text-default text-b-sm font-bold">Copy to clipboard </span>
      </Tooltip>
    </div>
  )
}
