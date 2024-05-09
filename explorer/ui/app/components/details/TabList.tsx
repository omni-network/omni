import { Link } from '@remix-run/react'
import React, { useEffect } from 'react'
import Tag from '../shared/tag'
import { hashShortener } from '~/lib/formatting'
import { dateFormatterXMsgPage } from './date.service'

export const TabList = ({ xMsgDetails }) => {
  console.log(xMsgDetails)

  return (
    <>
      <h5 className="text-default my-5 mt-5">XReceipt</h5>
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Timestamp</p>
        <p className='flex-1'>
          <span className="text-default"></span>{' '}
          {dateFormatterXMsgPage(new Date(xMsgDetails.receipt.timestamp))}
        </p>
      </div>
      {/* Source Address */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Source Address</p>
        <Link target="_blank" to="/" className="underline text-indigo-400 hidden lg:block">
          0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4
        </Link>
        <Link target="_blank" to="/" className="underline text-indigo-400 block lg:hidden">
          {hashShortener('0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4')}
        </Link>
      </div>
      {/* Relayer Address */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Relayer Address</p>
        <Link target="_blank" to="/" className="underline text-indigo-400 hidden lg:block">
          0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4
        </Link>
        <Link target="_blank" to="/" className="underline text-indigo-400 block lg:hidden">
          {hashShortener('0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4')}
        </Link>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Stream Offset Address</p>
        <p className="text-default">{xMsgDetails.offset}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Gas Used</p>
        <p className="text-default">{xMsgDetails.receipt.gasUsed}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Destination Gas Limit</p>
        <p className="text-default">{xMsgDetails.gasLimit}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Source Chain ID</p>
        <p className="text-default">{xMsgDetails.sourceChainID}</p>
      </div>
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Destination Chain ID</p>
        <p className="text-default">{xMsgDetails.destChainID}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Tx Hash</p>
        <Link
          target="_blank"
          to={xMsgDetails.receipt.txHashUrl}
          className="underline text-indigo-400 hidden lg:block"
        >
          {xMsgDetails.receipt.txHash}
        </Link>
        <Link
          target="_blank"
          to={xMsgDetails.receipt.txHashUrl}
          className="underline text-indigo-400 block lg:hidden"
        >
          {hashShortener(xMsgDetails.receipt.txHash)}
        </Link>
      </div>
      {/* Status */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48">Status</p>
        <div className="flex flex-col sm:flex-row items-start">
          <Tag status={xMsgDetails.status} />
          <p className="sm:ml-5">{xMsgDetails.receipt.revertReason}</p>
        </div>
      </div>
    </>
  )
}
