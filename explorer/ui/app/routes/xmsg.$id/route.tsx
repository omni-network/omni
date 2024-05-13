import { useLoaderData, useNavigate, useParams, useRouteError, useSearchParams } from '@remix-run/react'
import { BackBtn } from '~/components/details/BackBtn'
import Tag from '~/components/shared/tag'
import { From } from '~/components/details/From'
import { To } from '~/components/details/To'
import { TabList } from '~/components/details/TabList'
import { xmsg } from '~/components/queries/messages'
import { gqlClient } from '~/entry.server'
import { LoaderFunction, MetaFunction, json } from '@remix-run/node'
import { XMsg } from '~/graphql/graphql'

export const loader: LoaderFunction = async ({ request }) => {
  const pageUrl = request.url
  const parts = pageUrl.split('/')
  const XMsgIdDetails = parts[parts.length - 1].split('-')

  const url = new URL(request.url)
  const params: any = {}

  for (const [key, value] of url.searchParams) {
    params[key] = value
  }

  const variables = {
    sourceChainID: XMsgIdDetails[0],
    destChainID: XMsgIdDetails[1],
    offset: XMsgIdDetails[2],
  }

  const [xmsgRes] = await Promise.all([gqlClient.query(xmsg, variables)])
  const pollData = async () => {
    return json({
      xMsg: xmsgRes,
    })
  }

  return await pollData()
}

export const meta: MetaFunction = () => {
  return [
    { title: 'Omni Network Explorer' },
    { name: 'description', content: 'Omni Network Explorer' },
  ]
}

export default function Index() {
  const loaderData = useLoaderData<any>()

  const xMsgDetails:XMsg = loaderData.xMsg.data.xmsg

  const navigate = useNavigate()

  const onBackClickHandler = () => {
    navigate('/')
  }

  return (
    <div className="px-8 md:px-20">
      <div className="flex h-full w-full flex-col">
        {/* <Overview /> */}
        <div className={'h-5'}></div>
        <div className="w-full">
          <BackBtn onBackClickHandler={onBackClickHandler} />
          <h4 className="mt-5">
            XMsg <span className="p-2">/</span>{' '}
            <span className="text-default">{xMsgDetails?.displayID}</span>
          </h4>

          <div className="mt-5 p-4 w-full bg-raised rounded-lg">
            {/* Offset */}
            <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
              <p className="w-[150px] sm:w-48 text-sm">Offset</p>
              <p className="text-default">{xMsgDetails?.offset}</p>
            </div>
            {/* Status */}
            <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
              <p className="w-[150px] sm:w-48 text-sm">Status</p>
              <div className="flex flex-col sm:flex-row items-start">
                <Tag status={xMsgDetails?.status} />
                <p className="sm:ml-5">{xMsgDetails?.receipt?.revertReason}</p>
              </div>
            </div>
            {/* Data */}
            <div className="flex mt-5 pb-2 border-b-[1px] border-subtle border-solid">
              <p className="w-[150px] sm:w-48 text-sm">Data</p>
            </div>
            <From xMsgDetails={xMsgDetails} />
            <To xMsgDetails={xMsgDetails} />
            <TabList xMsgDetails={xMsgDetails} />
          </div>
        </div>
      </div>
    </div>
  )
}
