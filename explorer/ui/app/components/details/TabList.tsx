import { Link } from '@remix-run/react'
import React, { useEffect } from 'react'
import Tag from '../shared/tag'
import { dateFormatter, hashShortener } from '~/lib/formatting'
import { dateFormatterXMsgPage } from './date.service'
import { copyToClipboard } from '~/lib/utils'

export const TabList = ({ xMsgDetails }) => {
  return (
    <>
      <h6 className="text-default my-5 mt-20 text-lg">XReceipt</h6>
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Timestamp</p>
        <p className="flex-1">
          <span className="text-default">
            {dateFormatter(new Date(xMsgDetails?.block.timestamp))}{' '}
          </span>
          (
          {xMsgDetails?.receipt?.timestamp &&
            dateFormatterXMsgPage(new Date(xMsgDetails?.receipt?.timestamp))}
          )
        </p>
      </div>
      {/* Source Address */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Source Address</p>
        <Link target="_blank" to="/" className="underline text-indigo-400 hidden lg:block">
          {xMsgDetails?.txHash}
          <span className="icon-external-link" />
        </Link>
        <Link target="_blank" to="/" className="underline text-indigo-400 block lg:hidden">
          <span className="font-bold text-b-sm">{hashShortener(xMsgDetails?.txHash)}</span>
          <span className="icon-external-link" />
        </Link>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xMsgDetails?.txhash)}
        />
      </div>
      {/* Relayer Address */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Relayer Address</p>
        <Link target="_blank" to="/" className="underline text-indigo-400 hidden lg:block">
          {xMsgDetails?.to}
          <span className="icon-external-link" />
        </Link>
        <Link target="_blank" to="/" className="underline text-indigo-400 block lg:hidden">
          <span className="font-bold text-b-sm">{hashShortener(xMsgDetails?.to)}</span>
          <span className="icon-external-link" />
        </Link>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xMsgDetails?.to)}
        />
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Stream Offset Address</p>
        <p className="text-default">{xMsgDetails?.offset}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Gas Used</p>
        <p className="text-default">{xMsgDetails?.receipt?.gasUsed}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Destination Gas Limit</p>
        <p className="text-default">{xMsgDetails?.gasLimit}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Source Chain ID</p>
        <p className="text-default">{xMsgDetails?.sourceChain.chainID}</p>
      </div>
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Destination Chain ID</p>
        <p className="text-default">{xMsgDetails?.destChain.chainID}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Tx Hash</p>
        <Link
          target="_blank"
          to={xMsgDetails?.receipt?.txUrl}
          className="underline text-indigo-400 hidden lg:block"
        >
          {xMsgDetails?.receipt?.txHash}
          <span className="icon-external-link" />
        </Link>
        <Link
          target="_blank"
          to={xMsgDetails?.receipt?.txUrl}
          className="underline text-indigo-400 block lg:hidden"
        >
          <span className="font-bold text-b-sm">{hashShortener(xMsgDetails?.receipt?.txHash)}</span>
          <span className="icon-external-link" />
        </Link>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xMsgDetails?.receipt?.txHash)}
        />
      </div>
      {/* Status */}
      <div className="flex mt-5 pb-2 border-b-[1px] mb-3 border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Status</p>
        <div className="flex flex-col sm:flex-row items-start">
          <Tag status={xMsgDetails?.status} />
          <p className="sm:ml-5">{xMsgDetails?.receipt?.revertReason}</p>
        </div>
      </div>
    </>
  )
}
