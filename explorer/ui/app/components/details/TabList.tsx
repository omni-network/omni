import { Link } from '@remix-run/react'
import React, { useEffect } from 'react'
import Tag from '../shared/tag'
import { hashShortener } from '~/lib/formatting'
import { dateFormatterXMsgPage } from './date.service'

export const TabList = ({ xMsgDetails }) => {
  return (
    <>
      <h6 className="text-default my-5 mt-20 text-lg">XReceipt</h6>
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Timestamp</p>
        <p className='flex-1'>
          <span className="text-default"></span>{' '}
          {dateFormatterXMsgPage(new Date(xMsgDetails.receipt.timestamp))}
        </p>
      </div>
      {/* Source Address */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Source Address</p>
        <Link target="_blank" to="/" className="underline text-indigo-400 hidden lg:block">
          0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4
        </Link>
        <Link target="_blank" to="/" className="underline text-indigo-400 block lg:hidden">
          {hashShortener('0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4')}
        </Link>
      </div>
      {/* Relayer Address */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Relayer Address</p>
        <Link target="_blank" to="/" className="underline text-indigo-400 hidden lg:block">
          0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4
        </Link>
        <Link target="_blank" to="/" className="underline text-indigo-400 block lg:hidden">
          {hashShortener('0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4')}
        </Link>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Stream Offset Address</p>
        <p className="text-default">{xMsgDetails.offset}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Gas Used</p>
        <p className="text-default">{xMsgDetails.receipt.gasUsed}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Destination Gas Limit</p>
        <p className="text-default">{xMsgDetails.gasLimit}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Source Chain ID</p>
        <p className="text-default">{xMsgDetails.sourceChainID}</p>
      </div>
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Destination Chain ID</p>
        <p className="text-default">{xMsgDetails.destChainID}</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Tx Hash</p>
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
        <p className="w-[150px] sm:w-48 text-sm">Status</p>
        <div className="flex flex-col sm:flex-row items-start">
          <Tag status={xMsgDetails.status} />
          <p className="sm:ml-5">{xMsgDetails.receipt.revertReason}</p>
        </div>
      </div>
    </>
  )
}
