import type { LoaderFunction, MetaFunction } from '@remix-run/node'
import XBlockDataTable from '~/components/home/blockDataTable'
import XMsgDataTable from '~/components/home/messageDataTable'
import Overview from '~/components/home/overview'
import { json } from '@remix-run/node'
import { gqlClient } from '~/entry.server'
import { useFetcher, useRevalidator, useSearchParams } from '@remix-run/react'
import { useInterval } from '~/hooks/useInterval'
import { xblockcount } from '~/components/queries/block'
import { xmsgrange } from '~/components/queries/messages'
import { XMsg } from '~/graphql/graphql'
import { supportedchains } from '~/components/queries/chains'
import { mappedSourceChains } from '~/lib/sourceChains'

export const meta: MetaFunction = () => {
  return [
    { title: 'Omni Network Explorer' },
    { name: 'description', content: 'Omni Network Explorer' },
  ]
}

export const loader: LoaderFunction = async ({ request }) => {
  // const res = await gqlClient.query(xblockcount, {})

  console.log(request)
  const [xmsgRes, supportedChainsRes] = await Promise.all([
    gqlClient.query(xmsgrange, {
      from: '0x' + (0).toString(16),
      to: '0x' + (1000).toString(16),
    }),
    gqlClient.query(supportedchains, {}),
  ])

  const supportedChains = mappedSourceChains(supportedChainsRes.data?.supportedchains || [])
  const xmsgs = xmsgRes?.data?.xmsgrange ?? []

  // console.log('Supported chains', supportedChains)
  // console.log('xmsgData', xmsgRes?.data?.xmsgrange.length)

  const pollData = async () => {
    return json({
      // count: Number(res?.data?.xblockcount || '0x'),
      supportedChains,
      xmsgs,
    })
  }

  return await pollData()
}

export default function Index() {
  const revalidator = useRevalidator()

  // poll server every 5 seconds
  useInterval(() => {
    revalidator.revalidate()
  }, 10000)

  return (
    <div className="px-8 md:px-20">
      <div className="flex h-full w-full flex-col">
        {/* <Overview /> */}
        <div className={'h-20'}>

        </div>

        <div className="w-full">
          <XMsgDataTable />
        </div>
        <div className="grow"></div>
      </div>
    </div>
  )
}
