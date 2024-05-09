import { Link } from '@remix-run/react'
import React from 'react'
import { hashShortener } from '~/lib/formatting'

export const From = ({ xMsgDetails }) => {
  return (
    <>
      <h5 className="text-default my-5">From</h5>
      {/* Source Chain */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Source Chain</p>
        <p>
          <span className="text-default">Arbitrum </span> (36)
        </p>
      </div>
      {/* Source Address */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Source Address</p>
        <Link target='_blank' to={xMsgDetails.senderUrl} className="underline text-indigo-400 hidden lg:block">
          {xMsgDetails.sender}
        </Link>
        <Link target='_blank' to={xMsgDetails.senderUrl} className="underline text-indigo-400 block lg:hidden">
          {hashShortener(xMsgDetails.sender)}
        </Link>
      </div>
      {/* Tx Hash */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Tx Hash</p>
        <Link target='_blank' to={xMsgDetails.txHashUrl} className="underline text-indigo-400 hidden lg:block">
          {xMsgDetails.txHash}
        </Link>
        <Link target='_blank' to={xMsgDetails.txHashUrl} className="underline text-indigo-400 block lg:hidden">
          {hashShortener(xMsgDetails.txHash)}
        </Link>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Block Height</p>
        <p className="text-default">{xMsgDetails.block.height}</p>
      </div>
      {/* Block Hash */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Block Hash</p>
        <p className="text-default hidden lg:block">
          {xMsgDetails.block.hash}
        </p>
        <p className="text-default block lg:hidden">
          {hashShortener(xMsgDetails.block.hash)}
        </p>
      </div>
    </>
  )
}
