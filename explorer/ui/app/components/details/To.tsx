import { Link } from '@remix-run/react'
import React from 'react'
import { dateFormatter, hashShortener } from '~/lib/formatting'
import { dateFormatterXMsgPage } from './date.service'
import { copyToClipboard } from '~/lib/utils'

export const To = ({ xmsg }) => {
  return (
    <>
      <h6 className="text-default my-5 text-lg">To</h6>
      {/* Destination Chain */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Destination Chain</p>
        <div className="flex gap-2 items-center">
          <img className="w-5 h-5 rounded-full" src={xmsg.destChain.logoUrl} />
          <p className="text-default">{xmsg.destChain.name}</p>
        </div>
      </div>
      {/* Destination Address */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Destination Address</p>
        <Link
          target="_blank"
          to={xmsg?.toUrl}
          className="underline text-indigo-400 hidden lg:block"
        >
          {xmsg?.to}
          <span className="icon-external-link" />
        </Link>
        <Link
          target="_blank"
          to={xmsg?.toUrl}
          className="underline text-indigo-400 block lg:hidden"
        >
          <span className="font-bold text-b-sm">{hashShortener(xmsg?.to)}</span>
          <span className="icon-external-link" />
        </Link>
        <span
          data-tooltip-id="tooltip-clipboard"
          className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
          onClick={() => copyToClipboard(xmsg?.to)}
        />
      </div>
      {/* Tx Hash */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Tx Hash</p>
        {xmsg?.receipt ? (
          <>
          <Link
            target="_blank"
            to={xmsg.receipt.txUrl}
            className="underline text-indigo-400 hidden lg:block"
          >
            {xmsg.receipt.txHash}
            <span className="icon-external-link" />
          </Link>
          <Link
            target="_blank"
            to={xmsg.receipt.txUrl}
            className="underline text-indigo-400 block lg:hidden"
          >
            <span className="font-bold text-b-sm">{hashShortener(xmsg.receipt.txHash)}</span>
            <span className="icon-external-link" />
          </Link>
          <span
            data-tooltip-id="tooltip-clipboard"
            className="icon-copy cursor-pointer text-default hover:text-subtlest text-[16px] active:text-success transition-color ease-out duration-150"
            onClick={() => copyToClipboard(xmsg.receipt.txHash)}
          />
        </>
      ) : '--'}
      </div>
      {/* Gas Limit */}
      <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
        <p className="w-[150px] sm:w-48 text-sm">Gas Limit</p>
        <p className="text-default">{parseInt(xmsg?.gasLimit, 16)}</p>
      </div>
    </>
  )
}
