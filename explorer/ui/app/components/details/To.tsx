import { Link } from '@remix-run/react'
import React from 'react'
import { dateFormatter, hashShortener } from '~/lib/formatting'
import { dateFormatterXMsgPage } from './date.service';

export const To = ({ xMsgDetails }) => {
  return (
    <>
    <h5 className="text-default my-5">To</h5>
    {/* Destination Chain */}
    <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Destination Chain</p>
        <p>
          <span className="text-default">Optimism</span>(356)
        </p>
      </div>
      {/* Block Timestamp */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Block Timestamp</p>
        <p className='flex-1'>
          <span className="text-default"></span> {dateFormatterXMsgPage(new Date(xMsgDetails.block.timestamp))}
        </p>
      </div>
       {/* Destination Address */}
       <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Destination Address</p>
        <Link target='_blank' to={xMsgDetails.toUrl} className="underline text-indigo-400 hidden lg:block">
          {xMsgDetails.to}
        </Link>
        <Link target='_blank' to={xMsgDetails.toUrl} className="underline text-indigo-400 block lg:hidden">
          {hashShortener(xMsgDetails.to)}
        </Link>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Tx Hash</p>
        <Link target='_blank' to={xMsgDetails.txHashUrl} className="underline text-indigo-400 hidden lg:block">
          {xMsgDetails.receipt.txHash}
        </Link>
        <Link target='_blank' to={xMsgDetails.txHashUrl} className="underline text-indigo-400 block lg:hidden">
          {hashShortener(xMsgDetails.receipt.txHash)}
        </Link>
      </div>
      {/* Gas Limit */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Gas Limit</p>
        <p className="text-default">
          {xMsgDetails.gasLimit}
        </p>
      </div>
    </>
  )
}
