import { Link } from '@remix-run/react'
import React from 'react'
import { hashShortener } from '~/lib/formatting'

export const From = () => {
  return (
    <>
      <h5 className="text-default my-5">From</h5>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Source Chain</p>
        <p>
          <span className="text-default">Arbitrum </span> (36)
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
        <p className="w-48">Tx Hash</p>
        <Link to="/" className="underline text-indigo-400 hidden lg:block">
          0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4
        </Link>
        <Link to="/" className="underline text-indigo-400 block lg:hidden">
          {hashShortener('0x109f40f806567158aaad05e43afe240cf394608cacd0016466dfb24dce2927d4')}
        </Link>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Offset</p>
        <p className="text-default">2345</p>
      </div>
      {/* Offset */}
      <div className="flex mt-5 pb-2 border-b-2 border-gray-500 border-solid">
        <p className="w-48">Offset</p>
        <Link to="/" className="underline text-indigo-400">
          0xfc39076
        </Link>
      </div>
    </>
  )
}
