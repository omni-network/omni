import { Link } from '@remix-run/react'
import React from 'react'
import { hashShortener } from '~/lib/formatting'
import { copyToClipboard } from '~/lib/utils'

export const From = ({ xMsgDetails }) => {
  return (
    <>
      <h6 className="text-default my-5 text-lg">From</h6>
      {/* Source Chain */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Source Chain</p>
        <div className='flex gap-2 items-center'>
          <img className='w-5 h-5 ' src={xMsgDetails.sourceChain.logoUrl} />
          <p className="text-default">{xMsgDetails.sourceChain.name}</p>
        </div>
      </div>
      {/* Source Address */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Source Address</p>
        <Link
          target="_blank"
          to={xMsgDetails?.senderUrl}
          className="underline text-indigo-400 hidden lg:block"
        >
          {xMsgDetails?.sender}
          <span className="icon-external-link" />
        </Link>
        <Link
          target="_blank"
          to={xMsgDetails?.senderUrl}
          className="underline text-indigo-400 block lg:hidden"
        >
          <span className="font-bold text-b-sm">{hashShortener(xMsgDetails?.sender)}</span>
          <span className="icon-external-link" />
        </Link>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xMsgDetails?.sender)}
        />
      </div>
      {/* Tx Hash */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Tx Hash</p>
        <Link
          target="_blank"
          to={xMsgDetails?.txUrl}
          className="underline text-indigo-400 hidden lg:block"
        >
          {xMsgDetails?.txHash}
          <span className="icon-external-link" />
        </Link>
        <Link
          target="_blank"
          to={xMsgDetails?.txUrl}
          className="underline text-indigo-400 block lg:hidden"
        >
          <span className="font-bold text-b-sm">{hashShortener(xMsgDetails?.txHash)}</span>
          <span className="icon-external-link" />
        </Link>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xMsgDetails?.txHash)}
        />
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Block Height</p>
        <p className="text-default">{xMsgDetails?.block.height}</p>
      </div>
      {/* Block Hash */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Block Hash</p>
        <p className="text-default hidden lg:block">{xMsgDetails?.block.hash}</p>
        <p className="text-default block lg:hidden">{hashShortener(xMsgDetails?.block.hash)}</p>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xMsgDetails?.block.hash)}
        />
      </div>

    </>
  )
}
