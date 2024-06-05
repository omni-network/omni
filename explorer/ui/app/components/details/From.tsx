import { Link } from '@remix-run/react'
import React from 'react'
import { dateFormatter, hashShortener } from '~/lib/formatting'
import { dateFormatterXMsgPage } from './date.service'
import { copyToClipboard } from '~/lib/utils'

export const From = ({ xmsg }) => {
  return (
    <>
      <h6 className="text-default my-5 text-lg">From</h6>
      {/* Source Chain */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Source Chain</p>
        <div className='flex gap-2 items-center'>
          <img className='w-5 h-5 rounded-full' src={xmsg.sourceChain.logoUrl} />
          <p className="text-default">{xmsg.sourceChain.name}</p>
        </div>
      </div>
      {/* Source Address */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Source Address</p>
        <Link
          target="_blank"
          to={xmsg?.senderUrl}
          className="underline text-indigo-400 hidden lg:block"
        >
          {xmsg?.sender}
          <span className="icon-external-link" />
        </Link>
        <Link
          target="_blank"
          to={xmsg?.senderUrl}
          className="underline text-indigo-400 block lg:hidden"
        >
          <span className="font-bold text-b-sm">{hashShortener(xmsg?.sender)}</span>
          <span className="icon-external-link" />
        </Link>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xmsg?.sender)}
        />
      </div>
      {/* Tx Hash */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Tx Hash</p>
        <Link
          target="_blank"
          to={xmsg?.txUrl}
          className="underline text-indigo-400 hidden lg:block"
        >
          {xmsg?.txHash}
          <span className="icon-external-link" />
        </Link>
        <Link
          target="_blank"
          to={xmsg?.txUrl}
          className="underline text-indigo-400 block lg:hidden"
        >
          <span className="font-bold text-b-sm">{hashShortener(xmsg?.txHash)}</span>
          <span className="icon-external-link" />
        </Link>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xmsg?.txHash)}
        />
      </div>
      {/* Block Timestamp */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Block Timestamp</p>
        <p className="flex-1">
        {xmsg?.block?.timestamp ? (
          <>
            <span className="text-default">
              {dateFormatter(new Date(xmsg.block.timestamp))}{' '}
            </span>
            ({dateFormatterXMsgPage(new Date(xmsg.block.timestamp))})
          </>
        ) : null}
        </p>
      </div>
      {/* Block height */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Block Height</p>
        <p className="text-default">{parseInt(xmsg?.block.height)}</p>
      </div>
      {/* Block Hash */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Block Hash</p>
        <p className="text-default hidden lg:block">{xmsg?.block.hash}</p>
        <p className="text-default block lg:hidden">{hashShortener(xmsg?.block.hash)}</p>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xmsg?.block.hash)}
        />
      </div>

    </>
  )
}
