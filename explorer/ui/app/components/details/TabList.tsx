import { Link } from '@remix-run/react'
import React, { useEffect } from 'react'
import Tag from '../shared/tag'
import { dateFormatter, hashShortener } from '~/lib/formatting'
import { dateFormatterXMsgPage } from './date.service'
import { copyToClipboard } from '~/lib/utils'

export const TabList = ({ xmsg }) => {
  return (
    <>
      <h6 className="text-default my-5 mt-20 text-lg">XReceipt</h6>
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Timestamp</p>
        <p className="flex-1">
          {xmsg?.receipt ? (
          <>
            <span className="text-default">
              {dateFormatter(new Date(xmsg.receipt.timestamp))}{' '}
            </span>
            (
            {dateFormatterXMsgPage(new Date(xmsg.receipt.timestamp))}
            )
          </>
          ) : '--'}
        </p>
      </div>
      {/* Source Address */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Source Address</p>
        <Link target="_blank" to={xmsg?.senderUrl} className="underline text-indigo-400 hidden lg:block">
          {xmsg?.sender}
          <span className="icon-external-link" />
        </Link>
        <Link target="_blank" to={xmsg?.senderUrl} className="underline text-indigo-400 block lg:hidden">
          <span className="font-bold text-b-sm">{hashShortener(xmsg?.sender)}</span>
          <span className="icon-external-link" />
        </Link>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xmsg?.sender)}
        />
      </div>
      {/* Dest Address */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Dest Address</p>
        <Link target="_blank" to={xmsg?.toUrl} className="underline text-indigo-400 hidden lg:block">
          {xmsg?.to}
          <span className="icon-external-link" />
        </Link>
        <Link target="_blank" to={xmsg?.toUrl} className="underline text-indigo-400 block lg:hidden">
          <span className="font-bold text-b-sm">{hashShortener(xmsg?.to)}</span>
          <span className="icon-external-link" />
        </Link>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xmsg?.to)}
        />
      </div>
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Stream Offset</p>
        <p className="text-default">{parseInt(xmsg?.offset, 16)}</p>
      </div>
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Gas Used</p>
        <p className="text-default">{xmsg?.receipt ? parseInt(xmsg?.receipt?.gasUsed) : '--'}</p>
      </div>
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Destination Gas Limit</p>
        <p className="text-default">{parseInt(xmsg?.gasLimit)}</p>
      </div>
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Source Chain ID</p>
        <p className="text-default">{parseInt(xmsg?.sourceChain.chainID, 16)} ({xmsg?.sourceChain.chainID})</p>
      </div>
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Destination Chain ID</p>
        <p className="text-default">{parseInt(xmsg?.destChain.chainID, 16)} ({xmsg?.destChain.chainID})</p>
      </div>
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Tx Hash</p>
        {xmsg?.receipt ?
          <>
          <Link
            target="_blank"
            to={xmsg?.receipt?.txUrl}
            className="underline text-indigo-400 hidden lg:block"
          >
            {xmsg?.receipt?.txHash}
            <span className="icon-external-link" />
          </Link>
          <Link
            target="_blank"
            to={xmsg?.receipt?.txUrl}
            className="underline text-indigo-400 block lg:hidden"
          >
            <span className="font-bold text-b-sm">{hashShortener(xmsg?.receipt?.txHash)}</span>
            <span className="icon-external-link" />
          </Link>
          <span
            data-tooltip-id="tooltip-clipboard"
            className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
            onClick={() => copyToClipboard(xmsg?.receipt?.txHash)}
          />
        </> : '--'}
      </div>
      <div className="flex mt-5 pb-2 border-b-[1px] mb-3 border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Status</p>
        <div className="flex flex-col sm:flex-row items-start">
          <Tag status={xmsg?.status} />
          <p className="sm:ml-5">{xmsg?.receipt?.revertReason}</p>
        </div>
      </div>
    </>
  )
}
