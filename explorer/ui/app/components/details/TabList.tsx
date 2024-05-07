import { Link } from '@remix-run/react'
import React, { useEffect } from 'react'
import Tag from '../shared/tag'
import { hashShortener } from '~/lib/formatting'

export const TabList = () => {
  const ReceiptList = [1, 2, 3]
  const activeReceiptIndex = 0

  const checkActiveCss = (index) => {
    if (index === activeReceiptIndex) {

    }
    return 'active inline-flex active items-center justify-center p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 group'
  }

  useEffect(() => {}, [activeReceiptIndex])

  return (
    <>
      <h5 className="text-default my-5 mt-5">XReceipt</h5>
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Block Timestamp</p>
        <p>
          <span className="text-default">16 mins ago </span> (Apr 25 2024 00:05:23 AM +UTC)
        </p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Source Address</p>
        <Link to="/" className="underline text-indigo-400 hidden lg:block">
          0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4
        </Link>
        <Link to="/" className="underline text-indigo-400 block lg:hidden">
          {hashShortener('0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4')}
        </Link>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Relayer Address</p>
        <Link to="/" className="underline text-indigo-400 hidden lg:block">
          0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4
        </Link>
        <Link to="/" className="underline text-indigo-400 block lg:hidden">
          {hashShortener('0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4')}
        </Link>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Stream Offset Address</p>
        <p className="text-default">345</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Gas Used</p>
        <p className="text-default">28,909</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Gas Used</p>
        <p className="text-default">30,000</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Source Chain ID</p>
        <p className="text-default">34</p>
      </div>
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Destination Chain ID</p>
        <p className="text-default">3567</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Tx Hash</p>
        <Link to="/" className="underline text-indigo-400 hidden lg:block">
          0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4
        </Link>
        <Link to="/" className="underline text-indigo-400 block lg:hidden">
          {hashShortener('0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4')}
        </Link>
      </div>
      {/* Status */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Status</p>
        <Tag status="FAILED" />
        <p className="ml-5">"Reason for status"</p>
      </div>
    </>
  )
}
