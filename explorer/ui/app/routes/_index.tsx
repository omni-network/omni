import type { LoaderFunction, MetaFunction } from '@remix-run/node'
import XBlockDataTable from '~/components/home/blockDataTable'
import XMsgDataTable from '~/components/home/messageDataTable'
import Overview from '~/components/home/overview'
import { json } from '@remix-run/node'
import { gqlClient } from '~/entry.server'
import { useFetcher, useRevalidator } from '@remix-run/react'
import { useInterval } from '~/hooks/useInterval'
import { xblockcount } from '~/components/queries/block'
import { xmsgrange } from '~/components/queries/messages'
import { XMsg } from '~/graphql/graphql'
import { supportedchains } from '~/components/queries/chains'

export const meta: MetaFunction = () => {
  return [
    { title: 'Omni Network Explorer' },
    { name: 'description', content: 'Omni Network Explorer' },
  ]
}

export const loader: LoaderFunction = async ({ request, params, context }) => {
  // const res = await gqlClient.query(xblockcount, {})

  const [xmsgRes, supportedChainsRes] = await Promise.all([
    gqlClient.query(xmsgrange, {
      from: '0x' + (0).toString(16),
      to: '0x' + (1000).toString(16),
    }),
    gqlClient.query(supportedchains, {}),
  ])

  const supportedChains = supportedChainsRes.data?.supportedchains || []
  const xmsgs = xmsgRes?.data?.xmsgrange ?? []

  console.log('Supported chains', supportedChainsRes.data?.supportedchains)
  console.log('xmsgData', xmsgRes?.data?.xmsgrange.length)

  const pollData = async () => {
    return json({
      // count: Number(res?.data?.xblockcount || '0x'),
      xmsgs,
    })
  }

  return await pollData()
}

export default function Index() {
  const revalidator = useRevalidator()
  const fetcher = useFetcher()

  useInterval(() => {
    revalidator.revalidate()
  }, 5000)

  return (
    <div className="px-20">
      <div className="flex h-full w-full flex-col">
        <Overview />
        <div className="w-full">
          <XMsgDataTable />
        </div>
        <div className="grow"></div>
      </div>
    </div>
  )
}
