import { Link } from '@remix-run/react'
import React from 'react'
import { dateFormatter, hashShortener } from '~/lib/formatting'
import { dateFormatterXMsgPage } from './date.service';

export const To = ({ xMsgDetails }) => {
  return (
    <>
      <h6 className="text-default my-5 mt-20 text-lg">To</h6>
    {/* Destination Chain */}
    <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Destination Chain</p>
        <p>
          <span className="text-default">Optimism</span>(356)
        </p>
      </div>
      {/* Block Timestamp */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Block Timestamp</p>
        <p className='flex-1'>
          <span className="text-default"></span> {dateFormatterXMsgPage(new Date(xMsgDetails.block.timestamp))}
        </p>
      </div>
       {/* Destination Address */}
       <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Destination Address</p>
        <Link target='_blank' to={xMsgDetails.toUrl} className="underline text-indigo-400 hidden lg:block">
          {xMsgDetails.to}
        </Link>
        <Link target='_blank' to={xMsgDetails.toUrl} className="underline text-indigo-400 block lg:hidden">
          {hashShortener(xMsgDetails.to)}
        </Link>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Tx Hash</p>
        <Link target='_blank' to={xMsgDetails.txHashUrl} className="underline text-indigo-400 hidden lg:block">
          {xMsgDetails.receipt.txHash}
        </Link>
        <Link target='_blank' to={xMsgDetails.txHashUrl} className="underline text-indigo-400 block lg:hidden">
          {hashShortener(xMsgDetails.receipt.txHash)}
        </Link>
      </div>
      {/* Gas Limit */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Gas Limit</p>
        <p className="text-default">
          {xMsgDetails.gasLimit}
        </p>
      </div>
    </>
  )
}
